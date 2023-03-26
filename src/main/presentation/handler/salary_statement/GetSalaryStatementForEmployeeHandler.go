package salary_statement

import (
	"context"
	"fmt"
	"net/http"
	"usr/local/go/basicauth"
	"usr/local/go/db"
	"usr/local/go/src/main/domain-service/repository/administrator_repository"
	"usr/local/go/src/main/domain-service/repository/employee_repository"
	"usr/local/go/src/main/domain-service/repository/salary_statement_repository"

	"usr/local/go/src/main/infrastructure"
	"usr/local/go/swagger/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type GetSalaryStatementForEmployeeHandlerStruct struct {}

func (s *GetSalaryStatementForEmployeeHandlerStruct) Handle(params operations.GetEmployeeSalaryStatementParams) middleware.Responder {
	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
	dbInstance, err := db.CreateDbInstance(dbEnvironment)
	if err != nil {
		fmt.Printf("failed to create dbInstance. err:%s", err)
	}
	
	administratorExecuter := basicauth.Executer{Executer: "Employee"}

	employeeRepositoryStruct := infrastructure.NewEmployeeRepository(ctx, dbInstance)
	var employeeRepository employee_repository.EmployeeRepository = &employeeRepositoryStruct

	administratorRepositoryStruct := infrastructure.NewAdministratorRepository(ctx, dbInstance)
	var administratorRepository administrator_repository.AdministratorRepository = &administratorRepositoryStruct

	mailAddress, statusCode, err := basicauth.BasicAuth(employeeRepository, administratorRepository, administratorExecuter, params.HTTPRequest)

	if statusCode == http.StatusUnauthorized {
		return operations.NewGetEmployeeSalaryStatementUnauthorized().WithPayload(&operations.GetEmployeeSalaryStatementUnauthorizedBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusInternalServerError {
		return operations.NewGetEmployeeSalaryStatementInternalServerError().WithPayload(&operations.GetEmployeeSalaryStatementInternalServerErrorBody{
			Message: err.Error(),
		})
	}

	salaryStatementRepositoryStruct := infrastructure.NewSalaryStatementRepository(ctx, dbInstance)
	var salaryStatementRepository salary_statement_repository.SalaryStatementRepository = &salaryStatementRepositoryStruct

	result := salary_statement_application_service.

	response := operations.NewGetEmployeeProtectedOK().WithPayload(&operations.GetEmployeeProtectedOKBody{
		Message: "This is the protected handler",
	})

	return response
}