package basicauth_test

import (
	"context"
	"fmt"
	"os"
	"testing"
	"github.com/jokertennis/Payroll-Software/basicauth"
	"github.com/jokertennis/Payroll-Software/db"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/employee_repository"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/administrator_repository"
	"github.com/jokertennis/Payroll-Software/src/main/infrastructure"

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

	employeeRepositoryStruct := infrastructure.NewEmployeeRepository(ctx, dbInstance)
	var employeeRepository employee_repository.EmployeeRepository = &employeeRepositoryStruct

	administratorRepositoryStruct := infrastructure.NewAdministratorRepository(ctx, dbInstance)
	var administratorRepository administrator_repository.AdministratorRepository = &administratorRepositoryStruct

	for _, value := range cases {
		gotUserInformation, err := basicauth.GetExpectedUserInformation(employeeRepository, administratorRepository, value.email, value.executer)
		assert.Equal(t, value.expectedError, err)
		assert.Equal(t, value.expectedUserInformation, gotUserInformation)
	}
}
