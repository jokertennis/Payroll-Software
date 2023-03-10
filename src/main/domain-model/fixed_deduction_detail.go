package domainmodel

type FixedDeductionDetail struct {
	ID               uint32
	FixedDeductionID uint32
	Nominal          string
	Amount           int
}

// Not yet created value object for each attribute.
func NewFixedDeductionDetail(id uint32, fixedDeductionId uint32, nominal string, amount int) (*IndividualDeductionDetail, error) {
	return &IndividualDeductionDetail{id, fixedDeductionId, nominal, amount}, nil
}