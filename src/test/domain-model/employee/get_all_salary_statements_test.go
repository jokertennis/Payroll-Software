package employee

import (
	"fmt"
	"testing"
	"usr/local/go/src/main/domain-model/employee"
	"usr/local/go/src/main/domain-model/salary_statement"
	"usr/local/go/src/test/testtool"

	"github.com/stretchr/testify/assert"
)

func TestGetAllSalaryStatements(t *testing.T) {
	type fakesFunctions struct {
        FakeGetAllSalaryStatements func(employeeId uint32) ([]*salary_statement.SalaryStatement, error)
    }
	cases := map[string]struct {
		employee                       *employee.Employee
		fakesFunctions                 fakesFunctions
		expectedSalaryStatementsList   []*salary_statement.SalaryStatement
		expectedError                  error
	}{
		"Successfully get SalaryStatement.Error has not occurred.": {
			employee: &employee.Employee{
				ID: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement.SalaryStatement, error) {
					var salaryStatementsList []*salary_statement.SalaryStatement
					salaryStatement := salary_statement.SalaryStatement{
						ID: 1,
					}
					salaryStatementsList = append(salaryStatementsList, &salaryStatement)
                    return salaryStatementsList, nil
                },
			},
			expectedSalaryStatementsList: []*salary_statement.SalaryStatement{{ID: 1},},
			expectedError: nil,
		},
		"SalaryStatement that match the search criteria was not found.Error has not occurred.": {
			employee: &employee.Employee{
				ID: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement.SalaryStatement, error) {
                    return nil, nil
                },
			},
			expectedSalaryStatementsList: nil,
			expectedError: nil,
		},
		"error has occurred.": {
			employee: &employee.Employee{
				ID: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement.SalaryStatement, error) {
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