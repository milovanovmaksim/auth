package server

// ServerConfig интерфейс для работы с конфигурацией сервера.
type ServerConfig interface {
	Port() string
	Host() string
	Address() string
}
