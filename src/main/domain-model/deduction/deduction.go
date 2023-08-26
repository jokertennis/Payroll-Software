package deduction_domain_model

import (
	deduction_detail_domain_model "usr/local/go/src/main/domain-model/deduction_detail"
)

type DeductionType string

const (
	IndividualDeduction DeductionType = "Individual"
	FixedDeduction      DeductionType = "Fixed"
)

type Deduction struct {
	ID               uint32
	Amount           int
	Nominal          string
	DeductionType    DeductionType
	DeductionDetails []deduction_detail_domain_model.DeductionDetail
}

// Not yet created value object for each attribute.
func NewDeduction(id uint32, amount int, nominal string, deductionType DeductionType, deductionDetails []deduction_detail_domain_model.DeductionDetail) (Deduction, error) {
	return Deduction{id, amount, nominal, deductionType, deductionDetails}, nil
}
