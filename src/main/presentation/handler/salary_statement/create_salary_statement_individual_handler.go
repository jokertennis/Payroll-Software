package salary_statement

import (
	"context"
	"log"
	"net/http"
	"usr/local/go/basicauth"
	"usr/local/go/db"
	"usr/local/go/src/main/application-service/salary_statement_application_service"
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

	mailAddressOfAdministrator, statusCode, err := basicauth.BasicAuth(employeeRepository, administratorRepository, administratorExecuter, params.HTTPRequest)

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

	salaryStatementEntry := MappingSalaryStatementEntry(params.SalaryStatementRequest)

	result, statusCode, err := salary_statement_application_service.CreateSalaryStatementIndividualUseCase(administratorRepository, employeeRepository, salaryStatementRepository, mailAddressOfAdministrator, *params.SalaryStatementRequest.MailaddressOfEmployee, salaryStatementEntry)

	if statusCode == http.StatusBadRequest {
		return operations.NewPostAdministratorSalaryStatementIndividualBadRequest().WithPayload(&operations.PostAdministratorSalaryStatementIndividualBadRequestBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusUnauthorized {
		return operations.NewPostAdministratorSalaryStatementIndividualUnauthorized().WithPayload(&operations.PostAdministratorSalaryStatementIndividualUnauthorizedBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusInternalServerError {
		return operations.NewPostAdministratorSalaryStatementIndividualInternalServerError().WithPayload(&operations.PostAdministratorSalaryStatementIndividualInternalServerErrorBody{
			Message: err.Error(),
		})
	}

	return operations.NewPostAdministratorSalaryStatementIndividualCreated().WithPayload(&operations.PostAdministratorSalaryStatementIndividualCreatedBody{
		IDOfSalaryStatement: int64(result.SalaryStatementId),
	})
}

func MappingSalaryStatementEntry(params *models.SalaryStatementRequest) salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas {
	return salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas{
		IndividualEarningEntry: &salary_statement_repository.IndividualEarningEntry{
			Amount: int(*params.AmountOfEarning),
			Nominal: *params.NominalOfEarning,
			IndividualEarningDetailsEntry: []salary_statement_repository.IndividualEarningDetailEntry{
			},
		},
	}
}

func MappingIndividualEarningDetailsEntry(params []*models.SalaryStatementRequestIndividualEarningDetailsItems0) []salary_statement_repository.IndividualEarningDetailEntry {
	var individualEarningDetailsEntry []salary_statement_repository.IndividualEarningDetailEntry
	for _, value := range params {
		individualEarningDetailEntry := salary_statement_repository.IndividualEarningDetailEntry{
			Amount: int(value.AmountOfEarningDetail),
			Nominal: value.Nominal,
		}
		individualEarningDetailsEntry = append(individualEarningDetailsEntry, individualEarningDetailEntry)
	}
	return individualEarningDetailsEntry
}

func MappingIndividualDeductionDetailsEntry(params []*models.SalaryStatementRequestIndividualDeductionDetailsItems0) []salary_statement_repository.IndividualDeductionDetailEntry {
	var individualDeductionDetailsEntry []salary_statement_repository.IndividualDeductionDetailEntry
	for _, value := range params {
		individualDeductionDetailEntry := salary_statement_repository.IndividualDeductionDetailEntry{
			Amount: int(value.AmountOfDeductionDetail),
			Nominal: value.Nominal,
		}
		individualDeductionDetailsEntry = append(individualDeductionDetailsEntry, individualDeductionDetailEntry)
	}
	return individualDeductionDetailsEntry
}