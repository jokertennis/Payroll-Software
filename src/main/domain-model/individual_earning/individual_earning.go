package individual_earning_domain_model

import "usr/local/go/src/main/domain-model/individual_earning_detail"

type IndividualEarning struct {
	ID                          uint32
	Amount                      int
	Nominal                     string
	IndividualEarningDetails  []individual_earning_detail_domain_model.IndividualEarningDetail
}

// Not yet created value object for each attribute.
func NewIndividualEarning(id uint32, amount int, nominal string, individualEarningDetails []individual_earning_detail_domain_model.IndividualEarningDetail) (*IndividualEarning, error) {
	return &IndividualEarning{id, amount, nominal, individualEarningDetails}, nil
}