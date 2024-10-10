package tracker

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/contract"
	"github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/implementor"
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

func NewImplementor(conf ImplementorConfig) (*implementor.VehicleServiceImplementor, error) {
	baseUrl := fmt.Sprintf("%s:%d", conf.Host, conf.Port)
	if conf.Port == 0 {
		baseUrl = conf.Host
	}
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
		log.Println(baseUrl)
		logger.GetLogger().Error("vehicle-baseurl", logger.Field{Key: "baseurl", Value: baseUrl})
		return nil, err
	}
	return &implementor.VehicleServiceImplementor{
		GrpcConn: conn,
		Cli:      contract.NewVehicleServiceClient(conn),
	}, nil
}
