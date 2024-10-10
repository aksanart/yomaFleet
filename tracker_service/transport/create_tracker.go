// Code generated by proto-gen-svc-transport. DO NOT EDIT.
// Folder	: "github.com/aksanart/tracker_service/contract"
// File		: "contract/grpc.proto"

package transport

import (
	"context"

	"github.com/aksanart/tracker_service/contract"
)

func (t transport) CreateTracker(ctx context.Context, request *contract.CreateTrackerReq) (response *contract.DefaultResponse, err error) {
	// * Bridging to usecase related method
	return t.UseCase.CreateTracker(ctx, request)
}
