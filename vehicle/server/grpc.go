package server

import (
	"context"
	"fmt"
	"net"
	"regexp"
	"runtime/debug"

	"github.com/aksanart/vehicle/contract"
	"github.com/aksanart/vehicle/pkg/config"
	"github.com/aksanart/vehicle/pkg/logger"
	"github.com/aksanart/vehicle/transport"
	"github.com/aksanart/vehicle/usecase"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func RunGRPCServer(ctx context.Context, usecase *usecase.UseCase) *grpc.Server {
	grpcPort := fmt.Sprintf(":%s", config.GetConfig("grpc_port").GetString())
	grpcConn, err := net.Listen("tcp", grpcPort)
	if err != nil {
		logger.GetLogger().Fatal(fmt.Sprintf("failed to listen port: %v", err))
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			recovery.UnaryServerInterceptor(recovery.WithRecoveryHandlerContext(grpcRecoveryHandler)),
		),
		grpc.ChainStreamInterceptor(
			recovery.StreamServerInterceptor(recovery.WithRecoveryHandlerContext(grpcRecoveryHandler)),
		),
	)
	transport := transport.NewTransport(ctx, usecase)
	contract.RegisterVehicleServiceServer(grpcServer, transport)
	reflection.Register(grpcServer)

	healthServer := health.NewServer()
	healthpb.RegisterHealthServer(grpcServer, healthServer)
	for name := range grpcServer.GetServiceInfo() {
		healthServer.SetServingStatus(name, healthpb.HealthCheckResponse_SERVING)
	}
	healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)

	go grpcServer.Serve(grpcConn)
	logger.GetLogger().Info(fmt.Sprintf("server grpc listening at %v", grpcConn.Addr()))
	return grpcServer
}

func grpcRecoveryHandler(ctx context.Context, panic interface{}) error {
	newLineRegex := regexp.MustCompile(`\r?\n`)
	stackTrace := newLineRegex.ReplaceAllString(string(debug.Stack()), " ")
	logger.GetLogger().Error("panic happened", logger.ConvertMapToFields(map[string]interface{}{
		"panic_message":    panic,
		"panic_stacktrace": stackTrace,
	})...)
	return status.Errorf(codes.Internal, "server error happened")
}
