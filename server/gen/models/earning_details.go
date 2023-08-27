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

// EarningDetail is an object representing the database table.
type EarningDetail struct {
	ID        uint32    `boil:"id" json:"id" toml:"id" yaml:"id"`
	EarningID uint32    `boil:"earning_id" json:"earning_id" toml:"earning_id" yaml:"earning_id"`
	Nominal   string    `boil:"nominal" json:"nominal" toml:"nominal" yaml:"nominal"`
	Amount    int       `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	CreatedAt time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *earningDetailR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L earningDetailL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EarningDetailColumns = struct {
	ID        string
	EarningID string
	Nominal   string
	Amount    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "id",
	EarningID: "earning_id",
	Nominal:   "nominal",
	Amount:    "amount",
	CreatedAt: "created_at",
	UpdatedAt: "updated_at",
}

var EarningDetailTableColumns = struct {
	ID        string
	EarningID string
	Nominal   string
	Amount    string
	CreatedAt string
	UpdatedAt string
}{
	ID:        "earning_details.id",
	EarningID: "earning_details.earning_id",
	Nominal:   "earning_details.nominal",
	Amount:    "earning_details.amount",
	CreatedAt: "earning_details.created_at",
	UpdatedAt: "earning_details.updated_at",
}

// Generated where

var EarningDetailWhere = struct {
	ID        whereHelperuint32
	EarningID whereHelperuint32
	Nominal   whereHelperstring
	Amount    whereHelperint
	CreatedAt whereHelpertime_Time
	UpdatedAt whereHelpertime_Time
}{
	ID:        whereHelperuint32{field: "`earning_details`.`id`"},
	EarningID: whereHelperuint32{field: "`earning_details`.`earning_id`"},
	Nominal:   whereHelperstring{field: "`earning_details`.`nominal`"},
	Amount:    whereHelperint{field: "`earning_details`.`amount`"},
	CreatedAt: whereHelpertime_Time{field: "`earning_details`.`created_at`"},
	UpdatedAt: whereHelpertime_Time{field: "`earning_details`.`updated_at`"},
}

// EarningDetailRels is where relationship names are stored.
var EarningDetailRels = struct {
	Earning string
}{
	Earning: "Earning",
}

// earningDetailR is where relationships are stored.
type earningDetailR struct {
	Earning *Earning `boil:"Earning" json:"Earning" toml:"Earning" yaml:"Earning"`
}

// NewStruct creates a new relationship struct
func (*earningDetailR) NewStruct() *earningDetailR {
	return &earningDetailR{}
}

func (r *earningDetailR) GetEarning() *Earning {
	if r == nil {
		return nil
	}
	return r.Earning
}

// earningDetailL is where Load methods for each relationship are stored.
type earningDetailL struct{}

var (
	earningDetailAllColumns            = []string{"id", "earning_id", "nominal", "amount", "created_at", "updated_at"}
	earningDetailColumnsWithoutDefault = []string{"earning_id", "nominal", "amount", "created_at", "updated_at"}
	earningDetailColumnsWithDefault    = []string{"id"}
	earningDetailPrimaryKeyColumns     = []string{"id"}
	earningDetailGeneratedColumns      = []string{}
)

type (
	// EarningDetailSlice is an alias for a slice of pointers to EarningDetail.
	// This should almost always be used instead of []EarningDetail.
	EarningDetailSlice []*EarningDetail
	// EarningDetailHook is the signature for custom EarningDetail hook methods
	EarningDetailHook func(context.Context, boil.ContextExecutor, *EarningDetail) error

	earningDetailQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	earningDetailType                 = reflect.TypeOf(&EarningDetail{})
	earningDetailMapping              = queries.MakeStructMapping(earningDetailType)
	earningDetailPrimaryKeyMapping, _ = queries.BindMapping(earningDetailType, earningDetailMapping, earningDetailPrimaryKeyColumns)
	earningDetailInsertCacheMut       sync.RWMutex
	earningDetailInsertCache          = make(map[string]insertCache)
	earningDetailUpdateCacheMut       sync.RWMutex
	earningDetailUpdateCache          = make(map[string]updateCache)
	earningDetailUpsertCacheMut       sync.RWMutex
	earningDetailUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var earningDetailAfterSelectHooks []EarningDetailHook

var earningDetailBeforeInsertHooks []EarningDetailHook
var earningDetailAfterInsertHooks []EarningDetailHook

var earningDetailBeforeUpdateHooks []EarningDetailHook
var earningDetailAfterUpdateHooks []EarningDetailHook

var earningDetailBeforeDeleteHooks []EarningDetailHook
var earningDetailAfterDeleteHooks []EarningDetailHook

var earningDetailBeforeUpsertHooks []EarningDetailHook
var earningDetailAfterUpsertHooks []EarningDetailHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *EarningDetail) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningDetailAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *EarningDetail) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningDetailBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *EarningDetail) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningDetailAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *EarningDetail) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningDetailBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *EarningDetail) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningDetailAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *EarningDetail) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningDetailBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *EarningDetail) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningDetailAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *EarningDetail) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningDetailBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *EarningDetail) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range earningDetailAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEarningDetailHook registers your hook function for all future operations.
func AddEarningDetailHook(hookPoint boil.HookPoint, earningDetailHook EarningDetailHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		earningDetailAfterSelectHooks = append(earningDetailAfterSelectHooks, earningDetailHook)
	case boil.BeforeInsertHook:
		earningDetailBeforeInsertHooks = append(earningDetailBeforeInsertHooks, earningDetailHook)
	case boil.AfterInsertHook:
		earningDetailAfterInsertHooks = append(earningDetailAfterInsertHooks, earningDetailHook)
	case boil.BeforeUpdateHook:
		earningDetailBeforeUpdateHooks = append(earningDetailBeforeUpdateHooks, earningDetailHook)
	case boil.AfterUpdateHook:
		earningDetailAfterUpdateHooks = append(earningDetailAfterUpdateHooks, earningDetailHook)
	case boil.BeforeDeleteHook:
		earningDetailBeforeDeleteHooks = append(earningDetailBeforeDeleteHooks, earningDetailHook)
	case boil.AfterDeleteHook:
		earningDetailAfterDeleteHooks = append(earningDetailAfterDeleteHooks, earningDetailHook)
	case boil.BeforeUpsertHook:
		earningDetailBeforeUpsertHooks = append(earningDetailBeforeUpsertHooks, earningDetailHook)
	case boil.AfterUpsertHook:
		earningDetailAfterUpsertHooks = append(earningDetailAfterUpsertHooks, earningDetailHook)
	}
}

// One returns a single earningDetail record from the query.
func (q earningDetailQuery) One(ctx context.Context, exec boil.ContextExecutor) (*EarningDetail, error) {
	o := &EarningDetail{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for earning_details")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all EarningDetail records from the query.
func (q earningDetailQuery) All(ctx context.Context, exec boil.ContextExecutor) (EarningDetailSlice, error) {
	var o []*EarningDetail

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to EarningDetail slice")
	}

	if len(earningDetailAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all EarningDetail records in the query.
func (q earningDetailQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count earning_details rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q earningDetailQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if earning_details exists")
	}

	return count > 0, nil
}

// Earning pointed to by the foreign key.
func (o *EarningDetail) Earning(mods ...qm.QueryMod) earningQuery {
	queryMods := []qm.QueryMod{
		qm.Where("`id` = ?", o.EarningID),
	}

	queryMods = append(queryMods, mods...)

	return Earnings(queryMods...)
}

// LoadEarning allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (earningDetailL) LoadEarning(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEarningDetail interface{}, mods queries.Applicator) error {
	var slice []*EarningDetail
	var object *EarningDetail

	if singular {
		var ok bool
		object, ok = maybeEarningDetail.(*EarningDetail)
		if !ok {
			object = new(EarningDetail)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeEarningDetail)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeEarningDetail))
			}
		}
	} else {
		s, ok := maybeEarningDetail.(*[]*EarningDetail)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeEarningDetail)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeEarningDetail))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &earningDetailR{}
		}
		args = append(args, object.EarningID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &earningDetailR{}
			}

			for _, a := range args {
				if a == obj.EarningID {
					continue Outer
				}
			}

			args = append(args, obj.EarningID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`earnings`),
		qm.WhereIn(`earnings.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Earning")
	}

	var resultSlice []*Earning
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Earning")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for earnings")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for earnings")
	}

	if len(earningAfterSelectHooks) != 0 {
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
		object.R.Earning = foreign
		if foreign.R == nil {
			foreign.R = &earningR{}
		}
		foreign.R.EarningDetails = append(foreign.R.EarningDetails, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.EarningID == foreign.ID {
				local.R.Earning = foreign
				if foreign.R == nil {
					foreign.R = &earningR{}
				}
				foreign.R.EarningDetails = append(foreign.R.EarningDetails, local)
				break
			}
		}
	}

	return nil
}

