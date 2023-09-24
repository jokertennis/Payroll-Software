package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"github.com/jokertennis/Payroll-Software/server/gen/models"
	salary_statement_domain_model "github.com/jokertennis/Payroll-Software/src/main/domain-model/salary_statement"
	"github.com/jokertennis/Payroll-Software/src/main/domain-service/repository/salary_statement_repository"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type SalaryStatementRepository struct {
	ctx context.Context
	db  *sql.DB
}

func NewSalaryStatementRepository(ctx context.Context, db *sql.DB) SalaryStatementRepository {
	return SalaryStatementRepository{
		ctx: ctx,
		db:  db,
	}
}

func (r *SalaryStatementRepository) GetSalaryStatement(employeeId uint32, yearOfPayday int, monthOfPayday time.Month) (*salary_statement_domain_model.SalaryStatement, error) {
	start := time.Date(yearOfPayday, monthOfPayday, 1, 0, 0, 0, 0, time.UTC) // 年月の最初の日時
	end := start.AddDate(0, 1, 0)                                            // 年月の最後の日時
	salaryStatements, err := models.SalaryStatements(
		qm.Load(models.SalaryStatementRels.Earning),
		qm.Load(models.SalaryStatementRels.Deduction),
		qm.Load(models.SalaryStatementRels.Earning+"."+models.EarningRels.EarningDetails),
		qm.Load(models.SalaryStatementRels.Deduction+"."+models.DeductionRels.DeductionDetails),
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
		qm.Load(models.SalaryStatementRels.Earning),
		qm.Load(models.SalaryStatementRels.Deduction),
		qm.Load(models.SalaryStatementRels.Earning+"."+models.EarningRels.EarningDetails),
		qm.Load(models.SalaryStatementRels.Deduction+"."+models.DeductionRels.DeductionDetails),
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

func (r *SalaryStatementRepository) CreateSalaryStatement(salaryStatementEntry salary_statement_repository.SalaryStatementEntry) (salaryStatementId uint32, err error) {
	tx, err := r.db.BeginTx(r.ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create tx object.error:%s", err)
	}
	newEarning := &models.Earning{
		Nominal:     salaryStatementEntry.EarningEntry.Nominal,
		Amount:      salaryStatementEntry.EarningEntry.Amount,
		EarningType: string(salaryStatementEntry.EarningEntry.EarningType),
	}
	err = newEarning.Insert(r.ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to create earning.error:%s", err)
	}

	for _, value := range salaryStatementEntry.EarningEntry.EarningDetailsEntry {
		newEarningDetail := &models.EarningDetail{
			EarningID: newEarning.ID,
			Nominal:   value.Nominal,
			Amount:    value.Amount,
		}
		err = newEarningDetail.Insert(r.ctx, tx, boil.Infer())
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("failed to create earningDetail.error:%s", err)
		}
	}

	newDeduction := &models.Deduction{
		Nominal:       salaryStatementEntry.DeductionEntry.Nominal,
		Amount:        salaryStatementEntry.DeductionEntry.Amount,
		DeductionType: string(salaryStatementEntry.DeductionEntry.DeductionType),
	}
	err = newDeduction.Insert(r.ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to create deduction.error:%s", err)
	}

	for _, value := range salaryStatementEntry.DeductionEntry.DeductionDetailsEntry {
		newDeductionDetail := &models.DeductionDetail{
			DeductionID: newDeduction.ID,
			Nominal:     value.Nominal,
			Amount:      value.Amount,
		}
		err = newDeductionDetail.Insert(r.ctx, tx, boil.Infer())
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("failed to create deductionDetail.error:%s", err)
		}
	}

	newSalaryStatement := &models.SalaryStatement{
		EarningID:    newEarning.ID,
		DeductionID:  newDeduction.ID,
		EmployeeID:   salaryStatementEntry.EmployeeId,
		Nominal:      salaryStatementEntry.Nominal,
		Payday:       salaryStatementEntry.Payday,
		TargetPeriod: salaryStatementEntry.TargetPeriod,
	}
	err = newSalaryStatement.Insert(r.ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to create salaryStatement.error:%s", err)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return newSalaryStatement.ID, nil
}
