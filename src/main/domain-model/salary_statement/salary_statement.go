package salary_statement_domain_model

import (
	"fmt"
	"time"
	deduction_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/deduction"
	earning_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/earning"
)

type SalaryStatement struct {
	ID           uint32
	Earning      earning_domain_model.Earning
	Deduction    deduction_domain_model.Deduction
	EmployeeId   uint32
	Nominal      string
	Payday       time.Time
	TargetPeriod string
}

// Not yet created value object for each attribute.
func NewSalaryStatement(id uint32, earning earning_domain_model.Earning, deduction deduction_domain_model.Deduction, employeeId uint32, nominal string, payday time.Time, targetPeriod string) (*SalaryStatement, error) {
	salaryStatement := &SalaryStatement{id, earning, deduction, employeeId, nominal, payday, targetPeriod}
	err := salaryStatement.CheckDeductionAndEarningType()
	if err != nil {
		return nil, err
	}
	return salaryStatement, nil
}

func (s *SalaryStatement) CheckDeductionAndEarningType() error {
	if s.Deduction.DeductionType != deduction_domain_model.DeductionType(s.Earning.EarningType) {
		return fmt.Errorf("DeductionType and EarningType are not to be allowed to different.DeductionType:%s, EarningType:%s", s.Deduction.DeductionType, s.Earning.EarningType)
	}
	return nil
}
