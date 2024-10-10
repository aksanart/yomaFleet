package main

import (
	"context"
	"time"

	"github.com/aksanart/tracker_service/pkg/config"
	"github.com/aksanart/tracker_service/pkg/logger"
	repoUtil "github.com/aksanart/tracker_service/repository/util"
	"github.com/aksanart/tracker_service/server"
	"github.com/aksanart/tracker_service/usecase"
	"github.com/joho/godotenv"
)

func init() {
	godotenv.Load()
	config.LoadConfigMap()
	logger.LoadLogger()
	repoUtil.LoadRepository()
}

func main() {
	usecase := usecase.NewUsecase(repoUtil.GetRepo(), "")
	ctx, _ := context.WithCancel(context.Background())
	grpcServer := server.RunGRPCServer(ctx, usecase)
	restServer := server.RunRESTServer(ctx, usecase)

	wait := config.GracefulShutdown(ctx, 5*time.Second, map[string]config.Operation{
		"grpc": func(ctx context.Context) error {
			grpcServer.GracefulStop()
			return nil
		},
		"rest": func(ctx context.Context) error {
			return restServer.Shutdown(ctx)
		},
	})
	<-wait
}
