package employee

import (
	"fmt"
	"testing"
	"time"
	"usr/local/go/src/main/domain-model/employee"
	"usr/local/go/src/main/domain-model/salary_statement"
	"usr/local/go/src/main/domain-service/repository/salary_statement_repository"

	"github.com/stretchr/testify/assert"
)

type salaryStatementRepositoryMock struct {
	SalaryStatementRepository salary_statement_repository.SalaryStatementRepository
	FakeGetSalaryStatement func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement.SalaryStatement, error)
}

func (m *salaryStatementRepositoryMock) GetSalaryStatement(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement.SalaryStatement, error) {
    return m.FakeGetSalaryStatement(employeeId, yearOfPayday, monthOfPayday)
}

func (m *salaryStatementRepositoryMock) GetAllSalaryStatements(employeeId uint32) ([]*salary_statement.SalaryStatement, error) {
	return nil, nil
}

func TestGetSpecificSalaryStatement(t *testing.T) {
	type fakesFunctions struct {
        FakeGetSalaryStatement func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement.SalaryStatement, error)
    }
	cases := map[string]struct {
		employee                  *employee.Employee
		fakesFunctions            fakesFunctions
		expectedSalaryStatement   *salary_statement.SalaryStatement
		expectedError             error
	}{
		"Successfully get SalaryStatement.Error has not occurred.": {
			employee: &employee.Employee{
				ID: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement.SalaryStatement, error) {
                    salaryStatement := &salary_statement.SalaryStatement{ID: 1}
                    return salaryStatement, nil
                },
			},
			expectedSalaryStatement: &salary_statement.SalaryStatement{ID: 1},
			expectedError: nil,
		},
		"SalaryStatement that match the search criteria was not found.Error has not occurred.": {
			employee: &employee.Employee{
				ID: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement.SalaryStatement, error) {
                    return nil, nil
                },
			},
			expectedSalaryStatement: nil,
			expectedError: nil,
		},
		"error has occurred.": {
			employee: &employee.Employee{
				ID: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement.SalaryStatement, error) {
                    return nil, fmt.Errorf("cannot connect db.")
                },
			},
			expectedSalaryStatement: nil,
			expectedError: fmt.Errorf("failed to get salary statement.error:cannot connect db."),
		},
	}

	for _, value := range cases {
		salaryStatementRepository := &salaryStatementRepositoryMock{}
		salaryStatementRepository.FakeGetSalaryStatement = value.fakesFunctions.FakeGetSalaryStatement
		salaryStatement, err := value.employee.GetSpecificSalaryStatement(salaryStatementRepository, 2022, time.May)
		assert.Equal(t, value.expectedSalaryStatement, salaryStatement)
		assert.Equal(t, value.expectedError, err)
	}
}