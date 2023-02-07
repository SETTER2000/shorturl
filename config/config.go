package config

import (
	"fmt"
	"github.com/caarlos0/env/v7"
	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		Log  `yaml:"logger"`
	}
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}
	HTTP struct {
		Port string `env-required:"true" env:"SERVER_PORT" yaml:"port"`
		//Host          string `env-required:"true" env:"SERVER_HOST" yaml:"server_host"`
		BaseURL       string `env-required:"true" env:"SERVER_HOST" yaml:"base_url"`
		ServerAddress string `env-required:"true" env:"SERVER_ADDRESS" yaml:"server_address"`
	}

	Log struct {
		Level string `env-required:"true" yaml:"log_level"  env:"LOG_LEVEL"`
	}
)

var instance *Config

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	//err = cleanenv.ReadEnv(cfg)
	//if err != nil {
	//	return nil, err
	//}

	// caarlos0
	if err := env.Parse(&cfg); err != nil {
		fmt.Errorf("Error: %+v\n", err)
	}

	return cfg, nil
}
