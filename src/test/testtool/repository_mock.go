package testtool

import (
	"time"
	"usr/local/go/src/main/domain-model/employee"
	"usr/local/go/src/main/domain-model/salary_statement"
	"usr/local/go/src/main/domain-service/repository/employee_repository"
	"usr/local/go/src/main/domain-service/repository/salary_statement_repository"
)

type EmployeeRepositoryMock struct {
	EmployeeRepository           employee_repository.EmployeeRepository
	FakeGetEmployeeByMailAddress func(mailAddress string) (*employee.Employee, error)
}

func (m *EmployeeRepositoryMock) GetEmployeeByMailAddress(mailAddress string) (*employee.Employee, error) {
	return m.FakeGetEmployeeByMailAddress(mailAddress)
}

type SalaryStatementRepositoryMock struct {
	SalaryStatementRepository  salary_statement_repository.SalaryStatementRepository
	FakeGetSalaryStatement     func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement.SalaryStatement, error)
	FakeGetAllSalaryStatements func(employeeId uint32) ([]*salary_statement.SalaryStatement, error)
	FakeCreateSalaryStatement  func(salaryStatementEntry salary_statement.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error)
}

func (m *SalaryStatementRepositoryMock) GetSalaryStatement(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement.SalaryStatement, error) {
	return m.FakeGetSalaryStatement(employeeId, yearOfPayday, monthOfPayday)
}

func (m *SalaryStatementRepositoryMock) GetAllSalaryStatements(employeeId uint32) ([]*salary_statement.SalaryStatement, error) {
	return m.FakeGetAllSalaryStatements(employeeId)
}

func (m *SalaryStatementRepositoryMock) CreateSalaryStatement(salaryStatementEntry salary_statement.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
	return 
}