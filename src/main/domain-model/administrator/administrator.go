package administrator_domain_model

import (
	"errors"
	"fmt"
	"time"
	"usr/local/go/src/main/domain-model/employee"
	"usr/local/go/src/main/domain-service/repository/salary_statement_repository"
)

type Administrator struct {
	ID          uint32
	CompanyId   uint16
	Name        string
	MailAddress string
	Password    string
}

// Not yet created value object for each attribute.
func NewAdministrator(id uint32, companyId uint16, name string, mailAddress string, password string) (*Administrator, error) {
	return &Administrator{id, companyId, name, mailAddress, password}, nil
}

func (a *Administrator) SalaryStatementExists(salaryStatementRepository salary_statement_repository.SalaryStatementRepository, employee employee_domain_model.Employee, year int, month time.Month) (bool, error) {
	err := a.CheckEmployee(employee)
	if err != nil {
		return false, err
	}
	salaryStatement, err := salaryStatementRepository.GetSalaryStatement(employee.ID, year, month)
	if err != nil {
		return false, fmt.Errorf("failed to get salary statement.error:%s", err)
	}
	if salaryStatement == nil {
		return false, nil
	}
	return true, nil
}

func (a *Administrator) CreateSalaryStatementByUsingIndividualData(salaryStatementRepository salary_statement_repository.SalaryStatementRepository, employee employee_domain_model.Employee, salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
	err = a.CheckEmployee(employee)
	if err != nil {
		return 0, err
	}
	salaryStatementId, err = salaryStatementRepository.CreateSalaryStatement(salaryStatementEntry)
	if err != nil {
		return 0, fmt.Errorf("failed to create salary statement.error:%s", err)
	}
	return salaryStatementId, nil
}

func (a *Administrator) CheckEmployee(employee employee_domain_model.Employee) error {
	if a.CompanyId != employee.CompanyId {
		return errors.New("operations related to employee who belongs to a company different from your own are not possible")
	}
	return nil
}