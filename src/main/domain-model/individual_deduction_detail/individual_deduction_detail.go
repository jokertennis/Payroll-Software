package individual_deduction_detail

type IndividualDeductionDetail struct {
	ID                    uint32
	IndividualDeductionID uint32
	Nominal               string
	Amount                int
}

// Not yet created value object for each attribute.
func NewIndividualDeductionDetail(id uint32, individualDeductionId uint32, nominal string, amount int) (*IndividualDeductionDetail, error) {
	return &IndividualDeductionDetail{id, individualDeductionId, nominal, amount}, nil
}