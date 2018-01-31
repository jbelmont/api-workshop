package config

import (
	"os"
	"time"
)

type Config struct {
	ReadTimeout     time.Duration
	WriteTimeout    time.Duration
	ShutdownTimeout time.Duration
	DBTimeout       time.Duration
	APIHost         string
	MongoHost       string
	RedisHost       string
}

// New creates and returns a configuration
func New() Config {
	cfg := Config{
		ReadTimeout:     5 * time.Second,
		WriteTimeout:    10 * time.Second,
		ShutdownTimeout: 5 * time.Second,
		DBTimeout:       25 * time.Second,
		APIHost:         ":8080",
		MongoHost:       "mongodb://mongo:27017",
		RedisHost:       "redis:6379",
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
	return cfg
}
