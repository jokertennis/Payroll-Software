package domainmodel

type FixedEarning struct {
	ID                  uint32
	Amount              int
	Nominal             string
	FixedEarningDetails []FixedEarningDetail
}

// Not yet created value object for each attribute.
func NewFixedEarning(id uint32, amount int, nominal string, fixedEarningDetails []FixedEarningDetail) (*FixedEarning, error) {
	return &FixedEarning{id, amount, nominal, fixedEarningDetails}, nil
}
