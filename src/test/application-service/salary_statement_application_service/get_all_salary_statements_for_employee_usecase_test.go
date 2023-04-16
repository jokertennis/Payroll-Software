package salary_statement_application_service

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
	"usr/local/go/src/test/testtool"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestGetAllSalaryStatementsForEmployeeUseCase(t *testing.T) {
	type fakesFunctions struct {
		FakeGetEmployeeByMailAddress   func(mailAddress string) (*employee.Employee, error)
		FakeGetAllSalaryStatements func(employeeId uint32) ([]*salary_statement.SalaryStatement, error)
	}
	cases := map[string]struct {
		fakesFunctions     fakesFunctions
		expectedResult     *salary_statement_application_service.ResultOfGetAllSalaryStatementsForEmployee
		expectedStatusCode int
		expectedError      error
	}{
		"Successfully get Result data which has multiple salary statements.Error has not occurred.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee.Employee, error) {
					employee := &employee.Employee{ID: 1, Name: "従業員A"}
					return employee, nil
				},
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement.SalaryStatement, error) {
					var salaryStatementsList []*salary_statement.SalaryStatement
					salaryStatement1 := &salary_statement.SalaryStatement{
						ID: 1,
						IndividualEarning: &individual_earning.IndividualEarning{
							ID:      1,
							Amount:  300000,
							Nominal: "スタッフ支給総額",
							IndividualEarningDetails: []individual_earning_detail.IndividualEarningDetail{
								{ID: 1, IndividualEarningID: 1, Nominal: "スタッフ基本給", Amount: 250000},
								{ID: 2, IndividualEarningID: 1, Nominal: "スタッフ固定残業代", Amount: 50000},
							},
						},
						FixedEarning: nil,
						IndividualDeduction: &individual_deduction.IndividualDeduction{
							ID:      1,
							Amount:  15000,
							Nominal: "スタッフ控除総額",
							IndividualDeductionDetails: []individual_deduction_detail.IndividualDeductionDetail{
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

					salaryStatement2 := &salary_statement.SalaryStatement{
						ID:                2,
						IndividualEarning: nil,
						FixedEarning: &fixed_earning.FixedEarning{
							ID:      1,
							Amount:  300000,
							Nominal: "支給総額",
							FixedEarningDetails: []fixed_earning_detail.FixedEarningDetail{
								{ID: 1, FixedEarningID: 1, Nominal: "基本給", Amount: 250000},
								{ID: 2, FixedEarningID: 1, Nominal: "固定残業代", Amount: 50000},
							},
						},
						IndividualDeduction: nil,
						FixedDeduction: &fixed_deduction.FixedDeduction{
							ID:      1,
							Amount:  15000,
							Nominal: "控除総額",
							FixedDeductionDetails: []fixed_deduction_detail.FixedDeductionDetail{
								{ID: 1, FixedDeductionID: 1, Nominal: "所得税", Amount: 5000},
								{ID: 2, FixedDeductionID: 1, Nominal: "住民税", Amount: 10000},
							},
						},
						EmployeeId:   1,
						Nominal:      "2022年3月分給料明細",
						Payday:       time.Date(2022, time.March, 25, 12, 00, 00, 0, time.UTC),
						TargetPeriod: "2022年2月1日~2022年2月28日分",
					}

					salaryStatementsList = append(salaryStatementsList, salaryStatement1)
					salaryStatementsList = append(salaryStatementsList, salaryStatement2)
					return salaryStatementsList, nil
				},
			},
			expectedResult: &salary_statement_application_service.ResultOfGetAllSalaryStatementsForEmployee{
				NameOfEmployee: "従業員A",
				SalaryStatementForEmployeeList: []*salary_statement_application_service.SalaryStatementForEmployee{
					{
						Nominal: "2022年2月分給料明細",
						Payday: strfmt.DateTime(time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC)),
						TargetPeriod: "2022年1月1日~2022年1月31日分",
						AmountOfEarning: 300000,
						AmountOfDeduction: 15000,
						EarningDetails: []salary_statement_application_service.EarningDetailOfGetAllSalaryStatementsForEmployee{
							{Nominal: "スタッフ基本給", AmountOfEarningDetail: 250000},
							{Nominal: "スタッフ固定残業代", AmountOfEarningDetail: 50000},
						},
						DeductionDetails: []salary_statement_application_service.DeductionDetailOfGetAllSalaryStatementsForEmployee{
							{Nominal: "所得税", AmountOfDeductionDetail: 5000},
							{Nominal: "住民税", AmountOfDeductionDetail: 10000},
						},
					},{
						Nominal: "2022年3月分給料明細",
						Payday: strfmt.DateTime(time.Date(2022, time.March, 25, 12, 00, 00, 0, time.UTC)),
						TargetPeriod: "2022年2月1日~2022年2月28日分",
						AmountOfEarning: 300000,
						AmountOfDeduction: 15000,
						EarningDetails: []salary_statement_application_service.EarningDetailOfGetAllSalaryStatementsForEmployee{
							{Nominal: "基本給", AmountOfEarningDetail: 250000},
							{Nominal: "固定残業代", AmountOfEarningDetail: 50000},
						},
						DeductionDetails: []salary_statement_application_service.DeductionDetailOfGetAllSalaryStatementsForEmployee{
							{Nominal: "所得税", AmountOfDeductionDetail: 5000},
							{Nominal: "住民税", AmountOfDeductionDetail: 10000},
						},
					},
				},
			},
			expectedStatusCode: http.StatusOK,
			expectedError:      nil,
		},
		"Error has occurred when get employee.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee.Employee, error) {
					return nil, fmt.Errorf("failed to connect db.")
				},
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement.SalaryStatement, error) {
					return nil, nil
				},
			},
			expectedResult:     nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedError:      fmt.Errorf("InternalServerError:error:failed to connect db."),
		},
		"Employee is nil when get employee.Error has not occurred.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee.Employee, error) {
					return nil, nil
				},
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement.SalaryStatement, error) {
					return nil, nil
				},
			},
			expectedResult:     nil,
			expectedStatusCode: http.StatusUnauthorized,
			expectedError:      fmt.Errorf("unauthorized. Specified mailAddress is not found in registered user datas. MailAddress:employee@example.com"),
		},
		"Error has occurred when get all salary statements.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee.Employee, error) {
					employee := &employee.Employee{ID: 1, Name: "keven"}
					return employee, nil
				},
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement.SalaryStatement, error) {
					return nil, fmt.Errorf("failed to connect db.")
				},
			},
			expectedResult:     nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedError:      fmt.Errorf("internalServerError:error:failed to get all salary statements.error:failed to connect db."),
		},
		"Salary Statement is not found when get all salary statements.Error has not occurred.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee.Employee, error) {
					employee := &employee.Employee{ID: 1, Name: "keven"}
					return employee, nil
				},
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement.SalaryStatement, error) {
					return nil, nil
				},
			},
			expectedResult:     nil,
			expectedStatusCode: http.StatusOK,
			expectedError:      nil,
		},
		"Error has occurred when get deduction.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee.Employee, error) {
					employee := &employee.Employee{ID: 1, Name: "従業員A"}
					return employee, nil
				},
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement.SalaryStatement, error) {
					var salaryStatementsList []*salary_statement.SalaryStatement
					salaryStatement1 := &salary_statement.SalaryStatement{
						ID: 1,
						IndividualEarning: &individual_earning.IndividualEarning{
							ID:      1,
							Amount:  300000,
							Nominal: "スタッフ支給総額",
							IndividualEarningDetails: []individual_earning_detail.IndividualEarningDetail{
								{ID: 1, IndividualEarningID: 1, Nominal: "スタッフ基本給", Amount: 250000},
								{ID: 2, IndividualEarningID: 1, Nominal: "スタッフ固定残業代", Amount: 50000},
							},
						},
						FixedEarning: nil,
						IndividualDeduction: nil,
						FixedDeduction: nil,
						EmployeeId:     1,
						Nominal:        "2022年2月分給料明細",
						Payday:         time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
						TargetPeriod:   "2022年1月1日~2022年1月31日分",
					}
					salaryStatementsList = append(salaryStatementsList, salaryStatement1)
					return salaryStatementsList, nil
				},
			},
			expectedResult: nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedError:      fmt.Errorf("InternalServerError:error:neither IndividualDeduction nor FixedDeduction was not found.SalaryStatementId:1"),
		},
		"Error has occurred when get earning.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee.Employee, error) {
					employee := &employee.Employee{ID: 1, Name: "従業員A"}
					return employee, nil
				},
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement.SalaryStatement, error) {
					var salaryStatementsList []*salary_statement.SalaryStatement
					salaryStatement1 := &salary_statement.SalaryStatement{
						ID: 1,
						IndividualEarning: nil,
						FixedEarning: nil,
						IndividualDeduction: &individual_deduction.IndividualDeduction{
							ID:      1,
							Amount:  15000,
							Nominal: "スタッフ控除総額",
							IndividualDeductionDetails: []individual_deduction_detail.IndividualDeductionDetail{
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
					salaryStatementsList = append(salaryStatementsList, salaryStatement1)
					return salaryStatementsList, nil
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
		salaryStatementRepository.FakeGetAllSalaryStatements = value.fakesFunctions.FakeGetAllSalaryStatements
		result, statusCode, err := salary_statement_application_service.GetAllSalaryStatementsForEmployeeUseCase(employeeRepository, salaryStatementRepository, "employee@example.com")
		assert.Equal(t, value.expectedResult, result)
		assert.Equal(t, value.expectedStatusCode, statusCode)
		assert.Equal(t, value.expectedError, err)
	}
}