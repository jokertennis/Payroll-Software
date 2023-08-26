package earning_domain_model

import (
	earning_detail_domain_model "usr/local/go/src/main/domain-model/earning_detail"
)

type IndividualEarning struct {
	ID                       uint32
	Amount                   int
	Nominal                  string
	IndividualEarningDetails []earning_detail_domain_model.IndividualEarningDetail
}

// Not yet created value object for each attribute.
func NewIndividualEarning(id uint32, amount int, nominal string, individualEarningDetails []earning_detail_domain_model.IndividualEarningDetail) (*IndividualEarning, error) {
	return &IndividualEarning{id, amount, nominal, individualEarningDetails}, nil
}
