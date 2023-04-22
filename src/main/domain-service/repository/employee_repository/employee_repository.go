package employee_repository

import (
	"usr/local/go/src/main/domain-model/employee"
)

type EmployeeRepository interface {
	GetEmployeeByMailAddress(mailAddress string) (*employee_domain_model.Employee, error)
}