package salary_statement

import (
	"fmt"
	"time"
	"usr/local/go/src/main/domain-model/fixed_deduction"
	"usr/local/go/src/main/domain-model/fixed_earning"
	"usr/local/go/src/main/domain-model/individual_deduction"
	"usr/local/go/src/main/domain-model/individual_earning"
)

type SalaryStatement struct {
	ID                    uint32
	IndividualEarning     *individual_earning.IndividualEarning
	FixedEarning          *fixed_earning.FixedEarning
	IndividualDeduction   *individual_deduction.IndividualDeduction
	FixedDeduction        *fixed_deduction.FixedDeduction
	EmployeeId            uint32
	Nominal               string
	Payday                time.Time
	TargetPeriod          string
}

type SalaryStatementEntryByUsingIndividualDatas struct {
	IndividualEarning     *individual_earning.IndividualEarning
	IndividualDeduction   *individual_deduction.IndividualDeduction
	EmployeeId            uint32
	Nominal               string
	Payday                time.Time
	TargetPeriod          string
}

// Not yet created value object for each attribute.
func NewSalaryStatement(id uint32, individualEarning *individual_earning.IndividualEarning, fixedEarning *fixed_earning.FixedEarning, individualDeduction *individual_deduction.IndividualDeduction, fixedDeduction *fixed_deduction.FixedDeduction, employeeId uint32, nominal string, payday time.Time, targetPeriod string) (*SalaryStatement, error) {
	salaryStatement := &SalaryStatement{id, individualEarning, fixedEarning, individualDeduction, fixedDeduction, employeeId, nominal, payday, targetPeriod}
	err := salaryStatement.CheckDeductionAndEarning()
	if err != nil {
		return nil, err
	}
	return salaryStatement, nil
}

// Either IndividualDeduction or FixedDeduction must be nil.
// Either IndividualDeduction or FixedDeduction must be non-nil.
func (s *SalaryStatement) GetDeduction() (*individual_deduction.IndividualDeduction, *fixed_deduction.FixedDeduction, error) {
	err := s.CheckDeductionAndEarning()
	if err != nil {
		return nil, nil, err
	}
	if s.IndividualDeduction != nil {
		return s.IndividualDeduction, nil, nil
	}
		
	return nil, s.FixedDeduction, nil
}

// Either IndividualEarning or FixedEarning must be nil.
// Either IndividualEarning or FixedEarning must be non-nil.
func (s *SalaryStatement) GetEarning() (*individual_earning.IndividualEarning, *fixed_earning.FixedEarning, error) {
	err := s.CheckDeductionAndEarning()
	if err != nil {
		return nil, nil, err
	}
	if s.IndividualEarning != nil {
		return s.IndividualEarning, nil, nil
	}
	
	return nil, s.FixedEarning, nil
}

func (s *SalaryStatement) CheckDeductionAndEarning() error {
	if (s.FixedDeduction != nil && s.IndividualEarning != nil) {
		return fmt.Errorf("do not allow FixedDeduction and IndividualEarning to exist.SalaryStatementId:%d, FixedDeductionId:%d, IndividualEarningId:%d", s.ID, s.FixedDeduction.ID, s.IndividualEarning.ID)
	} else if (s.FixedEarning != nil && s.IndividualDeduction != nil) {
		return fmt.Errorf("do not allow FixedEarning and IndividualDeduction to exist.SalaryStatementId:%d, FixedEarningId:%d, IndividualDeductionId:%d", s.ID, s.FixedEarning.ID, s.IndividualDeduction.ID)
	} else if (s.FixedDeduction != nil && s.IndividualDeduction != nil) {
		return fmt.Errorf("both FixedDeduction and IndividualDeduction were found.SalaryStatementId:%d, FixedDeductionId:%d, IndividualDeductionId:%d", s.ID, s.FixedDeduction.ID, s.IndividualDeduction.ID)
	} else if (s.FixedEarning != nil && s.IndividualEarning != nil) {
		return fmt.Errorf("both FixedEarning and IndividualEarning were found.SalaryStatementId:%d, FixedEarningId:%d, IndividualEarningId:%d", s.ID, s.FixedEarning.ID, s.IndividualEarning.ID)
	} else if (s.FixedDeduction == nil && s.IndividualDeduction == nil) {
		return fmt.Errorf("neither IndividualDeduction nor FixedDeduction was not found.SalaryStatementId:%d", s.ID)
	} else if (s.FixedEarning == nil && s.IndividualEarning == nil) {
		return fmt.Errorf("neither IndividualEarning nor FixedEarning was not found.SalaryStatementId:%d", s.ID)
	}
	return nil
}