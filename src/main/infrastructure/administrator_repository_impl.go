package infrastructure

import (
	"context"
	"database/sql"
	"usr/local/go/server/gen/models"
	domainmodel "usr/local/go/src/main/domain-model"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AdministratorRepository struct {
	db *sql.DB
}

func NewAdministratorRepository(db *sql.DB) AdministratorRepository {
	return AdministratorRepository{db: db}
}

func (r *AdministratorRepository) GetAdministratorByMailAddress(ctx context.Context, mailAddress string) (*domainmodel.Administrator, error) {
	administrators, err := models.Administrators(qm.Where("mail_address=?", mailAddress)).All(ctx, r.db)
	if err != nil {
		return nil, err
	}
	if administrators == nil {
		return nil, nil
	}

	return MappingAdministratorDomainObject(administrators[0])
}

func MappingAdministratorDomainObject(m *models.Administrator) (*domainmodel.Administrator, error) {
	administrator, err := domainmodel.NewAdministrator(m.ID, m.CompanyID, m.Name, m.MailAddress, m.Password)
	if err != nil {
        return nil, err
    }
    return administrator, nil
}