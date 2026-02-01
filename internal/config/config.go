package config

import "os"

type Config struct {
	MongoURI string
	Port     string
}

func Load() *Config {
	uri := os.Getenv("MONGODB_URI")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return &Config{MongoURI: uri, Port: port}
}
