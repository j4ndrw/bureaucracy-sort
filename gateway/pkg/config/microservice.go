package config

import (
	"fmt"
	"os"
	"strings"
)

type MicroserviceConfig struct {
	Name string
}

type MicroservicesConfig struct {
	Approver MicroserviceConfig
	Notification MicroserviceConfig
	Sorting MicroserviceConfig
}

func NewMicroservice(microservice string) MicroserviceConfig{
	return MicroserviceConfig{
		Name: os.Getenv(fmt.Sprintf("%s_NAME", strings.ToUpper(microservice))),
	}
}
