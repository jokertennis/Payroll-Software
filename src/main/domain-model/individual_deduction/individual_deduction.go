package individual_deduction_domain_model

import (
	"usr/local/go/src/main/domain-model/individual_deduction_detail"
)

type IndividualDeduction struct {
	ID                          uint32
	Amount                      int
	Nominal                     string
	IndividualDeductionDetails  []individual_deduction_detail_domain_model.IndividualDeductionDetail
}

// Not yet created value object for each attribute.
func NewIndividualDeduction(id uint32, amount int, nominal string, individualDeductionDetails []individual_deduction_detail_domain_model.IndividualDeductionDetail) (*IndividualDeduction, error) {
	return &IndividualDeduction{id, amount, nominal, individualDeductionDetails}, nil
}