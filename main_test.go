package main

import (
	"context"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
	"usr/local/go/db"

	"github.com/stretchr/testify/assert"
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
	_ = db.CreateData(ctx, dbInstance)

	code := m.Run()

	os.Exit(code)
}

func TestUnprotectedHandler(t *testing.T) {
	cases := map[string]struct {
		email            string
		password         string
		expectStatusCode int
		expectResponse   string
	}{
		"execute by registered employee": {
			email:            "potter@example.com",
			password:         "testpass",
			expectStatusCode: http.StatusOK,
			expectResponse:   "This is the unprotected handler",
		},
		"execute by registered administrator": {
			email:            "test.administrator@example.com",
			password:         "testpass",
			expectStatusCode: http.StatusOK,
			expectResponse:   "This is the unprotected handler",
		},
		"Authentication header is not present": {
			email:            "",
			password:         "",
			expectStatusCode: http.StatusOK,
			expectResponse:   "This is the unprotected handler",
		},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/unprotected", nil)
	assert.Nil(t, err)

	for _, value := range cases {
		if value.email != "" && value.password != "" {
			req.SetBasicAuth(value.email, value.password)
		}
		res, err := client.Do(req)
		assert.Nil(t, err)
		defer res.Body.Close()
		assert.Equal(t, value.expectStatusCode, res.StatusCode)
		data, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.Equal(t, value.expectResponse, string(data))
	}
}

func TestProtectedHandlerForEmployee(t *testing.T) {
	cases := map[string]struct {
		email            string
		password         string
		expectStatusCode int
		expectResponse   string
	}{
		"execute by registered employee": {
			email:            "potter@example.com",
			password:         "testpass",
			expectStatusCode: http.StatusOK,
			expectResponse:   "This is the protected handler",
		},
		"execute by registered administrator": {
			email:            "test.administrator@example.com",
			password:         "testpass",
			expectStatusCode: http.StatusUnauthorized,
			expectResponse:   "Unauthorized.Specified mailAddress is not found in registered user datas.\n",
		},
		"Authentication header is not present": {
			email:            "",
			password:         "",
			expectStatusCode: http.StatusUnauthorized,
			expectResponse:   "Unauthorized.Please check to see if Authentication header is not present or invalid.\n",
		},
		"specify not existed user": {
			email:            "notfound@example.com",
			password:         "testpass",
			expectStatusCode: http.StatusUnauthorized,
			expectResponse:   "Unauthorized.Specified mailAddress is not found in registered user datas.\n",
		},
		"specify not incorrect password": {
			email:            "potter@example.com",
			password:         "incorrectpassword",
			expectStatusCode: http.StatusUnauthorized,
			expectResponse:   "Unauthorized.Please check to see if password is correct.\n",
		},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/protected", nil)
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
		assert.Equal(t, value.expectStatusCode, res.StatusCode)
		data, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.Equal(t, value.expectResponse, string(data))
	}
}

func TestProtectedHandlerForAdministrator(t *testing.T) {
	cases := map[string]struct {
		email            string
		password         string
		expectStatusCode int
		expectResponse   string
	}{
		"execute by registered administrator": {
			email:            "test.administrator@example.com",
			password:         "testpass",
			expectStatusCode: http.StatusOK,
			expectResponse:   "This is the protected handler",
		},
		"execute by registered employee": {
			email:            "potter@example.com",
			password:         "testpass",
			expectStatusCode: http.StatusUnauthorized,
			expectResponse:   "Unauthorized.Specified mailAddress is not found in registered user datas.\n",
		},
		"Authentication header is not present": {
			email:            "",
			password:         "",
			expectStatusCode: http.StatusUnauthorized,
			expectResponse:   "Unauthorized.Please check to see if Authentication header is not present or invalid.\n",
		},
		"specify not existed user": {
			email:            "notfound@example.com",
			password:         "testpass",
			expectStatusCode: http.StatusUnauthorized,
			expectResponse:   "Unauthorized.Specified mailAddress is not found in registered user datas.\n",
		},
		"specify not incorrect password": {
			email:            "test.administrator@example.com",
			password:         "incorrectpassword",
			expectStatusCode: http.StatusUnauthorized,
			expectResponse:   "Unauthorized.Please check to see if password is correct.\n",
		},
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/admin/protected", nil)
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
		assert.Equal(t, value.expectStatusCode, res.StatusCode)
		data, err := ioutil.ReadAll(res.Body)
		assert.Nil(t, err)
		assert.Equal(t, value.expectResponse, string(data))
	}
}
