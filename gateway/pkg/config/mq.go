package config

import (
	"os"
	"strconv"
)

type MessageQueueConfig struct {
	User     string
	Password string
	Host     string
	Port     int
}

func NewMessageQueueConfig() (MessageQueueConfig, error) {
	mqPort, err := strconv.Atoi(os.Getenv("MQ_PORT"))
	if err != nil {
		return MessageQueueConfig{}, err
	}

	return MessageQueueConfig{
		Host:     os.Getenv("MQ_HOST"),
		Port:     mqPort,
		User:     os.Getenv("MQ_USER"),
		Password: os.Getenv("MQ_PASS"),
	}, nil
}
