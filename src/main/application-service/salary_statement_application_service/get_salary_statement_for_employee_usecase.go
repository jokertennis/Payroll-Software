package salary_statement_application_service

import (
	"fmt"
	"net/http"
	"time"
	"usr/local/go/src/main/domain-model/fixed_deduction"
	"usr/local/go/src/main/domain-model/fixed_earning"
	"usr/local/go/src/main/domain-model/individual_deduction"
	"usr/local/go/src/main/domain-model/individual_earning"
	"usr/local/go/src/main/domain-service/repository/employee_repository"
	"usr/local/go/src/main/domain-service/repository/salary_statement_repository"

	"github.com/go-openapi/strfmt"
)

type ResultOfGetSalaryStatementForEmployee struct {
	Nominal           string
	Payday            strfmt.DateTime
	TargetPeriod      string
	AmountOfDeduction int
	NameOfEmployee    string
	AmountOfEarning   int
	EarningDetails    []EarningDetailOfGetSalaryStatementForEmployee
	DeductionDetails  []DeductionDetailOfGetSalaryStatementForEmployee
}

type EarningDetailOfGetSalaryStatementForEmployee struct {
	Nominal                  string
	AmountOfEarningDetail    int
}

type DeductionDetailOfGetSalaryStatementForEmployee struct {
	Nominal                   string
	AmountOfDeductionDetail   int
}

func GetSalaryStatementForEmployeeUseCase(employeeRepository employee_repository.EmployeeRepository, salaryStatementRepository salary_statement_repository.SalaryStatementRepository, mailAddress string, year int, month time.Month) (result *ResultOfGetSalaryStatementForEmployee, statusCode int, errorMessage error) {
	employee, err := employeeRepository.GetEmployeeByMailAddress(mailAddress)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("InternalServerError:error:%s", err)
	}
	// employee == nil will never be true.Because exist of employee with specified mailAddress was already checked at BasicAuth.
	if employee == nil {
		return nil, http.StatusUnauthorized, fmt.Errorf("unauthorized. Specified mailAddress is not found in registered user datas. MailAddress:%s", mailAddress)
	}

	salaryStatement, err := employee.GetSpecificSalaryStatement(salaryStatementRepository, year, month)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("internalServerError:error:%s", err)
	}

	if salaryStatement == nil {
		return nil, http.StatusNotFound, fmt.Errorf("notFound. SalaryStatement with specified year and month was not found in registered salary statement datas.UserMailAddress:%s, Year:%d, Month:%d", mailAddress, year, month)
	}

	individualDeduction, fixedDeduction, err := salaryStatement.GetDeduction()
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("InternalServerError:error:%s", err)
	}
	individualEarning, fixedEarning, err := salaryStatement.GetEarning()
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("InternalServerError:error:%s", err)
	}

	amountOfDeduction, deductionDetails := mappingAmountOfDeductionAndDeductionDetailOfGetSalaryStatementForEmployee(individualDeduction, fixedDeduction)
	amountOfEarning, earningDetails := mappingAmountOfEarningAndEarningDetailOfGetSalaryStatementForEmployee(individualEarning, fixedEarning)

	return MappingResultOfGetSalaryStatementForEmployee(salaryStatement.Nominal, strfmt.DateTime(salaryStatement.Payday), salaryStatement.TargetPeriod, amountOfDeduction, employee.Name, amountOfEarning, earningDetails, deductionDetails), http.StatusOK, nil
}

func mappingAmountOfDeductionAndDeductionDetailOfGetSalaryStatementForEmployee(individualDeduction *individual_deduction.IndividualDeduction, fixedDeduction *fixed_deduction.FixedDeduction) (int, []DeductionDetailOfGetSalaryStatementForEmployee) {
	var deductionDetails   []DeductionDetailOfGetSalaryStatementForEmployee
	var amountOfDeduction  int
	if individualDeduction != nil {
		amountOfDeduction = individualDeduction.Amount
		for _, value := range individualDeduction.IndividualDeductionDetails {
			deductionDetails = append(deductionDetails, DeductionDetailOfGetSalaryStatementForEmployee{
				Nominal: value.Nominal,
				AmountOfDeductionDetail: value.Amount,
			})
		}
		
	} else if fixedDeduction != nil {
		amountOfDeduction = fixedDeduction.Amount
		for _, value := range fixedDeduction.FixedDeductionDetails {
			deductionDetails = append(deductionDetails, DeductionDetailOfGetSalaryStatementForEmployee{
				Nominal: value.Nominal,
				AmountOfDeductionDetail: value.Amount,
			})
		}
	}
	return amountOfDeduction, deductionDetails
}

func mappingAmountOfEarningAndEarningDetailOfGetSalaryStatementForEmployee(individualEarning *individual_earning.IndividualEarning, fixedEarning *fixed_earning.FixedEarning) (int, []EarningDetailOfGetSalaryStatementForEmployee) {
	var earningDetails   []EarningDetailOfGetSalaryStatementForEmployee
	var amountOfEarning  int
	if individualEarning != nil {
		amountOfEarning = individualEarning.Amount
		for _, value := range individualEarning.IndividualEarningDetails {
			earningDetails = append(earningDetails, EarningDetailOfGetSalaryStatementForEmployee{
				Nominal: value.Nominal,
				AmountOfEarningDetail: value.Amount,
			})
		}
		
	} else if fixedEarning != nil {
		amountOfEarning = fixedEarning.Amount
		for _, value := range fixedEarning.FixedEarningDetails {
			earningDetails = append(earningDetails, EarningDetailOfGetSalaryStatementForEmployee{
				Nominal: value.Nominal,
				AmountOfEarningDetail: value.Amount,
			})
		}
	}
	return amountOfEarning, earningDetails
}

func MappingResultOfGetSalaryStatementForEmployee(nominal string, payday strfmt.DateTime, targetPeriod string, amountOfDeduction int, nameOfEmployee string, amountOfEarning int, earningDetails []EarningDetailOfGetSalaryStatementForEmployee, deductionDetails []DeductionDetailOfGetSalaryStatementForEmployee) *ResultOfGetSalaryStatementForEmployee {
	return &ResultOfGetSalaryStatementForEmployee{
		Nominal: nominal,
		Payday: payday,
		TargetPeriod: targetPeriod,
		AmountOfDeduction: amountOfDeduction,
		NameOfEmployee: nameOfEmployee,
		AmountOfEarning: amountOfEarning,
		EarningDetails: earningDetails,
		DeductionDetails: deductionDetails,
	}
}