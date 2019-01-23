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

	auth0ClientID := "abc1234"
	os.Setenv("AUTH_ZERO_CLIENT_ID", auth0ClientID)
	defer os.Unsetenv("AUTH_ZERO_CLIENT_ID")

	auth0RedirectURL := "http://faker.us"
	os.Setenv("AUTH_ZERO_REDIRECT_URL", auth0RedirectURL)
	defer os.Unsetenv("AUTH_ZERO_REDIRECT_URL")

	auth0URL := "http://auth0.org"
	os.Setenv("AUTH_ZERO_URL", auth0URL)
	defer os.Unsetenv("AUTH_ZERO_URL")

	auth0State := "12345state"
	os.Setenv("AUTH_ZERO_STATE", auth0State)
	defer os.Unsetenv("AUTH_ZERO_STATE")

	auth0CallbackURL := "http://cbsomewhere/callback"
	os.Setenv("AUTH_ZERO_CALLBACK_URL", auth0CallbackURL)
	defer os.Unsetenv("AUTH_ZERO_CALLBACK_URL")

	auth0Audience := "http://api-workshop.com"
	os.Setenv("AUTH_ZERO_AUDIENCE", auth0Audience)
	defer os.Unsetenv("AUTH_ZERO_AUDIENCE")

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
	if cfg.Auth0ClientID != auth0ClientID {
		t.Errorf("Expected %s but got %s", auth0ClientID, cfg.Auth0ClientID)
	}
	if cfg.Auth0RedirectURL != auth0RedirectURL {
		t.Errorf("Expected %s but got %s", auth0RedirectURL, cfg.Auth0RedirectURL)
	}
	if cfg.Auth0URL != auth0URL {
		t.Errorf("Expected %s but got %s", auth0URL, cfg.Auth0URL)
	}
	if cfg.Auth0State != auth0State {
		t.Errorf("Expected %s but got %s", auth0State, cfg.Auth0State)
	}
	if cfg.Auth0CallbackURL != auth0CallbackURL {
		t.Errorf("Expected %s but got %s", auth0CallbackURL, cfg.Auth0CallbackURL)
	}
	if auth0Audience := os.Getenv("AUTH_ZERO_AUDIENCE"); auth0Audience != "" {
		cfg.Auth0Audience = auth0Audience
	}
}
