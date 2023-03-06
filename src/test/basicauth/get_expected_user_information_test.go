package basicauth

import (
	"context"
	"fmt"
	"os"
	"testing"
	"usr/local/go/basicauth"
	"usr/local/go/db"
	"usr/local/go/src/main/domain-service/repository"
	"usr/local/go/src/main/infrastructure"

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

func TestGetExpectedUserInformation(t *testing.T) {
	cases := map[string]struct {
		email                   string
		executer                basicauth.Executer
		expectedError           error
		expectedUserInformation *basicauth.ExpectedUserInformation
	}{
		"get a registered employee as employee": {
			email:                   "potter@example.com",
			executer:                basicauth.Executer{Executer: "Employee"},
			expectedError:           nil,
			expectedUserInformation: &basicauth.ExpectedUserInformation{MailAddress: "potter@example.com", Password: "testpass"},
		},
		"get a registered administrator as administrator": {
			email:                   "test.administrator@example.com",
			executer:                basicauth.Executer{Executer: "Administrator"},
			expectedError:           nil,
			expectedUserInformation: &basicauth.ExpectedUserInformation{MailAddress: "test.administrator@example.com", Password: "testpass"},
		},
		"fail to get existed employee as administrator": {
			email:                   "potter@example.com",
			executer:                basicauth.Executer{Executer: "Administrator"},
			expectedError:           nil,
			expectedUserInformation: nil,
		},
		"fail to get existed administrator as employee": {
			email:                   "test.administrator@example.com",
			executer:                basicauth.Executer{Executer: "Employee"},
			expectedError:           nil,
			expectedUserInformation: nil,
		},
		"get not existed user": {
			email:                   "notfound@example.com",
			executer:                basicauth.Executer{Executer: "Employee"},
			expectedError:           nil,
			expectedUserInformation: nil,
		},
		"specify not supported executer": {
			email:                   "test.administrator@example.com",
			executer:                basicauth.Executer{Executer: "NotSupport"},
			expectedError:           fmt.Errorf("not support specified Executer:NotSupport"),
			expectedUserInformation: nil,
		},
	}

	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
	dbInstance, err := db.CreateDbInstance(dbEnvironment)
	assert.Nil(t, err)

	employeeRepositoryStruct := infrastructure.NewEmployeeRepository(dbInstance)
	var employeeRepository repository.EmployeeRepository = &employeeRepositoryStruct

	administratorRepositoryStruct := infrastructure.NewAdministratorRepository(dbInstance)
	var administratorRepository repository.AdministratorRepository = &administratorRepositoryStruct

	for _, value := range cases {
		gotUserInformation, err := basicauth.GetExpectedUserInformation(ctx, employeeRepository, administratorRepository, value.email, value.executer)
		assert.Equal(t, value.expectedError, err)
		assert.Equal(t, value.expectedUserInformation, gotUserInformation)
	}
}
