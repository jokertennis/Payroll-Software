package salarystatement

import (
	"context"
	"fmt"
	"net/http"
	"usr/local/go/basicauth"
	"usr/local/go/db"
	"usr/local/go/src/main/domain-service/repository"
	"usr/local/go/src/main/infrastructure"
	"usr/local/go/swagger/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type GetSalaryStatementForEmployeeHandler struct {}

func (s *GetSalaryStatementForEmployeeHandler) Handle(params operations.GetEmployeeSalaryStatementParams) middleware.Responder {
	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
	dbInstance, err := db.CreateDbInstance(dbEnvironment)
	if err != nil {
		fmt.Printf("failed to create dbInstance. err:%s", err)
	}
	
	administratorExecuter := basicauth.Executer{Executer: "Employee"}

	employeeRepositoryStruct := infrastructure.NewEmployeeRepository(dbInstance)
	var employeeRepository repository.EmployeeRepository = &employeeRepositoryStruct

	administratorRepositoryStruct := infrastructure.NewAdministratorRepository(dbInstance)
	var administratorRepository repository.AdministratorRepository = &administratorRepositoryStruct

	statusCode, err := basicauth.BasicAuth(ctx, employeeRepository, administratorRepository, administratorExecuter, params.HTTPRequest)

	if statusCode == http.StatusUnauthorized {
		return operations.NewGetEmployeeSalaryStatementUnauthorized().WithPayload(&operations.GetEmployeeSalaryStatementUnauthorizedBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusInternalServerError {
		return operations.NewGetEmployeeSalaryStatementInternalServerError().WithPayload(&operations.GetEmployeeSalaryStatementInternalServerErrorBody{
			Message: err.Error(),
		})
	}

	response := operations.NewGetEmployeeProtectedOK().WithPayload(&operations.GetEmployeeProtectedOKBody{
		Message: "This is the protected handler",
	})

	return response
}