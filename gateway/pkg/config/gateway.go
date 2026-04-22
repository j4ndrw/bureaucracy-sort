package config

import (
	"fmt"
	"os"
	"strconv"
)

type GatewayConfig struct {
	Host string
	Port int
}

func NewGatewayConfig() (GatewayConfig, error) {
	port, err := strconv.Atoi(os.Getenv("GATEWAY_PORT"))
	if err != nil {
		return GatewayConfig{}, err
	}
	return GatewayConfig{
		Host: os.Getenv("GATEWAY_HOST"),
		Port: port,
	}, nil
}

func newApiCluster(service string) string {
	return fmt.Sprintf("bureaucracy-sort.%s.service:8080/api", service)
}

func NewGateway(cfg *Config) (map[string]string, error) {
	if cfg == nil {
		return map[string]string{}, fmt.Errorf("No config defined!")
	}

	return map[string]string {
		"/api/notification": newApiCluster(cfg.Microservices.Notification.Name),
		"/api/approver": newApiCluster(cfg.Microservices.Approver.Name),
		"/api/sorting": newApiCluster(cfg.Microservices.Sorting.Name),
	}, nil
}
