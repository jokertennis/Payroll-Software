package salary_statement_domain_model_test

import (
	"fmt"
	"testing"
	"time"
	deduction_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/deduction"
	earning_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/earning"
	salary_statement_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/salary_statement"

	"github.com/stretchr/testify/assert"
)

func TestNewSalaryStatement(t *testing.T) {
	cases := map[string]struct {
		ID              uint32
		Earning         earning_domain_model.Earning
		Deduction       deduction_domain_model.Deduction
		EmployeeId      uint32
		Nominal         string
		Payday          time.Time
		TargetPeriod    string
		salaryStatement *salary_statement_domain_model.SalaryStatement
		expectedError   error
	}{
		"error has not occurred.": {
			ID:           1,
			Earning:      earning_domain_model.Earning{EarningType: earning_domain_model.FixedEarning},
			Deduction:    deduction_domain_model.Deduction{DeductionType: deduction_domain_model.FixedDeduction},
			EmployeeId:   1,
			Nominal:      "",
			Payday:       time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
			TargetPeriod: "2022年1月1日~2022年1月31日分",
			salaryStatement: &salary_statement_domain_model.SalaryStatement{
				ID:           1,
				Earning:      earning_domain_model.Earning{EarningType: earning_domain_model.FixedEarning},
				Deduction:    deduction_domain_model.Deduction{DeductionType: deduction_domain_model.FixedDeduction},
				EmployeeId:   1,
				Nominal:      "",
				Payday:       time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod: "2022年1月1日~2022年1月31日分",
			},
			expectedError: nil,
		},
		"error has occurred.": {
			ID:              1,
			Earning:         earning_domain_model.Earning{EarningType: earning_domain_model.FixedEarning},
			Deduction:       deduction_domain_model.Deduction{DeductionType: deduction_domain_model.IndividualDeduction},
			EmployeeId:      1,
			Nominal:         "",
			Payday:          time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
			TargetPeriod:    "2022年1月1日~2022年1月31日分",
			salaryStatement: nil,
			expectedError:   fmt.Errorf("DeductionType and EarningType are not to be allowed to different.DeductionType:Individual, EarningType:Fixed"),
		},
	}

	for _, value := range cases {
		salaryStatement, err := salary_statement_domain_model.NewSalaryStatement(
			value.ID,
			value.Earning,
			value.Deduction,
			value.EmployeeId,
			value.Nominal,
			value.Payday,
			value.TargetPeriod)
		assert.Equal(t, value.salaryStatement, salaryStatement)
		assert.Equal(t, value.expectedError, err)
	}
}
