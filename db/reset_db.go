package db

import (
	"context"
	"fmt"
	"usr/local/go/server/gen/models"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

func ResetDb(ctx context.Context, db boil.ContextExecutor) error {
	if _, err := models.SalaryStatements().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of salary_statement. err:%s", err)
	}

	if _, err := models.FixedDeductionDetails().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of fixed_deduction_detail. err:%s", err)
	}

	if _, err := models.FixedDeductions().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of fixed_deduction. err:%s", err)
	}

	if _, err := models.DeductionDetails().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of deduction_detail. err:%s", err)
	}

	if _, err := models.Deductions().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of deduction. err:%s", err)
	}

	if _, err := models.FixedEarningDetails().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of fixed_earning_detail. err:%s", err)
	}

	if _, err := models.FixedEarnings().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of fixed_earning. err:%s", err)
	}

	if _, err := models.EarningDetails().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of earning_detail. err:%s", err)
	}

	if _, err := models.Earnings().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of earning. err:%s", err)
	}

	if _, err := models.Employees().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of employee. err:%s", err)
	}

	if _, err := models.Administrators().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of administrator. err:%s", err)
	}

	if _, err := models.Companies().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of company. err:%s", err)
	}

	return nil
}
