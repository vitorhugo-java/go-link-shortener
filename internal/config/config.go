package config

import "os"

type Config struct {
	DBURL    string
	RedisURL string
	Port     string
	AppHost  string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	appHost := os.Getenv("APP_HOST")
	if appHost == "" {
		appHost = "localhost:" + port
	}
	return &Config{
		DBURL:    os.Getenv("DB_URL"),
		RedisURL: os.Getenv("REDIS_URL"),
		Port:     port,
		AppHost:  appHost,
	}
}
