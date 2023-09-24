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

func TestCheckDeductionAndEarning(t *testing.T) {
	cases := map[string]struct {
		salaryStatement *salary_statement_domain_model.SalaryStatement
		expectedError   error
	}{
		"Earning and Deduction type are different.": {
			salaryStatement: &salary_statement_domain_model.SalaryStatement{
				ID:           1,
				Earning:      earning_domain_model.Earning{EarningType: earning_domain_model.IndividualEarning},
				Deduction:    deduction_domain_model.Deduction{DeductionType: deduction_domain_model.FixedDeduction},
				EmployeeId:   1,
				Nominal:      "",
				Payday:       time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
				TargetPeriod: "2022年1月1日~2022年1月31日分",
			},
			expectedError: fmt.Errorf("DeductionType and EarningType are not to be allowed to different.DeductionType:Fixed, EarningType:Individual"),
		},
	}

	for _, value := range cases {
		err := value.salaryStatement.CheckDeductionAndEarningType()
		assert.Equal(t, value.expectedError, err)
	}
}
