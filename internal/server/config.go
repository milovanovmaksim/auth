package server

type ServerConfig interface {
	Port() string
	Host() string
	Address() string
}
