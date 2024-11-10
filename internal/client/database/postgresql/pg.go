package postgresql

import (
	"context"
	"fmt"
	"log"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/milovanovmaksim/auth/internal/client/database"
	"github.com/milovanovmaksim/auth/internal/client/database/prettier"
)

type key string

const (
	TxKey key = "tx"
)

// PostgreSQL представляет базу данных PostgreSQL.
type PostgreSQL struct {
	Pool *pgxpool.Pool
}

// newPostgreSQL создает новый PostgreSQL объект.
func newPostgreSQL(pool *pgxpool.Pool) PostgreSQL {
	return PostgreSQL{Pool: pool}
}

// Connect создает новый PostgreSQL объект и устанавливает соединение с PostgreSQL сервером.
func Connect(ctx context.Context, config database.DBConfig) (*PostgreSQL, error) {
	pool, err := pgxpool.Connect(ctx, config.Dsn())
	if err != nil {
		return nil, err
	}

	postgreSQL := newPostgreSQL(pool)
	return &postgreSQL, nil
}

// Close закрывает соединение с PostgreSQL сервером.
func (p *PostgreSQL) Close() {
	p.Pool.Close()
}

// Ping проверяет соединение с БД.
func (p *PostgreSQL) Ping(ctx context.Context) error {
	return p.Pool.Ping(ctx)
}

// ScanOneContext делегирует работу pgx.Row.Scan.
// Может быть использован в транзакциях при передачи контекста (context.Context) с ключом postgresql.TxKey и значением,
// удовлетворяющем интерфейсу pgx.Tx.
func (p *PostgreSQL) ScanOneContext(ctx context.Context, dest interface{}, q database.Query, args ...interface{}) error {
	row := p.QueryRowContext(ctx, q, args...)

	return row.Scan(dest)
}

// QueryContext делегирует работу (pgxpool.Pool).Query.
// Может быть использован в транзакциях при передачи контекста (context.Context) с ключом postgresql.TxKey и значением,
// удовлетворяющем интерфейсу pgx.Tx.
func (p *PostgreSQL) QueryContext(ctx context.Context, q database.Query, args ...interface{}) (pgx.Rows, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return p.Pool.Query(ctx, q.QueryRaw, args...)
}

// QueryRowContext делегирует работу (pgxpool.Pool).QueryRow.
// Может быть использован в транзакциях при передачи контекста (context.Context) с ключом postgresql.TxKey и значением,
// удовлетворяющем интерфейсу pgx.Tx.
func (p *PostgreSQL) QueryRowContext(ctx context.Context, q database.Query, args ...interface{}) pgx.Row {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.Pool.QueryRow(ctx, q.QueryRaw, args...)
}

// ScanAllContext делегирует работу pgxscan.ScanAll.
// Может быть использован в транзакциях при передачи контекста (context.Context) с ключом postgresql.TxKey и значением,
// удовлетворяющем интерфейсу pgx.Tx.
func (p *PostgreSQL) ScanAllContext(ctx context.Context, dest interface{}, q database.Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return pgxscan.ScanAll(dest, rows)
}

// ExecContext делегирует работу (pgxpool.Pool).Exec.
// Может быть использован в транзакциях при передачи контекста (context.Context) с ключом postgresql.TxKey и значением,
// удовлетворяющем интерфейсу pgx.Tx.
func (p *PostgreSQL) ExecContext(ctx context.Context, q database.Query, args ...interface{}) (pgconn.CommandTag, error) {
	logQuery(ctx, q, args...)

	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}
	return p.Pool.Exec(ctx, q.QueryRaw, args...)
}

// BeginTx делегирует работу pgxpool.Pool.BeginTx.
// Для подробной информации смотри (pgxpool.Pool).BeginTx.
func (p *PostgreSQL) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.Pool.BeginTx(ctx, txOptions)
}

// MakeContextTx возвращает копию родительского контекста с ключом postgresql.TxKey и значением pgx.Tx.
func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}


func logQuery(ctx context.Context, q database.Query, args ...interface{}) {
	prettyQuery := prettier.Pretty(q.QueryRaw, prettier.PlaceholderDollar, args...)
	log.Println(
		ctx,
		fmt.Sprintf("sql: %s", q.Name),
		fmt.Sprintf("query: %s", prettyQuery),
	)
}
