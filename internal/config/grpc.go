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

type GrpcConfig struct {
	port string
	host string
}

func NewGrpcConfig(port string, host string) GrpcConfig {
	return GrpcConfig{port, host}
}

func NewGrpcConfigFromEnv() (*GrpcConfig, error) {
	port := os.Getenv(grpcPort)
	if len(port) == 0 {
		return nil, fmt.Errorf("%s must be set", grpcPort)
	}

	host := os.Getenv(grpcHost)
	if len(host) == 0 {
		return nil, fmt.Errorf("%s must be set", grpcHost)
	}

	config := NewGrpcConfig(port, host)
	return &config, nil

}

func (cfg *GrpcConfig) Address() string {
	return net.JoinHostPort(cfg.host, cfg.port)
}
