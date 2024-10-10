package trackerservice

import (
	"context"

	"github.com/aksan/weplus/apigw/pkg/grpc_client/tracker/implementor"

	"github.com/aksan/weplus/apigw/pkg/grpc_client/tracker/contract"
	"google.golang.org/grpc/metadata"
)

type TrackerServiceclient struct {
	cli *implementor.TrackerServiceImplementor
}

func (i *TrackerServiceclient) HealthCheck(ctx context.Context, input *contract.EmptyRequest) (err error) {
	_, err = i.cli.HealthCheck(ctx, metadata.MD{}, input)
	return err
}

func (i *TrackerServiceclient) CreateTracker(ctx context.Context, input *contract.CreateTrackerReq) (resp *contract.CreateTrackerResponse, err error) {
	return i.cli.CreateTracker(ctx, metadata.MD{}, input)
}

func (i *TrackerServiceclient) UpdateTracker(ctx context.Context, input *contract.UpdateTrackerReq) (resp *contract.DefaultResponse, err error) {
	return i.cli.UpdateTracker(ctx, metadata.MD{}, input)
}

func (i *TrackerServiceclient) ListTracker(ctx context.Context, input *contract.EmptyRequest) (resp *contract.ListTrackerResponse, err error) {
	return i.cli.ListTracker(ctx, metadata.MD{}, input)
}

func (i *TrackerServiceclient) DetailTracker(ctx context.Context, input *contract.IDTrackerReq) (resp *contract.DetailTrackerResponse, err error) {
	return i.cli.DetailTracker(ctx, metadata.MD{}, input)
}

func (i *TrackerServiceclient) DeleteTracker(ctx context.Context, input *contract.IDTrackerReq) (resp *contract.DefaultResponse, err error) {
	return i.cli.DeleteTracker(ctx, metadata.MD{}, input)
}
