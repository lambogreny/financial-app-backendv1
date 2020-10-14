package database

import (
	"github.com/sirupsen/logrus"
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"


	"financial-app-backend/internal/config"
)
//te
func migrateDb(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return errors.Wrap(err, "connecting to database")
	}

	migrationSourse := fmt.Sprintf("file//%sinternal/database/migrations/", *config.DataDirectory)
	migrator, err := migrate.NewWithDatabaseInstance(migrationSourse, "postgres", driver)
	if err != nil {
		return errors.Wrap(err, "creating migrator")
	}

	if err := migrator.Up(); err != nil && err != migrate.ErrNoChange {
		return errors.Wrap(err, "executing migration")
	}

	Version, dirty, err := migrator.Version()
	if err != nil {
		return errors.Wrap(err, "getting migration version")
	}
	logrus.WithFields(logrus.Fields{
		"version": Version,
		"dirty":   dirty,
	}).Debug("Database migrated")

	return nil
}
