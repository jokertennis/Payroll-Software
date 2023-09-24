package employee_repository

import (
	"github.com/jokertennis/Payroll-Software/src/main/domain-model/employee"
)

type EmployeeRepository interface {
	GetEmployeeByMailAddress(mailAddress string) (*employee_domain_model.Employee, error)
}