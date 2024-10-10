package repointerface

import (
	"context"

	"github.com/aksan/weplus/apigw/pkg/grpc_client/tracker/contract"
)

type TrackerServiceIface interface {
	HealthCheck(context.Context, *contract.EmptyRequest) error
	CreateTracker(context.Context, *contract.CreateTrackerReq) (*contract.CreateTrackerResponse, error)
	UpdateTracker(context.Context, *contract.UpdateTrackerReq) (*contract.DefaultResponse, error)
	ListTracker(context.Context, *contract.EmptyRequest) (*contract.ListTrackerResponse, error)
	DetailTracker(context.Context, *contract.IDTrackerReq) (*contract.DetailTrackerResponse, error)
	DeleteTracker(context.Context, *contract.IDTrackerReq) (*contract.DefaultResponse, error)
}
