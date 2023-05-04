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

func TestNewSalaryStatement(t *testing.T) {
	cases := map[string]struct {
		ID                    uint32
	    IndividualEarning     *individual_earning_domain_model.IndividualEarning
	    FixedEarning          *fixed_earning_domain_model.FixedEarning
	    IndividualDeduction   *individual_deduction_domain_model.IndividualDeduction
	    FixedDeduction        *fixed_deduction_domain_model.FixedDeduction
	    EmployeeId            uint32
	    Nominal               string
	    Payday                time.Time
	    TargetPeriod          string
		salaryStatement       *salary_statement_domain_model.SalaryStatement
		expectedError         error
	}{
		"error has not occurred.": {
			ID: 1,
			IndividualEarning: &individual_earning_domain_model.IndividualEarning{ID: 1},
			FixedEarning: nil,
			IndividualDeduction: &individual_deduction_domain_model.IndividualDeduction{ID: 1},
			FixedDeduction: nil,
			EmployeeId: 1,
			Nominal: "",
			Payday: time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
			TargetPeriod: "2022年1月1日~2022年1月31日分",
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
		"error has occurred.": {
			ID: 1,
			IndividualEarning: &individual_earning_domain_model.IndividualEarning{ID: 1},
			FixedEarning: nil,
			IndividualDeduction: nil,
			FixedDeduction: &fixed_deduction_domain_model.FixedDeduction{ID: 1},
			EmployeeId: 1,
			Nominal: "",
			Payday: time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
			TargetPeriod: "2022年1月1日~2022年1月31日分",
			salaryStatement: nil,
			expectedError: fmt.Errorf("do not allow FixedDeduction and IndividualEarning to exist.SalaryStatementId:1, FixedDeductionId:1, IndividualEarningId:1"),
		},
	}

	for _, value := range cases {
		salaryStatement, err := salary_statement_domain_model.NewSalaryStatement(
			value.ID, 
			value.IndividualEarning, 
			value.FixedEarning, 
			value.IndividualDeduction, 
			value.FixedDeduction, 
			value.EmployeeId, 
			value.Nominal, 
			value.Payday, 
			value.TargetPeriod)
		assert.Equal(t, value.salaryStatement, salaryStatement)
		assert.Equal(t, value.expectedError, err)
	}
}