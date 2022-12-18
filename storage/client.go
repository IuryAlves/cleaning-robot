package storage

import (
	"context"
	"crypto/tls"
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/dialect/pgdialect"
	"time"
)

type Client struct {
	db *bun.DB
}

func New() *Client {
	db := GetDB()
	return &Client {
		db: db,
	}
}

type InsertRequest struct {
	Timestamp time.Time `json:"timestamp"`
	Commands int `json:"commands"`
	Result int `json:"result"`
	Duration time.Duration `json:"duration"`
}

func (c *Client) Insert(ctx context.Context, i InsertRequest) error {
	execution := &Executions{
		Timestamp: i.Timestamp,
		Commands: i.Commands,
		Result: i.Result,
		Duration: i.Duration,
	}
	_, err := c.db.NewInsert().Model(execution).Exec(ctx)
	return err
}

func GetDB() *bun.DB {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr("localhost:5432"),
		pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithUser("test"),
		pgdriver.WithPassword("test"),
		pgdriver.WithDatabase("test"),
		pgdriver.WithApplicationName("robot"),
		pgdriver.WithDialTimeout(5 * time.Second),
		pgdriver.WithReadTimeout(5 * time.Second),
		pgdriver.WithWriteTimeout(5 * time.Second))

	sqldb := sql.OpenDB(pgconn)
	return bun.NewDB(sqldb, pgdialect.New())
}

