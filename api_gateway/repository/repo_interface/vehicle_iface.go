package repointerface

import (
	"context"

	"github.com/aksan/weplus/apigw/pkg/grpc_client/vehicle/contract"
)

type VehicleServiceIface interface {
	HealthCheck(context.Context, *contract.EmptyRequest) error
	CreateVehicle(context.Context, *contract.CreateVehicleReq) (*contract.CreateVehicleResponse, error)
	UpdateVehicle(context.Context, *contract.UpdateVehicleReq) (*contract.DefaultResponse, error)
	ListVehicle(context.Context, *contract.ListVehicleReq) (*contract.ListVehicleResponse, error)
	DetailVehicle(context.Context, *contract.IDVehicleReq) (*contract.DetailVehicleResponse, error)
	DeleteVehicle(context.Context, *contract.IDVehicleReq) (*contract.DefaultResponse, error)
}
