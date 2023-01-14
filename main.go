package main

import (
	"context"
	"fmt"
	"usr/local/go/db"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	developDbEnvironment := db.DbEnvironment{Environment: "Develop"}
	developDbInstance, err := db.CreateDbInstance(developDbEnvironment)
	if err != nil {
		fmt.Printf("failed to create developDbInstance. err:%s", err)
	}

	// migrate develop_db
	migrateInstanceForDevelopDb, errorCreateMigrateInstance := db.CreateMigrateInstance(developDbInstance)
	if errorCreateMigrateInstance != nil {
		fmt.Printf("failed to create MigrateInstance for develop_db. err:%s", errorCreateMigrateInstance)
	}
	if err := migrateInstanceForDevelopDb.Up(); err != nil {
		fmt.Printf("failed to up. err:%s", err)
	}

	// create demo-data for manual test.
	if err := db.CreateData(ctx, developDbInstance); err != nil {
		fmt.Printf("failed to create demo data. err:%s", err)
	}

	// // create dbInstance which is used when accessing db.
	testDbEnvironment := db.DbEnvironment{Environment: "Test"}
	testDbInstance, err := db.CreateDbInstance(testDbEnvironment)
	if err != nil {
		fmt.Printf("failed to create testDbInstance. err:%s", err)
	}

	// // migrate test_db
	migrateInstanceForTestDb, errorCreateMigrateInstance := db.CreateMigrateInstance(testDbInstance)
	if errorCreateMigrateInstance != nil {
		fmt.Printf("failed to create MigrateInstance for test_db. err:%s", errorCreateMigrateInstance)
	}
	if err := migrateInstanceForTestDb.Up(); err != nil {
		fmt.Printf("failed to up. err:%s", err)
	}
}
