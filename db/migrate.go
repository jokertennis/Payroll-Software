package db

import (
	"database/sql"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/mattes/migrate/source/file"
)

// TODO:create as enum. only Develop.
type DbEnvironment struct {
	Environment string
}

func CreateDbInstance(dbEnvironment DbEnvironment) (*sql.DB, error) {
	switch dbEnvironment.Environment {
	case "Develop":
		// When we executed createDbInstance on local environment,localhost will be used.
		// On the other hand we executed createDbInstance inside container,db_container will be used.
		// Temporary change is not a problem.
		// We have to separate host_name depending on executed environment.
		db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/develop_db?parseTime=true")
		if err != nil {
			return nil, fmt.Errorf("sql.Open error. err:%s", err)
		}
		return db, nil
	}
	return nil, fmt.Errorf("not support specified DbEnviroment:%s", dbEnvironment.Environment)
}

// Reference : https://github.com/golang-migrate/migrate/tree/master/database/mysql#use-with-existing-client
func CreateMigrateInstance(db *sql.DB) (*migrate.Migrate, error) {
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