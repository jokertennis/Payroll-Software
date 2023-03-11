package domainmodel

import (
	"time"
)

type SalaryStatement struct {
	ID                    uint32
	IndividualEarning     IndividualEarning
	FixedEarning          FixedEarning
	IndividualDeduction   IndividualDeduction
	FixedDeduction        FixedDeduction
	EmployeeID            uint32
	Nominal               string
	Payday                time.Time
	TargetPeriod          string
}

// Not yet created value object for each attribute.
func NewSalaryStatement(id uint32, individualEarning IndividualEarning, fixedEarning FixedEarning, individualDeduction IndividualDeduction, fixedDeduction FixedDeduction, employeeId uint32, nominal string, payday time.Time, targetPeriod string) (*SalaryStatement, error) {
	// TODO: Implement not to allow both individualEarning and fixedEarning datas to be nil.
	// TODO: Implement not to allow both individualDeduction and fixedDeduction datas to be nil.
	return &SalaryStatement{id, individualEarning, fixedEarning, individualDeduction, fixedDeduction, employeeId, nominal, payday, targetPeriod}, nil
}