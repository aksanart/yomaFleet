package redisrepo

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisClient struct {
	conn *redis.Client
}

// GetData implements repointerface.RedisInterface.
func (r *RedisClient) GetData(ctx context.Context, key string, pointerData interface{}) (interface{}, error) {
	data := r.conn.Get(ctx, key).Val()
	if data == "" {
		return nil, errors.New("no data")
	}
	err := json.Unmarshal([]byte(data), pointerData)
	if err != nil {
		return nil, err
	}
	return pointerData, nil
}

func (r *RedisClient) GetStreamClient(ctx context.Context, streamName string) ([]redis.XStream, error) {
	return r.conn.XRead(ctx, &redis.XReadArgs{
		Streams: []string{streamName, "$"},
		Count:   2,
		Block:   time.Minute * 1, // timeout if not found
	}).Result()
}

func (r *RedisClient) RangeData(ctx context.Context, streamName string) ([]redis.XMessage, error) {
	return r.conn.XRange(ctx, streamName, "-", "+").Result()
}

func (r *RedisClient) Del(ctx context.Context, key string) error {
	return r.conn.Del(ctx, key).Err()
}

func (r *RedisClient) SetStreamData(ctx context.Context, key string, values interface{}) error {
	return r.conn.XAdd(ctx, &redis.XAddArgs{
		Stream: key,
		Values: values,
	}).Err()
}

func (r *RedisClient) ExpireStream(ctx context.Context, key string, ttl time.Duration) error {
	return r.conn.Expire(ctx, key, ttl).Err()
}
