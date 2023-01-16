package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Env  string `env: "TODO_ENV" envDefault:"dev"`
	Port int    `env: "PORT" envDefault:80"`
}

func New() (*Config, error) {
	ctg := &Config{}
	if err := env.Parse(ctg); err != nil {
		return nil, err
	}
	return ctg, nil
}
