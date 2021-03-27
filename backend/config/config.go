package config

import (
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	DB		DBConfig
}

func Load() *Config {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	return &Config{
		DB:	LoadDBConfig(),
	}
}