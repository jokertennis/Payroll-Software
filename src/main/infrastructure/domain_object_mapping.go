package infrastructure

import (
	"usr/local/go/server/gen/models"
	"usr/local/go/src/main/domain-model/administrator"
	"usr/local/go/src/main/domain-model/employee"
	"usr/local/go/src/main/domain-model/fixed_deduction"
	"usr/local/go/src/main/domain-model/fixed_deduction_detail"
	"usr/local/go/src/main/domain-model/fixed_earning"
	"usr/local/go/src/main/domain-model/fixed_earning_detail"
	"usr/local/go/src/main/domain-model/individual_deduction"
	"usr/local/go/src/main/domain-model/individual_deduction_detail"
	"usr/local/go/src/main/domain-model/individual_earning"
	"usr/local/go/src/main/domain-model/individual_earning_detail"
	"usr/local/go/src/main/domain-model/salary_statement"
)

func MappingEmployeeDomainObject(m *models.Employee) (*employee.Employee, error) {
	if m == nil {
		return nil, nil
	}
	employee, err := employee.NewEmployee(m.ID, m.CompanyID, m.Name, m.MailAddress, m.Password)
	if err != nil {
        return nil, err
    }
    return employee, nil
}

func MappingAdministratorDomainObject(m *models.Administrator) (*administrator.Administrator, error) {
	if m == nil {
		return nil, nil
	}
	administrator, err := administrator.NewAdministrator(m.ID, m.CompanyID, m.Name, m.MailAddress, m.Password)
	if err != nil {
        return nil, err
    }
    return administrator, nil
}

func MappingSalaryStatementDomainObject(m *models.SalaryStatement) (*salary_statement.SalaryStatement, error) {
	if m == nil {
		return nil, nil
	}
	var err error
	var mappingIndividualEarning *individual_earning.IndividualEarning
	if m.IndividualEarningID.Valid{
		mappingIndividualEarning, err = MappingIndividualEarning(m.R.IndividualEarning)
		if err != nil {
			return nil, err
		}
	} else {
		mappingIndividualEarning = nil
	}

	var mappingFixedEarning *fixed_earning.FixedEarning
	if m.FixedEarningID.Valid {
		mappingFixedEarning, err = MappingFixedEarning(m.R.FixedEarning)
		if err != nil {
			return nil, err
		}
	} else {
		mappingFixedEarning = nil
	}

	var mappingIndividualDeduction *individual_deduction.IndividualDeduction
	if m.IndividualDeductionID.Valid {
		mappingIndividualDeduction, err = MappingIndividualDeduction(m.R.IndividualDeduction)
		if err != nil {
			return nil, err
		}
	} else {
		mappingIndividualDeduction = nil
	}

	var mappingFixedDeduction *fixed_deduction.FixedDeduction
	if m.FixedDeductionID.Valid {
		mappingFixedDeduction, err = MappingFixedDeduction(m.R.FixedDeduction)
		if err != nil {
			return nil, err
		}
	} else {
		mappingFixedDeduction = nil
	}

	salaryStatement, err := salary_statement.NewSalaryStatement(m.ID, mappingIndividualEarning, mappingFixedEarning, mappingIndividualDeduction, mappingFixedDeduction, m.EmployeeID, m.Nominal, m.Payday, m.TargetPeriod)
	if err != nil {
        return nil, err
    }
    return salaryStatement, nil
}

func MappingIndividualEarning(m *models.IndividualEarning) (*individual_earning.IndividualEarning, error) {
	if m == nil {
		return nil, nil
	}
	var individualEarningDetailsList []individual_earning_detail.IndividualEarningDetail
	for _, indiindividualEarningDetail := range m.R.IndividualEarningDetails {
		mappingIndiindividualEarningDetail, err := MappingIndividualEarningDetail(indiindividualEarningDetail)
		if err != nil {
			return nil, err
		}
		individualEarningDetailsList = append(individualEarningDetailsList, *mappingIndiindividualEarningDetail)
	}
	individualEarning, err := individual_earning.NewIndividualEarning(m.ID, m.Amount, m.Nominal, individualEarningDetailsList)
	if err != nil {
		return nil, err
	}
	return individualEarning, nil
}

