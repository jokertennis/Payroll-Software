package deduction_detail_domain_model

type DeductionDetail struct {
	ID          uint32
	DeductionID uint32
	Nominal     string
	Amount      int
}

// Not yet created value object for each attribute.
func NewDeductionDetail(id uint32, deductionId uint32, nominal string, amount int) (DeductionDetail, error) {
	return DeductionDetail{id, deductionId, nominal, amount}, nil
}
