package salary_statement_application_service_test

import (
	"errors"
	"fmt"
	"net/http"
	"testing"
	"time"
	"usr/local/go/src/main/application-service/salary_statement_application_service"
	"usr/local/go/src/main/domain-model/administrator"
	"usr/local/go/src/main/domain-model/employee"
	"usr/local/go/src/main/domain-model/salary_statement"
	"usr/local/go/src/main/domain-service/repository/salary_statement_repository"
	"usr/local/go/src/testtool"

	"github.com/stretchr/testify/assert"
)

func TestCreateSalaryStatementIndividualUseCase(t *testing.T) {
	type fakesFunctions struct {
		FakeGetAdministratorByMailAddress func(mailAddress string) (*administrator_domain_model.Administrator, error)
		FakeGetEmployeeByMailAddress   func(mailAddress string) (*employee_domain_model.Employee, error)
		FakeGetSalaryStatement func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error)
		FakeCreateSalaryStatement func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error)
	}
	cases := map[string]struct {
		fakesFunctions     fakesFunctions
		expectedResult     *salary_statement_application_service.ResultOfCreateSalaryStatementIndividual
		expectedStatusCode int
		expectedError      error
	}{
		"Successfully create salary statement and get salary statement id. Error has not occurred.": {
			fakesFunctions: fakesFunctions{
				FakeGetAdministratorByMailAddress: func(mailAddress string) (*administrator_domain_model.Administrator, error) {
					administrator := &administrator_domain_model.Administrator{ID: 1, Name: "管理者A", CompanyId: 1}
					return administrator, nil
				},
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := &employee_domain_model.Employee{ID: 1, Name: "従業員A", CompanyId: 1}
					return employee, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
				FakeCreateSalaryStatement: func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
					return 1, nil
				},
			},
			expectedResult: &salary_statement_application_service.ResultOfCreateSalaryStatementIndividual{SalaryStatementId: 1},
			expectedStatusCode: http.StatusCreated,
			expectedError:      nil,
		},
		"Error has occurred when executed GetAdministratorByMailAddress()": {
			fakesFunctions: fakesFunctions{
				FakeGetAdministratorByMailAddress: func(mailAddress string) (*administrator_domain_model.Administrator, error) {
					return nil, errors.New("cannnot connect db.")
				},
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					return nil, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
				FakeCreateSalaryStatement: func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
					return 0, nil
				},
			},
			expectedResult: nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedError:      errors.New("InternalServerError:error:cannnot connect db."),
		},
		"Error has not occurred but administrator was nil when executed GetAdministratorByMailAddress()": {
			fakesFunctions: fakesFunctions{
				FakeGetAdministratorByMailAddress: func(mailAddress string) (*administrator_domain_model.Administrator, error) {
					return nil, nil
				},
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					return nil, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
				FakeCreateSalaryStatement: func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
					return 0, nil
				},
			},
			expectedResult: nil,
			expectedStatusCode: http.StatusUnauthorized,
			expectedError:      fmt.Errorf("unauthorized. Specified mailAddress is not found in registered user datas. MailAddress:%s", "administrator@example.com"),
		},
		"Error has occurred when executed GetEmployeeByMailAddress()": {
			fakesFunctions: fakesFunctions{
				FakeGetAdministratorByMailAddress: func(mailAddress string) (*administrator_domain_model.Administrator, error) {
					administrator := &administrator_domain_model.Administrator{ID: 1, Name: "管理者A", CompanyId: 1}
					return administrator, nil
				},
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					return nil, errors.New("cannnot connect db.")
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
				FakeCreateSalaryStatement: func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
					return 0, nil
				},
			},
			expectedResult: nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedError:      errors.New("InternalServerError:error:cannnot connect db."),
		},
		"Error has not occurred but employee was nil when executed GetEmployeeByMailAddress()": {
			fakesFunctions: fakesFunctions{
				FakeGetAdministratorByMailAddress: func(mailAddress string) (*administrator_domain_model.Administrator, error) {
					administrator := &administrator_domain_model.Administrator{ID: 1, Name: "管理者A", CompanyId: 1}
					return administrator, nil
				},
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					return nil, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
				FakeCreateSalaryStatement: func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
					return 0, nil
				},
			},
			expectedResult: nil,
			expectedStatusCode: http.StatusBadRequest,
			expectedError:      fmt.Errorf("badrequest. Specified mailAddress is not found in registered user datas. MailAddressOfEmployee:%s", "bob@example.com"),
		},
		"Error has occurred when id of company that employee belongs to is different from id of company that administrator belongs to.": {
			fakesFunctions: fakesFunctions{
				FakeGetAdministratorByMailAddress: func(mailAddress string) (*administrator_domain_model.Administrator, error) {
					administrator := &administrator_domain_model.Administrator{ID: 1, Name: "管理者A", CompanyId: 1}
					return administrator, nil
				},
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := &employee_domain_model.Employee{ID: 1, Name: "従業員A", CompanyId: 10}
					return employee, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
				FakeCreateSalaryStatement: func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
					return 0, nil
				},
			},
			expectedResult: nil,
			expectedStatusCode: http.StatusBadRequest,
			expectedError:      errors.New("operations related to employee who belongs to a company different from your own are not possible"),
		},
		"Error has occurred when executed SalaryStatementExists()": {
			fakesFunctions: fakesFunctions{
				FakeGetAdministratorByMailAddress: func(mailAddress string) (*administrator_domain_model.Administrator, error) {
					administrator := &administrator_domain_model.Administrator{ID: 1, Name: "管理者A", CompanyId: 1}
					return administrator, nil
				},
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := &employee_domain_model.Employee{ID: 1, Name: "従業員A", CompanyId: 1}
					return employee, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, errors.New("cannnot connect db.")
				},
				FakeCreateSalaryStatement: func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
					return 0, nil
				},
			},
			expectedResult: nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedError:      errors.New("InternalServerError:error:failed to get salary statement.error:cannnot connect db."),
		},
		"Error has not occurred but salary statement which has payday that match specified year and month was already existed": {
			fakesFunctions: fakesFunctions{
				FakeGetAdministratorByMailAddress: func(mailAddress string) (*administrator_domain_model.Administrator, error) {
					administrator := &administrator_domain_model.Administrator{ID: 1, Name: "管理者A", CompanyId: 1}
					return administrator, nil
				},
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := &employee_domain_model.Employee{ID: 1, Name: "従業員A", CompanyId: 1}
					return employee, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return &salary_statement_domain_model.SalaryStatement{ID: 1}, nil
				},
				FakeCreateSalaryStatement: func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
					return 0, nil
				},
			},
			expectedResult: nil,
			expectedStatusCode: http.StatusBadRequest,
			expectedError:      errors.New("badrequest. SalaryStatement which has payday that match year and month of specified payday was already existed"),
		},
		"Error has occurred when executed CreateSalaryStatementByUsingIndividualData()": {
			fakesFunctions: fakesFunctions{
				FakeGetAdministratorByMailAddress: func(mailAddress string) (*administrator_domain_model.Administrator, error) {
					administrator := &administrator_domain_model.Administrator{ID: 1, Name: "管理者A", CompanyId: 1}
					return administrator, nil
				},
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := &employee_domain_model.Employee{ID: 1, Name: "従業員A", CompanyId: 1}
					return employee, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
				FakeCreateSalaryStatement: func(salaryStatementEntry salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
					return 0, errors.New("cannnot connect db.")
				},
			},
			expectedResult: nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedError:      errors.New("internalServerError:error:failed to create salary statement.error:cannnot connect db."),
		},
	}

	for _, value := range cases {
		administratorRepository := &testtool.AdministratorRepositoryMock{}
		administratorRepository.FakeGetAdministratorByMailAddress = value.fakesFunctions.FakeGetAdministratorByMailAddress
		employeeRepository := &testtool.EmployeeRepositoryMock{}
		employeeRepository.FakeGetEmployeeByMailAddress = value.fakesFunctions.FakeGetEmployeeByMailAddress
		salaryStatementRepository := &testtool.SalaryStatementRepositoryMock{}
		salaryStatementRepository.FakeGetSalaryStatement = value.fakesFunctions.FakeGetSalaryStatement
		salaryStatementRepository.FakeCreateSalaryStatement = value.fakesFunctions.FakeCreateSalaryStatement
		result, statusCode, err := salary_statement_application_service.CreateSalaryStatementIndividualUseCase(administratorRepository, employeeRepository, salaryStatementRepository, "administrator@example.com", "bob@example.com", salary_statement_repository.SalaryStatementEntryByUsingIndividualDatas{})
		assert.Equal(t, value.expectedResult, result)
		assert.Equal(t, value.expectedStatusCode, statusCode)
		assert.Equal(t, value.expectedError, err)
	}
}