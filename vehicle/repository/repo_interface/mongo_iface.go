package repointerface

import (
	"context"

	"github.com/aksanart/vehicle/model"
)

type MongoInterface interface {
	HealthCheck(ctx context.Context) error
	FindAllVehilce(ctx context.Context, offest int) (res []*model.Vehicle, err error)
	Create(ctx context.Context, data *model.Vehicle) (id string, err error)
	Update(ctx context.Context, id string, data *model.Vehicle) (err error)
	Delete(ctx context.Context, id string) (err error)
	FindById(ctx context.Context, id string) (res *model.Vehicle, err error)
}
