// Code generated by proto-gen-go-kit-transport. DO NOT EDIT.
// * Source: "github.com/aksanart/tracker_service/contract/contract/grpc.proto"

package transport

import (
	"context"

	"github.com/aksanart/tracker_service/contract"
)

type UseCaseContract interface {
	HealthCheck(ctx context.Context, request *contract.EmptyRequest) (response *contract.DefaultResponse, err error)         // (unary)
	CreateTracker(ctx context.Context, request *contract.CreateTrackerReq) (response *contract.DefaultResponse, err error)   // (unary)
	UpdateTracker(ctx context.Context, request *contract.UpdateTrackerReq) (response *contract.DefaultResponse, err error)   // (unary)
	ListTracker(ctx context.Context, request *contract.EmptyRequest) (response *contract.ListTrackerResponse, err error)     // (unary)
	DetailTracker(ctx context.Context, request *contract.IDTrackerReq) (response *contract.DetailTrackerResponse, err error) // (unary)
	DeleteTracker(ctx context.Context, request *contract.IDTrackerReq) (response *contract.DefaultResponse, err error)       // (unary)
}

type transport struct {
	contract.UnimplementedTrackerServiceServer
	UseCase UseCaseContract
}

func NewTransport(ctx context.Context, useCase UseCaseContract) *transport {
	return &transport{
		UseCase: useCase,
	}
}
