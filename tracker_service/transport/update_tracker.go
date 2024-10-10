// Code generated by proto-gen-svc-transport. DO NOT EDIT.
// Folder	: "github.com/aksanart/tracker_service/contract"
// File		: "contract/grpc.proto"

package transport

import (
	"context"

	"github.com/aksanart/tracker_service/contract"
)

func (t transport) UpdateTracker(ctx context.Context, request *contract.UpdateTrackerReq) (response *contract.DefaultResponse, err error) {
	// * Bridging to usecase related method
	return t.UseCase.UpdateTracker(ctx, request)
}
