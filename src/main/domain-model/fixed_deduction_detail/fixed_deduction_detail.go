package fixed_deduction_detail_domain_model

type FixedDeductionDetail struct {
	ID               uint32
	FixedDeductionID uint32
	Nominal          string
	Amount           int
}

// Not yet created value object for each attribute.
func NewFixedDeductionDetail(id uint32, fixedDeductionId uint32, nominal string, amount int) (*FixedDeductionDetail, error) {
	return &FixedDeductionDetail{id, fixedDeductionId, nominal, amount}, nil
}