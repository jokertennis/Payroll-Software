package salary_statement_application_service_test

import (
	"fmt"
	"net/http"
	"testing"
	"time"
	"github.com/jokertennis/Payroll-Software/src/main/application-service/salary_statement_application_service"
	deduction_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/deduction"
	deduction_detail_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/deduction_detail"
	earning_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/earning"
	earning_detail_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/earning_detail"
	employee_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/employee"
	salary_statement_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/salary_statement"
	"github.com/jokertennis/Payroll-Software/src/testtool"

	"github.com/go-openapi/strfmt"
	"github.com/stretchr/testify/assert"
)

func TestGetAllSalaryStatementsForEmployeeUseCase(t *testing.T) {
	type fakesFunctions struct {
		FakeGetEmployeeByMailAddress func(mailAddress string) (*employee_domain_model.Employee, error)
		FakeGetAllSalaryStatements   func(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error)
	}
	cases := map[string]struct {
		fakesFunctions     fakesFunctions
		expectedResult     *salary_statement_application_service.ResultOfGetAllSalaryStatementsForEmployee
		expectedStatusCode int
		expectedError      error
	}{
		"Successfully get Result data which has multiple salary statements.Error has not occurred.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := employee_domain_model.Employee{ID: 1, Name: "従業員A"}
					return &employee, nil
				},
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error) {
					var salaryStatementsList []*salary_statement_domain_model.SalaryStatement
					salaryStatement1 := salary_statement_domain_model.SalaryStatement{
						ID: 1,
						Earning: earning_domain_model.Earning{
							ID:      1,
							Amount:  300000,
							Nominal: "スタッフ支給総額",
							EarningType: earning_domain_model.IndividualEarning,
							EarningDetails: []earning_detail_domain_model.EarningDetail{
								{ID: 1, EarningID: 1, Nominal: "スタッフ基本給", Amount: 250000},
								{ID: 2, EarningID: 1, Nominal: "スタッフ固定残業代", Amount: 50000},
							},
						},
						Deduction: deduction_domain_model.Deduction{
							ID:      1,
							Amount:  15000,
							Nominal: "スタッフ控除総額",
							DeductionType: deduction_domain_model.IndividualDeduction,
							DeductionDetails: []deduction_detail_domain_model.DeductionDetail{
								{ID: 1, DeductionID: 1, Nominal: "所得税", Amount: 5000},
								{ID: 2, DeductionID: 1, Nominal: "住民税", Amount: 10000},
							},
						},
						EmployeeId:     1,
						Nominal:        "2022年2月分給料明細",
						Payday:         time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC),
						TargetPeriod:   "2022年1月1日~2022年1月31日分",
					}

					salaryStatement2 := salary_statement_domain_model.SalaryStatement{
						ID:                2,
						Earning: earning_domain_model.Earning{
							ID:      2,
							Amount:  300000,
							Nominal: "支給総額",
							EarningType: earning_domain_model.FixedEarning,
							EarningDetails: []earning_detail_domain_model.EarningDetail{
								{ID: 3, EarningID: 1, Nominal: "基本給", Amount: 250000},
								{ID: 4, EarningID: 1, Nominal: "固定残業代", Amount: 50000},
							},
						},
						Deduction: deduction_domain_model.Deduction{
							ID:      2,
							Amount:  15000,
							Nominal: "控除総額",
							DeductionType: deduction_domain_model.FixedDeduction,
							DeductionDetails: []deduction_detail_domain_model.DeductionDetail{
								{ID: 3, DeductionID: 1, Nominal: "所得税", Amount: 5000},
								{ID: 4, DeductionID: 1, Nominal: "住民税", Amount: 10000},
							},
						},
						EmployeeId:   1,
						Nominal:      "2022年3月分給料明細",
						Payday:       time.Date(2022, time.March, 25, 12, 00, 00, 0, time.UTC),
						TargetPeriod: "2022年2月1日~2022年2月28日分",
					}

					salaryStatementsList = append(salaryStatementsList, &salaryStatement1)
					salaryStatementsList = append(salaryStatementsList, &salaryStatement2)
					return salaryStatementsList, nil
				},
			},
			expectedResult: &salary_statement_application_service.ResultOfGetAllSalaryStatementsForEmployee{
				NameOfEmployee: "従業員A",
				SalaryStatementForEmployeeList: []*salary_statement_application_service.SalaryStatementForEmployee{
					{
						Nominal:           "2022年2月分給料明細",
						Payday:            strfmt.DateTime(time.Date(2022, time.February, 25, 12, 00, 00, 0, time.UTC)),
						TargetPeriod:      "2022年1月1日~2022年1月31日分",
						AmountOfEarning:   300000,
						AmountOfDeduction: 15000,
						EarningDetails: []salary_statement_application_service.EarningDetailOfGetAllSalaryStatementsForEmployee{
							{Nominal: "スタッフ基本給", AmountOfEarningDetail: 250000},
							{Nominal: "スタッフ固定残業代", AmountOfEarningDetail: 50000},
						},
						DeductionDetails: []salary_statement_application_service.DeductionDetailOfGetAllSalaryStatementsForEmployee{
							{Nominal: "所得税", AmountOfDeductionDetail: 5000},
							{Nominal: "住民税", AmountOfDeductionDetail: 10000},
						},
					}, {
						Nominal:           "2022年3月分給料明細",
						Payday:            strfmt.DateTime(time.Date(2022, time.March, 25, 12, 00, 00, 0, time.UTC)),
						TargetPeriod:      "2022年2月1日~2022年2月28日分",
						AmountOfEarning:   300000,
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
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					return nil, fmt.Errorf("failed to connect db.")
				},
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error) {
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
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
			},
			expectedResult:     nil,
			expectedStatusCode: http.StatusUnauthorized,
			expectedError:      fmt.Errorf("unauthorized. Specified mailAddress is not found in registered user datas. MailAddress:employee@example.com"),
		},
		"Error has occurred when get all salary statements.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := employee_domain_model.Employee{ID: 1, Name: "keven"}
					return &employee, nil
				},
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error) {
					return nil, fmt.Errorf("failed to connect db.")
				},
			},
			expectedResult:     nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedError:      fmt.Errorf("internalServerError:error:failed to get all salary statements.error:failed to connect db."),
		},
		"Salary Statement is not found when get all salary statements.Error has not occurred.": {
			fakesFunctions: fakesFunctions{
				FakeGetEmployeeByMailAddress: func(mailAddress string) (*employee_domain_model.Employee, error) {
					employee := employee_domain_model.Employee{ID: 1, Name: "keven"}
					return &employee, nil
				},
				FakeGetAllSalaryStatements: func(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error) {
					return nil, nil
				},
			},
			expectedResult:     &salary_statement_application_service.ResultOfGetAllSalaryStatementsForEmployee{NameOfEmployee: "keven", SalaryStatementForEmployeeList: nil},
			expectedStatusCode: http.StatusOK,
			expectedError:      nil,
		},
	}

	for _, value := range cases {
		employeeRepository := testtool.EmployeeRepositoryMock{}
		employeeRepository.FakeGetEmployeeByMailAddress = value.fakesFunctions.FakeGetEmployeeByMailAddress
		salaryStatementRepository := testtool.SalaryStatementRepositoryMock{}
		salaryStatementRepository.FakeGetAllSalaryStatements = value.fakesFunctions.FakeGetAllSalaryStatements
		result, statusCode, err := salary_statement_application_service.GetAllSalaryStatementsForEmployeeUseCase(&employeeRepository, &salaryStatementRepository, "employee@example.com")
		assert.Equal(t, value.expectedResult, result)
		assert.Equal(t, value.expectedStatusCode, statusCode)
		assert.Equal(t, value.expectedError, err)
	}
}
