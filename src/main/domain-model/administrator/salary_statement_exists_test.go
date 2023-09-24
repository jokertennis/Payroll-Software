package administrator_domain_model_test

import (
	"errors"
	"testing"
	"time"
	"github.com/jokertennis/Payroll-Software/src/main/domain-model/employee"
	"github.com/jokertennis/Payroll-Software/src/main/domain-model/administrator"
	"github.com/jokertennis/Payroll-Software/src/main/domain-model/salary_statement"
	"github.com/jokertennis/Payroll-Software/src/testtool"

	"github.com/stretchr/testify/assert"
)

func TestSalaryStatementExists(t *testing.T) {
	type fakesFunctions struct {
        FakeGetSalaryStatement func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error)
    }
	cases := map[string]struct {
		administrator        administrator_domain_model.Administrator
		employee             employee_domain_model.Employee
		fakesFunctions       fakesFunctions
		expectedExistsBool   bool
		expectedError        error
	}{
		"Confirmed that salary statement exists.Error has not occurred.": {
			administrator: administrator_domain_model.Administrator{
				ID: 1,
				CompanyId: 1,
			},
			employee: employee_domain_model.Employee{
				ID: 1,
				CompanyId: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return &salary_statement_domain_model.SalaryStatement{ID: 1}, nil
				},
			},
			expectedExistsBool: true,
			expectedError: nil,
		},
		"Error has occurred when executed CheckEmployee()": {
			administrator: administrator_domain_model.Administrator{
				ID: 1,
				CompanyId: 1,
			},
			employee: employee_domain_model.Employee{
				ID: 1,
				CompanyId: 10,
			},
			fakesFunctions: fakesFunctions{
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
			},
			expectedExistsBool: false,
			expectedError: errors.New("operations related to employee who belongs to a company different from your own are not possible"),
		},
		"Error has occurred when executed GetSalaryStatement()": {
			administrator: administrator_domain_model.Administrator{
				ID: 1,
				CompanyId: 1,
			},
			employee: employee_domain_model.Employee{
				ID: 1,
				CompanyId: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, errors.New("cannnot connect db")
				},
			},
			expectedExistsBool: false,
			expectedError: errors.New("failed to get salary statement.error:cannnot connect db"),
		},
		"Error has not occurred but salary_statement was nil when executed GetSalaryStatement()": {
			administrator: administrator_domain_model.Administrator{
				ID: 1,
				CompanyId: 1,
			},
			employee: employee_domain_model.Employee{
				ID: 1,
				CompanyId: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
			},
			expectedExistsBool: false,
			expectedError: nil,
		},
	}

	for _, value := range cases {
		salaryStatementRepository := &testtool.SalaryStatementRepositoryMock{}
		salaryStatementRepository.FakeGetSalaryStatement = value.fakesFunctions.FakeGetSalaryStatement
		salaryStatementExists, err := value.administrator.SalaryStatementExists(salaryStatementRepository, value.employee, 2000, time.April)
		assert.Equal(t, value.expectedExistsBool, salaryStatementExists)
		assert.Equal(t, value.expectedError, err)
	}
}