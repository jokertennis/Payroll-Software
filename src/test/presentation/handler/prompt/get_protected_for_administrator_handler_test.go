package prompt

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"usr/local/go/db"
	"usr/local/go/server/gen/models"

	"github.com/stretchr/testify/assert"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func TestMain(m *testing.M) {
	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
	dbInstance, _ := db.CreateDbInstance(dbEnvironment)

	// executed reset of test db.
	_ = db.ResetDb(ctx, dbInstance)

	// executed creating test data.
	err := createDataForTestsInPromptPackage(ctx, dbInstance)
	if err != nil {
		fmt.Println(err)
	}

	code := m.Run()

	os.Exit(code)
}

// This function is used when we want to create test/demo data/
func createDataForTestsInPromptPackage(ctx context.Context, db boil.ContextExecutor) error {
	companies := []models.Company{
		{ID: 1, Name: "株式会社川田"},
		{ID: 2, Name: "株式会社kawata"}}

	for _, company := range companies {
		if err := company.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create demo datas of company. err:%s", err)
		}
	}

	employees := []models.Employee{
		{ID: 1, CompanyID: 1, Name: "Potter", MailAddress: "potter@example.com", Password: "testpass"}}

	for _, employee := range employees {
		if err := employee.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of employee. err:%s", err)
		}
	}

	administrators := []models.Administrator{
		{ID: 1, CompanyID: 1, Name: "TestAdministrator", MailAddress: "test.administrator@example.com", Password: "testpass"}}

	for _, administrator := range administrators {
		if err := administrator.Insert(ctx, db, boil.Infer()); err != nil {
			return fmt.Errorf("failed to create datas of administrator. err:%s", err)
		}
	}
	return nil
}

func TestProtectedHandlerForAdministrator(t *testing.T) {
	cases := map[string]struct {
		email              string
		password           string
		expectedStatusCode int
		expectedResponse   string
	}{
		"execute by registered administrator": {
			email:              "test.administrator@example.com",
			password:           "testpass",
			expectedStatusCode: http.StatusOK,
			expectedResponse:   "{\"message\":\"This is the protected handler\"}\n",
		},
		"execute by registered employee": {
			email:              "potter@example.com",
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
		"specify not existed user": {
			email:              "notfound@example.com",
			password:           "testpass",
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   "{\"message\":\"Unauthorized. Specified mailAddress is not found in registered user datas.\"}\n",
		},
		"specify not incorrect password": {
			email:              "test.administrator@example.com",
			password:           "incorrectpassword",
			expectedStatusCode: http.StatusUnauthorized,
			expectedResponse:   "{\"message\":\"Unauthorized. Please check to see if password is correct.\"}\n",
		},
	}
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/administrator/protected", nil)
	assert.Nil(t, err)

	for _, value := range cases {
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
