package fixed_earning_detail_domain_model

type FixedEarningDetail struct {
	ID             uint32
	FixedEarningID uint32
	Nominal        string
	Amount         int
}

// Not yet created value object for each attribute.
func NewFixedEarningDetail(id uint32, fixedEarningId uint32, nominal string, amount int) (*FixedEarningDetail, error) {
	return &FixedEarningDetail{id, fixedEarningId, nominal, amount}, nil
}
