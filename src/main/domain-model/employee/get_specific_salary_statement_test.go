package employee_domain_model_test

import (
	"fmt"
	"testing"
	"time"
	"github.com/jokertennis/Payroll-Software/src/main/domain-model/employee"
	"github.com/jokertennis/Payroll-Software/src/main/domain-model/salary_statement"
	"github.com/jokertennis/Payroll-Software/src/testtool"

	"github.com/stretchr/testify/assert"
)

func TestGetSpecificSalaryStatement(t *testing.T) {
	type fakesFunctions struct {
        FakeGetSalaryStatement func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error)
    }
	cases := map[string]struct {
		employee                  *employee_domain_model.Employee
		fakesFunctions            fakesFunctions
		expectedSalaryStatement   *salary_statement_domain_model.SalaryStatement
		expectedError             error
	}{
		"Successfully get SalaryStatement.Error has not occurred.": {
			employee: &employee_domain_model.Employee{
				ID: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
                    salaryStatement := &salary_statement_domain_model.SalaryStatement{ID: 1}
                    return salaryStatement, nil
                },
			},
			expectedSalaryStatement: &salary_statement_domain_model.SalaryStatement{ID: 1},
			expectedError: nil,
		},
		"SalaryStatement that match the search criteria was not found.Error has not occurred.": {
			employee: &employee_domain_model.Employee{
				ID: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
                    return nil, nil
                },
			},
			expectedSalaryStatement: nil,
			expectedError: nil,
		},
		"error has occurred.": {
			employee: &employee_domain_model.Employee{
				ID: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
                    return nil, fmt.Errorf("cannot connect db.")
                },
			},
			expectedSalaryStatement: nil,
			expectedError: fmt.Errorf("failed to get salary statement.error:cannot connect db."),
		},
	}

	for _, value := range cases {
		salaryStatementRepository := &testtool.SalaryStatementRepositoryMock{}
		salaryStatementRepository.FakeGetSalaryStatement = value.fakesFunctions.FakeGetSalaryStatement
		salaryStatement, err := value.employee.GetSpecificSalaryStatement(salaryStatementRepository, 2022, time.May)
		assert.Equal(t, value.expectedSalaryStatement, salaryStatement)
		assert.Equal(t, value.expectedError, err)
	}
}