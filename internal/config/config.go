package config

import "os"

type Config struct {
	Secret   string
	DBString string
	Port     string
}

func NewConfig() (*Config, error) {
	return &Config{
		Secret:   os.Getenv("SECRET"),
		DBString: os.Getenv("DB_STRING"),
		Port:     os.Getenv("PORT"),
	}, nil // TODO: check if env vars are empty and error in that case
}
