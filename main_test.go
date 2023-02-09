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
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/unprotected", nil)
	assert.Nil(t, err)
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "This is the unprotected handler", string(data))
}

func TestProtectedHandlerForEmployee(t *testing.T) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/protected", nil)
	assert.Nil(t, err)
	req.SetBasicAuth("potter@example.com", "testpass")
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "This is the protected handler", string(data))
}

func TestProtectedHandlerForAdministrator(t *testing.T) {
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/admin/protected", nil)
	assert.Nil(t, err)
	req.SetBasicAuth("test.administrator@example.com", "testpass")
	res, err := client.Do(req)
	assert.Nil(t, err)
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, "This is the protected handler", string(data))
}