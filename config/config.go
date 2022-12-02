package config

import (
	"log"

	"github.com/Netflix/go-env"
)

type Config struct {
	FileName string `env:"FILE_NAME"`

	RedisAddress  string `env:"REDIS_ADDRESS"`
	RedisPassword string `env:"REDIS_PASSWORD"`
}

func NewConfig() *Config {
	var cfg Config
	if _, err := env.UnmarshalFromEnviron(&cfg); err != nil {
		log.Fatal(err)
	}
	return &cfg
}
