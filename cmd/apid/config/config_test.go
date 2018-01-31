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
	os.Setenv("APIHost", apiHost)
	defer os.Unsetenv("APIHost")

	mongoHost := "mongodb://mongo:30015"
	os.Setenv("MongoHost", mongoHost)
	defer os.Unsetenv("MongoHost")

	redisHost := "redis:7005"
	os.Setenv("RedisHost", redisHost)
	defer os.Unsetenv("RedisHost")

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