// SetEarning of the earningDetail to the related item.
// Sets o.R.Earning to related.
// Adds o to related.R.EarningDetails.
func (o *EarningDetail) SetEarning(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Earning) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `earning_details` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"earning_id"}),
		strmangle.WhereClause("`", "`", 0, earningDetailPrimaryKeyColumns),
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

	o.EarningID = related.ID
	if o.R == nil {
		o.R = &earningDetailR{
			Earning: related,
		}
	} else {
		o.R.Earning = related
	}

	if related.R == nil {
		related.R = &earningR{
			EarningDetails: EarningDetailSlice{o},
		}
	} else {
		related.R.EarningDetails = append(related.R.EarningDetails, o)
	}

	return nil
}

// EarningDetails retrieves all the records using an executor.
func EarningDetails(mods ...qm.QueryMod) earningDetailQuery {
	mods = append(mods, qm.From("`earning_details`"))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"`earning_details`.*"})
	}

	return earningDetailQuery{q}
}

// FindEarningDetail retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEarningDetail(ctx context.Context, exec boil.ContextExecutor, iD uint32, selectCols ...string) (*EarningDetail, error) {
	earningDetailObj := &EarningDetail{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `earning_details` where `id`=?", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, earningDetailObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from earning_details")
	}

	if err = earningDetailObj.doAfterSelectHooks(ctx, exec); err != nil {
		return earningDetailObj, err
	}

	return earningDetailObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *EarningDetail) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no earning_details provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(earningDetailColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	earningDetailInsertCacheMut.RLock()
	cache, cached := earningDetailInsertCache[key]
	earningDetailInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			earningDetailAllColumns,
			earningDetailColumnsWithDefault,
			earningDetailColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(earningDetailType, earningDetailMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(earningDetailType, earningDetailMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `earning_details` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `earning_details` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `earning_details` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, earningDetailPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into earning_details")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == earningDetailMapping["id"] {
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
		return errors.Wrap(err, "models: unable to populate default values for earning_details")
	}

CacheNoHooks:
	if !cached {
		earningDetailInsertCacheMut.Lock()
		earningDetailInsertCache[key] = cache
		earningDetailInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the EarningDetail.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *EarningDetail) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	earningDetailUpdateCacheMut.RLock()
	cache, cached := earningDetailUpdateCache[key]
	earningDetailUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			earningDetailAllColumns,
			earningDetailPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update earning_details, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `earning_details` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, earningDetailPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(earningDetailType, earningDetailMapping, append(wl, earningDetailPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update earning_details row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for earning_details")
	}

	if !cached {
		earningDetailUpdateCacheMut.Lock()
		earningDetailUpdateCache[key] = cache
		earningDetailUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q earningDetailQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for earning_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for earning_details")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EarningDetailSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earningDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `earning_details` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earningDetailPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in earningDetail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all earningDetail")
	}
	return rowsAff, nil
}

var mySQLEarningDetailUniqueColumns = []string{
	"id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *EarningDetail) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no earning_details provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(earningDetailColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEarningDetailUniqueColumns, o)

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

	earningDetailUpsertCacheMut.RLock()
	cache, cached := earningDetailUpsertCache[key]
	earningDetailUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			earningDetailAllColumns,
			earningDetailColumnsWithDefault,
			earningDetailColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			earningDetailAllColumns,
			earningDetailPrimaryKeyColumns,
		)

		if !updateColumns.IsNone() && len(update) == 0 {
			return errors.New("models: unable to upsert earning_details, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "`earning_details`", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `earning_details` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(earningDetailType, earningDetailMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(earningDetailType, earningDetailMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for earning_details")
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
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == earningDetailMapping["id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(earningDetailType, earningDetailMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for earning_details")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.retQuery)
		fmt.Fprintln(writer, nzUniqueCols...)
	}
	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for earning_details")
	}

CacheNoHooks:
	if !cached {
		earningDetailUpsertCacheMut.Lock()
		earningDetailUpsertCache[key] = cache
		earningDetailUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single EarningDetail record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *EarningDetail) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no EarningDetail provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), earningDetailPrimaryKeyMapping)
	sql := "DELETE FROM `earning_details` WHERE `id`=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from earning_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for earning_details")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q earningDetailQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no earningDetailQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from earning_details")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for earning_details")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EarningDetailSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(earningDetailBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earningDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `earning_details` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earningDetailPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from earningDetail slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for earning_details")
	}

	if len(earningDetailAfterDeleteHooks) != 0 {
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
func (o *EarningDetail) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEarningDetail(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EarningDetailSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EarningDetailSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), earningDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `earning_details`.* FROM `earning_details` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, earningDetailPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EarningDetailSlice")
	}

	*o = slice

	return nil
}

// EarningDetailExists checks if the EarningDetail row exists.
func EarningDetailExists(ctx context.Context, exec boil.ContextExecutor, iD uint32) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `earning_details` where `id`=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if earning_details exists")
	}

	return exists, nil
}

// Exists checks if the EarningDetail row exists.
func (o *EarningDetail) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return EarningDetailExists(ctx, exec, o.ID)
}
