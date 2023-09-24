package salary_statement_application_service

import (
	"fmt"
	"net/http"
	deduction_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/deduction"
	earning_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/earning"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/employee_repository"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/salary_statement_repository"

	"github.com/go-openapi/strfmt"
)

type ResultOfGetAllSalaryStatementsForEmployee struct {
	NameOfEmployee                 string
	SalaryStatementForEmployeeList []*SalaryStatementForEmployee
}

type SalaryStatementForEmployee struct {
	Nominal           string
	Payday            strfmt.DateTime
	TargetPeriod      string
	AmountOfDeduction int
	AmountOfEarning   int
	EarningDetails    []EarningDetailOfGetAllSalaryStatementsForEmployee
	DeductionDetails  []DeductionDetailOfGetAllSalaryStatementsForEmployee
}

type EarningDetailOfGetAllSalaryStatementsForEmployee struct {
	Nominal               string
	AmountOfEarningDetail int
}

type DeductionDetailOfGetAllSalaryStatementsForEmployee struct {
	Nominal                 string
	AmountOfDeductionDetail int
}

func GetAllSalaryStatementsForEmployeeUseCase(employeeRepository employee_repository.EmployeeRepository, salaryStatementRepository salary_statement_repository.SalaryStatementRepository, mailAddress string) (result *ResultOfGetAllSalaryStatementsForEmployee, statusCode int, errorMessage error) {
	employee, err := employeeRepository.GetEmployeeByMailAddress(mailAddress)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("InternalServerError:error:%s", err)
	}
	// employee == nil will never be true.Because exist of employee with specified mailAddress was already checked at BasicAuth.
	if employee == nil {
		return nil, http.StatusUnauthorized, fmt.Errorf("unauthorized. Specified mailAddress is not found in registered user datas. MailAddress:%s", mailAddress)
	}

	salaryStatements, err := employee.GetAllSalaryStatements(salaryStatementRepository)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("internalServerError:error:%s", err)
	}

	if salaryStatements == nil {
		return &ResultOfGetAllSalaryStatementsForEmployee{NameOfEmployee: employee.Name, SalaryStatementForEmployeeList: nil}, http.StatusOK, nil
	}

	var salaryStatementForEmployeeList []*SalaryStatementForEmployee
	for _, salaryStatement := range salaryStatements {
		deduction := salaryStatement.Deduction
		earning := salaryStatement.Earning

		amountOfDeduction, deductionDetails := MappingAmountOfDeductionAndDeductionDetailOfGetAllSalaryStatementsForEmployee(deduction)
		amountOfEarning, earningDetails := MappingAmountOfEarningAndEarningDetailOfGetAllSalaryStatementsForEmployee(earning)

		salaryStatementForEmployee := MappingSalaryStatementForEmployee(salaryStatement.Nominal, strfmt.DateTime(salaryStatement.Payday), salaryStatement.TargetPeriod, amountOfDeduction, amountOfEarning, earningDetails, deductionDetails)
		salaryStatementForEmployeeList = append(salaryStatementForEmployeeList, salaryStatementForEmployee)
	}

	return MappingResultOfGetAllSalaryStatementsForEmployee(employee.Name, salaryStatementForEmployeeList), http.StatusOK, nil
}

func MappingAmountOfDeductionAndDeductionDetailOfGetAllSalaryStatementsForEmployee(deduction deduction_domain_model.Deduction) (int, []DeductionDetailOfGetAllSalaryStatementsForEmployee) {
	var deductionDetails []DeductionDetailOfGetAllSalaryStatementsForEmployee
	var amountOfDeduction int
	amountOfDeduction = deduction.Amount
	for _, value := range deduction.DeductionDetails {
		deductionDetails = append(deductionDetails, DeductionDetailOfGetAllSalaryStatementsForEmployee{
			Nominal:                 value.Nominal,
			AmountOfDeductionDetail: value.Amount,
		})
	}

	return amountOfDeduction, deductionDetails
}

func MappingAmountOfEarningAndEarningDetailOfGetAllSalaryStatementsForEmployee(earning earning_domain_model.Earning) (int, []EarningDetailOfGetAllSalaryStatementsForEmployee) {
	var earningDetails []EarningDetailOfGetAllSalaryStatementsForEmployee
	var amountOfEarning int
	amountOfEarning = earning.Amount
	for _, value := range earning.EarningDetails {
		earningDetails = append(earningDetails, EarningDetailOfGetAllSalaryStatementsForEmployee{
			Nominal:               value.Nominal,
			AmountOfEarningDetail: value.Amount,
		})
	}

	return amountOfEarning, earningDetails
}

func MappingSalaryStatementForEmployee(nominal string, payday strfmt.DateTime, targetPeriod string, amountOfDeduction int, amountOfEarning int, earningDetails []EarningDetailOfGetAllSalaryStatementsForEmployee, deductionDetails []DeductionDetailOfGetAllSalaryStatementsForEmployee) *SalaryStatementForEmployee {
	return &SalaryStatementForEmployee{
		Nominal:           nominal,
		Payday:            payday,
		TargetPeriod:      targetPeriod,
		AmountOfDeduction: amountOfDeduction,
		AmountOfEarning:   amountOfEarning,
		EarningDetails:    earningDetails,
		DeductionDetails:  deductionDetails,
	}
}

func MappingResultOfGetAllSalaryStatementsForEmployee(nameOfEmployee string, salaryStatementForEmployeeList []*SalaryStatementForEmployee) *ResultOfGetAllSalaryStatementsForEmployee {
	return &ResultOfGetAllSalaryStatementsForEmployee{
		NameOfEmployee:                 nameOfEmployee,
		SalaryStatementForEmployeeList: salaryStatementForEmployeeList,
	}
}
