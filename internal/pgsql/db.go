package pgsql

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgreSQL struct {
	Pool *pgxpool.Pool
}

func newPostgreSQL(pool *pgxpool.Pool) PostgreSQL {
	return PostgreSQL{Pool: pool}
}

func Connect(ctx context.Context, config *Config) (*PostgreSQL, error) {
	pool, err := pgxpool.Connect(ctx, config.Dsn())
	if err != nil {
		return nil, err
	}

	postgreSql := newPostgreSQL(pool)
	return &postgreSql, nil
}

func (p *PostgreSQL) GetPool() *pgxpool.Pool {
	return p.Pool
}

func (p *PostgreSQL) Close() {
	p.Pool.Close()
}
