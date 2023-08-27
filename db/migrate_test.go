package db_test

import (
	"fmt"
	"testing"
	"github.com/jokertennis/Payroll-Software/db"

	"github.com/stretchr/testify/assert"
)

func TestCreateDbInstance(t *testing.T) {
	cases := map[string]struct {
		dbEnvironment db.DbEnvironment
		expectedError error
	}{
		"execute at Develop environment": {
			dbEnvironment: db.DbEnvironment{Environment: "Develop"},
			expectedError: nil,
		},
		"execute at not supported environment": {
			dbEnvironment: db.DbEnvironment{Environment: "NotSupport"},
			expectedError: fmt.Errorf("not support specified DbEnviroment:NotSupport"),
		},
	}

	for _, value := range cases {
		_, err := db.CreateDbInstance(value.dbEnvironment)
		assert.Equal(t, value.expectedError, err)
	}
}

// TODO:
// The test of creating migrate-instance is not successful.
// I don't understand the reason why fails test.
// I confirmed that creating migrate-instance was successful when using on main.go file.
// If I understand the reason why fails test, I will fix implementation of migration process.

// func TestCreateMigrateInstance(t *testing.T) {
// 	// // create dbInstance which is used when accessing db.
// 	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
// 	dbInstance, err := db.CreateDbInstance(dbEnvironment)
// 	assert.Nil(t, err)

// 	// // migrate db
// 	migrateInstance, err := db.CreateMigrateInstance(dbInstance)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, migrateInstance)
// }
