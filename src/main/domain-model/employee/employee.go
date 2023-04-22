package employee_domain_model

import (
	"fmt"
	"time"
	"usr/local/go/src/main/domain-model/salary_statement"
	"usr/local/go/src/main/domain-service/repository/salary_statement_repository"
)

type Employee struct {
	ID          uint32
	CompanyId   uint16
	Name        string
	MailAddress string
	Password    string
}

// Not yet created value object for each attribute.
func NewEmployee(id uint32, companyId uint16, name string, mailAddress string, password string) (*Employee, error) {
	return &Employee{id, companyId, name, mailAddress, password}, nil
}

func (e *Employee) GetSpecificSalaryStatement(salaryStatementRepository salary_statement_repository.SalaryStatementRepository, year int, month time.Month) (_ *salary_statement.SalaryStatement, _ error) {
	salaryStatement, err := salaryStatementRepository.GetSalaryStatement(e.ID, year, month)
	if err != nil {
		return nil, fmt.Errorf("failed to get salary statement.error:%s", err)
	}
	if salaryStatement == nil {
		return nil, nil
	}
	return salaryStatement, nil
}

func (e *Employee) GetAllSalaryStatements(salaryStatementRepository salary_statement_repository.SalaryStatementRepository) (_ []*salary_statement.SalaryStatement, _ error) {
	salaryStatements, err := salaryStatementRepository.GetAllSalaryStatements(e.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to get all salary statements.error:%s", err)
	}
	if salaryStatements == nil {
		return nil, nil
	}
	return salaryStatements, nil
}