func MappingIndividualEarningDetail(m *models.IndividualEarningDetail) (*individual_earning_detail.IndividualEarningDetail, error) {
	if m == nil {
		return nil, nil
	}
	individualEarningDetail, err := individual_earning_detail.NewIndividualEarningDetail(m.ID, m.IndividualEarningID, m.Nominal, m.Amount)
	if err != nil {
		return nil, err
	}
	return individualEarningDetail, nil
}

func MappingFixedEarning(m *models.FixedEarning) (*fixed_earning.FixedEarning, error) {
	if m == nil {
		return nil, nil
	}
	var fixedEarningDetailsList []fixed_earning_detail.FixedEarningDetail
	for _, fixedEarningDetail := range m.R.FixedEarningDetails {
		mappingFixedEarningDetail, err := MappingFixedEarningDetail(fixedEarningDetail)
		if err != nil {
			return nil, err
		}
		fixedEarningDetailsList = append(fixedEarningDetailsList, *mappingFixedEarningDetail)
	}
	fixedEarning, err := fixed_earning.NewFixedEarning(m.ID, m.Amount, m.Nominal, fixedEarningDetailsList)
	if err != nil {
		return nil, err
	}
	return fixedEarning, nil
}

func MappingFixedEarningDetail(m *models.FixedEarningDetail) (*fixed_earning_detail.FixedEarningDetail, error) {
	if m == nil {
		return nil, nil
	}
	fixedEarningDetail, err := fixed_earning_detail.NewFixedEarningDetail(m.ID, m.FixedEarningID, m.Nominal, m.Amount)
	if err != nil {
		return nil, err
	}
	return fixedEarningDetail, nil
}

func MappingIndividualDeduction(m *models.IndividualDeduction) (*individual_deduction.IndividualDeduction, error) {
	if m == nil {
		return nil, nil
	}
	var individualDeductionDetailsList []individual_deduction_detail.IndividualDeductionDetail
	for _, indiindividualDeductionDetail := range m.R.IndividualDeductionDetails {
		mappingIndiindividualDeductionDetail, err := MappingIndividualDeductionDetail(indiindividualDeductionDetail)
		if err != nil {
			return nil, err
		}
		individualDeductionDetailsList = append(individualDeductionDetailsList, *mappingIndiindividualDeductionDetail)
	}
	individualDeduction, err := individual_deduction.NewIndividualDeduction(m.ID, m.Amount, m.Nominal, individualDeductionDetailsList)
	if err != nil {
		return nil, err
	}
	return individualDeduction, nil
}

func MappingIndividualDeductionDetail(m *models.IndividualDeductionDetail) (*individual_deduction_detail.IndividualDeductionDetail, error) {
	if m == nil {
		return nil, nil
	}
	individualDeductionDetail, err := individual_deduction_detail.NewIndividualDeductionDetail(m.ID, m.IndividualDeductionID, m.Nominal, m.Amount)
	if err != nil {
		return nil, err
	}
	return individualDeductionDetail, nil
}

func MappingFixedDeduction(m *models.FixedDeduction) (*fixed_deduction.FixedDeduction, error) {
	if m == nil {
		return nil, nil
	}
	var fixedDeductionDetailsList []fixed_deduction_detail.FixedDeductionDetail
	for _, fixedDeductionDetail := range m.R.FixedDeductionDetails {
		mappingFixedDeductionDetail, err := MappingFixedDeductionDetail(fixedDeductionDetail)
		if err != nil {
			return nil, err
		}
		fixedDeductionDetailsList = append(fixedDeductionDetailsList, *mappingFixedDeductionDetail)
	}
	fixedDeduction, err := fixed_deduction.NewFixedDeduction(m.ID, m.Amount, m.Nominal, fixedDeductionDetailsList)
	if err != nil {
		return nil, err
	}
	return fixedDeduction, nil
}

func MappingFixedDeductionDetail(m *models.FixedDeductionDetail) (*fixed_deduction_detail.FixedDeductionDetail, error) {
	if m == nil {
		return nil, nil
	}
	fixedDeductionDetail, err := fixed_deduction_detail.NewFixedDeductionDetail(m.ID, m.FixedDeductionID, m.Nominal, m.Amount)
	if err != nil {
		return nil, err
	}
	return fixedDeductionDetail, nil
}