package config_test

import (
	"os"
	"testing"

	"github.com/jbelmont/api-workshop/cmd/apid/config"
)

func TestNewConfig(t *testing.T) {
	cfg := config.New()

	if cfg.APIHost != ":8080" {
		t.Error("config should set default value of APIHost variable")
	}
	if cfg.MongoHost != "mongodb://mongo:27017" {
		t.Error("config should set default value for MongoHost variable")
	}
	if cfg.RedisHost != "redis:6379" {
		t.Error("config should set default value for RedisHost variable")
	}
}

func TestGetEnvironmentVariablesThatAreSet(t *testing.T) {
	apiHost := ":3000"
	os.Setenv("API_HOST", apiHost)
	defer os.Unsetenv("API_HOST")

	mongoHost := "mongodb://mongo:30015"
	os.Setenv("MONGO_HOST", mongoHost)
	defer os.Unsetenv("MONGO_HOST")

	redisHost := "redis:7005"
	os.Setenv("REDIS_HOST", redisHost)
	defer os.Unsetenv("REDIS_HOST")

	cfg := config.New()

	if cfg.APIHost != apiHost {
		t.Errorf("Expected %s but got %s", apiHost, cfg.APIHost)
	}
	if cfg.MongoHost != mongoHost {
		t.Errorf("Expected %s but got %s", mongoHost, cfg.MongoHost)
	}
	if cfg.RedisHost != redisHost {
		t.Errorf("Expected %s but got %s", redisHost, cfg.RedisHost)
	}
}
