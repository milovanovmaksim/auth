package postgresql

import (
	"context"

	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/milovanovmaksim/auth/internal/client/database"
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
	return p.Pool.Query(ctx, q.QueryRow, args...)
}

func (p *PostgreSQL) QueryRowContext(ctx context.Context, q database.Query, args ...interface{}) pgx.Row {
	return p.Pool.QueryRow(ctx, q.QueryRow, args...)
}

func (p *PostgreSQL) ScanAllContext(ctx context.Context, dest interface{}, q database.Query, args ...interface{}) error {
	rows, err := p.QueryContext(ctx, q, args...)
	if err != nil {
		return err
	}
	return pgxscan.ScanAll(dest, rows)
}

func (p *PostgreSQL) ExecContext(ctx context.Context, q database.Query, args ...interface{}) (pgconn.CommandTag, error) {
	return p.Pool.Exec(ctx, q.QueryRow, args...)
}