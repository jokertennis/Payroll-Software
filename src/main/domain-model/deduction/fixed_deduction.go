package deduction_domain_model

import (
	deduction_detail_domain_model "usr/local/go/src/main/domain-model/deduction_detail"
)

type FixedDeduction struct {
	ID                    uint32
	Amount                int
	Nominal               string
	FixedDeductionDetails []deduction_detail_domain_model.FixedDeductionDetail
}

// Not yet created value object for each attribute.
func NewFixedDeduction(id uint32, amount int, nominal string, fixedDeductionDetails []deduction_detail_domain_model.FixedDeductionDetail) (*FixedDeduction, error) {
	return &FixedDeduction{id, amount, nominal, fixedDeductionDetails}, nil
}
