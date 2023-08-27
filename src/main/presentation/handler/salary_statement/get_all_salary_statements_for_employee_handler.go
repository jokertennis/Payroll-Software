package salary_statement

import (
	"context"
	"log"
	"net/http"
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

type GetAllSalaryStatementsForEmployeeHandlerStruct struct {}

func (s *GetAllSalaryStatementsForEmployeeHandlerStruct) Handle(params operations.GetEmployeeSalaryStatementsParams) middleware.Responder {
	// create context
	ctx := context.Background()

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
		return operations.NewGetEmployeeSalaryStatementsUnauthorized().WithPayload(&operations.GetEmployeeSalaryStatementsUnauthorizedBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusInternalServerError {
		return operations.NewGetEmployeeSalaryStatementsInternalServerError().WithPayload(&operations.GetEmployeeSalaryStatementsInternalServerErrorBody{
			Message: err.Error(),
		})
	}

	salaryStatementRepositoryStruct := infrastructure.NewSalaryStatementRepository(ctx, dbInstance)
	var salaryStatementRepository salary_statement_repository.SalaryStatementRepository = &salaryStatementRepositoryStruct

	result, statusCode, err := salary_statement_application_service.GetAllSalaryStatementsForEmployeeUseCase(employeeRepository, salaryStatementRepository, mailAddress)

	if statusCode == http.StatusUnauthorized {
		return operations.NewGetEmployeeSalaryStatementsUnauthorized().WithPayload(&operations.GetEmployeeSalaryStatementsUnauthorizedBody{
			Message: err.Error(),
		})
	} else if statusCode == http.StatusInternalServerError {
		return operations.NewGetEmployeeSalaryStatementsInternalServerError().WithPayload(&operations.GetEmployeeSalaryStatementsInternalServerErrorBody{
			Message: err.Error(),
		})
	}

	return operations.NewGetEmployeeSalaryStatementsOK().WithPayload(&operations.GetEmployeeSalaryStatementsOKBody{
		NameOfEmployee: result.NameOfEmployee,
		SalaryStatements: mappingSalaryStatementsOfGetAllSalaryStatementsForEmployee(result.SalaryStatementForEmployeeList),
	})
}

func mappingSalaryStatementsOfGetAllSalaryStatementsForEmployee(salaryStatementForEmployeeList []*salary_statement_application_service.SalaryStatementForEmployee) []*operations.GetEmployeeSalaryStatementsOKBodySalaryStatementsItems0 {
	if salaryStatementForEmployeeList == nil {
		return nil
	}
	var response []*operations.GetEmployeeSalaryStatementsOKBodySalaryStatementsItems0
	for _, value := range salaryStatementForEmployeeList {
		response = append(response, &operations.GetEmployeeSalaryStatementsOKBodySalaryStatementsItems0{
			AmountOfDeduction: int32(value.AmountOfDeduction),
			AmountOfEarning: int32(value.AmountOfEarning),
			DeductionDetails: mappingDeductionDetailsOfGetAllSalaryStatementsForEmployee(value.DeductionDetails),
			EarningDetails: mappingEarningDetailsOfGetAllSalaryStatementsForEmployee(value.EarningDetails),
			Nominal: value.Nominal,
			Payday: value.Payday,
			TargetPeriod: value.TargetPeriod,
		})
	}
	return response
}

func mappingEarningDetailsOfGetAllSalaryStatementsForEmployee(earningDetailsOfResult []salary_statement_application_service.EarningDetailOfGetAllSalaryStatementsForEmployee) []*operations.GetEmployeeSalaryStatementsOKBodySalaryStatementsItems0EarningDetailsItems0 {
	var response []*operations.GetEmployeeSalaryStatementsOKBodySalaryStatementsItems0EarningDetailsItems0
	for _, value := range earningDetailsOfResult {
		response = append(response, &operations.GetEmployeeSalaryStatementsOKBodySalaryStatementsItems0EarningDetailsItems0{
			AmountOfEarningDetail: int32(value.AmountOfEarningDetail),
			Nominal: value.Nominal,
		})
	}
	return response
}

func mappingDeductionDetailsOfGetAllSalaryStatementsForEmployee(deductionDetailsOfResult []salary_statement_application_service.DeductionDetailOfGetAllSalaryStatementsForEmployee) []*operations.GetEmployeeSalaryStatementsOKBodySalaryStatementsItems0DeductionDetailsItems0 {
	var response []*operations.GetEmployeeSalaryStatementsOKBodySalaryStatementsItems0DeductionDetailsItems0
	for _, value := range deductionDetailsOfResult {
		response = append(response, &operations.GetEmployeeSalaryStatementsOKBodySalaryStatementsItems0DeductionDetailsItems0{
			AmountOfDeductionDetail: int32(value.AmountOfDeductionDetail),
			Nominal: value.Nominal,
		})
	}
	return response
}