package pgsql

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
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
func Connect(ctx context.Context, config *Config) (*PostgreSQL, error) {
	pool, err := pgxpool.Connect(ctx, config.Dsn())
	if err != nil {
		return nil, err
	}

	postgreSQL := newPostgreSQL(pool)
	return &postgreSQL, nil
}

// GetPool возвращает pgxpool.Pool.
func (p *PostgreSQL) GetPool() *pgxpool.Pool {
	return p.Pool
}

// Close закрывает соединение с PostgreSQL сервером.
func (p *PostgreSQL) Close() {
	p.Pool.Close()
}
