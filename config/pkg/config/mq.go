package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	amqp "github.com/rabbitmq/amqp091-go"
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

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Connect(cfg *Config) (*amqp.Connection, error) {
	if cfg == nil {
		return nil, fmt.Errorf("Config not provided - Aborting...")
	}
	return amqp.Dial(
		fmt.Sprintf(
			"amqp://%s:%s@%s:%d",
			cfg.MessageQueue.User,
			cfg.MessageQueue.Password,
			cfg.MessageQueue.Host,
			cfg.MessageQueue.Port,
		),
	)
}
