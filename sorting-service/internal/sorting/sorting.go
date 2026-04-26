package sorting

import (
	"fmt"
	"j4ndrw/bureaucracysort/config/pkg/mq"
	"log"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ScheduleSortConfig struct {
	RequireApproval *func()
	MqConn          *amqp.Connection
}

func WithRequiredApprovalCallback(cb func()) func(*ScheduleSortConfig) {
	return func(cfg *ScheduleSortConfig) {
		cfg.RequireApproval = &cb
	}
}

func WithMq(conn *amqp.Connection) func(*ScheduleSortConfig) {
	return func(cfg *ScheduleSortConfig) {
		cfg.MqConn = conn
	}
}

func ScheduleSort(
	input []int,
	options ...func(*ScheduleSortConfig),
) {
	id := uuid.New()
	cfg := &ScheduleSortConfig{}
	for _, option := range options {
		option(cfg)
	}

	if cfg.RequireApproval == nil {
		log.Fatal("This beats the purpose of this toy project - aborting...")
	}

	if cfg.MqConn == nil {
		log.Fatal("No message queue connection provided - aborting...")
	}

	ch, err := cfg.MqConn.Channel()
	if err != nil {
		log.Fatal(err.Error())
	}

	q, err := ch.QueueDeclare(
		fmt.Sprintf("sorting-approval-requests-%s", id),
		true,
		false,
		false,
		false,
		amqp.Table{
			amqp.QueueTypeArg: amqp.QueueTypeQuorum,
		},
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	mq.FailOnError(err, "Failed to register a consumer")

	var poll chan struct{}

	for i := range len(input) {
		for j := i; j < len(input); j++ {
			if input[i] >= input[j] {
				go func() {
					for msg := range msgs{
						approved := string(msg.Body) == "approved"

						if !approved {
							(*cfg.RequireApproval)()
						} else {
							close(poll)
							aux := input[i]
							input[i] = input[j]
							input[j] = aux
						}
					}
				}()
				<-poll
			}
		}
	}
}
