package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"usr/local/go/basicauth"
	"usr/local/go/db"
	"usr/local/go/src/main/infrastructure"
	"usr/local/go/src/main/domain-service/repository"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
	dbInstance, err := db.CreateDbInstance(dbEnvironment)
	if err != nil {
		fmt.Printf("failed to create dbInstance. err:%s", err)
	}

	// migrate db
	migrateInstance, errorCreateMigrateInstance := db.CreateMigrateInstance(dbInstance)
	if errorCreateMigrateInstance != nil {
		fmt.Printf("failed to create MigrateInstance. err:%s", errorCreateMigrateInstance)
	}
	if err := migrateInstance.Up(); err != nil {
		fmt.Printf("failed to up. err:%s", err)
	}

	// create test data.
	if err := db.CreateData(ctx, dbInstance); err != nil {
		fmt.Printf("failed to create test data. err:%s", err)
	}

	BuildHandler(ctx, dbInstance)

	http.ListenAndServe(":8080", nil)
}

func BuildHandler(ctx context.Context, dbInstance *sql.DB) {

	administratorExecuter := basicauth.Executer{Executer: "Administrator"}
	employeeExecuter := basicauth.Executer{Executer: "Employee"}

	employeeRepositoryStruct := infrastructure.NewEmployeeRepository(dbInstance)
	var employeeRepository repository.EmployeeRepository = &employeeRepositoryStruct

	administratorRepositoryStruct := infrastructure.NewAdministratorRepository(dbInstance)
	var administratorRepository repository.AdministratorRepository = &administratorRepositoryStruct

	http.HandleFunc("/unprotected", UnprotectedHandler)
	http.HandleFunc("/admin/protected", basicauth.BasicAuth(ctx, employeeRepository, administratorRepository, administratorExecuter, ProtectedHandlerForAdministrator))
	http.HandleFunc("/protected", basicauth.BasicAuth(ctx, employeeRepository, administratorRepository, employeeExecuter, ProtectedHandlerForEmployee))
}

func ProtectedHandlerForEmployee(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the protected handler")
}

func ProtectedHandlerForAdministrator(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the protected handler")
}

func UnprotectedHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is the unprotected handler")
}
