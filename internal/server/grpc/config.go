package grpc

import (
	"fmt"
	"net"
	"os"
)

const (
	grpcPort = "GRPC_PORT"
	grpcHost = "GRPC_HOST"
)

// Config представляет настройки grpc сервера.
type Config struct {
	port string
	host string
}

func newGrpcConfig(port string, host string) Config {
	return Config{port, host}
}

// NewGrpcConfigFromEnv создает новый объект GrpcConfig из .env файла.
func NewGrpcConfigFromEnv() (*Config, error) {
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
func (cfg *Config) Address() string {
	return net.JoinHostPort(cfg.Host(), cfg.Port())
}

// Host возвращает имя хоста grpc сервера.
func (cfg *Config) Host() string {
	return cfg.host
}

// Port возвращает порт grpc сервера.
func (cfg *Config) Port() string {
	return cfg.port
}
