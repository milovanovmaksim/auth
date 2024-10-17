package pgsql

import (
	"fmt"
	"os"
	"strconv"
)

const (
	pgUser         = "PG_USER"
	pgPassword     = "PG_PASSWORD"
	pgHost         = "PG_HOST"
	pgPort         = "PG_PORT"
	pgDatabaseName = "PG_DATABASE_NAME"
	sslMode        = "SSL_MODE"
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
		return nil, fmt.Errorf("%s must be set", pgPassword)
	}

	portAsString := os.Getenv(pgPort)
	if len(portAsString) == 0 {
		return nil, fmt.Errorf("%s must be set", pgPort)
	} else {
		port, err = strconv.ParseUint(portAsString, 0, 16)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %s as uint16", pgPort)
		}
	}

	host := os.Getenv(pgHost)
	if len(host) == 0 {
		return nil, fmt.Errorf("%s must be set", pgHost)
	}

	databaseName := os.Getenv(pgDatabaseName)
	if len(databaseName) == 0 {
		return nil, fmt.Errorf("%s must be set", pgDatabaseName)
	}

	sslMode := os.Getenv(sslMode)
	if len(sslMode) == 0 {
		return nil, fmt.Errorf("%s must be set", sslMode)
	}

	config := NewConfig(username, password, uint16(port), host, databaseName, sslMode)

	return &config, nil
}

func (c *Config) Dsn() string {
	return fmt.Sprintf("host=%s port=%v dbname=%s user=%s password=%s sslmode=%v", c.Host, c.Port, c.DatabaseName, c.Username, c.Password, c.SslMode)
}
