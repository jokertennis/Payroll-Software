package salary_statement_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
	"github.com/jokertennis/Payroll-Software/db"
	"github.com/jokertennis/Payroll-Software/server/gen/models"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/employee_repository"
	"github.com/jokertennis/Payroll-Software/src/main/infrastructure"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func createDataForTestOfCreateSalaryStatementHandler(ctx context.Context, db boil.ContextExecutor) error {
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
		{ID: 3, CompanyID: 2, Name: "Tom", MailAddress: "tom@example.com", Password: "testpass"}}

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

	earnings := []models.Earning{
		{ID: 1, Nominal: "エンジニアLv1家族手当0人", Amount: 300000},
		{ID: 2, Nominal: "エンジニアLv1家族手当1人", Amount: 350000}}

	for _, earning := range earnings {
		if err := earning.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of earning. err:%s", err)
		}
	}

	earning_details := []models.EarningDetail{
		{ID: 1, EarningID: 1, Nominal: "基本給", Amount: 250000},
		{ID: 2, EarningID: 1, Nominal: "固定残業代", Amount: 50000},
		{ID: 3, EarningID: 2, Nominal: "基本給", Amount: 250000},
		{ID: 4, EarningID: 2, Nominal: "固定残業代", Amount: 50000},
		{ID: 5, EarningID: 2, Nominal: "家族手当(1人分)", Amount: 50000}}

	for _, earning_detail := range earning_details {
		if err := earning_detail.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of earning_detail. err:%s", err)
		}
	}

	deductions := []models.Deduction{
		{ID: 1, Nominal: "エンジニアLv1家族手当0人控除", Amount: 60000},
		{ID: 2, Nominal: "エンジニアLv1家族手当1人控除", Amount: 70000}}

	for _, deduction := range deductions {
		if err := deduction.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of deduction. err:%s", err)
		}
	}

	deduction_details := []models.DeductionDetail{
		{ID: 1, DeductionID: 1, Nominal: "健康保険料", Amount: 12000},
		{ID: 2, DeductionID: 1, Nominal: "厚生年金保険料", Amount: 20000},
		{ID: 3, DeductionID: 1, Nominal: "雇用保険料", Amount: 300},
		{ID: 4, DeductionID: 1, Nominal: "所得税", Amount: 5000},
		{ID: 5, DeductionID: 1, Nominal: "住民税", Amount: 22700},
		{ID: 6, DeductionID: 2, Nominal: "健康保険料", Amount: 13000},
		{ID: 7, DeductionID: 2, Nominal: "厚生年金保険料", Amount: 21000},
		{ID: 8, DeductionID: 2, Nominal: "雇用保険料", Amount: 500},
		{ID: 9, DeductionID: 2, Nominal: "所得税", Amount: 6000},
		{ID: 10, DeductionID: 2, Nominal: "住民税", Amount: 29500}}

	for _, deduction_detail := range deduction_details {
		if err := deduction_detail.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of deduction_detail. err:%s", err)
		}
	}

	salary_statements := []models.SalaryStatement{
		{ID: 1, EarningID: null.Uint32{Uint32: 1, Valid: true}, DeductionID: null.Uint32{Uint32: 1, Valid: true}, EmployeeID: 1, Nominal: "2022年1月 給与明細", Payday: time.Date(2022, 1, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年12月1日~2022年12月31日"},
		{ID: 2, EarningID: null.Uint32{Uint32: 1, Valid: true}, DeductionID: null.Uint32{Uint32: 1, Valid: true}, EmployeeID: 1, Nominal: "2022年2月 給与明細", Payday: time.Date(2022, 2, 20, 0, 0, 0, 0, time.Local), TargetPeriod: "2022年1月1日~2022年1月31日"},
	}

	for _, salary_statement := range salary_statements {
		if err := salary_statement.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of salary_statement. err:%s", err)
		}
	}

	return nil
}

