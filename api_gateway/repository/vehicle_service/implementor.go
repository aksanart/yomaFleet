package vehicleservice

import (
	"context"

	"github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/implementor"

	"github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/contract"
	"google.golang.org/grpc/metadata"
)

type VehicleServiceclient struct {
	cli *implementor.VehicleServiceImplementor
}

func (i *VehicleServiceclient) HealthCheck(ctx context.Context, input *contract.EmptyRequest) (err error) {
	_, err = i.cli.HealthCheck(ctx, metadata.MD{}, input)
	return err
}

func (i *VehicleServiceclient) CreateVehicle(ctx context.Context, input *contract.CreateVehicleReq) (resp *contract.CreateVehicleResponse, err error) {
	return i.cli.CreateVehicle(ctx, metadata.MD{}, input)
}

func (i *VehicleServiceclient) UpdateVehicle(ctx context.Context, input *contract.UpdateVehicleReq) (resp *contract.DefaultResponse, err error) {
	return i.cli.UpdateVehicle(ctx, metadata.MD{}, input)
}

func (i *VehicleServiceclient) ListVehicle(ctx context.Context, input *contract.ListVehicleReq) (resp *contract.ListVehicleResponse, err error) {
	return i.cli.ListVehicle(ctx, metadata.MD{}, input)
}

func (i *VehicleServiceclient) DetailVehicle(ctx context.Context, input *contract.IDVehicleReq) (resp *contract.DetailVehicleResponse, err error) {
	return i.cli.DetailVehicle(ctx, metadata.MD{}, input)
}

func (i *VehicleServiceclient) DeleteVehicle(ctx context.Context, input *contract.IDVehicleReq) (resp *contract.DefaultResponse, err error) {
	return i.cli.DeleteVehicle(ctx, metadata.MD{}, input)
}
