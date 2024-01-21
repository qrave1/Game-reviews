package config

import "time"

type Config struct {
	Env string
	App
	DB
	Secret string
	Redis
}

type App struct {
	Address string
}

type Redis struct {
	Host     string
	TokenTTL time.Duration
}

type DB struct {
	DbUser string
	DbPass string
	DbHost string
	DbName string
}

func NewDefault() *Config {
	return &Config{
		Env: "PROD",
		App: App{
			Address: ":8080",
		},
		DB: DB{
			DbUser: "user",
			DbPass: "pass",
			DbHost: "localhost:7000",
			DbName: "gameReview",
		},
		Secret: "very-secret-key",
		Redis: Redis{
			Host:     "localhost:6379",
			TokenTTL: 24 * time.Hour,
		},
	}
}
