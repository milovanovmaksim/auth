package database

import (
	"context"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

// Handler - функция, которая выполняется в транзакции
type Handler func(ctx context.Context) error

// Client клиент для работы с БД
type Client interface {
	DB() DB
	Close() error
}

// NamedExecer интерфейс для работы с именованными запросами с помощью тегов в структурах
type NamedExecer interface {
	ScanOneContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
	ScanAllContext(ctx context.Context, dest interface{}, q Query, args ...interface{}) error
}

// QueryExecer интерфейс для работы с обычными запросами
type QueryExecer interface {
	ExecContext(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	QueryContext(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRowContext(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

// SQLExecer комбинирует NamedExecer и QueryExecer
type SQLExecer interface {
	NamedExecer
	QueryExecer
}

// Pinger интерфейс для проверки соединения с БД
type Pinger interface {
	Ping(ctx context.Context) error
}

// TxManager менеджер транзакций, который выполняет указанный пользователем обработчик в транзакции.
type TxManager interface {
	ReadCommitted(ctx context.Context, f Handler) error
}

// Transactor интерфейс для работы с транзакциями.
type Transactor interface {
	BeginTx(ctx context.Context, txOptinons pgx.TxOptions) (pgx.Tx, error)
}

// DB интерфейс для работы с БД
type DB interface {
	SQLExecer
	Transactor
	Pinger
	Close()
}
