package salary_statement_test

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"github.com/jokertennis/Payroll-Software/db"
	"github.com/jokertennis/Payroll-Software/server/gen/models"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func createDataForTestOfGetAllSalaryStatementsForEmployeeHandler(ctx context.Context, db boil.ContextExecutor) error {
	companies := []models.Company{
		{ID: 1, Name: "株式会社川田"},
		{ID: 2, Name: "株式会社kawata"},
		{ID: 3, Name: "株式会社追加"}}

	for _, company := range companies {
		if err := company.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of company. err:%s", err)
		}
	}

	employees := []models.Employee{
		{ID: 1, CompanyID: 1, Name: "Potter", MailAddress: "potter@example.com", Password: "testpass"},
		{ID: 2, CompanyID: 1, Name: "Joves", MailAddress: "joves@example.com", Password: "testpass"},
		{ID: 3, CompanyID: 1, Name: "Bob", MailAddress: "bob@example.com", Password: "testpass"},
		{ID: 4, CompanyID: 1, Name: "Yumi", MailAddress: "yumi@example.com", Password: "testpass"}}

	for _, employee := range employees {
		if err := employee.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of employee. err:%s", err)
		}
	}

	administrators := []models.Administrator{
		{ID: 1, CompanyID: 1, Name: "TestForAdministrator", MailAddress: "test.administrator@example.com", Password: "testpass"}}

	for _, administrator := range administrators {
		if err := administrator.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of administrator. err:%s", err)
		}
	}

	// individual_earnings := []models.IndividualEarning{
	// 	{ID: 1, Nominal: "支給", Amount: 300000},
	// 	{ID: 2, Nominal: "支給", Amount: 350000}}
	
	// for _, individual_earning := range individual_earnings {
	// 	if err := individual_earning.Insert(ctx, db, boil.Infer()); err != nil {
	// 		return fmt.Errorf("failed to create datas of individual_earning. err:%s", err)
	// 	}
	// }

	// individual_earning_details := []models.IndividualEarningDetail{
	// 	{ID: 1, IndividualEarningID: 1, Nominal: "基本給", Amount: 250000},
	// 	{ID: 2, IndividualEarningID: 1, Nominal: "住宅手当", Amount: 20000},
	// 	{ID: 3, IndividualEarningID: 1, Nominal: "通勤手当/非課税", Amount: 30000},
	// 	{ID: 4, IndividualEarningID: 2, Nominal: "基本給", Amount: 300000},
	// 	{ID: 5, IndividualEarningID: 2, Nominal: "残業手当", Amount: 50000}}

	// for _, individual_earning_detail := range individual_earning_details {
	// 	if err := individual_earning_detail.Insert(ctx, db, boil.Infer()); err != nil {
	// 		return fmt.Errorf("failed to create datas of individual_earning_detail. err:%s", err)
	// 	}
	// }

	// fixed_earnings := []models.FixedEarning{
	// 	{ID: 1, Nominal: "エンジニアLv1家族手当0人", Amount: 300000},
	// 	{ID: 2, Nominal: "エンジニアLv1家族手当1人", Amount: 350000}}

	// for _, fixed_earning := range fixed_earnings {
	// 	if err := fixed_earning.Insert(ctx, db, boil.Infer()); err != nil {
	// 		return fmt.Errorf("failed to create datas of fixed_earning. err:%s", err)
	// 	}
	// }

	// fixed_earning_details := []models.FixedEarningDetail{
	// 	{ID: 1, FixedEarningID: 1, Nominal: "基本給", Amount: 250000},
	// 	{ID: 2, FixedEarningID: 1, Nominal: "固定残業代", Amount: 50000},
	// 	{ID: 3, FixedEarningID: 2, Nominal: "基本給", Amount: 250000},
	// 	{ID: 4, FixedEarningID: 2, Nominal: "固定残業代", Amount: 50000},
	// 	{ID: 5, FixedEarningID: 2, Nominal: "家族手当(1人分)", Amount: 50000}}

	// for _, fixed_earning_detail := range fixed_earning_details {
	// 	if err := fixed_earning_detail.Insert(ctx, db, boil.Infer()); err != nil {
	// 		return fmt.Errorf("failed to create datas of fixed_earning_detail. err:%s", err)
	// 	}
	// }

	// individual_deductions := []models.IndividualDeduction{
	// 	{ID: 1, Amount: 60000, Nominal: "控除"},
	// 	{ID: 2, Amount: 70000, Nominal: "控除"}}
	
	// for _, individual_deduction := range individual_deductions {
	// 	if err := individual_deduction.Insert(ctx, db, boil.Infer()); err != nil {
	// 		return fmt.Errorf("failed to create datas of individual_deduction. err:%s", err)
	// 	}
	// }

	// individual_deduction_details := []models.IndividualDeductionDetail{
	// 	{ID: 1, IndividualDeductionID: 1, Nominal: "健康保険料", Amount: 12000},
	// 	{ID: 2, IndividualDeductionID: 1, Nominal: "厚生年金保険料", Amount: 20000},
	// 	{ID: 3, IndividualDeductionID: 1, Nominal: "雇用保険料", Amount: 300},
	// 	{ID: 4, IndividualDeductionID: 1, Nominal: "所得税", Amount: 5000},
	// 	{ID: 5, IndividualDeductionID: 1, Nominal: "住民税", Amount: 22700},
	// 	{ID: 6, IndividualDeductionID: 2, Nominal: "健康保険料", Amount: 13000},
	// 	{ID: 7, IndividualDeductionID: 2, Nominal: "厚生年金保険料", Amount: 21000},
	// 	{ID: 8, IndividualDeductionID: 2, Nominal: "雇用保険料", Amount: 500},
	// 	{ID: 9, IndividualDeductionID: 2, Nominal: "所得税", Amount: 6000},
	// 	{ID: 10, IndividualDeductionID: 2, Nominal: "住民税", Amount: 29500}}

	// for _, individual_deduction_detail := range individual_deduction_details {
	// 	if err := individual_deduction_detail.Insert(ctx, db, boil.Infer()); err != nil {
	// 		return fmt.Errorf("failed to create datas of individual_deduction_detail. err:%s", err)
	// 	}
	// }

	// fixed_deductions := []models.FixedDeduction{
	// 	{ID: 1, Nominal: "エンジニアLv1家族手当0人控除", Amount: 60000},
	// 	{ID: 2, Nominal: "エンジニアLv1家族手当1人控除", Amount: 70000}}

	// for _, fixed_deduction := range fixed_deductions {
	// 	if err := fixed_deduction.Insert(ctx, db, boil.Infer()); err != nil {
	// 		return fmt.Errorf("failed to create datas of fixed_deduction. err:%s", err)
	// 	}
	// }

	// fixed_deduction_details := []models.FixedDeductionDetail{
	// 	{ID: 1, FixedDeductionID: 1, Nominal: "健康保険料", Amount: 12000},
	// 	{ID: 2, FixedDeductionID: 1, Nominal: "厚生年金保険料", Amount: 20000},
	// 	{ID: 3, FixedDeductionID: 1, Nominal: "雇用保険料", Amount: 300},
	// 	{ID: 4, FixedDeductionID: 1, Nominal: "所得税", Amount: 5000},
	// 	{ID: 5, FixedDeductionID: 1, Nominal: "住民税", Amount: 22700},
	// 	{ID: 6, FixedDeductionID: 2, Nominal: "健康保険料", Amount: 13000},
	// 	{ID: 7, FixedDeductionID: 2, Nominal: "厚生年金保険料", Amount: 21000},
	// 	{ID: 8, FixedDeductionID: 2, Nominal: "雇用保険料", Amount: 500},
	// 	{ID: 9, FixedDeductionID: 2, Nominal: "所得税", Amount: 6000},
	// 	{ID: 10, FixedDeductionID: 2, Nominal: "住民税", Amount: 29500}}

	// for _, fixed_deduction_detail := range fixed_deduction_details {
	// 	if err := fixed_deduction_detail.Insert(ctx, db, boil.Infer()); err != nil {
	// 		return fmt.Errorf("failed to create datas of fixed_deduction_detail. err:%s", err)
	// 	}
	// }

	// salary_statements := []models.SalaryStatement{
	// 	{ID: 1, FixedEarningID: null.Uint32{Uint32: 1, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 1, Valid: true}, EmployeeID: 1, Nominal: "2022年1月 給与明細", Payday: time.Date(2022, 1, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
	// 	{ID: 2, FixedEarningID: null.Uint32{Uint32: 1, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 1, Valid: true}, EmployeeID: 1, Nominal: "2022年2月 給与明細", Payday: time.Date(2022, 2, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年1月1日~2022年1月31日"},
	// 	{ID: 3, FixedEarningID: null.Uint32{Uint32: 2, Valid: true}, FixedDeductionID: null.Uint32{Uint32: 2, Valid: true}, EmployeeID: 2, Nominal: "2022年2月 給与明細", Payday: time.Date(2022, 2, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年1月1日~2022年1月31日"},
	// 	{ID: 4, EmployeeID: 3, Nominal: "2010年2月 給与明細", Payday: time.Date(2010, 2, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2010年1月1日~2010年1月31日"},
	// 	{ID: 5, IndividualEarningID: null.Uint32{Uint32: 1, Valid: true}, IndividualDeductionID: null.Uint32{Uint32: 1, Valid: true}, EmployeeID: 2, Nominal: "2022年3月 給与明細", Payday: time.Date(2022, 3, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年2月1日~2022年2月28日"},
	// }

	// for _, salary_statement := range salary_statements {
	// 	if err := salary_statement.Insert(ctx, db, boil.Infer()); err != nil {
	// 		return fmt.Errorf("failed to create datas of salary_statement. err:%s", err)
	// 	}
	// }

	return nil
}

