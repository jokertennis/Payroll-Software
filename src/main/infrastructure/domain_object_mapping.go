package infrastructure

import (
	"github.com/jokertennis/Payroll-Software/server/gen/models"
	administrator_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/administrator"
	deduction_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/deduction"
	deduction_detail_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/deduction_detail"
	earning_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/earning"
	earning_detail_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/earning_detail"
	employee_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/employee"
	salary_statement_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/salary_statement"
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
	var mappingEarning *earning_domain_model.Earning
	mappingEarning, err = MappingEarning(m.R.Earning)
	if err != nil {
		return nil, err
	}

	var mappingDeduction *deduction_domain_model.Deduction
	mappingDeduction, err = MappingDeduction(m.R.Deduction)
	if err != nil {
		return nil, err
	}

	salaryStatement, err := salary_statement_domain_model.NewSalaryStatement(m.ID, *mappingEarning, *mappingDeduction, m.EmployeeID, m.Nominal, m.Payday, m.TargetPeriod)
	if err != nil {
		return nil, err
	}
	return salaryStatement, nil
}

func MappingEarning(m *models.Earning) (*earning_domain_model.Earning, error) {
	if m == nil {
		return nil, nil
	}
	var earningDetailsList []earning_detail_domain_model.EarningDetail
	for _, earningDetail := range m.R.EarningDetails {
		mappingEarningDetail, err := MappingEarningDetail(earningDetail)
		if err != nil {
			return nil, err
		}
		earningDetailsList = append(earningDetailsList, *mappingEarningDetail)
	}
	earning, err := earning_domain_model.NewEarning(m.ID, m.Amount, m.Nominal, earning_domain_model.EarningType(m.EarningType), earningDetailsList)
	if err != nil {
		return nil, err
	}
	return &earning, nil
}

func MappingEarningDetail(m *models.EarningDetail) (*earning_detail_domain_model.EarningDetail, error) {
	if m == nil {
		return nil, nil
	}
	earningDetail, err := earning_detail_domain_model.NewEarningDetail(m.ID, m.EarningID, m.Nominal, m.Amount)
	if err != nil {
		return nil, err
	}
	return &earningDetail, nil
}

func MappingDeduction(m *models.Deduction) (*deduction_domain_model.Deduction, error) {
	if m == nil {
		return nil, nil
	}
	var deductionDetailsList []deduction_detail_domain_model.DeductionDetail
	for _, deductionDetail := range m.R.DeductionDetails {
		mappingDeductionDetail, err := MappingDeductionDetail(deductionDetail)
		if err != nil {
			return nil, err
		}
		deductionDetailsList = append(deductionDetailsList, *mappingDeductionDetail)
	}
	deduction, err := deduction_domain_model.NewDeduction(m.ID, m.Amount, m.Nominal, deduction_domain_model.DeductionType(m.DeductionType), deductionDetailsList)
	if err != nil {
		return nil, err
	}
	return &deduction, nil
}

func MappingDeductionDetail(m *models.DeductionDetail) (*deduction_detail_domain_model.DeductionDetail, error) {
	if m == nil {
		return nil, nil
	}
	deductionDetail, err := deduction_detail_domain_model.NewDeductionDetail(m.ID, m.DeductionID, m.Nominal, m.Amount)
	if err != nil {
		return nil, err
	}
	return &deductionDetail, nil
}
