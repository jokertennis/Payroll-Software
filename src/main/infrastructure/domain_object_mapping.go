package infrastructure

import (
	"usr/local/go/server/gen/models"
	administrator_domain_model "usr/local/go/src/main/domain-model/administrator"
	deduction_domain_model "usr/local/go/src/main/domain-model/deduction"
	deduction_detail_domain_model "usr/local/go/src/main/domain-model/deduction_detail"
	earning_domain_model "usr/local/go/src/main/domain-model/earning"
	earning_detail_domain_model "usr/local/go/src/main/domain-model/earning_detail"
	employee_domain_model "usr/local/go/src/main/domain-model/employee"
	salary_statement_domain_model "usr/local/go/src/main/domain-model/salary_statement"
)

func MappingEmployeeDomainObject(m *models.Employee) (*employee_domain_model.Employee, error) {
	if m == nil {
		return nil, nil
	}
	employee, err := employee_domain_model.NewEmployee(m.ID, m.CompanyID, m.Name, m.MailAddress, m.Password)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func MappingAdministratorDomainObject(m *models.Administrator) (*administrator_domain_model.Administrator, error) {
	if m == nil {
		return nil, nil
	}
	administrator, err := administrator_domain_model.NewAdministrator(m.ID, m.CompanyID, m.Name, m.MailAddress, m.Password)
	if err != nil {
		return nil, err
	}
	return administrator, nil
}

func MappingSalaryStatementDomainObject(m *models.SalaryStatement) (*salary_statement_domain_model.SalaryStatement, error) {
	if m == nil {
		return nil, nil
	}
	var err error
	var mappingIndividualEarning *earning_domain_model.IndividualEarning
	if m.IndividualEarningID.Valid {
		mappingIndividualEarning, err = MappingIndividualEarning(m.R.IndividualEarning)
		if err != nil {
			return nil, err
		}
	} else {
		mappingIndividualEarning = nil
	}

	var mappingFixedEarning *earning_domain_model.FixedEarning
	if m.FixedEarningID.Valid {
		mappingFixedEarning, err = MappingFixedEarning(m.R.FixedEarning)
		if err != nil {
			return nil, err
		}
	} else {
		mappingFixedEarning = nil
	}

	var mappingIndividualDeduction *deduction_domain_model.IndividualDeduction
	if m.IndividualDeductionID.Valid {
		mappingIndividualDeduction, err = MappingIndividualDeduction(m.R.IndividualDeduction)
		if err != nil {
			return nil, err
		}
	} else {
		mappingIndividualDeduction = nil
	}

	var mappingFixedDeduction *deduction_domain_model.FixedDeduction
	if m.FixedDeductionID.Valid {
		mappingFixedDeduction, err = MappingFixedDeduction(m.R.FixedDeduction)
		if err != nil {
			return nil, err
		}
	} else {
		mappingFixedDeduction = nil
	}

	salaryStatement, err := salary_statement_domain_model.NewSalaryStatement(m.ID, mappingEarning, mappingDeduction, m.EmployeeID, m.Nominal, m.Payday, m.TargetPeriod)
	if err != nil {
		return nil, err
	}
	return salaryStatement, nil
}

func MappingEarning(m *models.IndividualEarning) (*earning_domain_model.Earning, error) {
	if m == nil {
		return nil, nil
	}
	var individualEarningDetailsList []earning_detail_domain_model.IndividualEarningDetail
	for _, indiindividualEarningDetail := range m.R.IndividualEarningDetails {
		mappingIndiindividualEarningDetail, err := MappingIndividualEarningDetail(indiindividualEarningDetail)
		if err != nil {
			return nil, err
		}
		individualEarningDetailsList = append(individualEarningDetailsList, *mappingIndiindividualEarningDetail)
	}
	individualEarning, err := earning_domain_model.NewIndividualEarning(m.ID, m.Amount, m.Nominal, individualEarningDetailsList)
	if err != nil {
		return nil, err
	}
	return individualEarning, nil
}

func MappingEarningDetail(m *models.IndividualEarningDetail) (*earning_detail_domain_model.EarningDetail, error) {
	if m == nil {
		return nil, nil
	}
	individualEarningDetail, err := earning_detail_domain_model.NewIndividualEarningDetail(m.ID, m.IndividualEarningID, m.Nominal, m.Amount)
	if err != nil {
		return nil, err
	}
	return individualEarningDetail, nil
}

func MappingDeduction(m *models.IndividualDeduction) (*deduction_domain_model.IndividualDeduction, error) {
	if m == nil {
		return nil, nil
	}
	var individualDeductionDetailsList []deduction_detail_domain_model.IndividualDeductionDetail
	for _, indiindividualDeductionDetail := range m.R.IndividualDeductionDetails {
		mappingIndiindividualDeductionDetail, err := MappingIndividualDeductionDetail(indiindividualDeductionDetail)
		if err != nil {
			return nil, err
		}
		individualDeductionDetailsList = append(individualDeductionDetailsList, *mappingIndiindividualDeductionDetail)
	}
	individualDeduction, err := deduction_domain_model.NewIndividualDeduction(m.ID, m.Amount, m.Nominal, individualDeductionDetailsList)
	if err != nil {
		return nil, err
	}
	return individualDeduction, nil
}

func MappingDeductionDetail(m *models.IndividualDeductionDetail) (*deduction_detail_domain_model.IndividualDeductionDetail, error) {
	if m == nil {
		return nil, nil
	}
	individualDeductionDetail, err := deduction_detail_domain_model.NewIndividualDeductionDetail(m.ID, m.IndividualDeductionID, m.Nominal, m.Amount)
	if err != nil {
		return nil, err
	}
	return individualDeductionDetail, nil
}
