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

func testDeductions(t *testing.T) {
	t.Parallel()

	query := Deductions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDeductionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
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

	count, err := Deductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDeductionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Deductions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Deductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDeductionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DeductionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Deductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDeductionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DeductionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Deduction exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DeductionExists to return true, but got false.")
	}
}

func testDeductionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	deductionFound, err := FindDeduction(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if deductionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDeductionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Deductions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDeductionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Deductions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDeductionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	deductionOne := &Deduction{}
	deductionTwo := &Deduction{}
	if err = randomize.Struct(seed, deductionOne, deductionDBTypes, false, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}
	if err = randomize.Struct(seed, deductionTwo, deductionDBTypes, false, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = deductionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = deductionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Deductions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDeductionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	deductionOne := &Deduction{}
	deductionTwo := &Deduction{}
	if err = randomize.Struct(seed, deductionOne, deductionDBTypes, false, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}
	if err = randomize.Struct(seed, deductionTwo, deductionDBTypes, false, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = deductionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = deductionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Deductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func deductionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Deduction) error {
	*o = Deduction{}
	return nil
}

func deductionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Deduction) error {
	*o = Deduction{}
	return nil
}

func deductionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Deduction) error {
	*o = Deduction{}
	return nil
}

func deductionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Deduction) error {
	*o = Deduction{}
	return nil
}

func deductionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Deduction) error {
	*o = Deduction{}
	return nil
}

func deductionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Deduction) error {
	*o = Deduction{}
	return nil
}

func deductionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Deduction) error {
	*o = Deduction{}
	return nil
}

func deductionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Deduction) error {
	*o = Deduction{}
	return nil
}

func deductionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Deduction) error {
	*o = Deduction{}
	return nil
}

func testDeductionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Deduction{}
	o := &Deduction{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, deductionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Deduction object: %s", err)
	}

	AddDeductionHook(boil.BeforeInsertHook, deductionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	deductionBeforeInsertHooks = []DeductionHook{}

	AddDeductionHook(boil.AfterInsertHook, deductionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	deductionAfterInsertHooks = []DeductionHook{}

	AddDeductionHook(boil.AfterSelectHook, deductionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	deductionAfterSelectHooks = []DeductionHook{}

	AddDeductionHook(boil.BeforeUpdateHook, deductionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	deductionBeforeUpdateHooks = []DeductionHook{}

	AddDeductionHook(boil.AfterUpdateHook, deductionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	deductionAfterUpdateHooks = []DeductionHook{}

	AddDeductionHook(boil.BeforeDeleteHook, deductionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	deductionBeforeDeleteHooks = []DeductionHook{}

	AddDeductionHook(boil.AfterDeleteHook, deductionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	deductionAfterDeleteHooks = []DeductionHook{}

	AddDeductionHook(boil.BeforeUpsertHook, deductionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	deductionBeforeUpsertHooks = []DeductionHook{}

	AddDeductionHook(boil.AfterUpsertHook, deductionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	deductionAfterUpsertHooks = []DeductionHook{}
}

func testDeductionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Deductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDeductionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(deductionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Deductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDeductionOneToOneSalaryStatementUsingSalaryStatement(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var foreign SalaryStatement
	var local Deduction

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &foreign, salaryStatementDBTypes, true, salaryStatementColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize SalaryStatement struct: %s", err)
	}
	if err := randomize.Struct(seed, &local, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&foreign.DeductionID, local.ID)
	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.SalaryStatement().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.DeductionID, foreign.DeductionID) {
		t.Errorf("want: %v, got %v", foreign.DeductionID, check.DeductionID)
	}

	ranAfterSelectHook := false
	AddSalaryStatementHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *SalaryStatement) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := DeductionSlice{&local}
	if err = local.L.LoadSalaryStatement(ctx, tx, false, (*[]*Deduction)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SalaryStatement == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.SalaryStatement = nil
	if err = local.L.LoadSalaryStatement(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.SalaryStatement == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testDeductionOneToOneSetOpSalaryStatementUsingSalaryStatement(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Deduction
	var b, c SalaryStatement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, deductionDBTypes, false, strmangle.SetComplement(deductionPrimaryKeyColumns, deductionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, salaryStatementDBTypes, false, strmangle.SetComplement(salaryStatementPrimaryKeyColumns, salaryStatementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, salaryStatementDBTypes, false, strmangle.SetComplement(salaryStatementPrimaryKeyColumns, salaryStatementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*SalaryStatement{&b, &c} {
		err = a.SetSalaryStatement(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.SalaryStatement != x {
			t.Error("relationship struct not set to correct value")
		}
		if x.R.Deduction != &a {
			t.Error("failed to append to foreign relationship struct")
		}

		if !queries.Equal(a.ID, x.DeductionID) {
			t.Error("foreign key was wrong value", a.ID)
		}

		zero := reflect.Zero(reflect.TypeOf(x.DeductionID))
		reflect.Indirect(reflect.ValueOf(&x.DeductionID)).Set(zero)

		if err = x.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.ID, x.DeductionID) {
			t.Error("foreign key was wrong value", a.ID, x.DeductionID)
		}

		if _, err = x.Delete(ctx, tx); err != nil {
			t.Fatal("failed to delete x", err)
		}
	}
}

func testDeductionOneToOneRemoveOpSalaryStatementUsingSalaryStatement(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Deduction
	var b SalaryStatement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, deductionDBTypes, false, strmangle.SetComplement(deductionPrimaryKeyColumns, deductionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, salaryStatementDBTypes, false, strmangle.SetComplement(salaryStatementPrimaryKeyColumns, salaryStatementColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetSalaryStatement(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveSalaryStatement(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.SalaryStatement().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.SalaryStatement != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(b.DeductionID) {
		t.Error("foreign key column should be nil")
	}

	if b.R.Deduction != nil {
		t.Error("failed to remove a from b's relationships")
	}
}

func testDeductionToManyDeductionDetails(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Deduction
	var b, c DeductionDetail

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, deductionDetailDBTypes, false, deductionDetailColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, deductionDetailDBTypes, false, deductionDetailColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.DeductionID = a.ID
	c.DeductionID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.DeductionDetails().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.DeductionID == b.DeductionID {
			bFound = true
		}
		if v.DeductionID == c.DeductionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DeductionSlice{&a}
	if err = a.L.LoadDeductionDetails(ctx, tx, false, (*[]*Deduction)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.DeductionDetails); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.DeductionDetails = nil
	if err = a.L.LoadDeductionDetails(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.DeductionDetails); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDeductionToManyAddOpDeductionDetails(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Deduction
	var b, c, d, e DeductionDetail

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, deductionDBTypes, false, strmangle.SetComplement(deductionPrimaryKeyColumns, deductionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*DeductionDetail{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, deductionDetailDBTypes, false, strmangle.SetComplement(deductionDetailPrimaryKeyColumns, deductionDetailColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*DeductionDetail{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddDeductionDetails(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.DeductionID {
			t.Error("foreign key was wrong value", a.ID, first.DeductionID)
		}
		if a.ID != second.DeductionID {
			t.Error("foreign key was wrong value", a.ID, second.DeductionID)
		}

		if first.R.Deduction != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Deduction != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.DeductionDetails[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.DeductionDetails[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.DeductionDetails().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}

func testDeductionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
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

func testDeductionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DeductionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDeductionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Deductions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	deductionDBTypes = map[string]string{`ID`: `mediumint`, `Amount`: `int`, `Nominal`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`}
	_                = bytes.MinRead
)

func testDeductionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(deductionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(deductionAllColumns) == len(deductionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Deductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDeductionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(deductionAllColumns) == len(deductionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Deduction{}
	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Deductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, deductionDBTypes, true, deductionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(deductionAllColumns, deductionPrimaryKeyColumns) {
		fields = deductionAllColumns
	} else {
		fields = strmangle.SetComplement(
			deductionAllColumns,
			deductionPrimaryKeyColumns,
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

	slice := DeductionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testDeductionsUpsert(t *testing.T) {
	t.Parallel()

	if len(deductionAllColumns) == len(deductionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLDeductionUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Deduction{}
	if err = randomize.Struct(seed, &o, deductionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Deduction: %s", err)
	}

	count, err := Deductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, deductionDBTypes, false, deductionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Deduction struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Deduction: %s", err)
	}

	count, err = Deductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
