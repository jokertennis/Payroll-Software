package administrator_domain_model_test

import (
	"errors"
	"testing"
	"usr/local/go/src/main/domain-model/administrator"
	"usr/local/go/src/main/domain-model/employee"
	"usr/local/go/src/main/domain-service/repository/salary_statement_repository"
	"usr/local/go/src/testtool"

	"github.com/stretchr/testify/assert"
)

func TestCreateSalaryStatementByUsingIndividualData(t *testing.T) {
	type fakesFunctions struct {
        FakeCreateSalaryStatement func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error)
    }
	cases := map[string]struct {
		administrator               administrator_domain_model.Administrator
		employee                    employee_domain_model.Employee
		fakesFunctions              fakesFunctions
		expectedSalaryStatementId   uint32
		expectedError               error
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
				FakeCreateSalaryStatement: func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
					return 10, nil
				},
			},
			expectedSalaryStatementId: 10,
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
				FakeCreateSalaryStatement: func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
					return 0, nil
				},
			},
			expectedSalaryStatementId: 0,
			expectedError: errors.New("operations related to employee who belongs to a company different from your own are not possible"),
		},
		"Error has occurred when executed CreateSalaryStatement()": {
			administrator: administrator_domain_model.Administrator{
				ID: 1,
				CompanyId: 1,
			},
			employee: employee_domain_model.Employee{
				ID: 1,
				CompanyId: 1,
			},
			fakesFunctions: fakesFunctions{
				FakeCreateSalaryStatement: func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
					return 0, errors.New("cannnot connect db")
				},
			},
			expectedSalaryStatementId: 0,
			expectedError: errors.New("failed to create salary statement.error:cannnot connect db"),
		},
	}

	for _, value := range cases {
		salaryStatementRepository := &testtool.SalaryStatementRepositoryMock{}
		salaryStatementRepository.FakeCreateSalaryStatement = value.fakesFunctions.FakeCreateSalaryStatement
		salaryStatementId, err := value.administrator.CreateSalaryStatementByUsingIndividualData(salaryStatementRepository, value.employee, salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas{})
		assert.Equal(t, value.expectedSalaryStatementId, salaryStatementId)
		assert.Equal(t, value.expectedError, err)
	}
}