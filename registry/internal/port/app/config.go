package app

import (
	"time"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	HTTP struct {
		Addr    string   `env:"ADDR" envDefault:":8000"`
		Clients []string `env:"CLIENTS"`
	} `envPrefix:"HTTP_"`

	DB struct {
		Driver string `env:"DRIVER"`
		DSN    string `env:"DSN"`
	} `envPrefix:"DB_"`

	ShutdownTimeout time.Duration `env:"SHUTDOWN_TIMEOUT" envDefault:"30s"`
}

func InitConfig(prefix string) (*Config, error) {
	var cfg Config

	opts := env.Options{
		Prefix: prefix,
	}

	err := env.ParseWithOptions(&cfg, opts)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
