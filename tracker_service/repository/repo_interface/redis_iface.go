package repointerface

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisInterface interface {
	GetData(ctx context.Context, key string, pointerData interface{}) (interface{}, error)
	GetStreamClient(ctx context.Context, streamName string) ([]redis.XStream, error)
	RangeData(ctx context.Context, streamName string) ([]redis.XMessage, error)
	Del(ctx context.Context, key string) error
	SetStreamData(ctx context.Context, key string, values interface{}) error
	ExpireStream(ctx context.Context, key string, ttl time.Duration) error
}
