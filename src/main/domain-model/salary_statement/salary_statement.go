package salary_statement

import (
	"time"
	"usr/local/go/src/main/domain-model/individual_earning"
	"usr/local/go/src/main/domain-model/fixed_earning"
	"usr/local/go/src/main/domain-model/individual_deduction"
	"usr/local/go/src/main/domain-model/fixed_deduction"
)

type SalaryStatement struct {
	ID                    uint32
	IndividualEarning     individual_earning.IndividualEarning
	FixedEarning          fixed_earning.FixedEarning
	IndividualDeduction   individual_deduction.IndividualDeduction
	FixedDeduction        fixed_deduction.FixedDeduction
	EmployeeId            uint32
	Nominal               string
	Payday                time.Time
	TargetPeriod          string
}

// Not yet created value object for each attribute.
func NewSalaryStatement(id uint32, individualEarning individual_earning.IndividualEarning, fixedEarning fixed_earning.FixedEarning, individualDeduction individual_deduction.IndividualDeduction, fixedDeduction fixed_deduction.FixedDeduction, employeeId uint32, nominal string, payday time.Time, targetPeriod string) (*SalaryStatement, error) {
	// TODO: Implement not to allow both individualEarning and fixedEarning datas to be nil.
	// TODO: Implement not to allow both individualDeduction and fixedDeduction datas to be nil.
	return &SalaryStatement{id, individualEarning, fixedEarning, individualDeduction, fixedDeduction, employeeId, nominal, payday, targetPeriod}, nil
}