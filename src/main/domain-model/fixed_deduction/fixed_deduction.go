package fixed_deduction_domain_model

import (
	"usr/local/go/src/main/domain-model/fixed_deduction_detail"
)

type FixedDeduction struct {
	ID                     uint32
	Amount                 int
	Nominal                string
	FixedDeductionDetails  []fixed_deduction_detail.FixedDeductionDetail
}

// Not yet created value object for each attribute.
func NewFixedDeduction(id uint32, amount int, nominal string, fixedDeductionDetails []fixed_deduction_detail.FixedDeductionDetail) (*FixedDeduction, error) {
	return &FixedDeduction{id, amount, nominal, fixedDeductionDetails}, nil
}