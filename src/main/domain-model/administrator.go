package domainmodel

type Administrator struct {
	ID          uint32
	CompanyId   uint16
	Name        string
	MailAddress string
	Password    string
}

// Not yet created value object for each attribute.
func NewAdministrator(id uint32, companyId uint16, name string, mailAddress string, password string) (*Administrator, error) {
	return &Administrator{id, companyId, name, mailAddress, password}, nil
}