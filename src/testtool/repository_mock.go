package testtool

import (
	"time"
	administrator_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/administrator"
	employee_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/employee"
	salary_statement_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/salary_statement"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/administrator_repository"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/employee_repository"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/salary_statement_repository"
)

type AdministratorRepositoryMock struct {
	AdministratorRepository           administrator_repository.AdministratorRepository
	FakeGetAdministratorByMailAddress func(mailAddress string) (*administrator_domain_model.Administrator, error)
}

func (m *AdministratorRepositoryMock) GetAdministratorByMailAddress(mailAddress string) (*administrator_domain_model.Administrator, error) {
	return m.FakeGetAdministratorByMailAddress(mailAddress)
}

type EmployeeRepositoryMock struct {
	EmployeeRepository           employee_repository.EmployeeRepository
	FakeGetEmployeeByMailAddress func(mailAddress string) (*employee_domain_model.Employee, error)
}

func (m *EmployeeRepositoryMock) GetEmployeeByMailAddress(mailAddress string) (*employee_domain_model.Employee, error) {
	return m.FakeGetEmployeeByMailAddress(mailAddress)
}

type SalaryStatementRepositoryMock struct {
	SalaryStatementRepository  salary_statement_repository.SalaryStatementRepository
	FakeGetSalaryStatement     func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error)
	FakeGetAllSalaryStatements func(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error)
	FakeCreateSalaryStatement  func(salaryStatementEntry salary_statement_repository.SalaryStatementEntry) (salaryStatementId uint32, err error)
}

func (m *SalaryStatementRepositoryMock) GetSalaryStatement(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
	return m.FakeGetSalaryStatement(employeeId, yearOfPayday, monthOfPayday)
}

func (m *SalaryStatementRepositoryMock) GetAllSalaryStatements(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error) {
	return m.FakeGetAllSalaryStatements(employeeId)
}

func (m *SalaryStatementRepositoryMock) CreateSalaryStatement(salaryStatementEntry salary_statement_repository.SalaryStatementEntry) (salaryStatementId uint32, err error) {
	return m.FakeCreateSalaryStatement(salaryStatementEntry)
}