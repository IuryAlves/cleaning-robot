package migrations

import (
	"context"
	"github.com/IuryAlves/cleaning-robot/storage"
	"github.com/uptrace/bun"
)

func init() {
	Migrations.MustRegister(func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewCreateTable().Model((*storage.Executions)(nil)).Exec(ctx)
		return err
	}, func(ctx context.Context, db *bun.DB) error {
		_, err := db.NewDropTable().Model((*storage.Executions)(nil)).Exec(ctx)
		return err
	})
}
