package tracker

import (
	"crypto/tls"
	"fmt"

	"github.com/aksan/weplus/apigw/pkg/grpc_client/tracker/contract"
	"github.com/aksan/weplus/apigw/pkg/grpc_client/tracker/implementor"
	"github.com/aksan/weplus/apigw/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
)

type ImplementorConfig struct {
	Host     string
	Port     int
	CertPath string
}

func NewImplementor(conf ImplementorConfig) (*implementor.TrackerServiceImplementor, error) {
	baseUrl := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	logger.GetLogger().Debug("tracker-baseurl", logger.Field{Key: "baseurl", Value: baseUrl})
	opts := []grpc.DialOption{}

	if conf.CertPath == "" && conf.Port == 443 {
		baseUrl = conf.Host
		opts = append(opts, grpc.WithTransportCredentials(
			credentials.NewTLS(&tls.Config{
				InsecureSkipVerify: true,
			}),
		))
	}

	if conf.CertPath == "" && conf.Port != 443 {
		opts = append(opts, grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		))
	}

	conn, err := grpc.NewClient(baseUrl, opts...)
	if err != nil {
		return nil, err
	}
	return &implementor.TrackerServiceImplementor{
		GrpcConn: conn,
		Cli:      contract.NewTrackerServiceClient(conn),
	}, nil
}
