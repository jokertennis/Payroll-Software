package administrator_repository

import (
	"github.com/jokertennis/Payroll-Software/src/main/domain-model/administrator"
)

type AdministratorRepository interface {
	GetAdministratorByMailAddress(mailAddress string) (*administrator_domain_model.Administrator, error)
}