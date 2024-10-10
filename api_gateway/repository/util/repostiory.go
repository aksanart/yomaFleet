package util

import (
	"log"

	"github.com/aksan/weplus/apigw/pkg/config"
	"github.com/aksan/weplus/apigw/repository"
	"github.com/aksan/weplus/apigw/repository/mongodb"
	redisrepo "github.com/aksan/weplus/apigw/repository/redis_repo"
	trackerservice "github.com/aksan/weplus/apigw/repository/tracker"
	vehicleservice "github.com/aksan/weplus/apigw/repository/vehicle_service"
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
		vehicleservice.NewVehicleServiceConfig(
			config.GetConfig("vehicle_host").GetString(),
			int(config.GetConfig("vehicle_port").GetInt()),
		),
		trackerservice.NewTrackerServiceConfig(
			config.GetConfig("tracker_host").GetString(),
			int(config.GetConfig("tracker_port").GetInt()),
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
