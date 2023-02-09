package config

import (
	"fmt"
	"github.com/caarlos0/env/v7"
	"github.com/ilyakaznacheev/cleanenv"
	"log"
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
		//Port string `envDefault:"8083"`
		// BASE_URL - базовый адрес результирующего сокращённого URL
		BaseURL string `env:"BASE_URL" envDefault:"http://localhost:8080"`
		// SERVER_ADDRESS - адрес запуска HTTP-сервера
		ServerAddress string `env:"SERVER_ADDRESS" envDefault:"localhost:8080"`
	}
	Storage struct {
		PathStorage string `env:"SHORTURL_DIR_STORAGE" envDefault:"/files"`
		FileStorage string `env:"FILE_STORAGE_PATH" envDefault:"storage.txt"`
	}
	Log struct {
		Level string `env-required:"true" yaml:"log_level"  env:"LOG_LEVEL"`
	}
)

var instance *Config

// NewConfig returns app config.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	dir := "internal/usecase/repo/files"
	createDir("SHORTURL_DIR_STORAGE", dir)
	os.Setenv("FILE_STORAGE_PATH", fmt.Sprintf("%s/%s", dir, "storage.txt"))

	err := cleanenv.ReadConfig("./config/config.yml", cfg)
	if err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	// caarlos0
	err = env.Parse(cfg)
	if err != nil {
		return nil, err
	}
	//log.Println(cfg)
	return cfg, nil
}

func createDir(dirNameEvn string, path string) {
	// создает каталог с именем path вместе со всеми необходимыми родительскими
	// элементами и возвращает nil или возвращает ошибку
	os.Setenv(dirNameEvn, path)
	dir := os.Getenv(dirNameEvn)
	if dir == " " {
		log.Fatalf("Missing environment variable %s! %s", dirNameEvn, dir)
	}

	err := os.MkdirAll(dir, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}
}
