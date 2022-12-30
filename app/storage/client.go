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
	"strconv"
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

func (c *Client) InsertExecution(ctx context.Context, execution Executions) (Executions, error) {
	_, err := c.db.NewInsert().Model(&execution).Exec(ctx)
	if err != nil {
		return Executions{}, err
	}
	return execution, nil
}

func (c *Client) GetDB() *bun.DB {
	return c.db
}

func WithPostgres() func(*Client) {
	return func(s *Client) {
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

		m := os.Getenv("DB_MAX_RETRIES")
		if m == "" {
			m = "10"
		}
		maxRetries, _ := strconv.Atoi(m)

		addr := strings.Join([]string{host, port}, ":")
		pgconn := pgdriver.NewConnector(
			pgdriver.WithNetwork("tcp"),
			pgdriver.WithAddr(addr),
			pgdriver.WithInsecure(true),
			pgdriver.WithUser(user),
			pgdriver.WithPassword(password),
			pgdriver.WithDatabase(dbName),
			pgdriver.WithApplicationName("robot"),
			pgdriver.WithDialTimeout(5*time.Second),
			pgdriver.WithReadTimeout(5*time.Second),
			pgdriver.WithWriteTimeout(5*time.Second),
		)

		for i := 0; i < maxRetries; i++ {
			sqldb := sql.OpenDB(pgconn)
			db := bun.NewDB(sqldb, pgdialect.New())
			err := db.Ping()
			if err == nil {
				s.db = db
				break
			}
			time.Sleep(5 * time.Second)
		}
	}
}
