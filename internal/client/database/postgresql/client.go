package postgresql

import (
	"github.com/milovanovmaksim/auth/internal/client/database"
)

type pgClient struct {
	pg database.DB
}

// DB возвращает объект, удовлетворяющий интерфейсу database.DB.
func (c *pgClient) DB() database.DB {
	return c.pg
}

// NewClient возвращает клиента для работы с БД PostgreSQL.
func NewClient(pg database.DB) database.Client {
	return &pgClient{pg}
}

// Close закрывает соединение с БД.
func (c *pgClient) Close() error {
	if c.pg != nil {
		c.pg.Close()
	}

	return nil
}
