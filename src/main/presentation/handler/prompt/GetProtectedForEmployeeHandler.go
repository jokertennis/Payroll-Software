package prompt

import (
	"context"
	"fmt"
	"net/http"
	"usr/local/go/basicauth"
	"usr/local/go/db"
	"usr/local/go/src/main/domain-service/repository/administrator_repository"
	"usr/local/go/src/main/domain-service/repository/employee_repository"
	"usr/local/go/src/main/infrastructure"
	"usr/local/go/swagger/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type GetEmployeeProtectedHandlerStruct struct {}

func (s *GetEmployeeProtectedHandlerStruct) Handle(params operations.GetEmployeeProtectedParams) middleware.Responder {
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

	statusCode, err := basicauth.BasicAuth(employeeRepository, administratorRepository, administratorExecuter, params.HTTPRequest)

	if statusCode == http.StatusUnauthorized {
		return operations.NewGetEmployeeProtectedUnauthorized().WithPayload(&operations.GetEmployeeProtectedUnauthorizedBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusInternalServerError {
		return operations.NewGetAdministratorProtectedInternalServerError().WithPayload(&operations.GetAdministratorProtectedInternalServerErrorBody{
			Message: err.Error(),
		})
	}

	response := operations.NewGetEmployeeProtectedOK().WithPayload(&operations.GetEmployeeProtectedOKBody{
		Message: "This is the protected handler",
	})

	return response
}