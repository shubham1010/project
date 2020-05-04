package migrations

import (
	"fmt"
	"os"
	"database/sql"

	log "github.com/sirupsen/logrus"
	"github.com/golang-migrate/migrate"
    "github.com/golang-migrate/migrate/database/mysql"
    _ "github.com/golang-migrate/migrate/source/file"
)

type MigrationLogger struct {
    verbose bool
}

func (ml *MigrationLogger) Verbose() bool {
    return ml.verbose
}

func MigrateDatabase(db *sql.DB) error {
    driver, err := mysql.WithInstance(db, &mysql.Config{})

    if err != nil {
        return err
    }

    dir, err := os.Getwd()

    if err != nil {
        log.Fatal(err)
    }

    migration, err := migrate.NewWithDatabaseInstance(
        fmt.Sprintf("file://%s/tables", dir),
        "mysql",
        driver,
    )

    if err != nil {
        return err
    }

    log.Info("[MIGRATION]: Applying database migrations")

    err = migration.Up()
    if err!=nil && err != migrate.ErrNoChange {
        return err
    }

    version, _, err := migration.Version()
    if err != nil {
        return err
    }

    log.Info("[MIGRATION]: Active database version: ", version)
    return nil

}
