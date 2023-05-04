package salary_statement_domain_model_test

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

func TestCheckDeductionAndEarning(t *testing.T) {
	cases := map[string]struct {
		salaryStatement *salary_statement_domain_model.SalaryStatement
		expectedError   error
	}{
		"both individualEarning and individualDeduction exist.": {
			salaryStatement: &salary_statement_domain_model.SalaryStatement{
				ID: 1,
				IndividualEarning: &individual_earning_domain_model.IndividualEarning{ID: 1},
				FixedEarning: nil,
				IndividualDeduction: &individual_deduction_domain_model.IndividualDeduction{ID: 1},
				FixedDeduction: nil,
				EmployeeId: 1,
				Nominal: "",
				Payday: time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod: "2022年1月1日~2022年1月31日分",
			},
			expectedError: nil,
		},
		"both fixedEarning and fixedDeduction exist.": {
			salaryStatement: &salary_statement_domain_model.SalaryStatement{
				ID: 1,
				IndividualEarning: nil,
				FixedEarning: &fixed_earning_domain_model.FixedEarning{ID: 1},
				IndividualDeduction: nil,
				FixedDeduction: &fixed_deduction_domain_model.FixedDeduction{ID: 1},
				EmployeeId: 1,
				Nominal: "",
				Payday: time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod: "2022年1月1日~2022年1月31日分",
			},
			expectedError: nil,
		},
		"both fixedDeduction and individualEarning exist": {
			salaryStatement: &salary_statement_domain_model.SalaryStatement{
				ID: 1,
				IndividualEarning: &individual_earning_domain_model.IndividualEarning{ID: 1},
				FixedEarning: nil,
				IndividualDeduction: nil,
				FixedDeduction: &fixed_deduction_domain_model.FixedDeduction{ID: 1},
				EmployeeId: 1,
				Nominal: "",
				Payday: time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod: "2022年1月1日~2022年1月31日分",
			},
			expectedError: fmt.Errorf("do not allow FixedDeduction and IndividualEarning to exist.SalaryStatementId:1, FixedDeductionId:1, IndividualEarningId:1"),
		},
		"both fixedEarning and individualDeduction exist": {
			salaryStatement: &salary_statement_domain_model.SalaryStatement{
				ID: 1,
				IndividualEarning: nil,
				FixedEarning: &fixed_earning_domain_model.FixedEarning{ID: 1},
				IndividualDeduction: &individual_deduction_domain_model.IndividualDeduction{ID: 1},
				FixedDeduction: nil,
				EmployeeId: 1,
				Nominal: "",
				Payday: time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod: "2022年1月1日~2022年1月31日分",
			},
			expectedError: fmt.Errorf("do not allow FixedEarning and IndividualDeduction to exist.SalaryStatementId:1, FixedEarningId:1, IndividualDeductionId:1"),
		},
		"both fixedDeduction and individualDeduction exist": {
			salaryStatement: &salary_statement_domain_model.SalaryStatement{
				ID: 1,
				IndividualEarning: nil,
				FixedEarning: nil,
				IndividualDeduction: &individual_deduction_domain_model.IndividualDeduction{ID: 1},
				FixedDeduction: &fixed_deduction_domain_model.FixedDeduction{ID: 1},
				EmployeeId: 1,
				Nominal: "",
				Payday: time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod: "2022年1月1日~2022年1月31日分",
			},
			expectedError: fmt.Errorf("both FixedDeduction and IndividualDeduction were found.SalaryStatementId:1, FixedDeductionId:1, IndividualDeductionId:1"),
		},
		"both fixedEarning and individualEarning exist": {
			salaryStatement: &salary_statement_domain_model.SalaryStatement{
				ID: 1,
				IndividualEarning: &individual_earning_domain_model.IndividualEarning{ID: 1},
				FixedEarning: &fixed_earning_domain_model.FixedEarning{ID: 1},
				IndividualDeduction: nil,
				FixedDeduction: nil,
				EmployeeId: 1,
				Nominal: "",
				Payday: time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod: "2022年1月1日~2022年1月31日分",
			},
			expectedError: fmt.Errorf("both FixedEarning and IndividualEarning were found.SalaryStatementId:1, FixedEarningId:1, IndividualEarningId:1"),
		},
		"both fixedDeduction and individualDeduction do not exist": {
			salaryStatement: &salary_statement_domain_model.SalaryStatement{
				ID: 1,
				IndividualEarning: nil,
				FixedEarning: &fixed_earning_domain_model.FixedEarning{ID: 1},
				IndividualDeduction: nil,
				FixedDeduction: nil,
				EmployeeId: 1,
				Nominal: "",
				Payday: time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod: "2022年1月1日~2022年1月31日分",
			},
			expectedError: fmt.Errorf("neither IndividualDeduction nor FixedDeduction was not found.SalaryStatementId:1"),
		},
		"both fixedEarning and individualEarning do not exist": {
			salaryStatement: &salary_statement_domain_model.SalaryStatement{
				ID: 1,
				IndividualEarning: nil,
				FixedEarning: nil,
				IndividualDeduction: &individual_deduction_domain_model.IndividualDeduction{ID: 1},
				FixedDeduction: nil,
				EmployeeId: 1,
				Nominal: "",
				Payday: time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod: "2022年1月1日~2022年1月31日分",
			},
			expectedError: fmt.Errorf("neither IndividualEarning nor FixedEarning was not found.SalaryStatementId:1"),
		},
	}

	for _, value := range cases {
		err := value.salaryStatement.CheckDeductionAndEarning()
		assert.Equal(t, value.expectedError, err)
	}
}