package domainmodel

type IndividualEarning struct {
	ID                          uint32
	Amount                      int
	Nominal                     string
	IndividualEarningDetails  []IndividualEarningDetail
}

// Not yet created value object for each attribute.
func NewIndividualEarning(id uint32, amount int, nominal string, individualEarningDetails []IndividualEarningDetail) (*IndividualEarning, error) {
	return &IndividualEarning{id, amount, nominal, individualEarningDetails}, nil
}