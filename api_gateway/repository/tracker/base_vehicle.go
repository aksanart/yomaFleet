package trackerservice

import (
	"context"
	"fmt"

	client "github.com/aksan/weplus/apigw/pkg/grpc_client/tracker"
	"github.com/aksan/weplus/apigw/pkg/grpc_client/tracker/contract"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"github.com/aksan/weplus/apigw/repository"
)

type trackerServiceConfig struct {
	host string
	port int
}

func NewTrackerServiceConfig(host string, port int) repository.RepoConf {
	return &trackerServiceConfig{
		host: host,
		port: port,
	}
}

func (conf *trackerServiceConfig) Init(r *repository.Repository) error {
	if conf.host == "" {
		return fmt.Errorf("trackerService repository host cannot be empty")
	}
	cli, err := client.NewImplementor(client.ImplementorConfig{
		Host: conf.host,
		Port: conf.port,
	})
	if err != nil {
		logger.GetLogger().Error("tracker-baseurl", logger.Field{Key: "cli", Value: cli})
		return err
	}
	client := &TrackerServiceclient{cli}
	if err := client.HealthCheck(context.Background(), &contract.EmptyRequest{}); err != nil {
		return fmt.Errorf("tracker repository failed to be created: healthcheck error, %s", err.Error())
	}
	r.TrackerService = client
	return nil
}

func (conf *trackerServiceConfig) GetRepoName() string {
	return "VehicleService"
}
