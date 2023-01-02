package db

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
)

// Reference : https://github.com/golang-migrate/migrate/tree/master/database/mysql#use-with-existing-client
func CreateMigrateInstance() (*migrate.Migrate, error) {
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
		"file://src/ddl",
		"mysql",
		driver,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create migration-instance. err:%s", err)
	}

	return m, nil
}

func CreateDemoData(migrateInstance migrate.Migrate) error {
	
}