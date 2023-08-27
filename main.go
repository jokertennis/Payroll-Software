package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"github.com/jokertennis/Payroll-Software/db"
	"github.com/jokertennis/Payroll-Software/src/main/presentation/handler/prompt"
	"github.com/jokertennis/Payroll-Software/src/main/presentation/handler/salary_statement"
	"github.com/jokertennis/Payroll-Software/swagger/restapi"
	"github.com/jokertennis/Payroll-Software/swagger/restapi/operations"

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
		log.Fatalf("failed to create dbInstance. err:%s", err)
	}

	// migrate db
	migrateInstance, errorCreateMigrateInstance := db.CreateMigrateInstance(dbInstance)
	if errorCreateMigrateInstance != nil {
		log.Fatalf("failed to create MigrateInstance. err:%s", errorCreateMigrateInstance)
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
		log.Fatalf("failed to create swaggerSpec Object. err:%s", err)
	}

	api := operations.NewSwaggerAPI(swaggerSpec)
	configureAPI(api, ctx, dbInstance)

	server := restapi.NewServer(api)
	server.Port = 8080
	// defer server.Shutdown()

	if err := server.Serve(); err != nil {
		log.Fatalf("failed to start server. err:%s", err)
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

	getSalaryStatementForEmployeeHandlerStruct := salary_statement.GetSalaryStatementForEmployeeHandlerStruct{}
	var getSalaryStatementForEmployeeHandler operations.GetEmployeeSalaryStatementHandler = &getSalaryStatementForEmployeeHandlerStruct
	api.GetEmployeeSalaryStatementHandler = getSalaryStatementForEmployeeHandler

	getAllSalaryStatementsForEmployeeHandlerStruct := salary_statement.GetAllSalaryStatementsForEmployeeHandlerStruct{}
	var getAllSalaryStatementsForEmployeeHandler operations.GetEmployeeSalaryStatementsHandler = &getAllSalaryStatementsForEmployeeHandlerStruct
	api.GetEmployeeSalaryStatementsHandler = getAllSalaryStatementsForEmployeeHandler

	createSalaryStatementHandlerStruct := salary_statement.CreateSalaryStatementHandlerStruct{}
	var createSalaryStatementHandler operations.PostAdministratorSalaryStatementHandler = &createSalaryStatementHandlerStruct
	api.PostAdministratorSalaryStatementHandler = createSalaryStatementHandler
}
