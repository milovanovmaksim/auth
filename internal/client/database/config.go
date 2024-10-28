package database

type DBConfig interface {
	Dsn() string
}
