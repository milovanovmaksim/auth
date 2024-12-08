package server

// Config интерфейс для работы с конфигурацией сервера.
type Config interface {
	Port() string
	Host() string
	Address() string
}
