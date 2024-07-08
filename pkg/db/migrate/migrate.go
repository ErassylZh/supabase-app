package migrate

import (
	"database/sql"
	"fmt"
	"github.com/pressly/goose/v3"
	"io/fs"
)

func SetBaseFS(baseFS fs.FS) {
	goose.SetBaseFS(baseFS)
}

func Up(db *sql.DB, migrationsDir string) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	goose.SetTableName(fmt.Sprintf("goose_db_version"))

	if err := goose.Up(db, migrationsDir); err != nil {
		return err
	}

	return nil
}
