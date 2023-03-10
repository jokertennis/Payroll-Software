package domainmodel

type IndividualDeduction struct {
	ID                          uint32
	Amount                      int
	Nominal                     string
	IndividualDeductionDetails  []IndividualDeductionDetail
}

// Not yet created value object for each attribute.
func NewIndividualDeduction(id uint32, amount int, nominal string, individualDeductionDetails []IndividualDeductionDetail) (*IndividualDeduction, error) {
	return &IndividualDeduction{id, amount, nominal, individualDeductionDetails}, nil
}