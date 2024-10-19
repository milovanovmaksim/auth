package config

import (
	"fmt"
	"net"
	"os"
)

const (
	grpcPort = "GRPC_PORT"
	grpcHost = "GRPC_HOST"
)

// GrpcConfig содержит настройки grpc сервера.
type GrpcConfig struct {
	port string
	host string
}

func newGrpcConfig(port string, host string) GrpcConfig {
	return GrpcConfig{port, host}
}

// NewGrpcConfigFromEnv создает новый объект GrpcConfig из .env файла.
func NewGrpcConfigFromEnv() (*GrpcConfig, error) {
	port := os.Getenv(grpcPort)
	if len(port) == 0 {
		return nil, fmt.Errorf("%s must be set", grpcPort)
	}

	host := os.Getenv(grpcHost)
	if len(host) == 0 {
		return nil, fmt.Errorf("%s must be set", grpcHost)
	}

	config := newGrpcConfig(port, host)
	return &config, nil

}

// Address возвращает адрес grpc сервера.
func (cfg *GrpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
