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
)

type ResultStruct struct {
	Payday            string
	TargetPeriod      string
	AmountOfDeduction int
	NameOfEmployee    string
	AmountOfEarning   int
	EarningDetails    []EarningDetail
	DeductionDetails  []DeductionDetail
}

type EarningDetail struct {
	Nominal                  string
	AmountOfEarningDetail    int
}

type DeductionDetail struct {
	Nominal                   string
	AmountOfDeductionDetail   int
}

func GetSalaryStatementForEmployeeUseCase(employeeRepository employee_repository.EmployeeRepository, salaryStatementRepository salary_statement_repository.SalaryStatementRepository, mailAddress string, year int, month time.Month) (result *ResultStruct, statusCode int, errorMessage error){
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
		return nil, http.StatusNotFound, fmt.Errorf("notFound. SalaryStatement with specified year and month was not found in registered salary statement datas.UserMailAddress:%s, Year:%d, Month:%d", employee.MailAddress, year, month)
	}

	individualDeduction, fixedDeduction, err := salaryStatement.GetDeduction()
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("InternalServerError:error:%s", err)
	}
	individualEarning, fixedEarning, err := salaryStatement.GetEarning()
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("InternalServerError:error:%s", err)
	}

	amountOfDeduction, deductionDetails := MappingAmountOfDeductionAndDeductionDetail(individualDeduction, fixedDeduction)
	amountOfEarning, earningDetails := MappingAmountOfEarningAndEarningDetail(individualEarning, fixedEarning)

	return MappingResultStruct(salaryStatement.Payday.String(), salaryStatement.TargetPeriod, amountOfDeduction, employee.Name, amountOfEarning, earningDetails, deductionDetails), http.StatusOK, nil
}

func MappingAmountOfDeductionAndDeductionDetail(individualDeduction *individual_deduction.IndividualDeduction, fixedDeduction *fixed_deduction.FixedDeduction) (int, []DeductionDetail) {
	var deductionDetails   []DeductionDetail
	var amountOfDeduction  int
	if individualDeduction != nil {
		amountOfDeduction = individualDeduction.Amount
		for _, value := range individualDeduction.IndividualDeductionDetails {
			deductionDetails = append(deductionDetails, DeductionDetail{
				Nominal: value.Nominal,
				AmountOfDeductionDetail: value.Amount,
			})
		}
		
	} else if fixedDeduction != nil {
		amountOfDeduction = fixedDeduction.Amount
		for _, value := range fixedDeduction.FixedDeductionDetails {
			deductionDetails = append(deductionDetails, DeductionDetail{
				Nominal: value.Nominal,
				AmountOfDeductionDetail: value.Amount,
			})
		}
	}
	return amountOfDeduction, deductionDetails
}

func MappingAmountOfEarningAndEarningDetail(individualEarning *individual_earning.IndividualEarning, fixedEarning *fixed_earning.FixedEarning) (int, []EarningDetail) {
	var earningDetails   []EarningDetail
	var amountOfEarning  int
	if individualEarning != nil {
		amountOfEarning = individualEarning.Amount
		for _, value := range individualEarning.IndividualEarningDetails {
			earningDetails = append(earningDetails, EarningDetail{
				Nominal: value.Nominal,
				AmountOfEarningDetail: value.Amount,
			})
		}
		
	} else if fixedEarning != nil {
		amountOfEarning = fixedEarning.Amount
		for _, value := range fixedEarning.FixedEarningDetails {
			earningDetails = append(earningDetails, EarningDetail{
				Nominal: value.Nominal,
				AmountOfEarningDetail: value.Amount,
			})
		}
	}
	return amountOfEarning, earningDetails
}

func MappingResultStruct(payday string, targetPeriod string, amountOfDeduction int, nameOfEmployee string, amountOfEarning int, earningDetails []EarningDetail, deductionDetails []DeductionDetail) *ResultStruct {
	return &ResultStruct{
		Payday: payday,
		TargetPeriod: targetPeriod,
		AmountOfDeduction: amountOfDeduction,
		NameOfEmployee: nameOfEmployee,
		AmountOfEarning: amountOfEarning,
		EarningDetails: earningDetails,
		DeductionDetails: deductionDetails,
	}
}