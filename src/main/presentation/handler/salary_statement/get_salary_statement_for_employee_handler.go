package salary_statement

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"
	"github.com/jokertennis/Payroll-Software/basicauth"
	"github.com/jokertennis/Payroll-Software/db"
	"github.com/jokertennis/Payroll-Software/src/main/application-service/salary_statement_application_service"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/administrator_repository"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/employee_repository"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/salary_statement_repository"

	"github.com/jokertennis/Payroll-Software/src/main/infrastructure"
	"github.com/jokertennis/Payroll-Software/swagger/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

type GetSalaryStatementForEmployeeHandlerStruct struct {}

func (s *GetSalaryStatementForEmployeeHandlerStruct) Handle(params operations.GetEmployeeSalaryStatementParams) middleware.Responder {
	// create context
	ctx := context.Background()

	queryParams := params.HTTPRequest.URL.Query()
	year, err := strconv.Atoi(queryParams.Get("year"))
	if (err != nil) {
		log.Fatalf("failed to change type of year.")
	}
	month, err := strconv.Atoi(queryParams.Get("month"))
	if (err != nil) {
		log.Fatalf("failed to change type of month.")
	}

	// create dbInstance which is used when accessing db.
	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
	dbInstance, err := db.CreateDbInstance(dbEnvironment)
	if err != nil {
		log.Fatalf("failed to create dbInstance. err:%s", err)
	}
	
	employeeExecuter := basicauth.Executer{Executer: "Employee"}

	employeeRepositoryStruct := infrastructure.NewEmployeeRepository(ctx, dbInstance)
	var employeeRepository employee_repository.EmployeeRepository = &employeeRepositoryStruct

	administratorRepositoryStruct := infrastructure.NewAdministratorRepository(ctx, dbInstance)
	var administratorRepository administrator_repository.AdministratorRepository = &administratorRepositoryStruct

	mailAddress, statusCode, err := basicauth.BasicAuth(employeeRepository, administratorRepository, employeeExecuter, params.HTTPRequest)

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

	result, statusCode, err := salary_statement_application_service.GetSalaryStatementForEmployeeUseCase(employeeRepository, salaryStatementRepository, mailAddress, year, time.Month(month))

	if statusCode == http.StatusUnauthorized {
		return operations.NewGetEmployeeSalaryStatementUnauthorized().WithPayload(&operations.GetEmployeeSalaryStatementUnauthorizedBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusNotFound {
		return operations.NewGetEmployeeSalaryStatementNotFound().WithPayload(&operations.GetEmployeeSalaryStatementNotFoundBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusInternalServerError {
		return operations.NewGetEmployeeSalaryStatementInternalServerError().WithPayload(&operations.GetEmployeeSalaryStatementInternalServerErrorBody{
			Message: err.Error(),
		})
	}

	return operations.NewGetEmployeeSalaryStatementOK().WithPayload(&operations.GetEmployeeSalaryStatementOKBody{
		Nominal: result.Nominal,
		Payday: result.Payday,
		TargetPeriod: result.TargetPeriod,
		AmountOfDeduction: int32(result.AmountOfDeduction),
		NameOfEmployee: result.NameOfEmployee,
		AmountOfEarning: int32(result.AmountOfEarning),
		EarningDetails: mappingEarningDetailsOfGetSalaryStatementForEmployee(result.EarningDetails),
		DeductionDetails: mappingDeductionDetailsOfGetSalaryStatementForEmployee(result.DeductionDetails),
	})
}

func mappingEarningDetailsOfGetSalaryStatementForEmployee(earningDetailsOfResult []salary_statement_application_service.EarningDetailOfGetSalaryStatementForEmployee) []*operations.GetEmployeeSalaryStatementOKBodyEarningDetailsItems0 {
	var response []*operations.GetEmployeeSalaryStatementOKBodyEarningDetailsItems0
	for _, value := range earningDetailsOfResult {
		response = append(response, &operations.GetEmployeeSalaryStatementOKBodyEarningDetailsItems0{
			AmountOfEarningDetail: int32(value.AmountOfEarningDetail),
			Nominal: value.Nominal,
		})
	}
	return response
}

func mappingDeductionDetailsOfGetSalaryStatementForEmployee(deductionDetailsOfResult []salary_statement_application_service.DeductionDetailOfGetSalaryStatementForEmployee) []*operations.GetEmployeeSalaryStatementOKBodyDeductionDetailsItems0 {
	var response []*operations.GetEmployeeSalaryStatementOKBodyDeductionDetailsItems0
	for _, value := range deductionDetailsOfResult {
		response = append(response, &operations.GetEmployeeSalaryStatementOKBodyDeductionDetailsItems0{
			AmountOfDeductionDetail: int32(value.AmountOfDeductionDetail),
			Nominal: value.Nominal,
		})
	}
	return response
}