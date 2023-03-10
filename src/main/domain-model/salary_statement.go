package domainmodel

import (
	"time"

	"github.com/volatiletech/null/v8"
)

type SalaryStatement struct {
	ID                    uint32
	IndividualEarningID   null.Uint32
	FixedEarningID        null.Uint32
	IndividualDeductionID null.Uint32
	FixedDeductionID      null.Uint32
	EmployeeID            uint32
	Nominal               string
	Payday                time.Time
	TargetPeriod          string
}

// Not yet created value object for each attribute.
func NewSalaryStatement(id uint32, companyId uint16, name string, mailAddress string, password string) (*Employee, error) {
	return &Employee{id, companyId, name, mailAddress, password}, nil
}