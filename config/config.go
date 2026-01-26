package config

import "github.com/caarlos0/env/v11"

type DB struct {
	Driver   string `env:"POSTGRES_DRIVER" envDefault:"postgres"`
	User     string `env:"POSTGRES_USER" envDefault:"user_test"`
	Password string `env:"POSTGRES_PASSWORD"  envDefault:"user123"`
	Port     string `env:"POSTGRES_PORT" envDefault:"5432"`
	Host     string `env:"POSTGRES_HOST"  envDefault:"localhost"`
	Name     string `env:"POSTGRES_NAME"  envDefault:"chopp_db"`
}

type Auth struct {
	SecretKey string `env:"SECRET_KEY" envDefault:"testeblablabna"`
}

type Config struct {
	DB   DB
	Auth Auth
}

func FromEnv() *Config {
	var c Config

	if err := env.Parse(&c.DB); err != nil {
		panic(err)
	}
	if err := env.Parse(&c.Auth); err != nil {
		panic(err)
	}

	return &c
}
