package infrastructure

import (
	"context"
	"database/sql"
	"usr/local/go/server/gen/models"
	domainmodel "usr/local/go/src/main/domain-model"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type EmployeeRepository struct {
	db *sql.DB
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return EmployeeRepository{db: db}
}

func (r *EmployeeRepository) GetEmployeeByMailAddress(ctx context.Context, mailAddress string) (*domainmodel.Employee, error) {
	employees, err := models.Employees(qm.Where("mail_address=?", mailAddress)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	if employees == nil {
		return nil, nil
	}

	return MappingEmployeeDomainObject(employees[0])
}

func MappingEmployeeDomainObject(m *models.Employee) (*domainmodel.Employee, error) {
	employee, err := domainmodel.NewEmployee(m.ID, m.CompanyID, m.Name, m.MailAddress, m.Password)
	if err != nil {
        return nil, err
    }
    return employee, nil
}