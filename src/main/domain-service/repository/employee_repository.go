package repository

import (
	"context"
	domainmodel "usr/local/go/src/main/domain-model"
)

type EmployeeRepository interface {
	GetEmployeeByMailAddress(ctx context.Context, mailAddress string) (*domainmodel.Employee, error)
}