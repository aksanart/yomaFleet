package config

import (
	"fmt"
	"os"
	"strings"
)

var appConfig map[string]data

var defaultConfig = map[string]interface{}{
	"check_healthy_repo": true,
	"app_name":           "tracker",
	"grpc_port":          6000,
	"rest_port":          8000,
	"log_level":          "INFO",
	"log_directory":      "",

	"mongo_host": "",
	"mongo_port": "",
	"mongo_user": "",
	"mongo_pass": "",
	"mongo_db":   "",

	"redis_host": "",
	"redis_port": "",
	"redis_pass": "",
	"redis_db":   0,

	"vehicle_host": "",
	"vehicle_port": 0,
	"tracker_host": "",
	"tracker_port": 0,

	"jwt_secret": "sC4LKS$DF9IR",
}

type data struct {
	value interface{}
}

func LoadConfigMap() {
	appConfig = LoadConfig(defaultConfig)
}

func GetConfig(key string) (val data) {
	if v, ok := appConfig[strings.ToLower(key)]; ok {
		val = v
	}
	return
}

func OsGetString(key, _default string) string {
	val := os.Getenv(key)
	if val == "" {
		return _default
	}
	return val
}

func LoadConfig(conf map[string]interface{}) map[string]data {
	retConfig := map[string]data{}
	for key, value := range conf {
		retConfig[key] = data{value: OsGetString(strings.ToUpper(key), fmt.Sprintf("%v", value))}
	}
	return retConfig
}
