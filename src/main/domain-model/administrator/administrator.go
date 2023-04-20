package administrator

import (
	"fmt"
	"usr/local/go/src/main/domain-model/salary_statement"
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

func CreateSalaryStatementByUsingIndividualData(salaryStatementRepository salary_statement_repository.SalaryStatementRepository, salaryStatementEntry salary_statement.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
	salaryStatementId, err = salaryStatementRepository.CreateSalaryStatement(salaryStatementEntry)
	if err != nil {
		return 0, fmt.Errorf("failed to create salary statement.error:%s", err)
	}
	return salaryStatementId, nil
}