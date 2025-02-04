package repositories

import (
	"authentication/src/api/config"
	"authentication/src/api/repositories/migrations"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/migrate"
)

func ConnectDb() *bun.DB {
	pgconn := pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(fmt.Sprintf("%s:%s", config.DbHostname, config.DbPort)),
		pgdriver.WithInsecure(true),
		//pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithUser(config.DbUsername),
		pgdriver.WithPassword(config.DbPassword),
		pgdriver.WithDatabase(config.DbDatabase),
		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	)

	sqldb := sql.OpenDB(pgconn)
	db := bun.NewDB(sqldb, pgdialect.New())
	connected := db.Ping()
	if connected != nil {
		log.Panic("Impossible to connect to db")
	}
	return db
}

type MigrationRunner struct {
	db *bun.DB
}

func NewMigrationRunner(db *bun.DB) *MigrationRunner {
	return &MigrationRunner{db: db}
}

func (mr *MigrationRunner) RunMigrations(ctx context.Context) error {
	migrator := migrate.NewMigrator(mr.db, migrations.Migrations)

	err := migrator.Init(ctx)
	if err != nil {
		return fmt.Errorf("failed to initialize migrator: %w", err)
	}

	group, err := migrator.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	if group.IsZero() {
		fmt.Printf("there are no new migrations to run\n")
		return nil
	}

	fmt.Printf("migrated to %s\n", group)
	return nil
}
