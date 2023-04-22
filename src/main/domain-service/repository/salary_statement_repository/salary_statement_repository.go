package salary_statement_repository

import (
	"time"
	"usr/local/go/src/main/domain-model/salary_statement"
)

type SalaryStatementEntryByUsingIndividualDatas struct {
	IndividualEarningEntry     *IndividualEarningEntry
	IndividualDeductionEntry   *IndividualDeductionEntry
	EmployeeId                 uint32
	Nominal                    string
	Payday                     time.Time
	TargetPeriod               string
}

type IndividualEarningEntry struct {
	Amount                         int
	Nominal                        string
	IndividualEarningDetailsEntry  []IndividualEarningDetailEntry
}

type IndividualEarningDetailEntry struct {
	Amount                         int
	Nominal                        string
}

type IndividualDeductionEntry struct {
	Amount                         int
	Nominal                        string
	IndividualDeductionDetailsEntry  []IndividualDeductionDetailEntry
}

type IndividualDeductionDetailEntry struct {
	Amount                         int
	Nominal                        string
}

type SalaryStatementRepository interface {
	GetSalaryStatement(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement.SalaryStatement, error)
	GetAllSalaryStatements(employeeId uint32) ([]*salary_statement.SalaryStatement, error)
	CreateSalaryStatement(salaryStatementEntry SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error)
}