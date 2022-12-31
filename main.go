package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	migrateInstance, errorCreateMigrateInstance := createMigrateInstance()
	if errorCreateMigrateInstance != nil {
		fmt.Printf("failed to create MigrateInstance. err:%s", errorCreateMigrateInstance)
	}

	if err := migrateInstance.Steps(2); err != nil {
		fmt.Printf("failed to step. err:%s", err)
	}

	if err := migrateInstance.Up(); err != nil {
		fmt.Printf("failed to up. err:%s", err)
	}
}

// Reference : https://github.com/golang-migrate/migrate/tree/master/database/mysql#use-with-existing-client
func createMigrateInstance() (*migrate.Migrate, error) {
	// https://github.com/go-sql-driver/mysql/#multistatements
	db, err := sql.Open("mysql", "root:password@tcp(db_container:3306)/develop_db")
	if err != nil {
		return nil, fmt.Errorf("sql.Open error. err:%s", err)
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create driver. err:%s", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"mysql",
		driver,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create migration-instance. err:%s", err)
	}

	return m, nil
}
