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

func testFixedDeductions(t *testing.T) {
	t.Parallel()

	query := FixedDeductions()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testFixedDeductionsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
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

	count, err := FixedDeductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFixedDeductionsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := FixedDeductions().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FixedDeductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFixedDeductionsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FixedDeductionSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := FixedDeductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testFixedDeductionsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := FixedDeductionExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if FixedDeduction exists: %s", err)
	}
	if !e {
		t.Errorf("Expected FixedDeductionExists to return true, but got false.")
	}
}

func testFixedDeductionsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	fixedDeductionFound, err := FindFixedDeduction(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if fixedDeductionFound == nil {
		t.Error("want a record, got nil")
	}
}

func testFixedDeductionsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = FixedDeductions().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testFixedDeductionsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := FixedDeductions().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testFixedDeductionsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	fixedDeductionOne := &FixedDeduction{}
	fixedDeductionTwo := &FixedDeduction{}
	if err = randomize.Struct(seed, fixedDeductionOne, fixedDeductionDBTypes, false, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}
	if err = randomize.Struct(seed, fixedDeductionTwo, fixedDeductionDBTypes, false, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = fixedDeductionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = fixedDeductionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FixedDeductions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testFixedDeductionsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	fixedDeductionOne := &FixedDeduction{}
	fixedDeductionTwo := &FixedDeduction{}
	if err = randomize.Struct(seed, fixedDeductionOne, fixedDeductionDBTypes, false, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}
	if err = randomize.Struct(seed, fixedDeductionTwo, fixedDeductionDBTypes, false, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = fixedDeductionOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = fixedDeductionTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FixedDeductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func fixedDeductionBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *FixedDeduction) error {
	*o = FixedDeduction{}
	return nil
}

func fixedDeductionAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *FixedDeduction) error {
	*o = FixedDeduction{}
	return nil
}

func fixedDeductionAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *FixedDeduction) error {
	*o = FixedDeduction{}
	return nil
}

func fixedDeductionBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FixedDeduction) error {
	*o = FixedDeduction{}
	return nil
}

func fixedDeductionAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *FixedDeduction) error {
	*o = FixedDeduction{}
	return nil
}

func fixedDeductionBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FixedDeduction) error {
	*o = FixedDeduction{}
	return nil
}

func fixedDeductionAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *FixedDeduction) error {
	*o = FixedDeduction{}
	return nil
}

func fixedDeductionBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FixedDeduction) error {
	*o = FixedDeduction{}
	return nil
}

func fixedDeductionAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *FixedDeduction) error {
	*o = FixedDeduction{}
	return nil
}

func testFixedDeductionsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &FixedDeduction{}
	o := &FixedDeduction{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FixedDeduction object: %s", err)
	}

	AddFixedDeductionHook(boil.BeforeInsertHook, fixedDeductionBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	fixedDeductionBeforeInsertHooks = []FixedDeductionHook{}

	AddFixedDeductionHook(boil.AfterInsertHook, fixedDeductionAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	fixedDeductionAfterInsertHooks = []FixedDeductionHook{}

	AddFixedDeductionHook(boil.AfterSelectHook, fixedDeductionAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	fixedDeductionAfterSelectHooks = []FixedDeductionHook{}

	AddFixedDeductionHook(boil.BeforeUpdateHook, fixedDeductionBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	fixedDeductionBeforeUpdateHooks = []FixedDeductionHook{}

	AddFixedDeductionHook(boil.AfterUpdateHook, fixedDeductionAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	fixedDeductionAfterUpdateHooks = []FixedDeductionHook{}

	AddFixedDeductionHook(boil.BeforeDeleteHook, fixedDeductionBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	fixedDeductionBeforeDeleteHooks = []FixedDeductionHook{}

	AddFixedDeductionHook(boil.AfterDeleteHook, fixedDeductionAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	fixedDeductionAfterDeleteHooks = []FixedDeductionHook{}

	AddFixedDeductionHook(boil.BeforeUpsertHook, fixedDeductionBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	fixedDeductionBeforeUpsertHooks = []FixedDeductionHook{}

	AddFixedDeductionHook(boil.AfterUpsertHook, fixedDeductionAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	fixedDeductionAfterUpsertHooks = []FixedDeductionHook{}
}

func testFixedDeductionsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FixedDeductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFixedDeductionsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(fixedDeductionColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := FixedDeductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testFixedDeductionToManyFixedDeductionDetails(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a FixedDeduction
	var b, c FixedDeductionDetail

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, fixedDeductionDetailDBTypes, false, fixedDeductionDetailColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, fixedDeductionDetailDBTypes, false, fixedDeductionDetailColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.FixedDeductionID = a.ID
	c.FixedDeductionID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.FixedDeductionDetails().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.FixedDeductionID == b.FixedDeductionID {
			bFound = true
		}
		if v.FixedDeductionID == c.FixedDeductionID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := FixedDeductionSlice{&a}
	if err = a.L.LoadFixedDeductionDetails(ctx, tx, false, (*[]*FixedDeduction)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FixedDeductionDetails); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.FixedDeductionDetails = nil
	if err = a.L.LoadFixedDeductionDetails(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.FixedDeductionDetails); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testFixedDeductionToManySalaryStatements(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a FixedDeduction
	var b, c SalaryStatement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
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

	queries.Assign(&b.FixedDeductionID, a.ID)
	queries.Assign(&c.FixedDeductionID, a.ID)
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
		if queries.Equal(v.FixedDeductionID, b.FixedDeductionID) {
			bFound = true
		}
		if queries.Equal(v.FixedDeductionID, c.FixedDeductionID) {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := FixedDeductionSlice{&a}
	if err = a.L.LoadSalaryStatements(ctx, tx, false, (*[]*FixedDeduction)(&slice), nil); err != nil {
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

func testFixedDeductionToManyAddOpFixedDeductionDetails(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a FixedDeduction
	var b, c, d, e FixedDeductionDetail

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, fixedDeductionDBTypes, false, strmangle.SetComplement(fixedDeductionPrimaryKeyColumns, fixedDeductionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*FixedDeductionDetail{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, fixedDeductionDetailDBTypes, false, strmangle.SetComplement(fixedDeductionDetailPrimaryKeyColumns, fixedDeductionDetailColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*FixedDeductionDetail{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddFixedDeductionDetails(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.FixedDeductionID {
			t.Error("foreign key was wrong value", a.ID, first.FixedDeductionID)
		}
		if a.ID != second.FixedDeductionID {
			t.Error("foreign key was wrong value", a.ID, second.FixedDeductionID)
		}

		if first.R.FixedDeduction != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.FixedDeduction != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.FixedDeductionDetails[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.FixedDeductionDetails[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.FixedDeductionDetails().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testFixedDeductionToManyAddOpSalaryStatements(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a FixedDeduction
	var b, c, d, e SalaryStatement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, fixedDeductionDBTypes, false, strmangle.SetComplement(fixedDeductionPrimaryKeyColumns, fixedDeductionColumnsWithoutDefault)...); err != nil {
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

		if !queries.Equal(a.ID, first.FixedDeductionID) {
			t.Error("foreign key was wrong value", a.ID, first.FixedDeductionID)
		}
		if !queries.Equal(a.ID, second.FixedDeductionID) {
			t.Error("foreign key was wrong value", a.ID, second.FixedDeductionID)
		}

		if first.R.FixedDeduction != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.FixedDeduction != &a {
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

func testFixedDeductionToManySetOpSalaryStatements(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a FixedDeduction
	var b, c, d, e SalaryStatement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, fixedDeductionDBTypes, false, strmangle.SetComplement(fixedDeductionPrimaryKeyColumns, fixedDeductionColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*SalaryStatement{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, salaryStatementDBTypes, false, strmangle.SetComplement(salaryStatementPrimaryKeyColumns, salaryStatementColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	err = a.SetSalaryStatements(ctx, tx, false, &b, &c)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.SalaryStatements().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	err = a.SetSalaryStatements(ctx, tx, true, &d, &e)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.SalaryStatements().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.FixedDeductionID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.FixedDeductionID) {
		t.Error("want c's foreign key value to be nil")
	}
	if !queries.Equal(a.ID, d.FixedDeductionID) {
		t.Error("foreign key was wrong value", a.ID, d.FixedDeductionID)
	}
	if !queries.Equal(a.ID, e.FixedDeductionID) {
		t.Error("foreign key was wrong value", a.ID, e.FixedDeductionID)
	}

	if b.R.FixedDeduction != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.FixedDeduction != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.FixedDeduction != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}
	if e.R.FixedDeduction != &a {
		t.Error("relationship was not added properly to the foreign struct")
	}

	if a.R.SalaryStatements[0] != &d {
		t.Error("relationship struct slice not set to correct value")
	}
	if a.R.SalaryStatements[1] != &e {
		t.Error("relationship struct slice not set to correct value")
	}
}

func testFixedDeductionToManyRemoveOpSalaryStatements(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a FixedDeduction
	var b, c, d, e SalaryStatement

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, fixedDeductionDBTypes, false, strmangle.SetComplement(fixedDeductionPrimaryKeyColumns, fixedDeductionColumnsWithoutDefault)...); err != nil {
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

	err = a.AddSalaryStatements(ctx, tx, true, foreigners...)
	if err != nil {
		t.Fatal(err)
	}

	count, err := a.SalaryStatements().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 4 {
		t.Error("count was wrong:", count)
	}

	err = a.RemoveSalaryStatements(ctx, tx, foreigners[:2]...)
	if err != nil {
		t.Fatal(err)
	}

	count, err = a.SalaryStatements().Count(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}
	if count != 2 {
		t.Error("count was wrong:", count)
	}

	if !queries.IsValuerNil(b.FixedDeductionID) {
		t.Error("want b's foreign key value to be nil")
	}
	if !queries.IsValuerNil(c.FixedDeductionID) {
		t.Error("want c's foreign key value to be nil")
	}

	if b.R.FixedDeduction != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if c.R.FixedDeduction != nil {
		t.Error("relationship was not removed properly from the foreign struct")
	}
	if d.R.FixedDeduction != &a {
		t.Error("relationship to a should have been preserved")
	}
	if e.R.FixedDeduction != &a {
		t.Error("relationship to a should have been preserved")
	}

	if len(a.R.SalaryStatements) != 2 {
		t.Error("should have preserved two relationships")
	}

	// Removal doesn't do a stable deletion for performance so we have to flip the order
	if a.R.SalaryStatements[1] != &d {
		t.Error("relationship to d should have been preserved")
	}
	if a.R.SalaryStatements[0] != &e {
		t.Error("relationship to e should have been preserved")
	}
}

func testFixedDeductionsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
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

func testFixedDeductionsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := FixedDeductionSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testFixedDeductionsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := FixedDeductions().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	fixedDeductionDBTypes = map[string]string{`ID`: `mediumint`, `Amount`: `int`, `Nominal`: `varchar`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`}
	_                     = bytes.MinRead
)

func testFixedDeductionsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(fixedDeductionPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(fixedDeductionAllColumns) == len(fixedDeductionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FixedDeductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testFixedDeductionsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(fixedDeductionAllColumns) == len(fixedDeductionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &FixedDeduction{}
	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := FixedDeductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, fixedDeductionDBTypes, true, fixedDeductionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(fixedDeductionAllColumns, fixedDeductionPrimaryKeyColumns) {
		fields = fixedDeductionAllColumns
	} else {
		fields = strmangle.SetComplement(
			fixedDeductionAllColumns,
			fixedDeductionPrimaryKeyColumns,
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

	slice := FixedDeductionSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testFixedDeductionsUpsert(t *testing.T) {
	t.Parallel()

	if len(fixedDeductionAllColumns) == len(fixedDeductionPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLFixedDeductionUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := FixedDeduction{}
	if err = randomize.Struct(seed, &o, fixedDeductionDBTypes, false); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FixedDeduction: %s", err)
	}

	count, err := FixedDeductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, fixedDeductionDBTypes, false, fixedDeductionPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize FixedDeduction struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert FixedDeduction: %s", err)
	}

	count, err = FixedDeductions().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}