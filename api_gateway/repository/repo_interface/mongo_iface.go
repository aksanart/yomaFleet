package repointerface

import (
	"context"

	"github.com/aksan/weplus/apigw/model"
)

type MongoInterface interface {
	HealthCheck(ctx context.Context) error
	Create(ctx context.Context, data *model.User) error
	Update(ctx context.Context, email string, data *model.User) (err error)
	Delete(ctx context.Context, email string) (err error)
	FindByEmail(ctx context.Context, email string) (res *model.User, err error)
}
