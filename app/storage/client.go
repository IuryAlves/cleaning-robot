package storage

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"os"
	"strings"
	"time"
)

type Client struct {
	db *bun.DB
}

func New(options ...func(*Client)) *Client {
	c := &Client{}
	for _, o := range options {
		o(c)
	}
	return c
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

func (c *Client) GetDB() *bun.DB {
	return c.db
}

func WithPostgres() func(*Client) {
	return func(s *Client ) {
		host := os.Getenv("DATABASE_HOST")
		if host == "" {
			fmt.Println(errors.New("env DATABASE_HOST is empty"))
			os.Exit(1)
		}
		port := os.Getenv("DATABASE_PORT")
		if port == "" {
			port = "5432"
		}
		user := os.Getenv("DATABASE_USER")
		if user == "" {
			fmt.Println(errors.New("env DATABASE_USER is empty"))
			os.Exit(1)
		}
		password := os.Getenv("DATABASE_PASSWORD")
		if password == "" {
			fmt.Println(errors.New("env DATABASE_PASSWORD is empty"))
			os.Exit(1)
		}
		dbName := os.Getenv("DATABASE_NAME")
		if dbName == "" {
			fmt.Println(errors.New("env DATABASE_NAME is empty"))
			os.Exit(1)
		}

		addr := strings.Join([]string{host, port}, ":")
		pgconn := pgdriver.NewConnector(
			pgdriver.WithNetwork("tcp"),
			pgdriver.WithAddr(addr),
			pgdriver.WithInsecure(true),
			pgdriver.WithUser(user),
			pgdriver.WithPassword(password),
			pgdriver.WithDatabase(dbName),
			pgdriver.WithApplicationName("robot"),
			pgdriver.WithDialTimeout(5 * time.Second),
			pgdriver.WithReadTimeout(5 * time.Second),
			pgdriver.WithWriteTimeout(5 * time.Second),
		)

		sqldb := sql.OpenDB(pgconn)
		s.db = bun.NewDB(sqldb, pgdialect.New())
	}
}

