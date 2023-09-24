// Code generated by SQLBoiler 4.15.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Earning is an object representing the database table.
type Earning struct {
	ID          uint32    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Nominal     string    `boil:"nominal" json:"nominal" toml:"nominal" yaml:"nominal"`
	Amount      int       `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	EarningType string    `boil:"earning_type" json:"earning_type" toml:"earning_type" yaml:"earning_type"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *earningR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L earningL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EarningColumns = struct {
	ID          string
	Nominal     string
	Amount      string
	EarningType string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	Nominal:     "nominal",
	Amount:      "amount",
	EarningType: "earning_type",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var EarningTableColumns = struct {
	ID          string
	Nominal     string
	Amount      string
	EarningType string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "earnings.id",
	Nominal:     "earnings.nominal",
	Amount:      "earnings.amount",
	EarningType: "earnings.earning_type",
	CreatedAt:   "earnings.created_at",
	UpdatedAt:   "earnings.updated_at",
}

// Generated where

var EarningWhere = struct {
	ID          whereHelperuint32
	Nominal     whereHelperstring
	Amount      whereHelperint
	EarningType whereHelperstring
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
}{
	ID:          whereHelperuint32{field: "`earnings`.`id`"},
	Nominal:     whereHelperstring{field: "`earnings`.`nominal`"},
	Amount:      whereHelperint{field: "`earnings`.`amount`"},
	EarningType: whereHelperstring{field: "`earnings`.`earning_type`"},
	CreatedAt:   whereHelpertime_Time{field: "`earnings`.`created_at`"},
	UpdatedAt:   whereHelpertime_Time{field: "`earnings`.`updated_at`"},
}

// EarningRels is where relationship names are stored.
var EarningRels = struct {
	EarningDetails   string
	SalaryStatements string
}{
	EarningDetails:   "EarningDetails",
	SalaryStatements: "SalaryStatements",
}

// earningR is where relationships are stored.
type earningR struct {
	EarningDetails   EarningDetailSlice   `boil:"EarningDetails" json:"EarningDetails" toml:"EarningDetails" yaml:"EarningDetails"`
	SalaryStatements SalaryStatementSlice `boil:"SalaryStatements" json:"SalaryStatements" toml:"SalaryStatements" yaml:"SalaryStatements"`
}

// NewStruct creates a new relationship struct
func (*earningR) NewStruct() *earningR {
	return &earningR{}
}

func (r *earningR) GetEarningDetails() EarningDetailSlice {
	if r == nil {
		return nil
	}
	return r.EarningDetails
}

func (r *earningR) GetSalaryStatements() SalaryStatementSlice {
	if r == nil {
		return nil
	}
	return r.SalaryStatements
}

// earningL is where Load methods for each relationship are stored.
type earningL struct{}

var (
	earningAllColumns            = []string{"id", "nominal", "amount", "earning_type", "created_at", "updated_at"}
	earningColumnsWithoutDefault = []string{"nominal", "amount", "earning_type", "created_at", "updated_at"}
	earningColumnsWithDefault    = []string{"id"}
	earningPrimaryKeyColumns     = []string{"id"}
	earningGeneratedColumns      = []string{}
)

type (
	// EarningSlice is an alias for a slice of pointers to Earning.
	// This should almost always be used instead of []Earning.
	EarningSlice []*Earning
	// EarningHook is the signature for custom Earning hook methods
	EarningHook func(context.Context, boil.ContextExecutor, *Earning) error

	earningQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	earningType                 = reflect.TypeOf(&Earning{})
	earningMapping              = queries.MakeStructMapping(earningType)
	earningPrimaryKeyMapping, _ = queries.BindMapping(earningType, earningMapping, earningPrimaryKeyColumns)
	earningInsertCacheMut       sync.RWMutex
	earningInsertCache          = make(map[string]insertCache)
	earningUpdateCacheMut       sync.RWMutex
	earningUpdateCache          = make(map[string]updateCache)
	earningUpsertCacheMut       sync.RWMutex
	earningUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var earningAfterSelectHooks []EarningHook

var earningBeforeInsertHooks []EarningHook
var earningAfterInsertHooks []EarningHook

var earningBeforeUpdateHooks []EarningHook
var earningAfterUpdateHooks []EarningHook

var earningBeforeDeleteHooks []EarningHook
var earningAfterDeleteHooks []EarningHook

var earningBeforeUpsertHooks []EarningHook
var earningAfterUpsertHooks []EarningHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Earning) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Earning) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Earning) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Earning) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Earning) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Earning) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Earning) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Earning) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Earning) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEarningHook registers your hook function for all future operations.
func AddEarningHook(hookPoint boil.HookPoint, earningHook EarningHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		earningAfterSelectHooks = append(earningAfterSelectHooks, earningHook)
	case boil.BeforeInsertHook:
		earningBeforeInsertHooks = append(earningBeforeInsertHooks, earningHook)
	case boil.AfterInsertHook:
		earningAfterInsertHooks = append(earningAfterInsertHooks, earningHook)
	case boil.BeforeUpdateHook:
		earningBeforeUpdateHooks = append(earningBeforeUpdateHooks, earningHook)
	case boil.AfterUpdateHook:
		earningAfterUpdateHooks = append(earningAfterUpdateHooks, earningHook)
	case boil.BeforeDeleteHook:
		earningBeforeDeleteHooks = append(earningBeforeDeleteHooks, earningHook)
	case boil.AfterDeleteHook:
		earningAfterDeleteHooks = append(earningAfterDeleteHooks, earningHook)
	case boil.BeforeUpsertHook:
		earningBeforeUpsertHooks = append(earningBeforeUpsertHooks, earningHook)
	case boil.AfterUpsertHook:
		earningAfterUpsertHooks = append(earningAfterUpsertHooks, earningHook)
	}
}

