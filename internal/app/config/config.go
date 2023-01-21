package config

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	IsDebug *bool `yaml:"is_debug" env-required:"true"`
	Listen  struct {
		Type   string `yaml:"type" env-default:"port"`
		BindIP string `yaml:"bind_ip" env-default:"127.0.0.1"`
		Port   string `yaml:"port" env-default:"8081"`
	} `yaml:"listen"`
}

var instance *Config
var once sync.Once

// GetConfig singleton
// выполниться один раз
func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
		// cleanenv.ReadConfig - парсит конфиг,
		// ещё он может частично обновить переменные конфига, переменными окружения os
		// cleanenv.ReadEnv - может получить все окружение os
		if err := cleanenv.ReadConfig("config.yml", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)
			log.Fatalf(help)
			log.Fatal(err)
		}
	})
	return instance
}
