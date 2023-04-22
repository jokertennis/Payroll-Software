package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
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

func (r *SalaryStatementRepository) GetSalaryStatement(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
	start := time.Date(yearOfPayday, monthOfPayday, 1, 0, 0, 0, 0, time.UTC) // 年月の最初の日時
	end := start.AddDate(0, 1, 0) // 年月の最後の日時
	salaryStatements, err := models.SalaryStatements(
		qm.Load(models.SalaryStatementRels.IndividualEarning),
		qm.Load(models.SalaryStatementRels.FixedEarning),
		qm.Load(models.SalaryStatementRels.IndividualDeduction),
		qm.Load(models.SalaryStatementRels.FixedDeduction),
		qm.Load(models.SalaryStatementRels.IndividualEarning+"."+models.IndividualEarningRels.IndividualEarningDetails),
		qm.Load(models.SalaryStatementRels.FixedEarning+"."+models.FixedEarningRels.FixedEarningDetails),
		qm.Load(models.SalaryStatementRels.IndividualDeduction+"."+models.IndividualDeductionRels.IndividualDeductionDetails),
		qm.Load(models.SalaryStatementRels.FixedDeduction+"."+models.FixedDeductionRels.FixedDeductionDetails),
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
	if len(salaryStatements) >= 2 {
		return nil, fmt.Errorf("multiple salary statement were found.range of payday:%s~%s", start, end)
	}
	return MappingSalaryStatementDomainObject(salaryStatements[0])
}

func (r *SalaryStatementRepository) GetAllSalaryStatements(employeeId uint32) ([]*salary_statement_domain_model.SalaryStatement, error) {
	salaryStatements, err := models.SalaryStatements(
		qm.Load(models.SalaryStatementRels.IndividualEarning),
		qm.Load(models.SalaryStatementRels.FixedEarning),
		qm.Load(models.SalaryStatementRels.IndividualDeduction),
		qm.Load(models.SalaryStatementRels.FixedDeduction),
		qm.Load(models.SalaryStatementRels.IndividualEarning+"."+models.IndividualEarningRels.IndividualEarningDetails),
		qm.Load(models.SalaryStatementRels.FixedEarning+"."+models.FixedEarningRels.FixedEarningDetails),
		qm.Load(models.SalaryStatementRels.IndividualDeduction+"."+models.IndividualDeductionRels.IndividualDeductionDetails),
		qm.Load(models.SalaryStatementRels.FixedDeduction+"."+models.FixedDeductionRels.FixedDeductionDetails),
		qm.Where("salary_statements.employee_id=?", employeeId),
	).All(r.ctx, r.db)
	if err != nil {
		return nil, err
	}
	if salaryStatements == nil {
		return nil, nil
	}

	var salaryStatementDomainObjectList []*salary_statement_domain_model.SalaryStatement 
	for _, salaryStatement := range salaryStatements {
		salaryStatementDomainObject, err := MappingSalaryStatementDomainObject(salaryStatement)
		if err != nil {
			return nil, err
		}
		salaryStatementDomainObjectList = append(salaryStatementDomainObjectList, salaryStatementDomainObject)
	}
	return salaryStatementDomainObjectList, nil
}