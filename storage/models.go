package storage

import (
	"github.com/uptrace/bun"
	"time"
)

type Executions struct {
	bun.BaseModel `bun:"table:executions"`

	ID	 int64  `bun:",pk,autoincrement"`
	Timestamp time.Time
	Commands int
	Result int
	Duration time.Duration
}
