package fixed_earning

import "usr/local/go/src/main/domain-model/fixed_earning_detail"

type FixedEarning struct {
	ID                  uint32
	Amount              int
	Nominal             string
	FixedEarningDetails []fixed_earning_detail.FixedEarningDetail
}

// Not yet created value object for each attribute.
func NewFixedEarning(id uint32, amount int, nominal string, fixedEarningDetails []fixed_earning_detail.FixedEarningDetail) (*FixedEarning, error) {
	return &FixedEarning{id, amount, nominal, fixedEarningDetails}, nil
}
