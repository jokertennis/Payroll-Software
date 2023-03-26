package infrastructure

import (
	"context"
	"database/sql"
	"time"
	"usr/local/go/server/gen/models"
	"usr/local/go/src/main/domain-model/salary_statement"

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
		qm.Where("employee_id=?", employeeId),
		qm.Expr(qm.And("payday >= ?", start),
		qm.Or("payday < ?", end)),
		qm.Load(models.SalaryStatementRels.IndividualEarning),
		qm.Load(models.SalaryStatementRels.FixedEarning),
		qm.Load(models.SalaryStatementRels.IndividualDeduction),
		qm.Load(models.SalaryStatementRels.FixedDeduction),
		qm.Load(models.SalaryStatementRels.Employee),
	).All(r.ctx, r.db)
	if err != nil {
		return nil, err
	}
	if salaryStatements == nil {
		return nil, nil
	}

	return MappingSalaryStatementDomainObject(salaryStatements[0])
}