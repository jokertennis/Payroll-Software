package repository

import (
	"context"
	"time"
	domainmodel "usr/local/go/src/main/domain-model"
)

type SalaryStatementRepository interface {
	GetSalaryStatement(ctx context.Context, employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*domainmodel.SalaryStatement, error)
}