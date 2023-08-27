package salary_statement_repository

import (
	"time"
	deduction_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/deduction"
	earning_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/earning"
	salary_statement_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/salary_statement"
)

type SalaryStatementEntry struct {
	EarningEntry   EarningEntry
	DeductionEntry DeductionEntry
	EmployeeId     uint32
	Nominal        string
	Payday         time.Time
	TargetPeriod   string
}

type EarningEntry struct {
	Amount              int
	Nominal             string
	EarningType         earning_domain_model.EarningType
	EarningDetailsEntry []EarningDetailEntry
}

type EarningDetailEntry struct {
	Amount  int
	Nominal string
}

type DeductionEntry struct {
	Amount                int
	Nominal               string
	DeductionType         deduction_domain_model.DeductionType
	DeductionDetailsEntry []DeductionDetailEntry
}

type DeductionDetailEntry struct {
	Amount  int
	Nominal string
}

type SalaryStatementRepository interface {
	GetSalaryStatement(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error)
	GetAllSalaryStatements(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error)
	CreateSalaryStatement(salaryStatementEntry SalaryStatementEntry) (salaryStatementId uint32, err error)
}
