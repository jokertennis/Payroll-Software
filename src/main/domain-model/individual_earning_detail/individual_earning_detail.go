package individual_earning_detail_domain_model

type IndividualEarningDetail struct {
	ID                    uint32
	IndividualEarningID   uint32
	Nominal               string
	Amount                int
}

// Not yet created value object for each attribute.
func NewIndividualEarningDetail(id uint32, individualEarningId uint32, nominal string, amount int) (*IndividualEarningDetail, error) {
	return &IndividualEarningDetail{id, individualEarningId, nominal, amount}, nil
}