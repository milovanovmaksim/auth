package pgsql

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const (
	pgUser = "PG_USER"
	pgPassword = "PG_PASSWORD"
)

type Config struct {
	Username     string
	Password     string
	Host         string
	DatabaseName string
	SslMode      string
	Port         uint16
}

func NewConfig(username string, password string, port uint16, host string, databaseName string, sslMode string) Config {
	return Config{username, password, host, databaseName, sslMode, port}
}

func NewConfigFromEnv() (*Config, error) {
	var port uint64
	var err error

	username := os.Getenv(pgUser)
	if len(username) == 0 {
		return nil, fmt.Errorf("%s must be set", pgUser)
	}

	password := os.Getenv(pgPassword)
	if len(password) == 0 {
		return nil, errors.New("PG_PASSWORD must be set")
	}

	portAsString := os.Getenv("PG_PORT")
	if len(portAsString) == 0 {
		return nil, errors.New("PG_PORT must be set")
	} else {
		port, err = strconv.ParseUint(portAsString, 0, 16)
		if err != nil {
			return nil, errors.New("failed to parse PG_PORT")
		}
	}

	host := os.Getenv("PG_HOST")
	if len(host) == 0 {
		return nil, errors.New("PG_HOST must be set")
	}

	databaseName := os.Getenv("PG_DATABASE_NAME")
	if len(databaseName) == 0 {
		return nil, errors.New("PG_DATABASE_NAME must be set")
	}

	sslMode := os.Getenv("SSL_MODE")
	if len(sslMode) == 0 {
		return nil, errors.New("SSL_MODE must be set")
	}

	config := NewConfig(username, password, uint16(port), host, databaseName, sslMode)

	return &config, nil
}

func (c *Config) Dsn() string {
	return fmt.Sprintf("host=%s port=%v dbname=%s user=%s password=%s sslmode=%v", c.Host, c.Port, c.DatabaseName, c.Username, c.Password, c.SslMode)
}
