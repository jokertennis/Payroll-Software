package earning_domain_model

import (
	earning_detail_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/earning_detail"
)

type EarningType string

const (
	IndividualEarning EarningType = "Individual"
	FixedEarning      EarningType = "Fixed"
)

type Earning struct {
	ID             uint32
	Amount         int
	Nominal        string
	EarningType    EarningType
	EarningDetails []earning_detail_domain_model.EarningDetail
}

// Not yet created value object for each attribute.
func NewEarning(id uint32, amount int, nominal string, earningType EarningType, earningDetails []earning_detail_domain_model.EarningDetail) (Earning, error) {
	return Earning{id, amount, nominal, earningType, earningDetails}, nil
}
