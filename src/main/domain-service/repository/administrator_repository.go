package repository

import (
	"context"
	domainmodel "usr/local/go/src/main/domain-model"
)

type AdministratorRepository interface {
	GetAdministratorByMailAddress(ctx context.Context, mailAddress string) (*domainmodel.Administrator, error)
}