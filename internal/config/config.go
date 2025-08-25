package config

import (
	"log"
	"os"
)

type Config struct {
	Port     string
	MongoURI string
	MongoDB  string
}

func Load() *Config {
	return &Config{
		Port:     getEnv("PORT", "8080"),
		MongoURI: mustEnv("MONGO_URI"),
		MongoDB:  getEnv("MONGO_DB", "kirana"),
	}

}

func getEnv(k, d string) string {
	v := os.Getenv(k)
	if v != "" {
		return v
	}
	return d

}

func mustEnv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("missing env %s", k)
	}
	return v

}
