package infrastructure

import (
	"context"
	"database/sql"
	"usr/local/go/server/gen/models"
	"usr/local/go/src/main/domain-model/administrator"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type AdministratorRepository struct {
	ctx context.Context
	db *sql.DB
}

func NewAdministratorRepository(ctx context.Context, db *sql.DB) AdministratorRepository {
	return AdministratorRepository{
		ctx: ctx,
		db: db,
	}
}

func (r *AdministratorRepository) GetAdministratorByMailAddress(mailAddress string) (*administrator.Administrator, error) {
	administrators, err := models.Administrators(qm.Where("mail_address=?", mailAddress)).All(r.ctx, r.db)
	if err != nil {
		return nil, err
	}
	if administrators == nil {
		return nil, nil
	}

	return MappingAdministratorDomainObject(administrators[0])
}