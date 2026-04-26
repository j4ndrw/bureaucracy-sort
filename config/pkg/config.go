package config

type Config struct {
	MessageQueue MessageQueueConfig
	DB           DBConfig
	Microservices MicroservicesConfig
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

	return Config{
		MessageQueue: mqConfig,
		DB: dbConfig,
		Microservices: MicroservicesConfig{
			Approver: NewMicroservice("approver"),
			Notification: NewMicroservice("notification"),
			Sorting: NewMicroservice("sorting"),
		},
	}, nil
}
