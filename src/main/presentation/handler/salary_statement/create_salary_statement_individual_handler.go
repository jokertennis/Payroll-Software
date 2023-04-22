package salary_statement

import (
	"context"
	"log"
	"net/http"
	"usr/local/go/basicauth"
	"usr/local/go/db"
	"usr/local/go/src/main/application-service/salary_statement_application_service"
	"usr/local/go/src/main/domain-model/individual_earning"
	"usr/local/go/src/main/domain-model/individual_earning_detail"
	"usr/local/go/src/main/domain-model/salary_statement"
	"usr/local/go/src/main/domain-service/repository/administrator_repository"
	"usr/local/go/src/main/domain-service/repository/employee_repository"
	"usr/local/go/src/main/domain-service/repository/salary_statement_repository"

	"usr/local/go/src/main/infrastructure"
	"usr/local/go/swagger/models"
	"usr/local/go/swagger/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type CreateSalaryStatementIndividualHandlerStruct struct {}

func (s *CreateSalaryStatementIndividualHandlerStruct) Handle(params operations.PostAdministratorSalaryStatementIndividualParams) middleware.Responder {
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

	mailAddress, statusCode, err := basicauth.BasicAuth(employeeRepository, administratorRepository, administratorExecuter, params.HTTPRequest)

	if statusCode == http.StatusUnauthorized {
		return operations.NewPostAdministratorSalaryStatementIndividualUnauthorized().WithPayload(&operations.PostAdministratorSalaryStatementIndividualUnauthorizedBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusInternalServerError {
		return operations.NewPostAdministratorSalaryStatementIndividualInternalServerError().WithPayload(&operations.PostAdministratorSalaryStatementIndividualInternalServerErrorBody{
			Message: err.Error(),
		})
	}

	salaryStatementRepositoryStruct := infrastructure.NewSalaryStatementRepository(ctx, dbInstance)
	var salaryStatementRepository salary_statement_repository.SalaryStatementRepository = &salaryStatementRepositoryStruct

	result, statusCode, err := salary_statement_application_service.GetAllSalaryStatementsForEmployeeUseCase(employeeRepository, salaryStatementRepository, mailAddress)

	if statusCode == http.StatusUnauthorized {
		return operations.NewPostAdministratorSalaryStatementIndividualUnauthorized().WithPayload(&operations.PostAdministratorSalaryStatementIndividualUnauthorizedBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusInternalServerError {
		return operations.NewPostAdministratorSalaryStatementIndividualInternalServerError().WithPayload(&operations.PostAdministratorSalaryStatementIndividualInternalServerErrorBody{
			Message: err.Error(),
		})
	}

	return operations.NewPostAdministratorSalaryStatementIndividualCreated().WithPayload(&operations.PostAdministratorSalaryStatementIndividualCreatedBody{
		IDOfSalaryStatement: result.NameOfEmployee,
	})
}

func mappingSalaryStatementEntryByUsingIndividualDatas(params models.SalaryStatementRequest) salary_statement.SalaryStatementEntryByUsingIndividualDatas{
	return &salary_statement.SalaryStatementEntryByUsingIndividualDatas{
		IndividualEarning: &individual_earning.IndividualEarning{
			Amount: int(*params.AmountOfEarning),
			Nominal: *params.NominalOfEarning,
			IndividualEarningDetails: &individual_earning_detail.IndividualEarningDetail,
		},
	}
}