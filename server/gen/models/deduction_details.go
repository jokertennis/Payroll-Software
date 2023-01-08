// Code generated by SQLBoiler 4.14.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

// DeductionDetail is an object representing the database table.
type DeductionDetail struct {
	ID          uint32    `boil:"id" json:"id" toml:"id" yaml:"id"`
	DeductionID uint32    `boil:"deduction_id" json:"deduction_id" toml:"deduction_id" yaml:"deduction_id"`
	Nominal     string    `boil:"nominal" json:"nominal" toml:"nominal" yaml:"nominal"`
	Amount      int       `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	CreatedAt   time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt   time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *deductionDetailR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L deductionDetailL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var DeductionDetailColumns = struct {
	ID          string
	DeductionID string
	Nominal     string
	Amount      string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "id",
	DeductionID: "deduction_id",
	Nominal:     "nominal",
	Amount:      "amount",
	CreatedAt:   "created_at",
	UpdatedAt:   "updated_at",
}

var DeductionDetailTableColumns = struct {
	ID          string
	DeductionID string
	Nominal     string
	Amount      string
	CreatedAt   string
	UpdatedAt   string
}{
	ID:          "deduction_details.id",
	DeductionID: "deduction_details.deduction_id",
	Nominal:     "deduction_details.nominal",
	Amount:      "deduction_details.amount",
	CreatedAt:   "deduction_details.created_at",
	UpdatedAt:   "deduction_details.updated_at",
}

// Generated where

type whereHelperuint32 struct{ field string }

func (w whereHelperuint32) EQ(x uint32) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperuint32) NEQ(x uint32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperuint32) LT(x uint32) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperuint32) LTE(x uint32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperuint32) GT(x uint32) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperuint32) GTE(x uint32) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperuint32) IN(slice []uint32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperuint32) NIN(slice []uint32) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

var DeductionDetailWhere = struct {
	ID          whereHelperuint32
	DeductionID whereHelperuint32
	Nominal     whereHelperstring
	Amount      whereHelperint
	CreatedAt   whereHelpertime_Time
	UpdatedAt   whereHelpertime_Time
}{
	ID:          whereHelperuint32{field: "`deduction_details`.`id`"},
	DeductionID: whereHelperuint32{field: "`deduction_details`.`deduction_id`"},
	Nominal:     whereHelperstring{field: "`deduction_details`.`nominal`"},
	Amount:      whereHelperint{field: "`deduction_details`.`amount`"},
	CreatedAt:   whereHelpertime_Time{field: "`deduction_details`.`created_at`"},
	UpdatedAt:   whereHelpertime_Time{field: "`deduction_details`.`updated_at`"},
}

// DeductionDetailRels is where relationship names are stored.
var DeductionDetailRels = struct {
	Deduction string
}{
	Deduction: "Deduction",
}

// deductionDetailR is where relationships are stored.
type deductionDetailR struct {
	Deduction *Deduction `boil:"Deduction" json:"Deduction" toml:"Deduction" yaml:"Deduction"`
}

// NewStruct creates a new relationship struct
func (*deductionDetailR) NewStruct() *deductionDetailR {
	return &deductionDetailR{}
}

func (r *deductionDetailR) GetDeduction() *Deduction {
	if r == nil {
		return nil
	}
	return r.Deduction
}

// deductionDetailL is where Load methods for each relationship are stored.
type deductionDetailL struct{}

var (
	deductionDetailAllColumns            = []string{"id", "deduction_id", "nominal", "amount", "created_at", "updated_at"}
	deductionDetailColumnsWithoutDefault = []string{"deduction_id", "nominal", "amount", "created_at", "updated_at"}
	deductionDetailColumnsWithDefault    = []string{"id"}
	deductionDetailPrimaryKeyColumns     = []string{"id"}
	deductionDetailGeneratedColumns      = []string{}
)

type (
	// DeductionDetailSlice is an alias for a slice of pointers to DeductionDetail.
	// This should almost always be used instead of []DeductionDetail.
	DeductionDetailSlice []*DeductionDetail
	// DeductionDetailHook is the signature for custom DeductionDetail hook methods
	DeductionDetailHook func(context.Context, boil.ContextExecutor, *DeductionDetail) error

	deductionDetailQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	deductionDetailType                 = reflect.TypeOf(&DeductionDetail{})
	deductionDetailMapping              = queries.MakeStructMapping(deductionDetailType)
	deductionDetailPrimaryKeyMapping, _ = queries.BindMapping(deductionDetailType, deductionDetailMapping, deductionDetailPrimaryKeyColumns)
	deductionDetailInsertCacheMut       sync.RWMutex
	deductionDetailInsertCache          = make(map[string]insertCache)
	deductionDetailUpdateCacheMut       sync.RWMutex
	deductionDetailUpdateCache          = make(map[string]updateCache)
	deductionDetailUpsertCacheMut       sync.RWMutex
	deductionDetailUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var deductionDetailAfterSelectHooks []DeductionDetailHook

var deductionDetailBeforeInsertHooks []DeductionDetailHook
var deductionDetailAfterInsertHooks []DeductionDetailHook

var deductionDetailBeforeUpdateHooks []DeductionDetailHook
var deductionDetailAfterUpdateHooks []DeductionDetailHook

var deductionDetailBeforeDeleteHooks []DeductionDetailHook
var deductionDetailAfterDeleteHooks []DeductionDetailHook

var deductionDetailBeforeUpsertHooks []DeductionDetailHook
var deductionDetailAfterUpsertHooks []DeductionDetailHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DeductionDetail) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deductionDetailAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DeductionDetail) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deductionDetailBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DeductionDetail) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deductionDetailAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DeductionDetail) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deductionDetailBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DeductionDetail) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deductionDetailAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DeductionDetail) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deductionDetailBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DeductionDetail) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deductionDetailAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DeductionDetail) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deductionDetailBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DeductionDetail) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range deductionDetailAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDeductionDetailHook registers your hook function for all future operations.
func AddDeductionDetailHook(hookPoint boil.HookPoint, deductionDetailHook DeductionDetailHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		deductionDetailAfterSelectHooks = append(deductionDetailAfterSelectHooks, deductionDetailHook)
	case boil.BeforeInsertHook:
		deductionDetailBeforeInsertHooks = append(deductionDetailBeforeInsertHooks, deductionDetailHook)
	case boil.AfterInsertHook:
		deductionDetailAfterInsertHooks = append(deductionDetailAfterInsertHooks, deductionDetailHook)
	case boil.BeforeUpdateHook:
		deductionDetailBeforeUpdateHooks = append(deductionDetailBeforeUpdateHooks, deductionDetailHook)
	case boil.AfterUpdateHook:
		deductionDetailAfterUpdateHooks = append(deductionDetailAfterUpdateHooks, deductionDetailHook)
	case boil.BeforeDeleteHook:
		deductionDetailBeforeDeleteHooks = append(deductionDetailBeforeDeleteHooks, deductionDetailHook)
	case boil.AfterDeleteHook:
		deductionDetailAfterDeleteHooks = append(deductionDetailAfterDeleteHooks, deductionDetailHook)
	case boil.BeforeUpsertHook:
		deductionDetailBeforeUpsertHooks = append(deductionDetailBeforeUpsertHooks, deductionDetailHook)
	case boil.AfterUpsertHook:
		deductionDetailAfterUpsertHooks = append(deductionDetailAfterUpsertHooks, deductionDetailHook)
	}
}

// One returns a single deductionDetail record from the query.
func (q deductionDetailQuery) One(ctx context.Context, exec boil.ContextExecutor) (*DeductionDetail, error) {
	o := &DeductionDetail{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for deduction_details")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all DeductionDetail records from the query.
func (q deductionDetailQuery) All(ctx context.Context, exec boil.ContextExecutor) (DeductionDetailSlice, error) {
	var o []*DeductionDetail

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DeductionDetail slice")
	}

	if len(deductionDetailAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all DeductionDetail records in the query.
func (q deductionDetailQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count deduction_details rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q deductionDetailQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if deduction_details exists")
	}

	return count > 0, nil
}

// Deduction pointed to by the foreign key.
func (o *DeductionDetail) Deduction(mods ...qm.QueryMod) deductionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.DeductionID),
	}

	queryMods = append(queryMods, mods...)

	return Deductions(queryMods...)
}

// LoadDeduction allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (deductionDetailL) LoadDeduction(ctx context.Context, e boil.ContextExecutor, singular bool, maybeDeductionDetail interface{}, mods queries.Applicator) error {
	var slice []*DeductionDetail
	var object *DeductionDetail

	if singular {
		var ok bool
		object, ok = maybeDeductionDetail.(*DeductionDetail)
		if !ok {
			object = new(DeductionDetail)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeDeductionDetail)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeDeductionDetail))
			}
		}
	} else {
		s, ok := maybeDeductionDetail.(*[]*DeductionDetail)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeDeductionDetail)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeDeductionDetail))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &deductionDetailR{}
		}
		args = append(args, object.DeductionID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &deductionDetailR{}
			}

			for _, a := range args {
				if a == obj.DeductionID {
					continue Outer
				}
			}

			args = append(args, obj.DeductionID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`deductions`),
		qm.WhereIn(`deductions.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Deduction")
	}

	var resultSlice []*Deduction
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Deduction")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for deductions")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for deductions")
	}

	if len(deductionAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Deduction = foreign
		if foreign.R == nil {
			foreign.R = &deductionR{}
		}
		foreign.R.DeductionDetails = append(foreign.R.DeductionDetails, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.DeductionID == foreign.ID {
				local.R.Deduction = foreign
				if foreign.R == nil {
					foreign.R = &deductionR{}
				}
				foreign.R.DeductionDetails = append(foreign.R.DeductionDetails, local)
				break
			}
		}
	}

	return nil
}

