package infrastructure

import (
	"usr/local/go/server/gen/models"
	domainmodel "usr/local/go/src/main/domain-model"
)

func MappingEmployeeDomainObject(m *models.Employee) (*domainmodel.Employee, error) {
	employee, err := domainmodel.NewEmployee(m.ID, m.CompanyID, m.Name, m.MailAddress, m.Password)
	if err != nil {
        return nil, err
    }
    return employee, nil
}

func MappingAdministratorDomainObject(m *models.Administrator) (*domainmodel.Administrator, error) {
	administrator, err := domainmodel.NewAdministrator(m.ID, m.CompanyID, m.Name, m.MailAddress, m.Password)
	if err != nil {
        return nil, err
    }
    return administrator, nil
}

func MappingSalaryStatementDomainObject(m *models.SalaryStatement) (*domainmodel.SalaryStatement, error) {
	mappingIndividualEarning, err := MappingIndividualEarning(m.R.IndividualEarning)
	if err != nil {
		return nil, err
	}
	mappingFixedEarning, err := MappingFixedEarning(m.R.FixedEarning)
	if err != nil {
		return nil, err
	}
	mappingIndividualDeduction, err := MappingIndividualDeduction(m.R.IndividualDeduction)
	if err != nil {
		return nil, err
	}
	mappingFixedDeduction, err := MappingFixedDeduction(m.R.FixedDeduction)
	if err != nil {
		return nil, err
	}

	salaryStatement, err := domainmodel.NewSalaryStatement(m.ID, *mappingIndividualEarning, *mappingFixedEarning, *mappingIndividualDeduction, *mappingFixedDeduction, m.EmployeeID, m.Nominal, m.Payday, m.TargetPeriod)
	if err != nil {
        return nil, err
    }
    return salaryStatement, nil
}

func MappingIndividualEarning(m *models.IndividualEarning) (*domainmodel.IndividualEarning, error) {
	var individualEarningDetailsList []domainmodel.IndividualEarningDetail
	for _, indiindividualEarningDetail := range m.R.IndividualEarningDetails {
		mappingIndiindividualEarningDetail, err := MappingIndividualEarningDetail(indiindividualEarningDetail)
		if err != nil {
			return nil, err
		}
		individualEarningDetailsList = append(individualEarningDetailsList, *mappingIndiindividualEarningDetail)
	}
	individualEarning, err := domainmodel.NewIndividualEarning(m.ID, m.Amount, m.Nominal, individualEarningDetailsList)
	if err != nil {
		return nil, err
	}
	return individualEarning, nil
}

func MappingIndividualEarningDetail(m *models.IndividualEarningDetail) (*domainmodel.IndividualEarningDetail, error) {
	individualEarningDetail, err := domainmodel.NewIndividualEarningDetail(m.ID, m.IndividualEarningID, m.Nominal, m.Amount)
	if err != nil {
		return nil, err
	}
	return individualEarningDetail, nil
}

func MappingFixedEarning(m *models.FixedEarning) (*domainmodel.FixedEarning, error) {
	var fixedEarningDetailsList []domainmodel.FixedEarningDetail
	for _, fixedEarningDetail := range m.R.FixedEarningDetails {
		mappingFixedEarningDetail, err := MappingFixedEarningDetail(fixedEarningDetail)
		if err != nil {
			return nil, err
		}
		fixedEarningDetailsList = append(fixedEarningDetailsList, *mappingFixedEarningDetail)
	}
	fixedEarning, err := domainmodel.NewFixedEarning(m.ID, m.Amount, m.Nominal, fixedEarningDetailsList)
	if err != nil {
		return nil, err
	}
	return fixedEarning, nil
}

func MappingFixedEarningDetail(m *models.FixedEarningDetail) (*domainmodel.FixedEarningDetail, error) {
	fixedEarningDetail, err := domainmodel.NewFixedEarningDetail(m.ID, m.FixedEarningID, m.Nominal, m.Amount)
	if err != nil {
		return nil, err
	}
	return fixedEarningDetail, nil
}

func MappingIndividualDeduction(m *models.IndividualDeduction) (*domainmodel.IndividualDeduction, error) {
	var individualDeductionDetailsList []domainmodel.IndividualDeductionDetail
	for _, indiindividualDeductionDetail := range m.R.IndividualDeductionDetails {
		mappingIndiindividualDeductionDetail, err := MappingIndividualDeductionDetail(indiindividualDeductionDetail)
		if err != nil {
			return nil, err
		}
		individualDeductionDetailsList = append(individualDeductionDetailsList, *mappingIndiindividualDeductionDetail)
	}
	individualDeduction, err := domainmodel.NewIndividualDeduction(m.ID, m.Amount, m.Nominal, individualDeductionDetailsList)
	if err != nil {
		return nil, err
	}
	return individualDeduction, nil
}

func MappingIndividualDeductionDetail(m *models.IndividualDeductionDetail) (*domainmodel.IndividualDeductionDetail, error) {
	individualDeductionDetail, err := domainmodel.NewIndividualDeductionDetail(m.ID, m.IndividualDeductionID, m.Nominal, m.Amount)
	if err != nil {
		return nil, err
	}
	return individualDeductionDetail, nil
}

func MappingFixedDeduction(m *models.FixedDeduction) (*domainmodel.FixedDeduction, error) {
	var fixedDeductionDetailsList []domainmodel.FixedDeductionDetail
	for _, fixedDeductionDetail := range m.R.FixedDeductionDetails {
		mappingFixedDeductionDetail, err := MappingFixedDeductionDetail(fixedDeductionDetail)
		if err != nil {
			return nil, err
		}
		fixedDeductionDetailsList = append(fixedDeductionDetailsList, *mappingFixedDeductionDetail)
	}
	fixedDeduction, err := domainmodel.NewFixedDeduction(m.ID, m.Amount, m.Nominal, fixedDeductionDetailsList)
	if err != nil {
		return nil, err
	}
	return fixedDeduction, nil
}

func MappingFixedDeductionDetail(m *models.FixedDeductionDetail) (*domainmodel.FixedDeductionDetail, error) {
	fixedDeductionDetail, err := domainmodel.NewFixedDeductionDetail(m.ID, m.FixedDeductionID, m.Nominal, m.Amount)
	if err != nil {
		return nil, err
	}
	return fixedDeductionDetail, nil
}