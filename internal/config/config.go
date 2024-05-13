package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string `env:"PORT" env-default:"8080"`
	}
	Database struct {
		URL string `env:"DATABASE_URL"`
	}
	JWT struct {
		Secret string `env:"JWT_SECRET"`
	}
}

func LoadConfig() *Config {
	var cfg Config

	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	return &cfg
}
