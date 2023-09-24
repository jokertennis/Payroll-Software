package employee_domain_model_test

import (
	"fmt"
	"testing"
	"github.com/jokertennis/Payroll-Software/src/main/domain-model/employee"
	"github.com/jokertennis/Payroll-Software/src/main/domain-model/salary_statement"
	"github.com/jokertennis/Payroll-Software/src/testtool"

	"github.com/stretchr/testify/assert"
)

func TestGetAllSalaryStatements(t *testing.T) {
	type fakesFunctions struct {
        FakeGetAllSalaryStatements func(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error)
    }
	cases := map[string]struct {
		employee                       *employee_domain_model.Employee
		fakesFunctions                 fakesFunctions
		expectedSalaryStatementsList   []*salary_statement_domain_model.SalaryStatement
		expectedError                  error
	}{
		"Successfully get SalaryStatement.Error has not occurred.": {
			employee: &employee_domain_model.Employee{
				ID: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error) {
					var salaryStatementsList []*salary_statement_domain_model.SalaryStatement
					salaryStatement := salary_statement_domain_model.SalaryStatement{
						ID: 1,
					}
					salaryStatementsList = append(salaryStatementsList, &salaryStatement)
                    return salaryStatementsList, nil
                },
			},
			expectedSalaryStatementsList: []*salary_statement_domain_model.SalaryStatement{{ID: 1},},
			expectedError: nil,
		},
		"SalaryStatement that match the search criteria was not found.Error has not occurred.": {
			employee: &employee_domain_model.Employee{
				ID: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error) {
                    return nil, nil
                },
			},
			expectedSalaryStatementsList: nil,
			expectedError: nil,
		},
		"error has occurred.": {
			employee: &employee_domain_model.Employee{
				ID: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error) {
                    return nil, fmt.Errorf("cannot connect db.")
                },
			},
			expectedSalaryStatementsList: nil,
			expectedError: fmt.Errorf("failed to get all salary statements.error:cannot connect db."),
		},
	}

	for _, value := range cases {
		salaryStatementRepository := &testtool.SalaryStatementRepositoryMock{}
		salaryStatementRepository.FakeGetAllSalaryStatements = value.fakesFunctions.FakeGetAllSalaryStatements
		salaryStatements, err := value.employee.GetAllSalaryStatements(salaryStatementRepository)
		assert.Equal(t, value.expectedSalaryStatementsList, salaryStatements)
		assert.Equal(t, value.expectedError, err)
	}
}