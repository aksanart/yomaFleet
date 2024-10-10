package repointerface

import (
	"context"

	"github.com/aksanart/tracker_service/model"
)

type MongoInterface interface {
	HealthCheck(ctx context.Context) error
	FindAllTracker(ctx context.Context) (res []*model.Tracker, err error)
	Create(ctx context.Context, data *model.Tracker) error
	Update(ctx context.Context, id string, data *model.Tracker) (err error)
	Delete(ctx context.Context, id string) (err error)
	FindById(ctx context.Context, id string) (res *model.Tracker, err error)
	FindByVehicleId(ctx context.Context, vehicleId string) (res *model.Tracker, err error)
	UpdateLocation(ctx context.Context, vehicleId string, data *model.Tracker) error
}
