package infrastructure

import (
	"context"
	"database/sql"
	"usr/local/go/server/gen/models"
	"usr/local/go/src/main/domain-model/employee"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type EmployeeRepository struct {
	ctx context.Context
	db *sql.DB
}

func NewEmployeeRepository(ctx context.Context, db *sql.DB) EmployeeRepository {
	return EmployeeRepository{
		ctx: ctx,
		db: db,
	}
}

func (r *EmployeeRepository) GetEmployeeByMailAddress(mailAddress string) (*employee_domain_model.Employee, error) {
	employees, err := models.Employees(qm.Where("mail_address=?", mailAddress)).All(r.ctx, r.db)
	if err != nil {
		return nil, err
	}
	if employees == nil {
		return nil, nil
	}

	return MappingEmployeeDomainObject(employees[0])
}