package redisrepo

import (
	"context"
	"fmt"

	"github.com/aksanart/tracker_service/repository"
	"github.com/go-redis/redis/v8"
)

type redisConf struct {
	RedisHost     string
	RedisPort     string
	RedisPassword string
	RedisDB       int
}

func NewRedisConfig(host, port, password string, db int) repository.RepoConf {
	return &redisConf{
		RedisHost:     host,
		RedisPort:     port,
		RedisPassword: password,
		RedisDB:       db,
	}
}

func (*redisConf) GetRepoName() string {
	return "Repo Redis"
}

// Connect implements repository.RepoConf
func (r *redisConf) Init(repo *repository.Repository) error {
	if err := r.Validate(); err != nil {
		return err
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", r.RedisHost, r.RedisPort),
		Password: r.RedisPassword,
		DB:       r.RedisDB,
	})
	if client == nil {
		return fmt.Errorf("error connect to reporedis")
	}

	if err := client.Ping(context.Background()).Err(); err != nil {
		return err
	}
	repo.Redis = &RedisClient{
		conn: client,
	}

	return nil
}

// Validate implements repository.RepoConf
func (rc *redisConf) Validate() error {
	if rc.RedisHost == "" {
		return fmt.Errorf("reporedis host can't be empty")
	} else if rc.RedisPort == "" {
		return fmt.Errorf("reporedis port can't be empty")
	}
	return nil
}
