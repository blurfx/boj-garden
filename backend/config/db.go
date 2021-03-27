package config

import "os"

type DBConfig struct {
	User		string
	Password	string
	Name		string
	Host		string
	Port		string
}

func LoadDBConfig() DBConfig {
	return DBConfig{
		User:		os.Getenv("DB_USER"),
		Password:	os.Getenv("DB_PASSWORD"),
		Name:		os.Getenv("DB_NAME"),
		Host:		os.Getenv("DB_HOST"),
		Port: 		os.Getenv("DB_PORT"),
	}
}