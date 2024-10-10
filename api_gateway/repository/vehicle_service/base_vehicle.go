package vehicleservice

import (
	"context"
	"fmt"

	client "github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle"
	"github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/contract"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"github.com/aksan/weplus/apigw/repository"
)

type vehicleServiceConfig struct {
	host string
	port int
}

func NewVehicleServiceConfig(host string, port int) repository.RepoConf {
	return &vehicleServiceConfig{
		host: host,
		port: port,
	}
}

func (conf *vehicleServiceConfig) Init(r *repository.Repository) error {
	if conf.host == "" {
		return fmt.Errorf("vehicleService repository host cannot be empty")
	}
	cli, err := client.NewImplementor(client.ImplementorConfig{
		Host: conf.host,
		Port: conf.port,
	})
	if err != nil {
		logger.GetLogger().Error("vehicle-baseurl", logger.Field{Key: "cli", Value: cli})
		return err
	}
	client := &VehicleServiceclient{cli}
	if err := client.HealthCheck(context.Background(), &contract.EmptyRequest{}); err != nil {
		return fmt.Errorf("vehicle repository failed to be created: healthcheck error, %s", err.Error())
	}
	r.VehicleService = client
	return nil
}

func (conf *vehicleServiceConfig) GetRepoName() string {
	return "VehicleService"
}
