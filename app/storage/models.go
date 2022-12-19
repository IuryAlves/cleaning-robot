package storage

import (
	"github.com/uptrace/bun"
	"time"
)

type Executions struct {
	bun.BaseModel `bun:"table:executions"`

	ID        int64         `bun:",pk,autoincrement" json:"id"`
	Timestamp time.Time     `json:"timestamp"`
	Commands  int           `json:"commands"`
	Result    int           `json:"result"`
	Duration  time.Duration `json:"duration"`
}
