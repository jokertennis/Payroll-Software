package salary_statement_application_service

import (
	"errors"
	"fmt"
	"net/http"
	"usr/local/go/src/main/domain-service/repository/administrator_repository"
	"usr/local/go/src/main/domain-service/repository/employee_repository"
	"usr/local/go/src/main/domain-service/repository/salary_statement_repository"
)

type ResultOfCreateSalaryStatementIndividual struct {
	SalaryStatementId int
}

func CreateSalaryStatementIndividualUseCase(administratorRepository administrator_repository.AdministratorRepository, employeeRepository employee_repository.EmployeeRepository, salaryStatementRepository salary_statement_repository.SalaryStatementRepository, mailAddressOfAdministrator string, mailAddressOfEmployee string, salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (result *ResultOfCreateSalaryStatementIndividual, statusCode int, errorMessage error) {
	administrator, err := administratorRepository.GetAdministratorByMailAddress(mailAddressOfAdministrator)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("InternalServerError:error:%s", err)
	}
	// administrator == nil will never be true.Because exist of administrator with specified mailAddress was already checked at BasicAuth.
	if administrator == nil {
		return nil, http.StatusUnauthorized, fmt.Errorf("unauthorized. Specified mailAddress is not found in registered user datas. MailAddress:%s", mailAddressOfAdministrator)
	}

	employee, err := employeeRepository.GetEmployeeByMailAddress(mailAddressOfEmployee)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("InternalServerError:error:%s", err)
	}
	if employee == nil {
		return nil, http.StatusBadRequest, fmt.Errorf("badrequest. Specified mailAddress is not found in registered user datas. MailAddressOfEmployee:%s", mailAddressOfEmployee)
	}

	salaryStatementEntry.EmployeeId = employee.ID
	
	err = administrator.CheckEmployee(*employee)
	if err != nil {
		return nil, http.StatusBadRequest, errors.New("operations related to employee who belongs to a company different from your own are not possible")
	}

	salaryStatementExists, err := administrator.SalaryStatementExists(salaryStatementRepository, employee, salaryStatementEntry.Payday.Year(), salaryStatementEntry.Payday.Month())
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("InternalServerError:error:%s", err)
	}
	if salaryStatementExists {
		return nil, http.StatusBadRequest, fmt.Errorf("badrequest. SalaryStatement which has payday that match year and month of specified payday was already existed.")
	}

	salaryStatementId, err := administrator.CreateSalaryStatementByUsingIndividualData(salaryStatementRepository, employee, salaryStatementEntry)
	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("internalServerError:error:%s", err)
	}

	return &ResultOfCreateSalaryStatementIndividual{SalaryStatementId: int(salaryStatementId)}, http.StatusCreated, nil
}