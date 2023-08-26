package salary_statement_domain_model_test

import (
	"fmt"
	"testing"
	"time"
	deduction_domain_model "usr/local/go/src/main/domain-model/deduction"
	earning_domain_model "usr/local/go/src/main/domain-model/earning"
	salary_statement_domain_model "usr/local/go/src/main/domain-model/salary_statement"

	"github.com/stretchr/testify/assert"
)

func TestGetEarning(t *testing.T) {
	cases := map[string]struct {
		salaryStatement   *salary_statement_domain_model.SalaryStatement
		individualEarning *earning_domain_model.IndividualEarning
		fixedEarning      *earning_domain_model.FixedEarning
		expectedError     error
	}{
		"individualEarning exists.": {
			salaryStatement: &salary_statement_domain_model.SalaryStatement{
				ID:                  1,
				IndividualEarning:   &earning_domain_model.IndividualEarning{ID: 1},
				FixedEarning:        nil,
				IndividualDeduction: &deduction_domain_model.IndividualDeduction{ID: 1},
				FixedDeduction:      nil,
				EmployeeId:          1,
				Nominal:             "",
				Payday:              time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod:        "2022年1月1日~2022年1月31日分",
			},
			individualEarning: &earning_domain_model.IndividualEarning{ID: 1},
			fixedEarning:      nil,
			expectedError:     nil,
		},
		"fixedEarning exist.": {
			salaryStatement: &salary_statement_domain_model.SalaryStatement{
				ID:                  1,
				IndividualEarning:   nil,
				FixedEarning:        &earning_domain_model.FixedEarning{ID: 1},
				IndividualDeduction: nil,
				FixedDeduction:      &deduction_domain_model.FixedDeduction{ID: 1},
				EmployeeId:          1,
				Nominal:             "",
				Payday:              time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod:        "2022年1月1日~2022年1月31日分",
			},
			individualEarning: nil,
			fixedEarning:      &earning_domain_model.FixedEarning{ID: 1},
			expectedError:     nil,
		},
		"error has occurred.": {
			salaryStatement: &salary_statement_domain_model.SalaryStatement{
				ID:                  1,
				IndividualEarning:   nil,
				FixedEarning:        &earning_domain_model.FixedEarning{ID: 1},
				IndividualDeduction: &deduction_domain_model.IndividualDeduction{ID: 1},
				FixedDeduction:      nil,
				EmployeeId:          1,
				Nominal:             "",
				Payday:              time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod:        "2022年1月1日~2022年1月31日分",
			},
			individualEarning: nil,
			fixedEarning:      nil,
			expectedError:     fmt.Errorf("do not allow FixedEarning and IndividualDeduction to exist.SalaryStatementId:1, FixedEarningId:1, IndividualDeductionId:1"),
		},
	}

	for _, value := range cases {
		individualEarning, fixedEarning, err := value.salaryStatement.GetEarning()
		assert.Equal(t, value.individualEarning, individualEarning)
		assert.Equal(t, value.fixedEarning, fixedEarning)
		assert.Equal(t, value.expectedError, err)
	}
}
