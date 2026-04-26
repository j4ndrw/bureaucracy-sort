package mq

import (
	"fmt"
	"j4ndrw/bureaucracysort/config/pkg/config"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func Connect(cfg *config.Config) (*amqp.Connection, error) {
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