// One returns a single earning record from the query.
func (q earningQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Earning, error) {
	o := &Earning{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for earnings")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Earning records from the query.
func (q earningQuery) All(ctx context.Context, exec boil.ContextExecutor) (EarningSlice, error) {
	var o []*Earning

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Earning slice")
	}

	if len(earningAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Earning records in the query.
func (q earningQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count earnings rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q earningQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if earnings exists")
	}

	return count > 0, nil
}

// EarningDetails retrieves all the earning_detail's EarningDetails with an executor.
func (o *Earning) EarningDetails(mods ...qm.QueryMod) earningDetailQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`earning_details`.`earning_id`=?", o.ID),
	)

	return EarningDetails(queryMods...)
}

// SalaryStatements retrieves all the salary_statement's SalaryStatements with an executor.
func (o *Earning) SalaryStatements(mods ...qm.QueryMod) salaryStatementQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`salary_statements`.`earning_id`=?", o.ID),
	)

	return SalaryStatements(queryMods...)
}

// LoadEarningDetails allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (earningL) LoadEarningDetails(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEarning interface{}, mods queries.Applicator) error {
	var slice []*Earning
	var object *Earning

	if singular {
		var ok bool
		object, ok = maybeEarning.(*Earning)
		if !ok {
			object = new(Earning)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEarning)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEarning))
			}
		}
	} else {
		s, ok := maybeEarning.(*[]*Earning)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEarning)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEarning))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &earningR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &earningR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`earning_details`),
		qm.WhereIn(`earning_details.earning_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load earning_details")
	}

	var resultSlice []*EarningDetail
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice earning_details")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on earning_details")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for earning_details")
	}

	if len(earningDetailAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.EarningDetails = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &earningDetailR{}
			}
			foreign.R.Earning = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.EarningID {
				local.R.EarningDetails = append(local.R.EarningDetails, foreign)
				if foreign.R == nil {
					foreign.R = &earningDetailR{}
				}
				foreign.R.Earning = local
				break
			}
		}
	}

	return nil
}

// LoadSalaryStatements allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (earningL) LoadSalaryStatements(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEarning interface{}, mods queries.Applicator) error {
	var slice []*Earning
	var object *Earning

	if singular {
		var ok bool
		object, ok = maybeEarning.(*Earning)
		if !ok {
			object = new(Earning)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEarning)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEarning))
			}
		}
	} else {
		s, ok := maybeEarning.(*[]*Earning)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEarning)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEarning))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &earningR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &earningR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`salary_statements`),
		qm.WhereIn(`salary_statements.earning_id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load salary_statements")
	}

	var resultSlice []*SalaryStatement
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice salary_statements")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on salary_statements")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for salary_statements")
	}

	if len(salaryStatementAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.SalaryStatements = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &salaryStatementR{}
			}
			foreign.R.Earning = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.EarningID {
				local.R.SalaryStatements = append(local.R.SalaryStatements, foreign)
				if foreign.R == nil {
					foreign.R = &salaryStatementR{}
				}
				foreign.R.Earning = local
				break
			}
		}
	}

	return nil
}

