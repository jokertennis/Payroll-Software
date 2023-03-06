package domainmodel

type Employee struct {
	ID          uint32
	CompanyId   uint16
	Name        string
	MailAddress string
	Password    string
}

// Not yet created value object for each attribute.
func NewEmployee(id uint32, companyId uint16, name string, mailAddress string, password string) (*Employee, error) {
	return &Employee{id, companyId, name, mailAddress, password}, nil
}