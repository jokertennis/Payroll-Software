package basicauth

import (
	"context"
	"fmt"
	"os"
	"testing"
	"usr/local/go/basicauth"
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

func TestGetExpectedUserInformationForEmployeeTest(t *testing.T) {
	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
	dbInstance, err := db.CreateDbInstance(dbEnvironment)
	assert.Nil(t, err)

	expectedUserInformation, err := basicauth.GetExpectedUserInformation(ctx, dbInstance, "potter@example.com", basicauth.Executer{Executer: "Employee"})
	assert.Nil(t, err)
	assert.NotNil(t, expectedUserInformation)
}

func TestGetExpectedUserInformationForAdministratorTest(t *testing.T) {
	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
	dbInstance, err := db.CreateDbInstance(dbEnvironment)
	assert.Nil(t, err)

	expectedUserInformation, err := basicauth.GetExpectedUserInformation(ctx, dbInstance, "test.administrator@example.com", basicauth.Executer{Executer: "Administrator"})
	assert.Nil(t, err)
	assert.NotNil(t, expectedUserInformation)
}

func TestGetExpectedUserInformationForNotExistUserTest(t *testing.T) {
	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
	dbInstance, err := db.CreateDbInstance(dbEnvironment)
	assert.Nil(t, err)

	expectedUserInformation, err := basicauth.GetExpectedUserInformation(ctx, dbInstance, "potter@example.com", basicauth.Executer{Executer: "Administrator"})
	assert.Nil(t, err)
	assert.Nil(t, expectedUserInformation)
}

func TestGetExpectedUserInformationForNotExistUserTestOk(t *testing.T) {
	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
	dbInstance, err := db.CreateDbInstance(dbEnvironment)
	assert.Nil(t, err)

	expectedUserInformation, err := basicauth.GetExpectedUserInformation(ctx, dbInstance, "potter@example.com", basicauth.Executer{Executer: "NotSupport"})
	assert.Equal(t, fmt.Errorf("not support specified Executer:NotSupport"), err)
	assert.Nil(t, expectedUserInformation)
}