// TODO: add a test to verify behavior when null is specified for non-nullable data.
// Now, when null is specified for non-nullable data, statuscode 602 is returned.
// we want to return 400 statuscode when this situation has occurred.
// I will add tests after modifying it to return 400 instead of 602 statuscode.
// Corresponding ticket: https://github.com/jokertennis/Payroll-Software/issues/49
func TestCreateSalaryStatementHandler(t *testing.T) {
	cases := map[string]struct {
		requestBody        string
		email              string
		password           string
		expectedStatusCode int
		expectedResponse   string
	}{
		"registered administrator successfully create salary statements.": {
			requestBody: `{
				"mailaddressOfEmployee": "joves@example.com",
				"nominal": "2022年2月分給料明細",
				"payday": "2022-02-20T00:00:00.000Z",
				"target_period": "2022年1月1日~2022年1月31日",
				"nominal_of_earning": "支給総額",
				"amount_of_earning": 300000,
				"_earning_details": [
				  {
					"nominal": "基本給",
					"amount_of_earning_detail": 250000
				  },
				  {
					"nominal": "残業代",
					"amount_of_earning_detail": 50000
				  }
				],
				"nominal_of_deduction": "控除総額",
				"amount_of_deduction": 50000,
				"_deduction_details": [
				  {
					"nominal": "住民税",
					"amount_of_deduction_detail": 10000
				  },
				  {
					"nominal": "所得税",
					"amount_of_deduction_detail": 15000
				  },
				  {
					"nominal": "健康保険料",
					"amount_of_deduction_detail": 5000
				  },
				  {
					"nominal": "厚生年金保険料",
					"amount_of_deduction_detail": 20000
				  }
				]
			  }`,
			email:              "test.administrator@example.com",
			password:           "testpass",
			expectedStatusCode: http.StatusCreated,
			expectedResponse:   "",
		},
		"badrequest: employee of specified mail address has already salary statement which has payday of specified year and month.": {
			requestBody: `{
				"mailaddressOfEmployee": "potter@example.com",
				"nominal": "2022年2月分給料明細",
				"payday": "2022-02-20T00:00:00.000Z",
				"target_period": "2022年1月1日~2022年1月31日",
				"nominal_of_earning": "支給総額",
				"amount_of_earning": 300000,
				"_earning_details": [
				  {
					"nominal": "基本給",
					"amount_of_earning_detail": 250000
				  },
				  {
					"nominal": "残業代",
					"amount_of_earning_detail": 50000
				  }
				],
				"nominal_of_deduction": "控除総額",
				"amount_of_deduction": 50000,
				"_deduction_details": [
				  {
					"nominal": "住民税",
					"amount_of_deduction_detail": 10000
				  },
				  {
					"nominal": "所得税",
					"amount_of_deduction_detail": 15000
				  },
				  {
					"nominal": "健康保険料",
					"amount_of_deduction_detail": 5000
				  },
				  {
					"nominal": "厚生年金保険料",
					"amount_of_deduction_detail": 20000
				  }
				]
			  }`,
			email:              "test.administrator@example.com",
			password:           "testpass",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   "{\"message\":\"badrequest. SalaryStatement which has payday that match year and month of specified payday was already existed\"}\n",
		},
		"badrequest: specified not existed mailaddressOfEmployee": {
			requestBody: `{
				"mailaddressOfEmployee": "notfound@example.com",
				"nominal": "2022年2月分給料明細",
				"payday": "2022-02-20T00:00:00.000Z",
				"target_period": "2022年1月1日~2022年1月31日",
				"nominal_of_earning": "支給総額",
				"amount_of_earning": 300000,
				"_earning_details": [
				  {
					"nominal": "基本給",
					"amount_of_earning_detail": 250000
				  },
				  {
					"nominal": "残業代",
					"amount_of_earning_detail": 50000
				  }
				],
				"nominal_of_deduction": "控除総額",
				"amount_of_deduction": 50000,
				"_deduction_details": [
				  {
					"nominal": "住民税",
					"amount_of_deduction_detail": 10000
				  },
				  {
					"nominal": "所得税",
					"amount_of_deduction_detail": 15000
				  },
				  {
					"nominal": "健康保険料",
					"amount_of_deduction_detail": 5000
				  },
				  {
					"nominal": "厚生年金保険料",
					"amount_of_deduction_detail": 20000
				  }
				]
			  }`,
			email:              "test.administrator@example.com",
			password:           "testpass",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   "{\"message\":\"badrequest. Specified mailAddress is not found in registered user datas. MailAddressOfEmployee:notfound@example.com\"}\n",
		},
		"badrequest: specified mailaddress of employee that belongs to a company different from administrator's own": {
			requestBody: `{
				"mailaddressOfEmployee": "tom@example.com",
				"nominal": "2022年2月分給料明細",
				"payday": "2022-02-20T00:00:00.000Z",
				"target_period": "2022年1月1日~2022年1月31日",
				"nominal_of_earning": "支給総額",
				"amount_of_earning": 300000,
				"_earning_details": [
				  {
					"nominal": "基本給",
					"amount_of_earning_detail": 250000
				  },
				  {
					"nominal": "残業代",
					"amount_of_earning_detail": 50000
				  }
				],
				"nominal_of_deduction": "控除総額",
				"amount_of_deduction": 50000,
				"_deduction_details": [
				  {
					"nominal": "住民税",
					"amount_of_deduction_detail": 10000
				  },
				  {
					"nominal": "所得税",
					"amount_of_deduction_detail": 15000
				  },
				  {
					"nominal": "健康保険料",
					"amount_of_deduction_detail": 5000
				  },
				  {
					"nominal": "厚生年金保険料",
					"amount_of_deduction_detail": 20000
				  }
				]
			  }`,
			email:              "test.administrator@example.com",
			password:           "testpass",
			expectedStatusCode: http.StatusBadRequest,
			expectedResponse:   "{\"message\":\"operations related to employee who belongs to a company different from your own are not possible\"}\n",
		},
		"execute by registered employee": {
			requestBody: `{
				"mailaddressOfEmployee": "joves@example.com",
				"nominal": "2022年2月分給料明細",
				"payday": "2022-02-20T00:00:00.000Z",
				"target_period": "2022年1月1日~2022年1月31日",
				"nominal_of_earning": "支給総額",
				"amount_of_earning": 300000,
				"_earning_details": [
				  {
					"nominal": "基本給",
					"amount_of_earning_detail": 250000
				  },
				  {
					"nominal": "残業代",
					"amount_of_earning_detail": 50000
				  }
				],
				"nominal_of_deduction": "控除総額",
				"amount_of_deduction": 50000,
				"_deduction_details": [
				  {
					"nominal": "住民税",
					"amount_of_deduction_detail": 10000
				  },
				  {
					"nominal": "所得税",
					"amount_of_deduction_detail": 15000
				  },
				  {
					"nominal": "健康保険料",
					"amount_of_deduction_detail": 5000
				  },
				  {
					"nominal": "厚生年金保険料",
					"amount_of_deduction_detail": 20000
				  }
				]
			  }`,
			email:              "potter@example.com",
			password:           "testpass",
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   "{\"message\":\"Unauthorized. Specified mailAddress is not found in registered user datas.\"}\n",
		},
		"Authentication header is not present": {
			requestBody: `{
				"mailaddressOfEmployee": "joves@example.com",
				"nominal": "2022年2月分給料明細",
				"payday": "2022-02-20T00:00:00.000Z",
				"target_period": "2022年1月1日~2022年1月31日",
				"nominal_of_earning": "支給総額",
				"amount_of_earning": 300000,
				"_earning_details": [
				  {
					"nominal": "基本給",
					"amount_of_earning_detail": 250000
				  },
				  {
					"nominal": "残業代",
					"amount_of_earning_detail": 50000
				  }
				],
				"nominal_of_deduction": "控除総額",
				"amount_of_deduction": 50000,
				"_deduction_details": [
				  {
					"nominal": "住民税",
					"amount_of_deduction_detail": 10000
				  },
				  {
					"nominal": "所得税",
					"amount_of_deduction_detail": 15000
				  },
				  {
					"nominal": "健康保険料",
					"amount_of_deduction_detail": 5000
				  },
				  {
					"nominal": "厚生年金保険料",
					"amount_of_deduction_detail": 20000
				  }
				]
			  }`,
			email:              "",
			password:           "",
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   "{\"message\":\"Unauthorized. Please check to see if Authentication header is not present or invalid.\"}\n",
		},
		"specify not incorrect password": {
			requestBody: `{
				"mailaddressOfEmployee": "joves@example.com",
				"nominal": "2022年2月分給料明細",
				"payday": "2022-02-20T00:00:00.000Z",
				"target_period": "2022年1月1日~2022年1月31日",
				"nominal_of_earning": "支給総額",
				"amount_of_earning": 300000,
				"_earning_details": [
				  {
					"nominal": "基本給",
					"amount_of_earning_detail": 250000
				  },
				  {
					"nominal": "残業代",
					"amount_of_earning_detail": 50000
				  }
				],
				"nominal_of_deduction": "控除総額",
				"amount_of_deduction": 50000,
				"_deduction_details": [
				  {
					"nominal": "住民税",
					"amount_of_deduction_detail": 10000
				  },
				  {
					"nominal": "所得税",
					"amount_of_deduction_detail": 15000
				  },
				  {
					"nominal": "健康保険料",
					"amount_of_deduction_detail": 5000
				  },
				  {
					"nominal": "厚生年金保険料",
					"amount_of_deduction_detail": 20000
				  }
				]
			  }`,
			email:              "test.administrator@example.com",
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
	err := createDataForTestOfCreateSalaryStatementHandler(ctx, dbInstance)
	if err != nil {
		fmt.Println(err)
	}

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	for _, value := range cases {
		req, err := http.NewRequest(http.MethodPost, "http://localhost:8080/administrator/salary_statement_", bytes.NewBufferString(value.requestBody))
		assert.Nil(t, err)

		if value.email == "" && value.password == "" {
			req.Header = http.Header{}
		} else {
			req.SetBasicAuth(value.email, value.password)
		}

		req.Header.Set("Content-Type", "application/json")

		res, err := client.Do(req)
		assert.Nil(t, err)
		defer res.Body.Close()
		assert.Equal(t, value.expectedStatusCode, res.StatusCode)
		if value.expectedStatusCode == http.StatusCreated {
			// get the latest salary statement data.
			createdLatestData, err := models.SalaryStatements(
				qm.Load(models.SalaryStatementRels.Earning),
				qm.Load(models.SalaryStatementRels.Deduction),
				qm.Load(models.SalaryStatementRels.Earning+"."+models.EarningRels.EarningDetails),
				qm.Load(models.SalaryStatementRels.Deduction+"."+models.DeductionRels.DeductionDetails),
				qm.OrderBy("id DESC"),
				qm.Limit(1),
			).One(ctx, dbInstance)
			assert.Nil(t, err)
			createdLatestSalaryStatementDomainObject, err := infrastructure.MappingSalaryStatementDomainObject(createdLatestData)
			assert.Nil(t, err)
			// convert string of request body to json.
			var expectedData map[string]interface{}
			err = json.Unmarshal([]byte(value.requestBody), &expectedData)
			assert.Nil(t, err)
			// check employee_id
			employeeRepositoryStruct := infrastructure.NewEmployeeRepository(ctx, dbInstance)
			var employeeRepository employee_repository.EmployeeRepository = &employeeRepositoryStruct
			employee, err := employeeRepository.GetEmployeeByMailAddress(expectedData["mailaddressOfEmployee"].(string))
			assert.Nil(t, err)
			assert.Equal(t, employee.ID, createdLatestSalaryStatementDomainObject.EmployeeId)
			// check other datas
			assert.Equal(t, expectedData["nominal"].(string), createdLatestSalaryStatementDomainObject.Nominal)
			assert.Equal(t, expectedData["target_period"].(string), createdLatestSalaryStatementDomainObject.TargetPeriod)
			assert.Equal(t, int(expectedData["amount_of_earning"].(float64)), createdLatestSalaryStatementDomainObject.Earning.Amount)
			assert.Equal(t, expectedData["nominal_of_earning"].(string), createdLatestSalaryStatementDomainObject.Earning.Nominal)
			assert.Equal(t, int(expectedData["amount_of_deduction"].(float64)), createdLatestSalaryStatementDomainObject.Deduction.Amount)
			assert.Equal(t, expectedData["nominal_of_deduction"].(string), createdLatestSalaryStatementDomainObject.Deduction.Nominal)
			// TODO: check payday

			var arrayOfNominalOfEarningDetail []string
			var arrayOfAmountOfEarningDetail []int
			for _, value := range createdLatestSalaryStatementDomainObject.Earning.EarningDetails {
				arrayOfNominalOfEarningDetail = append(arrayOfNominalOfEarningDetail, value.Nominal)
				arrayOfAmountOfEarningDetail = append(arrayOfAmountOfEarningDetail, value.Amount)
			}
			earningDetails := expectedData["_earning_details"].([]interface{})
			for _, detail := range earningDetails {
				detailMap := detail.(map[string]interface{})
				assert.Contains(t, arrayOfNominalOfEarningDetail, detailMap["nominal"].(string))
				assert.Contains(t, arrayOfAmountOfEarningDetail, int(detailMap["amount_of_earning_detail"].(float64)))
			}

			var arrayOfNominalOfDeductionDetail []string
			var arrayOfAmountOfDeductionDetail []int
			for _, value := range createdLatestSalaryStatementDomainObject.Deduction.DeductionDetails {
				arrayOfNominalOfDeductionDetail = append(arrayOfNominalOfDeductionDetail, value.Nominal)
				arrayOfAmountOfDeductionDetail = append(arrayOfAmountOfDeductionDetail, value.Amount)
			}
			deductionDetails := expectedData["_deduction_details"].([]interface{})
			for _, detail := range deductionDetails {
				detailMap := detail.(map[string]interface{})
				assert.Contains(t, arrayOfNominalOfDeductionDetail, detailMap["nominal"].(string))
				assert.Contains(t, arrayOfAmountOfDeductionDetail, int(detailMap["amount_of_deduction_detail"].(float64)))
			}
		} else {
			data, err := ioutil.ReadAll(res.Body)
			assert.Nil(t, err)
			assert.Equal(t, value.expectedResponse, string(data))
			// 1. confirm that response body is expected value.
		}
	}
}
