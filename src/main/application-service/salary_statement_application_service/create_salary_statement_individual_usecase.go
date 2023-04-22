package salary_statement_application_service

import (
	"fmt"
	"net/http"
	"usr/local/go/src/main/domain-service/repository/employee_repository"
	"usr/local/go/src/main/domain-service/repository/salary_statement_repository"

	"github.com/go-openapi/strfmt"
)

type ResultOfCreateSalaryStatementIndividual struct {
	IdOfSalaryStatement int
}

func CreateSalaryStatementIndividualUseCase(employeeRepository employee_repository.EmployeeRepository, salaryStatementRepository salary_statement_repository.SalaryStatementRepository, ) (result *ResultOfCreateSalaryStatementIndividual, statusCode int, errorMessage error) {
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

	return nil, nil, nil
}