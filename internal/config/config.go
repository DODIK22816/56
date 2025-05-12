package config

import (
	"github.com/DODIK22816/final5/internal/orchestrator/delivery/rest/servers"
	"github.com/DODIK22816/final5/pkg/db/cache"
	"github.com/DODIK22816/final5/pkg/db/postgres"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	postgres.PostgresConfig
	cache.RedisConfig

	CalculatorServerConfig servers.CalculatorServerConfig
	ComputingPower         int    `env:"COMPUTING_POWER" env-default:"10"`
	Port                   string `env:"PORT" env-default:"8080"`
	OrchestratorHost       string `env:"ORCHESTRATOR_HOST" env-default:"localhost"`
}

func NewConfig() (*Config, error) {
	cfg := Config{}
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
