package migration

import (
	"database/sql"

	"github.com/pressly/goose"
	"github.com/sirupsen/logrus"
)

func NewMigration(db *sql.DB, log *logrus.Logger, dir string) error {
	log.Info("migration has started")
	if err := goose.Up(db, dir); err != nil {
		return err
	}
	log.Info("migration has finished")

	return nil
}
