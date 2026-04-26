package config

import (
	"os"
	"strconv"
)

type DBConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

func NewDBConfig() (DBConfig, error) {
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		return DBConfig{}, err
	}

	return DBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     dbPort,
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB"),
	}, nil
}
