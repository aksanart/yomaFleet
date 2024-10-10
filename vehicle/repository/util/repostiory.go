package util

import (
	"log"

	"github.com/aksanart/vehicle/pkg/config"
	"github.com/aksanart/vehicle/repository"
	"github.com/aksanart/vehicle/repository/mongodb"
	redisrepo "github.com/aksanart/vehicle/repository/redis_repo"
)

var repo *repository.Repository

func LoadRepository() {
	repoList, err := repository.NewRepository([]repository.RepoConf{
		// postgree.NewDatabaseConf(
		// 	config.GetConfig("db_host").GetString(),
		// 	config.GetConfig("db_username").GetString(),
		// 	config.GetConfig("db_password").GetString(),
		// 	config.GetConfig("db_ssl_mode").GetString(),
		// 	config.GetConfig("db_port").GetString(),
		// 	config.GetConfig("db_name").GetString(),
		// ),
		mongodb.NewMongoDB(
			config.GetConfig("mongo_host").GetString(),
			config.GetConfig("mongo_port").GetString(),
			config.GetConfig("mongo_user").GetString(),
			config.GetConfig("mongo_pass").GetString(),
			config.GetConfig("mongo_db").GetString(),
		),
		redisrepo.NewRedisConfig(
			config.GetConfig("redis_host").GetString(),
			config.GetConfig("redis_port").GetString(),
			config.GetConfig("redis_pass").GetString(),
			int(config.GetConfig("redis_db").GetInt()),
		),
	})
	if err != nil {
		log.Fatalf("cannot initiate repository, with error: %v", err)
	}
	repo = repoList
}

func GetRepo() *repository.Repository {
	return repo
}
