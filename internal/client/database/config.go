package database

// DBConfig интерфейс для работы с настройками БД.
type DBConfig interface {
	Dsn() string
}
