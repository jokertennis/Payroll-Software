package db_test

import (
	"context"
	"testing"
	"github.com/jokertennis/Payroll-Software/db"
	"github.com/jokertennis/Payroll-Software/server/gen/models"

	"github.com/stretchr/testify/assert"
)

// confirm that creating test/demo data is successful and reset of db is successful.
func TestCreateDataAndResetDb(t *testing.T) {
	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	dbEnvironment := db.DbEnvironment{Environment: "Develop"}
	dbInstance, err := db.CreateDbInstance(dbEnvironment)
	assert.Nil(t, err)

	// executed reset of db.
	err = db.ResetDb(ctx, dbInstance)
	assert.Nil(t, err)

	err = db.CreateData(ctx, dbInstance)
	assert.Nil(t, err)

	// confirm that creating test data is successful
	boolCompanyExist, err := models.CompanyExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.True(t, boolCompanyExist)

	boolEmployeeExist, err := models.EmployeeExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.True(t, boolEmployeeExist)

	boolAdministratorExist, err := models.AdministratorExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.True(t, boolAdministratorExist)

	boolEarningExist, err := models.EarningExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.True(t, boolEarningExist)

	boolEarningDetailExist, err := models.EarningDetailExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.True(t, boolEarningDetailExist)

	boolDeductionExist, err := models.DeductionExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.True(t, boolDeductionExist)

	boolDeductionDetailExist, err := models.DeductionDetailExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.True(t, boolDeductionDetailExist)

	boolSalaryStatementExist, err := models.SalaryStatementExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.True(t, boolSalaryStatementExist)

	// executed reset of test db.
	err = db.ResetDb(ctx, dbInstance)
	assert.Nil(t, err)

	// confirm that reset of test db is successful.
	boolCompanyExist, err = models.CompanyExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.False(t, boolCompanyExist)

	boolEmployeeExist, err = models.EmployeeExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.False(t, boolEmployeeExist)

	boolAdministratorExist, err = models.AdministratorExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.False(t, boolAdministratorExist)

	boolEarningExist, err = models.EarningExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.False(t, boolEarningExist)

	boolEarningDetailExist, err = models.EarningDetailExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.False(t, boolEarningDetailExist)

	boolDeductionExist, err = models.DeductionExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.False(t, boolDeductionExist)

	boolDeductionDetailExist, err = models.DeductionDetailExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.False(t, boolDeductionDetailExist)

	boolSalaryStatementExist, err = models.SalaryStatementExists(ctx, dbInstance, 1)
	assert.Nil(t, err)
	assert.False(t, boolSalaryStatementExist)
}
