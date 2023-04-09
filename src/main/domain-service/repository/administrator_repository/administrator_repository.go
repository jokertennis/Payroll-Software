package administrator_repository

import (
	"usr/local/go/src/main/domain-model/administrator"
)

type AdministratorRepository interface {
	GetAdministratorByMailAddress(mailAddress string) (*administrator.Administrator, error)
}