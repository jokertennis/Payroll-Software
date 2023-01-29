// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testEmployees(t *testing.T) {
	t.Parallel()

	query := Employees()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEmployeesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Employees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmployeesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Employees().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Employees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmployeesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmployeeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Employees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEmployeesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EmployeeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Employee exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EmployeeExists to return true, but got false.")
	}
}

func testEmployeesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	employeeFound, err := FindEmployee(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if employeeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEmployeesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Employees().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEmployeesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Employees().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEmployeesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	employeeOne := &Employee{}
	employeeTwo := &Employee{}
	if err = randomize.Struct(seed, employeeOne, employeeDBTypes, false, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}
	if err = randomize.Struct(seed, employeeTwo, employeeDBTypes, false, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = employeeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = employeeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Employees().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEmployeesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	employeeOne := &Employee{}
	employeeTwo := &Employee{}
	if err = randomize.Struct(seed, employeeOne, employeeDBTypes, false, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}
	if err = randomize.Struct(seed, employeeTwo, employeeDBTypes, false, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = employeeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = employeeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Employees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func employeeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Employee) error {
	*o = Employee{}
	return nil
}

func employeeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Employee) error {
	*o = Employee{}
	return nil
}

func employeeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Employee) error {
	*o = Employee{}
	return nil
}

func employeeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Employee) error {
	*o = Employee{}
	return nil
}

func employeeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Employee) error {
	*o = Employee{}
	return nil
}

func employeeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Employee) error {
	*o = Employee{}
	return nil
}

func employeeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Employee) error {
	*o = Employee{}
	return nil
}

func employeeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Employee) error {
	*o = Employee{}
	return nil
}

func employeeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Employee) error {
	*o = Employee{}
	return nil
}

func testEmployeesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Employee{}
	o := &Employee{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, employeeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Employee object: %s", err)
	}

	AddEmployeeHook(boil.BeforeInsertHook, employeeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	employeeBeforeInsertHooks = []EmployeeHook{}

	AddEmployeeHook(boil.AfterInsertHook, employeeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	employeeAfterInsertHooks = []EmployeeHook{}

	AddEmployeeHook(boil.AfterSelectHook, employeeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	employeeAfterSelectHooks = []EmployeeHook{}

	AddEmployeeHook(boil.BeforeUpdateHook, employeeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	employeeBeforeUpdateHooks = []EmployeeHook{}

	AddEmployeeHook(boil.AfterUpdateHook, employeeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	employeeAfterUpdateHooks = []EmployeeHook{}

	AddEmployeeHook(boil.BeforeDeleteHook, employeeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	employeeBeforeDeleteHooks = []EmployeeHook{}

	AddEmployeeHook(boil.AfterDeleteHook, employeeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	employeeAfterDeleteHooks = []EmployeeHook{}

	AddEmployeeHook(boil.BeforeUpsertHook, employeeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	employeeBeforeUpsertHooks = []EmployeeHook{}

	AddEmployeeHook(boil.AfterUpsertHook, employeeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	employeeAfterUpsertHooks = []EmployeeHook{}
}

func testEmployeesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Employees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmployeesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(employeeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Employees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEmployeeToManySalaryStatements(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Employee
	var b, c SalaryStatement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, salaryStatementDBTypes, false, salaryStatementColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, salaryStatementDBTypes, false, salaryStatementColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.EmployeeID = a.ID
	c.EmployeeID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.SalaryStatements().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.EmployeeID == b.EmployeeID {
			bFound = true
		}
		if v.EmployeeID == c.EmployeeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := EmployeeSlice{&a}
	if err = a.L.LoadSalaryStatements(ctx, tx, false, (*[]*Employee)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SalaryStatements); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.SalaryStatements = nil
	if err = a.L.LoadSalaryStatements(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.SalaryStatements); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testEmployeeToManyAddOpSalaryStatements(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Employee
	var b, c, d, e SalaryStatement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, employeeDBTypes, false, strmangle.SetComplement(employeePrimaryKeyColumns, employeeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*SalaryStatement{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, salaryStatementDBTypes, false, strmangle.SetComplement(salaryStatementPrimaryKeyColumns, salaryStatementColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*SalaryStatement{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddSalaryStatements(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.EmployeeID {
			t.Error("foreign key was wrong value", a.ID, first.EmployeeID)
		}
		if a.ID != second.EmployeeID {
			t.Error("foreign key was wrong value", a.ID, second.EmployeeID)
		}

		if first.R.Employee != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Employee != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.SalaryStatements[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.SalaryStatements[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.SalaryStatements().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testEmployeeToOneCompanyUsingCompany(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Employee
	var foreign Company

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, employeeDBTypes, false, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, companyDBTypes, false, companyColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Company struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.CompanyID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Company().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddCompanyHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Company) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := EmployeeSlice{&local}
	if err = local.L.LoadCompany(ctx, tx, false, (*[]*Employee)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Company == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Company = nil
	if err = local.L.LoadCompany(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Company == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testEmployeeToOneSetOpCompanyUsingCompany(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Employee
	var b, c Company

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, employeeDBTypes, false, strmangle.SetComplement(employeePrimaryKeyColumns, employeeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, companyDBTypes, false, strmangle.SetComplement(companyPrimaryKeyColumns, companyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, companyDBTypes, false, strmangle.SetComplement(companyPrimaryKeyColumns, companyColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Company{&b, &c} {
		err = a.SetCompany(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Company != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Employees[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.CompanyID != x.ID {
			t.Error("foreign key was wrong value", a.CompanyID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.CompanyID))
		reflect.Indirect(reflect.ValueOf(&a.CompanyID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.CompanyID != x.ID {
			t.Error("foreign key was wrong value", a.CompanyID, x.ID)
		}
	}
}

func testEmployeesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEmployeesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EmployeeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEmployeesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Employees().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	employeeDBTypes = map[string]string{`ID`: `mediumint`, `CompanyID`: `smallint`, `Name`: `varchar`, `MailAddress`: `varchar`, `Password`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`}
	_               = bytes.MinRead
)

func testEmployeesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(employeePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(employeeAllColumns) == len(employeePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Employees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEmployeesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(employeeAllColumns) == len(employeePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Employee{}
	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Employees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, employeeDBTypes, true, employeePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(employeeAllColumns, employeePrimaryKeyColumns) {
		fields = employeeAllColumns
	} else {
		fields = strmangle.SetComplement(
			employeeAllColumns,
			employeePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := EmployeeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEmployeesUpsert(t *testing.T) {
	t.Parallel()

	if len(employeeAllColumns) == len(employeePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLEmployeeUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Employee{}
	if err = randomize.Struct(seed, &o, employeeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Employee: %s", err)
	}

	count, err := Employees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, employeeDBTypes, false, employeePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Employee struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Employee: %s", err)
	}

	count, err = Employees().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
