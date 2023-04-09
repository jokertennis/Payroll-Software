package prompt

import (
	"context"
	"log"
	"net/http"
	"usr/local/go/basicauth"
	"usr/local/go/db"
	"usr/local/go/src/main/domain-service/repository/administrator_repository"
	"usr/local/go/src/main/domain-service/repository/employee_repository"
	"usr/local/go/src/main/infrastructure"
	"usr/local/go/swagger/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type GetAdministratorProtectedHandlerStruct struct {}

func (s *GetAdministratorProtectedHandlerStruct) Handle(params operations.GetAdministratorProtectedParams) middleware.Responder {
	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
	dbInstance, err := db.CreateDbInstance(dbEnvironment)
	if err != nil {
		log.Fatalf("failed to create dbInstance. err:%s", err)
	}
	
	administratorExecuter := basicauth.Executer{Executer: "Administrator"}

	employeeRepositoryStruct := infrastructure.NewEmployeeRepository(ctx, dbInstance)
	var employeeRepository employee_repository.EmployeeRepository = &employeeRepositoryStruct

	administratorRepositoryStruct := infrastructure.NewAdministratorRepository(ctx, dbInstance)
	var administratorRepository administrator_repository.AdministratorRepository = &administratorRepositoryStruct

	_, statusCode, err := basicauth.BasicAuth(employeeRepository, administratorRepository, administratorExecuter, params.HTTPRequest)

	if statusCode == http.StatusUnauthorized{
		return operations.NewGetAdministratorProtectedUnauthorized().WithPayload(&operations.GetAdministratorProtectedUnauthorizedBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusInternalServerError {
		return operations.NewGetAdministratorProtectedInternalServerError().WithPayload(&operations.GetAdministratorProtectedInternalServerErrorBody{
			Message: err.Error(),
		})
	}

	response := operations.NewGetAdministratorProtectedOK().WithPayload(&operations.GetAdministratorProtectedOKBody{
		Message: "This is the protected handler",
	})

	return response
}