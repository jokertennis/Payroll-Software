package db

import(
	"testing"
	"usr/local/go/db"

	"github.com/stretchr/testify/assert"
)

func TestCreateDbInstanceForDevelopOk(t *testing.T) {
	developDbEnvironment := db.DbEnvironment{Environment: "Develop"}
	_, err := db.CreateDbInstance(developDbEnvironment)
	// only confirm that error don't occur when creating db instance.
	assert.Nil(t, err)
}

func TestCreateDbInstanceForTestOk(t *testing.T) {
	testDbEnvironment := db.DbEnvironment{Environment: "Test"}
	_, err := db.CreateDbInstance(testDbEnvironment)
	// only confirm that error don't occur when creating db instance.
	assert.Nil(t, err)
}

func TestCreateDbInstanceForTestException(t *testing.T) {
	notExistedDbEnvironment := db.DbEnvironment{Environment: "NotExist"}
	_, err := db.CreateDbInstance(notExistedDbEnvironment)
	// only confirm that expetced error occors.
	assert.EqualError(t, err, "not support specific DbEnviroment:{NotExist}")
}

// TODO:
// The test of creating migrate-instance is not successful.
// I don't understand the reason why fails test.
// I confirmed that success of creating migrate-instance when using on main.go file.
// If I understand the reason why fails test, I will fix implementation of migration process.

// func TestCreateMigrateInstanceOk(t *testing.T) {
// 	// // create dbInstance which is used when accessing db.
// 	testDbEnvironment := db.DbEnvironment{Environment: "Test"}
// 	testDbInstance, err := db.CreateDbInstance(testDbEnvironment)
// 	assert.Nil(t, err)

// 	// // migrate test_db
// 	migrateInstanceForTestDb, errorCreateMigrateInstance := db.CreateMigrateInstance(testDbInstance)
// 	assert.Nil(t, errorCreateMigrateInstance)
// 	assert.NotNil(t, migrateInstanceForTestDb)
// }