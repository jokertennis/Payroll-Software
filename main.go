package main

import (
	"context"
	"database/sql"
	"fmt"
	"usr/local/go/db"
	"usr/local/go/src/main/presentation/handler/prompt"
	"usr/local/go/swagger/restapi"
	"usr/local/go/swagger/restapi/operations"

	"github.com/go-openapi/loads"
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

	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		fmt.Printf("failed to create swaggerSpec Object. err:%s", err)
	}

	api := operations.NewSwaggerAPI(swaggerSpec)
	configureAPI(api, ctx, dbInstance)

	server := restapi.NewServer(api)
	server.Port = 8080
	// defer server.Shutdown()

	if err := server.Serve(); err != nil {
		fmt.Printf("failed to start server. err:%s", err)
	}
}

func configureAPI(api *operations.SwaggerAPI, ctx context.Context, dbInstance *sql.DB) {
	getUnprotectedHandlerStruct := prompt.GetUnprotectedHandlerStruct{}
	var getUnprotectedHandler operations.GetUnprotectedHandler = &getUnprotectedHandlerStruct
	api.GetUnprotectedHandler = getUnprotectedHandler

	getEmployeeProtectedHandlerStruct := prompt.GetEmployeeProtectedHandlerStruct{}
	var getEmployeeProtectedHandler operations.GetEmployeeProtectedHandler = &getEmployeeProtectedHandlerStruct
	api.GetEmployeeProtectedHandler = getEmployeeProtectedHandler

	getAdministratorProtectedHandlerStruct := prompt.GetAdministratorProtectedHandlerStruct{}
	var getAdministratorProtectedHandler operations.GetAdministratorProtectedHandler = &getAdministratorProtectedHandlerStruct
	api.GetAdministratorProtectedHandler = getAdministratorProtectedHandler
}
