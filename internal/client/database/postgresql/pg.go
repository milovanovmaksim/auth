package postgresql

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/milovanovmaksim/auth/internal/client/database"
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

func (p *PostgreSQL) Ping(ctx context.Context) error {
	return p.Pool.Ping(ctx)
}

func (p *PostgreSQL) ScanOneContext(ctx context.Context, dest interface{}, q database.Query, args ...interface{}) error {
	row := p.QueryRowContext(ctx, q, args...)

	return row.Scan(dest)
}

func (p *PostgreSQL) QueryContext(ctx context.Context, q database.Query, args ...interface{}) (pgx.Rows, error) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Query(ctx, q.QueryRaw, args...)
	}

	return p.Pool.Query(ctx, q.QueryRaw, args...)
}

func (p *PostgreSQL) QueryRowContext(ctx context.Context, q database.Query, args ...interface{}) pgx.Row {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.QueryRow(ctx, q.QueryRaw, args...)
	}

	return p.Pool.QueryRow(ctx, q.QueryRaw, args...)
}

func (p *PostgreSQL) ScanAllContext(ctx context.Context, dest interface{}, q database.Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return pgxscan.ScanAll(dest, rows)
}

func (p *PostgreSQL) ExecContext(ctx context.Context, q database.Query, args ...interface{}) (pgconn.CommandTag, error) {
	tx, ok := ctx.Value(TxKey).(pgx.Tx)
	if ok {
		return tx.Exec(ctx, q.QueryRaw, args...)
	}
	return p.Pool.Exec(ctx, q.QueryRaw, args...)
}

func (p *PostgreSQL) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.Pool.BeginTx(ctx, txOptions)
}

func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}
