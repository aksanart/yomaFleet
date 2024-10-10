package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"

	"github.com/aksan/weplus/apigw/contract"
	"github.com/aksan/weplus/apigw/pkg/config"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"github.com/aksan/weplus/apigw/transport"
	"github.com/aksan/weplus/apigw/usecase"
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
	contract.RegisterApiGatewayHandlerServer(ctx, gwMux, transport)
	go restServer.Serve(restConn)
	logger.GetLogger().Info(fmt.Sprintf("server rest listening at %v", restConn.Addr()))
	return restServer
}
