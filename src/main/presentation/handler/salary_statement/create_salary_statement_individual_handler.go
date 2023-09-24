package salary_statement

import (
	"context"
	"log"
	"net/http"
	"time"
	"github.com/jokertennis/Payroll-Software/basicauth"
	"github.com/jokertennis/Payroll-Software/db"
	"github.com/jokertennis/Payroll-Software/src/main/application-service/salary_statement_application_service"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/administrator_repository"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/employee_repository"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/salary_statement_repository"

	"github.com/jokertennis/Payroll-Software/src/main/infrastructure"
	"github.com/jokertennis/Payroll-Software/swagger/models"
	"github.com/jokertennis/Payroll-Software/swagger/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type CreateSalaryStatementHandlerStruct struct {}

func (s *CreateSalaryStatementHandlerStruct) Handle(params operations.PostAdministratorSalaryStatementParams) middleware.Responder {
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
		return operations.NewPostAdministratorSalaryStatementUnauthorized().WithPayload(&operations.PostAdministratorSalaryStatementUnauthorizedBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusInternalServerError {
		return operations.NewPostAdministratorSalaryStatementInternalServerError().WithPayload(&operations.PostAdministratorSalaryStatementInternalServerErrorBody{
			Message: err.Error(),
		})
	}

	salaryStatementRepositoryStruct := infrastructure.NewSalaryStatementRepository(ctx, dbInstance)
	var salaryStatementRepository salary_statement_repository.SalaryStatementRepository = &salaryStatementRepositoryStruct

	salaryStatementEntry := MappingSalaryStatementEntry(params.SalaryStatementRequest)

	result, statusCode, err := salary_statement_application_service.CreateSalaryStatementUseCase(administratorRepository, employeeRepository, salaryStatementRepository, mailAddressOfAdministrator, *params.SalaryStatementRequest.MailaddressOfEmployee, salaryStatementEntry)

	if statusCode == http.StatusBadRequest {
		return operations.NewPostAdministratorSalaryStatementBadRequest().WithPayload(&operations.PostAdministratorSalaryStatementBadRequestBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusUnauthorized {
		return operations.NewPostAdministratorSalaryStatementUnauthorized().WithPayload(&operations.PostAdministratorSalaryStatementUnauthorizedBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusInternalServerError {
		return operations.NewPostAdministratorSalaryStatementInternalServerError().WithPayload(&operations.PostAdministratorSalaryStatementInternalServerErrorBody{
			Message: err.Error(),
		})
	}

	return operations.NewPostAdministratorSalaryStatementCreated().WithPayload(&operations.PostAdministratorSalaryStatementCreatedBody{
		IDOfSalaryStatement: int64(result.SalaryStatementId),
	})
}

func MappingSalaryStatementEntry(params *models.SalaryStatementRequest) salary_statement_repository.SalaryStatementEntry {
	return salary_statement_repository.SalaryStatementEntry{
		EarningEntry: salary_statement_repository.EarningEntry{
			Amount: int(*params.AmountOfEarning),
			Nominal: *params.NominalOfEarning,
			EarningDetailsEntry: MappingEarningDetailsEntry(params.EarningDetails),
		},
		DeductionEntry: salary_statement_repository.DeductionEntry{
			Amount: int(*params.AmountOfDeduction),
			Nominal: *params.NominalOfDeduction,
			DeductionDetailsEntry: MappingDeductionDetailsEntry(params.DeductionDetails),
		},
		// params don't have employee id, then set to 0 temporarily
		EmployeeId: 0,
		Nominal: *params.Nominal,
		Payday: time.Time(*params.Payday),
		TargetPeriod: *params.TargetPeriod,
	}
}

func MappingEarningDetailsEntry(params []*models.SalaryStatementRequestEarningDetailsItems0) []salary_statement_repository.EarningDetailEntry {
	var EarningDetailsEntry []salary_statement_repository.EarningDetailEntry
	for _, value := range params {
		EarningDetailEntry := salary_statement_repository.EarningDetailEntry{
			Amount: int(value.AmountOfEarningDetail),
			Nominal: value.Nominal,
		}
		EarningDetailsEntry = append(EarningDetailsEntry, EarningDetailEntry)
	}
	return EarningDetailsEntry
}

func MappingDeductionDetailsEntry(params []*models.SalaryStatementRequestDeductionDetailsItems0) []salary_statement_repository.DeductionDetailEntry {
	var DeductionDetailsEntry []salary_statement_repository.DeductionDetailEntry
	for _, value := range params {
		DeductionDetailEntry := salary_statement_repository.DeductionDetailEntry{
			Amount: int(value.AmountOfDeductionDetail),
			Nominal: value.Nominal,
		}
		DeductionDetailsEntry = append(DeductionDetailsEntry, DeductionDetailEntry)
	}
	return DeductionDetailsEntry
}