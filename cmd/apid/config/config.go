package config

import (
	"os"
	"time"
)

// Config holds environment variables
type Config struct {
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
	ShutdownTimeout  time.Duration
	DBTimeout        time.Duration
	APIHost          string
	MongoHost        string
	RedisHost        string
	Auth0ClientID    string
	Auth0RedirectURL string
	Auth0URL         string
	Auth0State       string
	Auth0CallbackURL string
	Auth0Audience    string
}

// New creates and returns a configuration
func New() Config {
	cfg := Config{
		ReadTimeout:      5 * time.Second,
		WriteTimeout:     10 * time.Second,
		ShutdownTimeout:  5 * time.Second,
		DBTimeout:        25 * time.Second,
		APIHost:          ":8080",
		MongoHost:        "mongodb://mongo:27017",
		RedisHost:        "redis:6379",
		Auth0ClientID:    "",
		Auth0RedirectURL: "",
		Auth0URL:         "",
		Auth0State:       "",
		Auth0CallbackURL: "",
		Auth0Audience:    "",
	}

	// Populate environment variables if set
	if apiHost := os.Getenv("API_HOST"); apiHost != "" {
		cfg.APIHost = apiHost
	}
	if mongoHost := os.Getenv("MONGO_HOST"); mongoHost != "" {
		cfg.MongoHost = mongoHost
	}
	if redisHost := os.Getenv("REDIS_HOST"); redisHost != "" {
		cfg.RedisHost = redisHost
	}
	if auth0ClientID := os.Getenv("AUTH_ZERO_CLIENT_ID"); auth0ClientID != "" {
		cfg.Auth0ClientID = auth0ClientID
	}
	if auth0RedirectURL := os.Getenv("AUTH_ZERO_REDIRECT_URL"); auth0RedirectURL != "" {
		cfg.Auth0RedirectURL = auth0RedirectURL
	}
	if auth0URL := os.Getenv("AUTH_ZERO_URL"); auth0URL != "" {
		cfg.Auth0URL = auth0URL
	}
	if auth0State := os.Getenv("AUTH_ZERO_STATE"); auth0State != "" {
		cfg.Auth0State = auth0State
	}
	if auth0CallbackURL := os.Getenv("AUTH_ZERO_CALLBACK_URL"); auth0CallbackURL != "" {
		cfg.Auth0CallbackURL = auth0CallbackURL
	}
	if auth0Audience := os.Getenv("AUTH_ZERO_AUDIENCE"); auth0Audience != "" {
		cfg.Auth0Audience = auth0Audience
	}
	return cfg
}
