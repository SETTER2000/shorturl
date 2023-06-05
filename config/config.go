// Package config - конфигурация, настройка всего сервиса.
package config

import (
	"flag"
	"fmt"
	"github.com/SETTER2000/shorturl/pkg/log/logger"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config структура содержит всё окружение проекта
// App - переменные окружения для приложения.
// HTTP - окружения для сервера.
// Storage - окружения для хранилищ.
// Cookie - окружения для куки.
// Log - окружения для логирования
type (
	Config struct {
		App     `json:"app"`
		HTTP    `json:"http"`
		Storage `json:"storage"`
		Cookie  `json:"cookie"`
		Log     `json:"logger"`
	}

	App struct {
		// Название сервиса
		Name string `env-required:"true" json:"name"    env:"APP_NAME"`
		// Версия сервиса
		Version string `env-required:"true" json:"version" env:"APP_VERSION"`
		// Имя файла конфигурации должно задаваться через флаг -c/-config или переменную окружения CONFIG
		ConfigFileName string `env:"CONFIG"`
	}

	HTTP struct {
		// При передаче флага -s или переменной окружения ENABLE_HTTPS сервер запуститься с
		// помощью метода http.ListenAndServeTLS или tls.Listen.
		EnableHTTPS bool `env:"ENABLE_HTTPS"`
		// передать строковое представление бесклассовой адресации (CIDR)
		TrustedSubnet string `env:"TRUSTED_SUBNET"`
		// BASE_URL - базовый адрес результирующего сокращённого URL
		BaseURL string `json:"base_url" env:"BASE_URL" `
		// SERVER_ADDRESS - адрес запуска HTTP-сервера
		ServerAddress string `json:"server_address" env:"SERVER_ADDRESS"`
		// SERVER_DOMAIN - доменное имя сервера
		ServerDomain string `env:"SERVER_DOMAIN" json:"server_domain"`
		CertsDir     string `env:"CERTS_DIR"`
		CertFile     string `env:"CERT_NAME_FILE"`
		KeyFile      string `env:"CERT_KEY_FILE"`
	}

	Storage struct {
		// FILE_STORAGE_PATH путь до файла с сокращёнными URL (директории не создаёт)
		FileStorage string `env:"FILE_STORAGE_PATH"`
		// Строка с адресом подключения к БД, например для PostgreSQL (драйвер pgx): postgres://username:password@localhost:5432/database_name
		ConnectDB string `env:"DATABASE_DSN"`
	}

	Cookie struct {
		// ACCESS_TOKEN_NAME - содержит наименование для куки доступа, по умолчанию access_token
		// Например куки будет выглядеть так:
		// access_token=5d9470be88997d3a200126e686ac7dab0190db8b341ba40e5c6cccf1e6ba66a08f05717dece9; Path=/;
		AccessTokenName string `env-required:"true" json:"access_token_name" env:"ACCESS_TOKEN_NAME" envDefault:"access_token"`
		// SECRET_KEY ключ шифрования для куки
		SecretKey string `env-required:"true" json:"secret_key" env:"SECRET_KEY" envDefault:"RtsynerpoGIYdab_s234r"` // cookie encryp application
		//ExpirationTime  time.Time `env-required:"true" yaml:"expiration_time" env:"EXPIRATION_TIME"`
	}

	Log struct {
		// LOG_LEVEL переменная окружения, содержит значение уровня логирования проекта
		Level string `env-required:"true" yaml:"log_level"  json:"log_level"  env:"LOG_LEVEL"`
	}
)

// Config .
var instance *Config

// NewConfig (singleton) возвращает инициализированную структуру конфига.
func NewConfig() (*Config, error) {
	cfg := &Config{}
	log := logger.New("debug")
	log.Info("read application configuration")
	cfg.App.ConfigFileName = "config/config.json"
	checkConfigFile(&cfg.App.ConfigFileName, cfg)

	// StringVar flags
	flag.StringVar(&cfg.App.ConfigFileName, "c", "", "configuration file name")
	flag.StringVar(&cfg.HTTP.BaseURL, "b", cfg.HTTP.BaseURL, "the base address of the resulting shortened URL")
	flag.StringVar(&cfg.Storage.ConnectDB, "d", "", "dsn connect string urlExample PostgreSQL: postgres://username:password@localhost:5432/database_name")
	flag.StringVar(&cfg.Storage.FileStorage, "f", "storage.txt", "path to file with abbreviated URLs")
	flag.StringVar(&cfg.HTTP.ServerAddress, "a", cfg.HTTP.ServerAddress, "host to listen on")
	flag.BoolVar(&cfg.HTTP.EnableHTTPS, "s", false, "start server with https protocol")
	flag.StringVar(&cfg.HTTP.TrustedSubnet, "t", "", "you can pass Classless Addressing Representation (CIDR) strings")
	flag.StringVar(&cfg.HTTP.CertsDir, "cd", "certs", "certificate directory")
	flag.StringVar(&cfg.HTTP.CertFile, "cc", "dev.crt", "name file certificate")
	flag.StringVar(&cfg.HTTP.KeyFile, "ck", "dev_rsa.key", "name file key certificate")
	flag.StringVar(&cfg.HTTP.ServerDomain, "dm", "rooder.ru", "server domain name")
	flag.StringVar(&cfg.Cookie.SecretKey, "sk", "RtsynerpoGIYdab_s234r", "cookie secret key")

	// Usage .
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Shorturl Version %s %v\nUsage : Project Shorturl - URL Shortener Server\n", os.Args[0], cfg.App.Version)
		flag.PrintDefaults()
	}

	// Parse .
	flag.Parse()
	checkConfigFile(&cfg.App.ConfigFileName, cfg)

	return cfg, nil
}

func checkConfigFile(path *string, cfg interface{}) {
	env := os.Getenv("CONFIG")
	if len(env) < 1 && len(*path) > 0 {
		env = *path
	} else if len(env) < 1 {
		return
	}
	// ReadConfig делает следующее:
	// разобрать файл конфигурации в соответствии с форматом
	// YAML | JSON (yaml тег в данном случае);
	// читает переменные среды и перезаписывает значения из файла
	// значениями, которые были найдены в среде (env тег);
	// если на первых двух шагах значение не найдено,
	// поле будет заполнено значением по умолчанию (env-default тег), если оно задано.
	if err := cleanenv.ReadConfig(env, cfg); err != nil {
		panic(err)
	}
}
