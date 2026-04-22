package config

type Config struct {
	MessageQueue MessageQueueConfig
	DB           DBConfig
	Microservices MicroservicesConfig
	Gateway GatewayConfig
}

func Load() (Config, error) {
	mqConfig, err := NewMessageQueueConfig()
	if err != nil {
		return Config{}, err
	}

	dbConfig, err := NewDBConfig()
	if err != nil {
		return Config{}, err
	}

	gatewayConfig, err := NewGatewayConfig()
	if err != nil {
		return Config{}, err
	}

	return Config{
		MessageQueue: mqConfig,
		DB: dbConfig,
		Gateway: gatewayConfig,
		Microservices: MicroservicesConfig{
			Approver: NewMicroservice("approver"),
			Notification: NewMicroservice("notification"),
			Sorting: NewMicroservice("sorting"),
		},
	}, nil
}
