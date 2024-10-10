package main

import (
	"context"
	"time"

	"github.com/aksan/weplus/apigw/pkg/config"
	"github.com/aksan/weplus/apigw/pkg/logger"
	repoUtil "github.com/aksan/weplus/apigw/repository/util"
	"github.com/aksan/weplus/apigw/server"
	"github.com/aksan/weplus/apigw/usecase"
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
