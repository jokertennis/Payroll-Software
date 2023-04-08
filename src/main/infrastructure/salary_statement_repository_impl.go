package infrastructure

import (
	"context"
	"database/sql"
	"log"
	"time"
	"usr/local/go/server/gen/models"
	"usr/local/go/src/main/domain-model/fixed_deduction"
	"usr/local/go/src/main/domain-model/fixed_earning"
	"usr/local/go/src/main/domain-model/individual_deduction"
	"usr/local/go/src/main/domain-model/individual_earning"
	"usr/local/go/src/main/domain-model/salary_statement"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SalaryStatementRepository struct {
	ctx context.Context
	db *sql.DB
}

func NewSalaryStatementRepository(ctx context.Context, db *sql.DB) SalaryStatementRepository {
	return SalaryStatementRepository{
		ctx: ctx,
		db: db,
	}
}

func (r *SalaryStatementRepository) GetSalaryStatement(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement.SalaryStatement, error) {
	start := time.Date(yearOfPayday, monthOfPayday, 1, 0, 0, 0, 0, time.UTC) // 年月の最初の日時
	end := start.AddDate(0, 1, 0) // 年月の最後の日時
	salaryStatements, err := models.SalaryStatements(
		qm.Where("salary_statements.employee_id=?", employeeId),
		qm.And("salary_statements.payday >= ?", start),
		qm.And("salary_statements.payday < ?", end),
	).All(r.ctx, r.db)
	if err != nil {
		return nil, err
	}
	if salaryStatements == nil {
		return nil, nil
	}
	salaryStatement := salaryStatements[0]
	var individualEarning *individual_earning.IndividualEarning
	if salaryStatement.IndividualEarningID.Valid {
		individualEarning, err = GetIndividualEarning(r.ctx, r.db, salaryStatement.IndividualEarningID)
		if err != nil {
			log.Fatalf("failed to get individual earning.%s", err)
		}
	} else {
		individualEarning = nil
	}

	var fixedEarning *fixed_earning.FixedEarning
	if salaryStatement.FixedEarningID.Valid {
		fixedEarning, err = GetFixedEarning(r.ctx, r.db, salaryStatement.FixedEarningID)
		if err != nil {
			log.Fatalf("failed to get fixed earning.%s", err)
		}
	} else {
		fixedEarning = nil
	}

	var individualDeduction *individual_deduction.IndividualDeduction
	if salaryStatement.IndividualDeductionID.Valid {
		individualDeduction, err = GetIndividualDeduction(r.ctx, r.db, salaryStatement.IndividualDeductionID)
		if err != nil {
			log.Fatalf("failed to get individual deduction.%s", err)
		}
	} else {
		individualDeduction = nil
	}

	var fixedDeduction *fixed_deduction.FixedDeduction
	if salaryStatement.FixedDeductionID.Valid {
		fixedDeduction, err = GetFixedDeduction(r.ctx, r.db, salaryStatement.FixedDeductionID)
		if err != nil {
			log.Fatalf("failed to get fixed deduction.%s", err)
		}
	} else {
		fixedDeduction = nil
	}

	salaryStatementDomainObject, err := salary_statement.NewSalaryStatement(salaryStatement.ID, individualEarning, fixedEarning, individualDeduction, fixedDeduction, salaryStatement.EmployeeID, salaryStatement.Nominal, salaryStatement.Payday, salaryStatement.TargetPeriod)
	if err != nil {
        return nil, err
    }
    return salaryStatementDomainObject, nil
}

func GetIndividualEarning(ctx context.Context, db *sql.DB, individualEarningId null.Uint32) (*individual_earning.IndividualEarning, error){
	individualEarning, err := models.IndividualEarnings(
		qm.Load(models.IndividualEarningRels.IndividualEarningDetails),
		qm.Where("individual_earnings.id=?", individualEarningId),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}
	if individualEarning == nil {
		return nil, nil
	}

	return MappingIndividualEarning(individualEarning)
}

func GetFixedEarning(ctx context.Context, db *sql.DB, fixedEarningId null.Uint32) (*fixed_earning.FixedEarning, error) {
	fixedEarning, err := models.FixedEarnings(
		qm.Load(models.FixedEarningRels.FixedEarningDetails),
		qm.Where("fixed_earnings.id=?", fixedEarningId),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}
	if fixedEarning == nil {
		return nil, nil
	}

	return MappingFixedEarning(fixedEarning)
}

func GetIndividualDeduction(ctx context.Context, db *sql.DB, individualDeductionId null.Uint32) (*individual_deduction.IndividualDeduction, error){
	individualDeduction, err := models.IndividualDeductions(
		qm.Load(models.IndividualDeductionRels.IndividualDeductionDetails),
		qm.Where("individual_deductions.id=?", individualDeductionId),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}
	if individualDeduction == nil {
		return nil, nil
	}

	return MappingIndividualDeduction(individualDeduction)
}

func GetFixedDeduction(ctx context.Context, db *sql.DB, fixedDeductionId null.Uint32) (*fixed_deduction.FixedDeduction, error) {
	fixedDeduction, err := models.FixedDeductions(
		qm.Load(models.FixedDeductionRels.FixedDeductionDetails),
		qm.Where("fixed_deductions.id=?", fixedDeductionId),
	).One(ctx, db)
	if err != nil {
		return nil, err
	}
	if fixedDeduction == nil {
		return nil, nil
	}

	return MappingFixedDeduction(fixedDeduction)
}