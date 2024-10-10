
package implementor

import (
	"context"
	"github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type VehicleServiceImplementor struct {
	GrpcConn *grpc.ClientConn
	Cli      contract.VehicleServiceClient
}

func (i *VehicleServiceImplementor) Close() error {
	if i.GrpcConn != nil {
		return i.GrpcConn.Close()
	}
	return nil

}

func (i *VehicleServiceImplementor) GetDefaultMetadata(ctx context.Context) metadata.MD {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return md
	} else {
		return metadata.MD{}
	}
}
