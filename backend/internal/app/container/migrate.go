package container

import (
	"context"
	"errors"
	"fmt"

	"github.com/tosaken1116/spino_cup_2024/backend/migrations"
	"github.com/uptrace/bun/migrate"
)

func (a *App) Migrate(ctx context.Context) (err error) {
	migrator := migrate.NewMigrator(a.db.DB, migrations.Migrations)
	if err := migrator.Init(ctx); err != nil {
		return err
	}

	if err := migrator.Lock(ctx); err != nil {
		return err
	}
	defer func() {
		if _err := migrator.Unlock(ctx); _err != nil {
			err = errors.Join(err, _err)
		}
	}()

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return err
	}
	if group.IsZero() {
		fmt.Printf("there are no new migrations to run (database is up to date)\n")
		return nil
	}

	return nil
}
