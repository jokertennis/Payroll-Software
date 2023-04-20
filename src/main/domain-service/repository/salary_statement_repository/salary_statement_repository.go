package salary_statement_repository

import (
	"time"
	"usr/local/go/src/main/domain-model/salary_statement"
)

type SalaryStatementRepository interface {
	GetSalaryStatement(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement.SalaryStatement, error)
	GetAllSalaryStatements(employeeId uint32) ([]*salary_statement.SalaryStatement, error)
	CreateSalaryStatement(salaryStatementEntry salary_statement.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error)
}