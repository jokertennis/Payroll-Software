package infrastructure

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"usr/local/go/server/gen/models"
	"usr/local/go/src/main/domain-model/salary_statement"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
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

func (r *SalaryStatementRepository) GetAllSalaryStatements(employeeId uint32) ([]*salary_statement.SalaryStatement, error) {
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

	var salaryStatementDomainObjectList []*salary_statement.SalaryStatement 
	for _, salaryStatement := range salaryStatements {
		salaryStatementDomainObject, err := MappingSalaryStatementDomainObject(salaryStatement)
		if err != nil {
			return nil, err
		}
		salaryStatementDomainObjectList = append(salaryStatementDomainObjectList, salaryStatementDomainObject)
	}
	return salaryStatementDomainObjectList, nil
}

func (r *SalaryStatementRepository) CreateSalaryStatement(salaryStatementEntry salary_statement.SalaryStatementEntryByUsingIndividualDatas) (salaryStatementId uint32, err error) {
	tx, err := r.db.BeginTx(r.ctx, nil)
	if err != nil {
		return 0, fmt.Errorf("failed to create tx object.error:%s", err)
	}
	newIndividualEarning := &models.IndividualEarning{
		Nominal: salaryStatementEntry.IndividualEarning.Nominal,
		Amount: salaryStatementEntry.IndividualEarning.Amount,
	}
	err = newIndividualEarning.Insert(r.ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to create individualEarning.error:%s", err)
	}

	for _, value := range salaryStatementEntry.IndividualEarning.IndividualEarningDetails {
		newIndividualEarningDetail := &models.IndividualEarningDetail{
			IndividualEarningID: newIndividualEarning.ID,
			Nominal: value.Nominal,
			Amount: value.Amount,
		}
		err = newIndividualEarningDetail.Insert(r.ctx, tx, boil.Infer())
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("failed to create individualEarningDetail.error:%s", err)
		}
	}

	newIndividualDeduction := &models.IndividualDeduction{
		Nominal: salaryStatementEntry.IndividualDeduction.Nominal,
		Amount: salaryStatementEntry.IndividualDeduction.Amount,
	}
	err = newIndividualDeduction.Insert(r.ctx, tx, boil.Infer())
	if err != nil {
		tx.Rollback()
		return 0, fmt.Errorf("failed to create individualDeduction.error:%s", err)
	}

	for _, value := range salaryStatementEntry.IndividualDeduction.IndividualDeductionDetails {
		newIndividualDeductionDetail := &models.IndividualDeductionDetail{
			IndividualDeductionID: newIndividualDeduction.ID,
			Nominal: value.Nominal,
			Amount: value.Amount,
		}
		err = newIndividualDeductionDetail.Insert(r.ctx, tx, boil.Infer())
		if err != nil {
			tx.Rollback()
			return 0, fmt.Errorf("failed to create individualDeductionDetail.error:%s", err)
		}
	}

	newSalaryStatement := &models.SalaryStatement{
		IndividualEarningID: null.Uint32From(newIndividualEarning.ID),
		IndividualDeductionID: null.Uint32From(newIndividualDeduction.ID),
		EmployeeID: salaryStatementEntry.EmployeeId,
		Nominal: salaryStatementEntry.Nominal,
		Payday: salaryStatementEntry.Payday,
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