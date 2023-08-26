package earning_domain_model

import (
	earning_detail_domain_model "usr/local/go/src/main/domain-model/earning_detail"
)

type FixedEarning struct {
	ID                  uint32
	Amount              int
	Nominal             string
	FixedEarningDetails []earning_detail_domain_model.FixedEarningDetail
}

// Not yet created value object for each attribute.
func NewFixedEarning(id uint32, amount int, nominal string, fixedEarningDetails []earning_detail_domain_model.FixedEarningDetail) (*FixedEarning, error) {
	return &FixedEarning{id, amount, nominal, fixedEarningDetails}, nil
}
