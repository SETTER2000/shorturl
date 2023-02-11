package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v7"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type (
	Config struct {
		App     `yaml:"app"`
		HTTP    `yaml:"http"`
		Storage `yaml:"storage"`
		Log     `yaml:"logger"`
	}
	App struct {
		Name    string `env-required:"true" yaml:"name"    env:"APP_NAME"`
		Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
	}
	HTTP struct {
		// BASE_URL - базовый адрес результирующего сокращённого URL
		BaseURL string `env:"BASE_URL"`
		// SERVER_ADDRESS - адрес запуска HTTP-сервера
		ServerAddress string `env:"SERVER_ADDRESS"`
	}
	Storage struct {
		// FILE_STORAGE_PATH путь до файла с сокращёнными URL (директории не создаёт)
		FileStorage string `env:"FILE_STORAGE_PATH"`
	}
	Log struct {
		Level string `env-required:"true" yaml:"log_level"  env:"LOG_LEVEL"`
	}
)

var instance *Config

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}

	// yaml
	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	// flags
	flag.StringVar(&cfg.HTTP.ServerAddress, "a", "localhost:8080", "host to listen on")
	flag.StringVar(&cfg.HTTP.BaseURL, "b", "http://localhost:8080", "the base address of the resulting shortened URL")
	flag.StringVar(&cfg.Storage.FileStorage, "f", "storage.txt", "path to file with abbreviated URLs")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Version of %s\n%v\nUsage : Project Shorturl - URL Shortener Server\n", os.Args[0], cfg.App.Version)
		flag.PrintDefaults()
	}
	flag.Parse()

	// environ
	err = env.Parse(cfg) // caarlos0
	if err != nil {
		return nil, err
	}

	//log.Println(cfg)
	return cfg, nil
}
