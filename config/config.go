package config

import (
	"fmt"
	"github.com/caarlos0/env/v7"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
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
		Port string `envDefault:"8080"`
		// BASE_URL - базовый адрес результирующего сокращённого URL
		BaseURL string `env:"BASE_URL" envDefault:"http://localhost"`
		// SERVER_ADDRESS - адрес запуска HTTP-сервера
		ServerAddress string `env:"SERVER_ADDRESS" envDefault:"localhost"`
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
	//
	//err = cleanenv.ReadEnv(cfg)
	//if err != nil {
	//	return nil, err
	//}

	//caarlos0
	err = env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	log.Printf("CONFIGS:: %s", cfg)
	return cfg, nil
}