func TestGetAllSalaryStatementsForEmployeeHandler(t *testing.T) {
	cases := map[string]struct {
		email              string
		password           string
		expectedStatusCode int
		expectedResponse   string
	}{
		"registered employee successfully get multiple salary statements.": {
			email:              "joves@example.com",
			password:           "testpass",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "{\"name_of_employee\":\"Joves\",\"salary_statements\":[{\"amount_of_deduction\":70000,\"amount_of_earning\":350000,\"deduction_details\":[{\"amount_of_deduction_detail\":13000,\"nominal\":\"健康保険料\"},{\"amount_of_deduction_detail\":21000,\"nominal\":\"厚生年金保険料\"},{\"amount_of_deduction_detail\":500,\"nominal\":\"雇用保険料\"},{\"amount_of_deduction_detail\":6000,\"nominal\":\"所得税\"},{\"amount_of_deduction_detail\":29500,\"nominal\":\"住民税\"}],\"earning_details\":[{\"amount_of_earning_detail\":250000,\"nominal\":\"基本給\"},{\"amount_of_earning_detail\":50000,\"nominal\":\"固定残業代\"},{\"amount_of_earning_detail\":50000,\"nominal\":\"家族手当(1人分)\"}],\"nominal\":\"2022年2月 給与明細\",\"payday\":\"2022-02-19T00:00:00.000Z\",\"target_period\":\"2022年1月1日~2022年1月31日\"},{\"amount_of_deduction\":60000,\"amount_of_earning\":300000,\"deduction_details\":[{\"amount_of_deduction_detail\":12000,\"nominal\":\"健康保険料\"},{\"amount_of_deduction_detail\":20000,\"nominal\":\"厚生年金保険料\"},{\"amount_of_deduction_detail\":300,\"nominal\":\"雇用保険料\"},{\"amount_of_deduction_detail\":5000,\"nominal\":\"所得税\"},{\"amount_of_deduction_detail\":22700,\"nominal\":\"住民税\"}],\"earning_details\":[{\"amount_of_earning_detail\":250000,\"nominal\":\"基本給\"},{\"amount_of_earning_detail\":20000,\"nominal\":\"住宅手当\"},{\"amount_of_earning_detail\":30000,\"nominal\":\"通勤手当/非課税\"}],\"nominal\":\"2022年3月 給与明細\",\"payday\":\"2022-03-19T00:00:00.000Z\",\"target_period\":\"2022年2月1日~2022年2月28日\"}]}\n",
		},
		"registered employee that don't have any salary statement.": {
			email:              "yumi@example.com",
			password:           "testpass",
		 expectedStatusCode:    http.StatusOK,
		 expectedResponse:      "{\"name_of_employee\":\"Yumi\",\"salary_statements\":null}\n",
		},
		"internal server error has occurred.": {
			email:              "bob@example.com",
			password:           "testpass",
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponse:   "{\"message\":\"internalServerError:error:failed to get all salary statements.error:neither IndividualDeduction nor FixedDeduction was not found.SalaryStatementId:4\"}\n",
		},
		"execute by registered administrator": {
			email:              "test.administrator@example.com",
			password:           "testpass",
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   "{\"message\":\"Unauthorized. Specified mailAddress is not found in registered user datas.\"}\n",
		},
		"Authentication header is not present": {
			email:              "",
			password:           "",
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   "{\"message\":\"Unauthorized. Please check to see if Authentication header is not present or invalid.\"}\n",
		},
		"specify not incorrect password": {
			email:              "potter@example.com",
			password:           "incorrectpassword",
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   "{\"message\":\"Unauthorized. Please check to see if password is correct.\"}\n",
		},
	}

	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
	dbInstance, _ := db.CreateDbInstance(dbEnvironment)

	// executed reset of test db.
	_ = db.ResetDb(ctx, dbInstance)

	// executed creating test data.
	err := createDataForTestOfGetAllSalaryStatementsForEmployeeHandler(ctx, dbInstance)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	for _, value := range cases {
		
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/employee/salary_statements", nil)
		assert.Nil(t, err)

		if value.email == "" && value.password == "" {
			req.Header = http.Header{}
		} else {
			req.SetBasicAuth(value.email, value.password)
		}
		res, err := client.Do(req)
		assert.Nil(t, err)
		defer res.Body.Close()
		assert.Equal(t, value.expectedStatusCode, res.StatusCode)
		data, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.Equal(t, value.expectedResponse, string(data))
	}
}
