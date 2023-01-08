package db

import (
	"context"
	"testing"
	"usr/local/go/db"
	"usr/local/go/server/gen/models"

	"github.com/stretchr/testify/assert"
)

// confirm that creating test data is successful and reset of test db is successful.
func TestCreateTestDataAndResetDb(t *testing.T) {
	// create context
	ctx := context.Background()

	// create dbInstance which is used when accessing db.
	testDbEnvironment := db.DbEnvironment{Environment: "Test"}
	testDbInstance, err := db.CreateDbInstance(testDbEnvironment)
	assert.Nil(t, err)
	err = db.CreateTestData(ctx, testDbInstance)
	assert.Nil(t, err)
	
	// confirm that creating test data is successful
	boolCompanyExist, err := models.CompanyExists(ctx, testDbInstance, 1)
	assert.True(t, boolCompanyExist)

	boolEmployeeExist, err := models.EmployeeExists(ctx, testDbInstance, 1)
	assert.True(t, boolEmployeeExist)

	boolEarningExist, err := models.EarningExists(ctx, testDbInstance, 1)
	assert.True(t, boolEarningExist)

	boolEarningDetailExist, err := models.EarningDetailExists(ctx, testDbInstance, 1)
	assert.True(t, boolEarningDetailExist)

	boolFixedEarningExist, err := models.FixedEarningExists(ctx, testDbInstance, 1)
	assert.True(t, boolFixedEarningExist)

	boolFixedEarningDetailExist, err := models.FixedEarningDetailExists(ctx, testDbInstance, 1)
	assert.True(t, boolFixedEarningDetailExist)

	boolDeductionExist, err := models.DeductionExists(ctx, testDbInstance, 1)
	assert.True(t, boolDeductionExist)

	boolDeductionDetailExist, err := models.DeductionDetailExists(ctx, testDbInstance, 1)
	assert.True(t, boolDeductionDetailExist)

	boolFixedDeductionExist, err := models.FixedDeductionExists(ctx, testDbInstance, 1)
	assert.True(t, boolFixedDeductionExist)

	boolFixedDeductionDetailExist, err := models.FixedDeductionDetailExists(ctx, testDbInstance, 1)
	assert.True(t, boolFixedDeductionDetailExist)

	boolSalaryStatementExist, err := models.SalaryStatementExists(ctx, testDbInstance, 1)
	assert.True(t, boolSalaryStatementExist)

	// executed reset of test db.
	err = db.ResetDb(ctx, testDbInstance)
	assert.Nil(t, err)

	// confirm that reset of test db is successful.
	boolCompanyExist, err = models.CompanyExists(ctx, testDbInstance, 1)
	assert.False(t, boolCompanyExist)

	boolEmployeeExist, err = models.EmployeeExists(ctx, testDbInstance, 1)
	assert.False(t, boolEmployeeExist)

	boolEarningExist, err = models.EarningExists(ctx, testDbInstance, 1)
	assert.False(t, boolEarningExist)

	boolEarningDetailExist, err = models.EarningDetailExists(ctx, testDbInstance, 1)
	assert.False(t, boolEarningDetailExist)

	boolFixedEarningExist, err = models.FixedEarningExists(ctx, testDbInstance, 1)
	assert.False(t, boolFixedEarningExist)

	boolFixedEarningDetailExist, err = models.FixedEarningDetailExists(ctx, testDbInstance, 1)
	assert.False(t, boolFixedEarningDetailExist)

	boolDeductionExist, err = models.DeductionExists(ctx, testDbInstance, 1)
	assert.False(t, boolDeductionExist)

	boolDeductionDetailExist, err = models.DeductionDetailExists(ctx, testDbInstance, 1)
	assert.False(t, boolDeductionDetailExist)

	boolFixedDeductionExist, err = models.FixedDeductionExists(ctx, testDbInstance, 1)
	assert.False(t, boolFixedDeductionExist)

	boolFixedDeductionDetailExist, err = models.FixedDeductionDetailExists(ctx, testDbInstance, 1)
	assert.False(t, boolFixedDeductionDetailExist)

	boolSalaryStatementExist, err = models.SalaryStatementExists(ctx, testDbInstance, 1)
	assert.False(t, boolSalaryStatementExist)
}