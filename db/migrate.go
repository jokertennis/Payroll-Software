package db

import (
	"context"
	"database/sql"
	"fmt"
	"usr/local/go/server/gen/models"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/mattes/migrate/source/file"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// TODO:create as enum. only Develop and Test
type DbEnvironment struct {
	Environment string
}

func CreateDbInstance(dbEnvironment DbEnvironment) (*sql.DB, error) {
	switch dbEnvironment.Environment {
	case "Develop":
		// When we executed createDbInstance on local environment,localhost will be used.
		// On the other hand we executed createDbInstance inside container,db_container will be used.
		// Temporary change is not a problem.
		// We have to separate host_name depending on executed environment.
		db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/develop_db")
		if err != nil {
			return nil, fmt.Errorf("sql.Open error. err:%s", err)
		}
		return db, nil
	case "Test":
		db, err := sql.Open("mysql", "root:password@tcp(localhost:3306)/test_db")
		if err != nil {
			return nil, fmt.Errorf("sql.Open error. err:%s", err)
		}
		return db, nil
	}
	return nil, fmt.Errorf("not support specific DbEnviroment:%s", dbEnvironment)
}

// Reference : https://github.com/golang-migrate/migrate/tree/master/database/mysql#use-with-existing-client
func CreateMigrateInstance(db *sql.DB) (*migrate.Migrate, error) {
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to create driver. err:%s", err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://src/ddl",
		"mysql",
		driver,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create migration-instance. err:%s", err)
	}

	return m, nil
}

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

	if _, err := models.IndividualDeductionDetails().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of individual_deduction_detail. err:%s", err)
	}

	if _, err := models.IndividualDeductions().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of individual_deduction. err:%s", err)
	}

	if _, err := models.FixedEarningDetails().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of fixed_earning_detail. err:%s", err)
	}

	if _, err := models.FixedEarnings().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of fixed_earning. err:%s", err)
	}

	if _, err := models.IndividualEarningDetails().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of individual_earning_detail. err:%s", err)
	}

	if _, err := models.IndividualEarnings().DeleteAll(ctx, db); err != nil {
		return fmt.Errorf("failed to delete datas of individual_earning. err:%s", err)
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
