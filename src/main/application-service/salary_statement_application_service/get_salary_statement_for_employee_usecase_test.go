package salary_statement_application_service_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"
	"usr/local/go/src/main/application-service/salary_statement_application_service"
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
	"usr/local/go/src/testtool"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestGetSalaryStatementForEmployeeUseCase(t *testing.T) {
	type fakesFunctions struct {
		FakeGetEmployeeByMailAddress   func(mailAddress string) (*employee_domain_model.Employee, error)
		FakeGetSalaryStatement func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error)
	}
	cases := map[string]struct {
		fakesFunctions     fakesFunctions
		expectedResult     *salary_statement_application_service.ResultOfGetSalaryStatementForEmployee
		expectedStatusCode int
		expectedError      error
	}{
		"Successfully get Result data when salary statement has individual earning and individual deduction.Error has not occurred.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := &employee_domain_model.Employee{ID: 1, Name: "従業員A"}
					return employee, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					salaryStatement := &salary_statement_domain_model.SalaryStatement{
						ID: 1,
						IndividualEarning: &individual_earning_domain_model.IndividualEarning{
							ID:      1,
							Amount:  300000,
							Nominal: "スタッフ支給総額",
							IndividualEarningDetails: []individual_earning_detail_domain_model.IndividualEarningDetail{
								{ID: 1, IndividualEarningID: 1, Nominal: "スタッフ基本給", Amount: 250000},
								{ID: 2, IndividualEarningID: 1, Nominal: "スタッフ固定残業代", Amount: 50000},
							},
						},
						FixedEarning: nil,
						IndividualDeduction: &individual_deduction_domain_model.IndividualDeduction{
							ID:      1,
							Amount:  15000,
							Nominal: "スタッフ控除総額",
							IndividualDeductionDetails: []individual_deduction_detail_domain_model.IndividualDeductionDetail{
								{ID: 1, IndividualDeductionID: 1, Nominal: "所得税", Amount: 5000},
								{ID: 2, IndividualDeductionID: 1, Nominal: "住民税", Amount: 10000},
							},
						},
						FixedDeduction: nil,
						EmployeeId:     1,
						Nominal:        "2022年2月分給料明細",
						Payday:         time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
						TargetPeriod:   "2022年1月1日~2022年1月31日分",
					}
					return salaryStatement, nil
				},
			},
			expectedResult: &salary_statement_application_service.ResultOfGetSalaryStatementForEmployee{
				Nominal:           "2022年2月分給料明細",
				Payday:            strfmt.DateTime(time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC)),
				TargetPeriod:      "2022年1月1日~2022年1月31日分",
				AmountOfDeduction: 15000,
				NameOfEmployee:    "従業員A",
				AmountOfEarning:   300000,
				EarningDetails: []salary_statement_application_service.EarningDetailOfGetSalaryStatementForEmployee{
					{Nominal: "スタッフ基本給", AmountOfEarningDetail: 250000},
					{Nominal: "スタッフ固定残業代", AmountOfEarningDetail: 50000},
				},
				DeductionDetails: []salary_statement_application_service.DeductionDetailOfGetSalaryStatementForEmployee{
					{Nominal: "所得税", AmountOfDeductionDetail: 5000},
					{Nominal: "住民税", AmountOfDeductionDetail: 10000},
				},
			},
			expectedStatusCode: http.StatusOK,
			expectedError:      nil,
		},
		"Successfully get Result data when salary statement has fixed earning and fixed deduction.Error has not occurred.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := &employee_domain_model.Employee{ID: 1, Name: "keven"}
					return employee, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					salaryStatement := &salary_statement_domain_model.SalaryStatement{
						ID:                1,
						IndividualEarning: nil,
						FixedEarning: &fixed_earning_domain_model.FixedEarning{
							ID:      1,
							Amount:  300000,
							Nominal: "支給総額",
							FixedEarningDetails: []fixed_earning_detail_domain_model.FixedEarningDetail{
								{ID: 1, FixedEarningID: 1, Nominal: "基本給", Amount: 250000},
								{ID: 2, FixedEarningID: 1, Nominal: "固定残業代", Amount: 50000},
							},
						},
						IndividualDeduction: nil,
						FixedDeduction: &fixed_deduction_domain_model.FixedDeduction{
							ID:      1,
							Amount:  15000,
							Nominal: "控除総額",
							FixedDeductionDetails: []fixed_deduction_detail_domain_model.FixedDeductionDetail{
								{ID: 1, FixedDeductionID: 1, Nominal: "所得税", Amount: 5000},
								{ID: 2, FixedDeductionID: 1, Nominal: "住民税", Amount: 10000},
							},
						},
						EmployeeId:   1,
						Nominal:      "2022年2月分給料明細",
						Payday:       time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
						TargetPeriod: "2022年1月1日~2022年1月31日分",
					}
					return salaryStatement, nil
				},
			},
			expectedResult: &salary_statement_application_service.ResultOfGetSalaryStatementForEmployee{
				Nominal:           "2022年2月分給料明細",
				Payday:            strfmt.DateTime(time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC)),
				TargetPeriod:      "2022年1月1日~2022年1月31日分",
				AmountOfDeduction: 15000,
				NameOfEmployee:    "keven",
				AmountOfEarning:   300000,
				EarningDetails: []salary_statement_application_service.EarningDetailOfGetSalaryStatementForEmployee{
					{Nominal: "基本給", AmountOfEarningDetail: 250000},
					{Nominal: "固定残業代", AmountOfEarningDetail: 50000},
				},
				DeductionDetails: []salary_statement_application_service.DeductionDetailOfGetSalaryStatementForEmployee{
					{Nominal: "所得税", AmountOfDeductionDetail: 5000},
					{Nominal: "住民税", AmountOfDeductionDetail: 10000},
				},
			},
			expectedStatusCode: http.StatusOK,
			expectedError:      nil,
		},
		"Error has occurred when get employee.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					return nil, fmt.Errorf("failed to connect db.")
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
			},
			expectedResult:     nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedError:      fmt.Errorf("InternalServerError:error:failed to connect db."),
		},
		"Employee is nil when get employee.Error has not occurred.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					return nil, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
			},
			expectedResult:     nil,
			expectedStatusCode: http.StatusUnauthorized,
			expectedError:      fmt.Errorf("unauthorized. Specified mailAddress is not found in registered user datas. MailAddress:employee@example.com"),
		},
		"Error has occurred when get salary statement.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := &employee_domain_model.Employee{ID: 1, Name: "keven"}
					return employee, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, fmt.Errorf("failed to connect db.")
				},
			},
			expectedResult:     nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedError:      fmt.Errorf("internalServerError:error:failed to get salary statement.error:failed to connect db."),
		},
		"Salary Statement is nil when get salary statement.Error has not occurred.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := &employee_domain_model.Employee{ID: 1, Name: "keven"}
					return employee, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
			},
			expectedResult:     nil,
			expectedStatusCode: http.StatusNotFound,
			expectedError:      fmt.Errorf("notFound. SalaryStatement with specified year and month was not found in registered salary statement datas.UserMailAddress:employee@example.com, Year:2022, Month:2"),
		},
		"Error has occurred when get deduction.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := &employee_domain_model.Employee{ID: 1, Name: "keven"}
					return employee, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					salaryStatement := &salary_statement_domain_model.SalaryStatement{
						ID:                1,
						IndividualEarning: nil,
						FixedEarning: &fixed_earning_domain_model.FixedEarning{
							ID:      1,
							Amount:  300000,
							Nominal: "支給総額",
							FixedEarningDetails: []fixed_earning_detail_domain_model.FixedEarningDetail{
								{ID: 1, FixedEarningID: 1, Nominal: "基本給", Amount: 250000},
								{ID: 2, FixedEarningID: 1, Nominal: "固定残業代", Amount: 50000},
							},
						},
						IndividualDeduction: nil,
						FixedDeduction: nil,
						EmployeeId:   1,
						Nominal:      "2022年2月分給料明細",
						Payday:       time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
						TargetPeriod: "2022年1月1日~2022年1月31日分",
					}
					return salaryStatement, nil
				},
			},
			expectedResult: nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedError:      fmt.Errorf("InternalServerError:error:neither IndividualDeduction nor FixedDeduction was not found.SalaryStatementId:1"),
		},
		"Error has occurred when get earning.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := &employee_domain_model.Employee{ID: 1, Name: "keven"}
					return employee, nil
				},
				FakeGetSalaryStatement: func(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
					salaryStatement := &salary_statement_domain_model.SalaryStatement{
						ID:                1,
						IndividualEarning: nil,
						FixedEarning: nil,
						IndividualDeduction: nil,
						FixedDeduction: &fixed_deduction_domain_model.FixedDeduction{
							ID:      1,
							Amount:  15000,
							Nominal: "控除総額",
							FixedDeductionDetails: []fixed_deduction_detail_domain_model.FixedDeductionDetail{
								{ID: 1, FixedDeductionID: 1, Nominal: "所得税", Amount: 5000},
								{ID: 2, FixedDeductionID: 1, Nominal: "住民税", Amount: 10000},
							},
						},
						EmployeeId:   1,
						Nominal:      "2022年2月分給料明細",
						Payday:       time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
						TargetPeriod: "2022年1月1日~2022年1月31日分",
					}
					return salaryStatement, nil
				},
			},
			expectedResult: nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedError:      fmt.Errorf("InternalServerError:error:neither IndividualEarning nor FixedEarning was not found.SalaryStatementId:1"),
		},
	}

	for _, value := range cases {
		employeeRepository := &testtool.EmployeeRepositoryMock{}
		employeeRepository.FakeGetEmployeeByMailAddress = value.fakesFunctions.FakeGetEmployeeByMailAddress
		salaryStatementRepository := &testtool.SalaryStatementRepositoryMock{}
		salaryStatementRepository.FakeGetSalaryStatement = value.fakesFunctions.FakeGetSalaryStatement
		result, statusCode, err := salary_statement_application_service.GetSalaryStatementForEmployeeUseCase(employeeRepository, salaryStatementRepository, "employee@example.com", 2022, time.February)
		assert.Equal(t, value.expectedResult, result)
		assert.Equal(t, value.expectedStatusCode, statusCode)
		assert.Equal(t, value.expectedError, err)
	}
}