// SetDeduction of the deductionDetail to the related item.
// Sets o.R.Deduction to related.
// Adds o to related.R.DeductionDetails.
func (o *DeductionDetail) SetDeduction(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Deduction) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `deduction_details` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"deduction_id"}),
		strmangle.WhereClause("`", "`", 0, deductionDetailPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DeductionID = related.ID
	if o.R == nil {
		o.R = &deductionDetailR{
			Deduction: related,
		}
	} else {
		o.R.Deduction = related
	}

	if related.R == nil {
		related.R = &deductionR{
			DeductionDetails: DeductionDetailSlice{o},
		}
	} else {
		related.R.DeductionDetails = append(related.R.DeductionDetails, o)
	}

	return nil
}

// DeductionDetails retrieves all the records using an executor.
func DeductionDetails(mods ...qm.QueryMod) deductionDetailQuery {
	mods = append(mods, qm.From("`deduction_details`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`deduction_details`.*"})
	}

	return deductionDetailQuery{q}
}

// FindDeductionDetail retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDeductionDetail(ctx context.Context, exec boil.ContextExecutor, iD uint32, selectCols ...string) (*DeductionDetail, error) {
	deductionDetailObj := &DeductionDetail{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `deduction_details` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, deductionDetailObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from deduction_details")
	}

	if err = deductionDetailObj.doAfterSelectHooks(ctx, exec); err != nil {
		return deductionDetailObj, err
	}

	return deductionDetailObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *DeductionDetail) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no deduction_details provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(deductionDetailColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	deductionDetailInsertCacheMut.RLock()
	cache, cached := deductionDetailInsertCache[key]
	deductionDetailInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			deductionDetailAllColumns,
			deductionDetailColumnsWithDefault,
			deductionDetailColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(deductionDetailType, deductionDetailMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(deductionDetailType, deductionDetailMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `deduction_details` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `deduction_details` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `deduction_details` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, deductionDetailPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into deduction_details")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == deductionDetailMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for deduction_details")
	}

CacheNoHooks:
	if !cached {
		deductionDetailInsertCacheMut.Lock()
		deductionDetailInsertCache[key] = cache
		deductionDetailInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the DeductionDetail.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *DeductionDetail) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	deductionDetailUpdateCacheMut.RLock()
	cache, cached := deductionDetailUpdateCache[key]
	deductionDetailUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			deductionDetailAllColumns,
			deductionDetailPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update deduction_details, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `deduction_details` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, deductionDetailPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(deductionDetailType, deductionDetailMapping, append(wl, deductionDetailPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update deduction_details row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for deduction_details")
	}

	if !cached {
		deductionDetailUpdateCacheMut.Lock()
		deductionDetailUpdateCache[key] = cache
		deductionDetailUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q deductionDetailQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for deduction_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for deduction_details")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DeductionDetailSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deductionDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `deduction_details` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, deductionDetailPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in deductionDetail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all deductionDetail")
	}
	return rowsAff, nil
}

var mySQLDeductionDetailUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *DeductionDetail) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no deduction_details provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(deductionDetailColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLDeductionDetailUniqueColumns, o)

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

	deductionDetailUpsertCacheMut.RLock()
	cache, cached := deductionDetailUpsertCache[key]
	deductionDetailUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			deductionDetailAllColumns,
			deductionDetailColumnsWithDefault,
			deductionDetailColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			deductionDetailAllColumns,
			deductionDetailPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert deduction_details, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`deduction_details`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `deduction_details` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(deductionDetailType, deductionDetailMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(deductionDetailType, deductionDetailMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for deduction_details")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == deductionDetailMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(deductionDetailType, deductionDetailMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for deduction_details")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for deduction_details")
	}

CacheNoHooks:
	if !cached {
		deductionDetailUpsertCacheMut.Lock()
		deductionDetailUpsertCache[key] = cache
		deductionDetailUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single DeductionDetail record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DeductionDetail) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no DeductionDetail provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), deductionDetailPrimaryKeyMapping)
	sql := "DELETE FROM `deduction_details` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from deduction_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for deduction_details")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q deductionDetailQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no deductionDetailQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from deduction_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for deduction_details")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DeductionDetailSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(deductionDetailBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deductionDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `deduction_details` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, deductionDetailPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from deductionDetail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for deduction_details")
	}

	if len(deductionDetailAfterDeleteHooks) != 0 {
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
func (o *DeductionDetail) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindDeductionDetail(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DeductionDetailSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := DeductionDetailSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), deductionDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `deduction_details`.* FROM `deduction_details` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, deductionDetailPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DeductionDetailSlice")
	}

	*o = slice

	return nil
}

// DeductionDetailExists checks if the DeductionDetail row exists.
func DeductionDetailExists(ctx context.Context, exec boil.ContextExecutor, iD uint32) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `deduction_details` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if deduction_details exists")
	}

	return exists, nil
}

// Exists checks if the DeductionDetail row exists.
func (o *DeductionDetail) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return DeductionDetailExists(ctx, exec, o.ID)
}
