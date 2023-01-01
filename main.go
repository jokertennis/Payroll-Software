package main

import (
	"fmt"
	"usr/local/go/db"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	migrateInstance, errorCreateMigrateInstance := db.CreateMigrateInstance()
	if errorCreateMigrateInstance != nil {
		fmt.Printf("failed to create MigrateInstance. err:%s", errorCreateMigrateInstance)
	}

	if err := migrateInstance.Up(); err != nil {
		fmt.Printf("failed to up. err:%s", err)
	}
}