// AddEarningDetails adds the given related objects to the existing relationships
// of the earning, optionally inserting them as new records.
// Appends related to o.R.EarningDetails.
// Sets related.R.Earning appropriately.
func (o *Earning) AddEarningDetails(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*EarningDetail) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.EarningID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `earning_details` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"earning_id"}),
				strmangle.WhereClause("`", "`", 0, earningDetailPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.EarningID = o.ID
		}
	}

	if o.R == nil {
		o.R = &earningR{
			EarningDetails: related,
		}
	} else {
		o.R.EarningDetails = append(o.R.EarningDetails, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &earningDetailR{
				Earning: o,
			}
		} else {
			rel.R.Earning = o
		}
	}
	return nil
}

// AddSalaryStatements adds the given related objects to the existing relationships
// of the earning, optionally inserting them as new records.
// Appends related to o.R.SalaryStatements.
// Sets related.R.Earning appropriately.
func (o *Earning) AddSalaryStatements(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*SalaryStatement) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.EarningID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `salary_statements` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"earning_id"}),
				strmangle.WhereClause("`", "`", 0, salaryStatementPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.EarningID = o.ID
		}
	}

	if o.R == nil {
		o.R = &earningR{
			SalaryStatements: related,
		}
	} else {
		o.R.SalaryStatements = append(o.R.SalaryStatements, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &salaryStatementR{
				Earning: o,
			}
		} else {
			rel.R.Earning = o
		}
	}
	return nil
}

// Earnings retrieves all the records using an executor.
func Earnings(mods ...qm.QueryMod) earningQuery {
	mods = append(mods, qm.From("`earnings`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`earnings`.*"})
	}

	return earningQuery{q}
}

// FindEarning retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEarning(ctx context.Context, exec boil.ContextExecutor, iD uint32, selectCols ...string) (*Earning, error) {
	earningObj := &Earning{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `earnings` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, earningObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from earnings")
	}

	if err = earningObj.doAfterSelectHooks(ctx, exec); err != nil {
		return earningObj, err
	}

	return earningObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Earning) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no earnings provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(earningColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	earningInsertCacheMut.RLock()
	cache, cached := earningInsertCache[key]
	earningInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			earningAllColumns,
			earningColumnsWithDefault,
			earningColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(earningType, earningMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(earningType, earningMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `earnings` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `earnings` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `earnings` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, earningPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into earnings")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint32(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == earningMapping["id"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ID,
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, identifierCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for earnings")
	}

CacheNoHooks:
	if !cached {
		earningInsertCacheMut.Lock()
		earningInsertCache[key] = cache
		earningInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Earning.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Earning) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	earningUpdateCacheMut.RLock()
	cache, cached := earningUpdateCache[key]
	earningUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			earningAllColumns,
			earningPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update earnings, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `earnings` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, earningPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(earningType, earningMapping, append(wl, earningPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update earnings row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for earnings")
	}

	if !cached {
		earningUpdateCacheMut.Lock()
		earningUpdateCache[key] = cache
		earningUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q earningQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for earnings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for earnings")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EarningSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earningPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `earnings` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earningPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in earning slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all earning")
	}
	return rowsAff, nil
}

var mySQLEarningUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Earning) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no earnings provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(earningColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEarningUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	earningUpsertCacheMut.RLock()
	cache, cached := earningUpsertCache[key]
	earningUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			earningAllColumns,
			earningColumnsWithDefault,
			earningColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			earningAllColumns,
			earningPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert earnings, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`earnings`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `earnings` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(earningType, earningMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(earningType, earningMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for earnings")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ID = uint32(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == earningMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(earningType, earningMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for earnings")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for earnings")
	}

CacheNoHooks:
	if !cached {
		earningUpsertCacheMut.Lock()
		earningUpsertCache[key] = cache
		earningUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Earning record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Earning) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Earning provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), earningPrimaryKeyMapping)
	sql := "DELETE FROM `earnings` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from earnings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for earnings")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q earningQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no earningQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from earnings")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for earnings")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EarningSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(earningBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earningPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `earnings` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earningPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from earning slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for earnings")
	}

	if len(earningAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Earning) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEarning(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EarningSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EarningSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earningPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `earnings`.* FROM `earnings` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earningPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EarningSlice")
	}

	*o = slice

	return nil
}

// EarningExists checks if the Earning row exists.
func EarningExists(ctx context.Context, exec boil.ContextExecutor, iD uint32) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `earnings` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if earnings exists")
	}

	return exists, nil
}

// Exists checks if the Earning row exists.
func (o *Earning) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EarningExists(ctx, exec, o.ID)
}
