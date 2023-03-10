package domainmodel

type FixedDeduction struct {
	ID                     uint32
	Amount                 int
	Nominal                string
	FixedDeductionDetails  []FixedDeductionDetail
}

// Not yet created value object for each attribute.
func NewFixedDeduction(id uint32, amount int, nominal string, fixedDeductionDetails []FixedDeductionDetail) (*FixedDeduction, error) {
	return &FixedDeduction{id, amount, nominal, fixedDeductionDetails}, nil
}