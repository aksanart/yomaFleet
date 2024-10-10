
package implementor

import (
	"context"
	"github.com/aksan/weplus/apigw/pkg/grpc_client/tracker/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type TrackerServiceImplementor struct {
	GrpcConn *grpc.ClientConn
	Cli      contract.TrackerServiceClient
}

func (i *TrackerServiceImplementor) Close() error {
	if i.GrpcConn != nil {
		return i.GrpcConn.Close()
	}
	return nil

}

func (i *TrackerServiceImplementor) GetDefaultMetadata(ctx context.Context) metadata.MD {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		return md
	} else {
		return metadata.MD{}
	}
}
