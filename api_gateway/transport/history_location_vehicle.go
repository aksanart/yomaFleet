// Code generated by proto-gen-svc-transport. DO NOT EDIT.
// Folder	: "github.com/aksan/weplus/apigw/contract"
// File		: "contract/grpc.proto"

package transport

import (
	"context"

	"github.com/aksan/weplus/apigw/contract"
)

func (t transport) HistoryLocationVehicle(ctx context.Context, request *contract.HistoryLocationVehicleReq) (response *contract.HistoryLocationVehicleResponse, err error) {
	// * Bridging to usecase related method
	return t.UseCase.HistoryLocationVehicle(ctx, request)
}
