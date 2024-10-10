// Code generated by proto-gen-go-lib-contract . DO NOT EDIT.
// Source: "github.com/aksanart/vehicle/contract"

package contract

import (
	"context"
)

type VehicleService interface {
	HealthCheck(context.Context, *EmptyRequest) (*DefaultResponse, error)
	CreateVehicle(context.Context, *CreateVehicleReq) (*CreateVehicleResponse, error)
	UpdateVehicle(context.Context, *UpdateVehicleReq) (*DefaultResponse, error)
	ListVehicle(context.Context, *ListVehicleReq) (*ListVehicleResponse, error)
	DetailVehicle(context.Context, *IDVehicleReq) (*DetailVehicleResponse, error)
	DeleteVehicle(context.Context, *IDVehicleReq) (*DefaultResponse, error)
}
