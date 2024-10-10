package implementor

import (
	"context"

	"github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/contract"
	"google.golang.org/grpc/metadata"
)

func (i *VehicleServiceImplementor) HealthCheck(ctx context.Context, md metadata.MD, input *contract.EmptyRequest) (resp *contract.DefaultResponse, err error) {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return i.Cli.HealthCheck(ctx, input)
}

func (i *VehicleServiceImplementor) CreateVehicle(ctx context.Context, md metadata.MD, input *contract.CreateVehicleReq) (resp *contract.CreateVehicleResponse, err error) {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return i.Cli.CreateVehicle(ctx, input)
}

func (i *VehicleServiceImplementor) UpdateVehicle(ctx context.Context, md metadata.MD, input *contract.UpdateVehicleReq) (resp *contract.DefaultResponse, err error) {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return i.Cli.UpdateVehicle(ctx, input)
}

func (i *VehicleServiceImplementor) ListVehicle(ctx context.Context, md metadata.MD, input *contract.ListVehicleReq) (resp *contract.ListVehicleResponse, err error) {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return i.Cli.ListVehicle(ctx, input)
}

func (i *VehicleServiceImplementor) DetailVehicle(ctx context.Context, md metadata.MD, input *contract.IDVehicleReq) (resp *contract.DetailVehicleResponse, err error) {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return i.Cli.DetailVehicle(ctx, input)
}

func (i *VehicleServiceImplementor) DeleteVehicle(ctx context.Context, md metadata.MD, input *contract.IDVehicleReq) (resp *contract.DefaultResponse, err error) {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return i.Cli.DeleteVehicle(ctx, input)
}
