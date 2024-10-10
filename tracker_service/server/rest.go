package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"

	"github.com/aksanart/tracker_service/contract"
	"github.com/aksanart/tracker_service/pkg/config"
	"github.com/aksanart/tracker_service/pkg/logger"
	"github.com/aksanart/tracker_service/transport"
	"github.com/aksanart/tracker_service/usecase"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

func RunRESTServer(ctx context.Context, usecase *usecase.UseCase) *http.Server {
	restPort := fmt.Sprintf(":%s", config.GetConfig("rest_port").GetString())
	restConn, err := net.Listen("tcp", restPort)
	if err != nil {
		logger.GetLogger().Fatal(fmt.Sprintf("failed to listen port: %v", err))
	}

	gwMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &CustomMarshaler{}),
		runtime.WithErrorHandler(ErrorCustomFormat),
		runtime.WithIncomingHeaderMatcher(CustomMatcherMrg),
	)
	gwMux.HandlePath("GET", "/debug/pprof", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) { pprof.Index(w, r) })
	gwMux.HandlePath("GET", "/debug/pprof/{function}", func(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
		switch pathParams["function"] {
		case "cmdline":
			pprof.Cmdline(w, r)
		case "profile":
			pprof.Profile(w, r)
		case "symbol":
			pprof.Symbol(w, r)
		case "trace":
			pprof.Trace(w, r)
		default:
			pprof.Index(w, r)
		}
	})

	restServer := &http.Server{
		Addr:    "localhost" + restPort,
		Handler: gwMux,
	}
	transport := transport.NewTransport(ctx, usecase)
	contract.RegisterTrackerServiceHandlerServer(ctx, gwMux, transport)
	go restServer.Serve(restConn)
	logger.GetLogger().Info(fmt.Sprintf("server rest listening at %v", restConn.Addr()))
	return restServer
}
