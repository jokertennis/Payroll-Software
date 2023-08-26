package earning_detail_domain_model

type EarningDetail struct {
	ID        uint32
	EarningID uint32
	Nominal   string
	Amount    int
}

// Not yet created value object for each attribute.
func NewEarningDetail(id uint32, earningId uint32, nominal string, amount int) (EarningDetail, error) {
	return EarningDetail{id, earningId, nominal, amount}, nil
}
