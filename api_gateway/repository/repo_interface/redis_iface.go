package repointerface

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisInterface interface {
	GetData(ctx context.Context, key string, pointerData interface{}) error
	GetStreamClient(ctx context.Context, streamName string) ([]redis.XStream, error)
	RangeData(ctx context.Context, streamName string) ([]redis.XMessage, error)
	Del(ctx context.Context, key string) error
	Set(ctx context.Context, key string, data []byte, ttl time.Duration) error
	SetStreamData(ctx context.Context, key string, values interface{}) error
	Get(ctx context.Context, key string) (string)
}
