// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

func testEarningDetails(t *testing.T) {
	t.Parallel()

	query := EarningDetails()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testEarningDetailsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
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

	count, err := EarningDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEarningDetailsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := EarningDetails().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EarningDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEarningDetailsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EarningDetailSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := EarningDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testEarningDetailsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := EarningDetailExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if EarningDetail exists: %s", err)
	}
	if !e {
		t.Errorf("Expected EarningDetailExists to return true, but got false.")
	}
}

func testEarningDetailsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	earningDetailFound, err := FindEarningDetail(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if earningDetailFound == nil {
		t.Error("want a record, got nil")
	}
}

func testEarningDetailsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = EarningDetails().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testEarningDetailsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := EarningDetails().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testEarningDetailsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	earningDetailOne := &EarningDetail{}
	earningDetailTwo := &EarningDetail{}
	if err = randomize.Struct(seed, earningDetailOne, earningDetailDBTypes, false, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}
	if err = randomize.Struct(seed, earningDetailTwo, earningDetailDBTypes, false, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = earningDetailOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = earningDetailTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EarningDetails().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testEarningDetailsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	earningDetailOne := &EarningDetail{}
	earningDetailTwo := &EarningDetail{}
	if err = randomize.Struct(seed, earningDetailOne, earningDetailDBTypes, false, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}
	if err = randomize.Struct(seed, earningDetailTwo, earningDetailDBTypes, false, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = earningDetailOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = earningDetailTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EarningDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func earningDetailBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *EarningDetail) error {
	*o = EarningDetail{}
	return nil
}

func earningDetailAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *EarningDetail) error {
	*o = EarningDetail{}
	return nil
}

func earningDetailAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *EarningDetail) error {
	*o = EarningDetail{}
	return nil
}

func earningDetailBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EarningDetail) error {
	*o = EarningDetail{}
	return nil
}

func earningDetailAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *EarningDetail) error {
	*o = EarningDetail{}
	return nil
}

func earningDetailBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EarningDetail) error {
	*o = EarningDetail{}
	return nil
}

func earningDetailAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *EarningDetail) error {
	*o = EarningDetail{}
	return nil
}

func earningDetailBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EarningDetail) error {
	*o = EarningDetail{}
	return nil
}

func earningDetailAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *EarningDetail) error {
	*o = EarningDetail{}
	return nil
}

func testEarningDetailsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &EarningDetail{}
	o := &EarningDetail{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, earningDetailDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EarningDetail object: %s", err)
	}

	AddEarningDetailHook(boil.BeforeInsertHook, earningDetailBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	earningDetailBeforeInsertHooks = []EarningDetailHook{}

	AddEarningDetailHook(boil.AfterInsertHook, earningDetailAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	earningDetailAfterInsertHooks = []EarningDetailHook{}

	AddEarningDetailHook(boil.AfterSelectHook, earningDetailAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	earningDetailAfterSelectHooks = []EarningDetailHook{}

	AddEarningDetailHook(boil.BeforeUpdateHook, earningDetailBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	earningDetailBeforeUpdateHooks = []EarningDetailHook{}

	AddEarningDetailHook(boil.AfterUpdateHook, earningDetailAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	earningDetailAfterUpdateHooks = []EarningDetailHook{}

	AddEarningDetailHook(boil.BeforeDeleteHook, earningDetailBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	earningDetailBeforeDeleteHooks = []EarningDetailHook{}

	AddEarningDetailHook(boil.AfterDeleteHook, earningDetailAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	earningDetailAfterDeleteHooks = []EarningDetailHook{}

	AddEarningDetailHook(boil.BeforeUpsertHook, earningDetailBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	earningDetailBeforeUpsertHooks = []EarningDetailHook{}

	AddEarningDetailHook(boil.AfterUpsertHook, earningDetailAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	earningDetailAfterUpsertHooks = []EarningDetailHook{}
}

func testEarningDetailsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EarningDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEarningDetailsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(earningDetailColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := EarningDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testEarningDetailToOneEarningUsingEarning(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local EarningDetail
	var foreign Earning

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, earningDetailDBTypes, false, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, earningDBTypes, false, earningColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Earning struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.EarningID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Earning().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddEarningHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Earning) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := EarningDetailSlice{&local}
	if err = local.L.LoadEarning(ctx, tx, false, (*[]*EarningDetail)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Earning == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Earning = nil
	if err = local.L.LoadEarning(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Earning == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testEarningDetailToOneSetOpEarningUsingEarning(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a EarningDetail
	var b, c Earning

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, earningDetailDBTypes, false, strmangle.SetComplement(earningDetailPrimaryKeyColumns, earningDetailColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, earningDBTypes, false, strmangle.SetComplement(earningPrimaryKeyColumns, earningColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, earningDBTypes, false, strmangle.SetComplement(earningPrimaryKeyColumns, earningColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Earning{&b, &c} {
		err = a.SetEarning(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Earning != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.EarningDetails[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.EarningID != x.ID {
			t.Error("foreign key was wrong value", a.EarningID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.EarningID))
		reflect.Indirect(reflect.ValueOf(&a.EarningID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.EarningID != x.ID {
			t.Error("foreign key was wrong value", a.EarningID, x.ID)
		}
	}
}

func testEarningDetailsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
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

func testEarningDetailsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := EarningDetailSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testEarningDetailsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := EarningDetails().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	earningDetailDBTypes = map[string]string{`ID`: `mediumint`, `EarningID`: `mediumint`, `Nominal`: `varchar`, `Amount`: `int`, `CreatedAt`: `timestamp`, `UpdatedAt`: `timestamp`}
	_                    = bytes.MinRead
)

func testEarningDetailsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(earningDetailPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(earningDetailAllColumns) == len(earningDetailPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EarningDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testEarningDetailsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(earningDetailAllColumns) == len(earningDetailPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &EarningDetail{}
	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := EarningDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, earningDetailDBTypes, true, earningDetailPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(earningDetailAllColumns, earningDetailPrimaryKeyColumns) {
		fields = earningDetailAllColumns
	} else {
		fields = strmangle.SetComplement(
			earningDetailAllColumns,
			earningDetailPrimaryKeyColumns,
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

	slice := EarningDetailSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testEarningDetailsUpsert(t *testing.T) {
	t.Parallel()

	if len(earningDetailAllColumns) == len(earningDetailPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}
	if len(mySQLEarningDetailUniqueColumns) == 0 {
		t.Skip("Skipping table with no unique columns to conflict on")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := EarningDetail{}
	if err = randomize.Struct(seed, &o, earningDetailDBTypes, false); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EarningDetail: %s", err)
	}

	count, err := EarningDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, earningDetailDBTypes, false, earningDetailPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize EarningDetail struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert EarningDetail: %s", err)
	}

	count, err = EarningDetails().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
