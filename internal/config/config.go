package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type ENVConfig struct {
	Host string `env:"LOCAL_HOST" env-default:"localhost"`
	Port string `env:"LOCAL_PORT" env-default:"8080"`
	JWT  string `env:"JWT_SECRET" env-default:"secret"`
}

type DBConfig struct {
	// TODO: postgres
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-default:"5432"`
	Name     string `env:"DB_NAME" env-default:"postgres"`
	User     string `env:"DB_USER" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"postgres"`
}

func LoadDBConfig() (*DBConfig, error) {
	var cfg DBConfig

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	return &cfg, nil
}

func LoadEnvConfig() (*ENVConfig, error) {
	var cfg ENVConfig

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	return &cfg, nil
}
