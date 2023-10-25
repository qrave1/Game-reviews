package config

import (
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env string `yaml:"env"`
	App `yaml:"app"`
	DB  `yaml:"db"`
}

type App struct {
	Address string `yaml:"address"`
}

type DB struct {
	DbUser string `yaml:"user"`
	DbPass string `yaml:"pass"`
	DbHost string `yaml:"host"`
	DbName string `yaml:"name"`
}

func MustLoad() *Config {
	path := "./config/config.yml"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("config file doesn't exist: %s", path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("can't read config: %s", err)
	}

	return &cfg
}
