// Code generated by proto-gen-svc-transport. DO NOT EDIT.
// Folder	: "github.com/aksanart/vehicle/contract"
// File		: "contract/grpc.proto"

package transport

import (
	"context"
	"github.com/aksanart/vehicle/contract"
)

func (t transport) DetailVehicle(ctx context.Context, request *contract.IDVehicleReq) (response *contract.DetailVehicleResponse, err error) {
	// * Bridging to usecase related method
	return t.UseCase.DetailVehicle(ctx, request)
}
