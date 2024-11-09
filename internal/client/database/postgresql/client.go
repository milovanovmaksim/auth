package postgresql

import (
	"github.com/milovanovmaksim/auth/internal/client/database"
)

type pgClient struct {
	pg database.DB
}

func (c *pgClient) DB() database.DB {
	return c.pg
}

func NewClient(pg database.DB) database.Client {
	return &pgClient{pg}
}

func (c *pgClient) Close() error {
	if c.pg != nil {
		c.pg.Close()
	}

	return nil
}
