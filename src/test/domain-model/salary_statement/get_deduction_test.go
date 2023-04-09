package salary_statement

import (
	"fmt"
	"testing"
	"time"
	"usr/local/go/src/main/domain-model/fixed_deduction"
	"usr/local/go/src/main/domain-model/fixed_earning"
	"usr/local/go/src/main/domain-model/individual_deduction"
	"usr/local/go/src/main/domain-model/individual_earning"
	"usr/local/go/src/main/domain-model/salary_statement"

	"github.com/stretchr/testify/assert"
)

func TestGetDeduction(t *testing.T) {
	cases := map[string]struct {
		salaryStatement     *salary_statement.SalaryStatement
		individualDeduction *individual_deduction.IndividualDeduction
		fixedDeduction      *fixed_deduction.FixedDeduction
		expectedError       error
	}{
		"individualDeduction exists.": {
			salaryStatement: &salary_statement.SalaryStatement{
				ID: 1,
				IndividualEarning: &individual_earning.IndividualEarning{ID: 1},
				FixedEarning: nil,
				IndividualDeduction: &individual_deduction.IndividualDeduction{ID: 1},
				FixedDeduction: nil,
				EmployeeId: 1,
				Nominal: "",
				Payday: time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod: "2022年1月1日~2022年1月31日分",
			},
			individualDeduction: &individual_deduction.IndividualDeduction{ID: 1},
			fixedDeduction: nil,
			expectedError: nil,
		},
		"fixedDeduction exist.": {
			salaryStatement: &salary_statement.SalaryStatement{
				ID: 1,
				IndividualEarning: nil,
				FixedEarning: &fixed_earning.FixedEarning{ID: 1},
				IndividualDeduction: nil,
				FixedDeduction: &fixed_deduction.FixedDeduction{ID: 1},
				EmployeeId: 1,
				Nominal: "",
				Payday: time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod: "2022年1月1日~2022年1月31日分",
			},
			individualDeduction: nil,
			fixedDeduction: &fixed_deduction.FixedDeduction{ID: 1},
			expectedError: nil,
		},
		"error has occurred.": {
			salaryStatement: &salary_statement.SalaryStatement{
				ID: 1,
				IndividualEarning: nil,
				FixedEarning: &fixed_earning.FixedEarning{ID: 1},
				IndividualDeduction: &individual_deduction.IndividualDeduction{ID: 1},
				FixedDeduction: nil,
				EmployeeId: 1,
				Nominal: "",
				Payday: time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod: "2022年1月1日~2022年1月31日分",
			},
			individualDeduction: nil,
			fixedDeduction: nil,
			expectedError: fmt.Errorf("do not allow FixedEarning and IndividualDeduction to exist.SalaryStatementId:1, FixedEarningId:1, IndividualDeductionId:1"),
		},
	}

	for _, value := range cases {
		individualDeduction, fixedDeduction, err := value.salaryStatement.GetDeduction()
		assert.Equal(t, value.individualDeduction, individualDeduction)
		assert.Equal(t, value.fixedDeduction, fixedDeduction)
		assert.Equal(t, value.expectedError, err)
	}
}