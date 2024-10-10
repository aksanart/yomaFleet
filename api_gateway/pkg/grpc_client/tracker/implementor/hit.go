package implementor

import (
	"context"

	"github.com/aksan/weplus/apigw/pkg/grpc_client/tracker/contract"
	"google.golang.org/grpc/metadata"
)

func (i *TrackerServiceImplementor) HealthCheck(ctx context.Context, md metadata.MD, input *contract.EmptyRequest) (resp *contract.DefaultResponse, err error) {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return i.Cli.HealthCheck(ctx, input)
}

func (i *TrackerServiceImplementor) CreateTracker(ctx context.Context, md metadata.MD, input *contract.CreateTrackerReq) (resp *contract.CreateTrackerResponse, err error) {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return i.Cli.CreateTracker(ctx, input)
}

func (i *TrackerServiceImplementor) UpdateTracker(ctx context.Context, md metadata.MD, input *contract.UpdateTrackerReq) (resp *contract.DefaultResponse, err error) {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return i.Cli.UpdateTracker(ctx, input)
}

func (i *TrackerServiceImplementor) ListTracker(ctx context.Context, md metadata.MD, input *contract.EmptyRequest) (resp *contract.ListTrackerResponse, err error) {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return i.Cli.ListTracker(ctx, input)
}

func (i *TrackerServiceImplementor) DetailTracker(ctx context.Context, md metadata.MD, input *contract.IDTrackerReq) (resp *contract.DetailTrackerResponse, err error) {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return i.Cli.DetailTracker(ctx, input)
}

func (i *TrackerServiceImplementor) DeleteTracker(ctx context.Context, md metadata.MD, input *contract.IDTrackerReq) (resp *contract.DefaultResponse, err error) {
	ctx = metadata.NewOutgoingContext(ctx, md)
	return i.Cli.DeleteTracker(ctx, input)
}
