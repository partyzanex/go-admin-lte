// Code generated by SQLBoiler 4.7.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package postgres

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

// AuthToken is an object representing the database table.
type AuthToken struct {
	ID        int64     `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID    int64     `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	Token     string    `boil:"token" json:"token" toml:"token" yaml:"token"`
	Type      string    `boil:"type" json:"type" toml:"type" yaml:"type"`
	DTExpired time.Time `boil:"dt_expired" json:"dt_expired" toml:"dt_expired" yaml:"dt_expired"`
	DTCreated time.Time `boil:"dt_created" json:"dt_created" toml:"dt_created" yaml:"dt_created"`

	R *authTokenR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L authTokenL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AuthTokenColumns = struct {
	ID        string
	UserID    string
	Token     string
	Type      string
	DTExpired string
	DTCreated string
}{
	ID:        "id",
	UserID:    "user_id",
	Token:     "token",
	Type:      "type",
	DTExpired: "dt_expired",
	DTCreated: "dt_created",
}

var AuthTokenTableColumns = struct {
	ID        string
	UserID    string
	Token     string
	Type      string
	DTExpired string
	DTCreated string
}{
	ID:        "auth_token.id",
	UserID:    "auth_token.user_id",
	Token:     "auth_token.token",
	Type:      "auth_token.type",
	DTExpired: "auth_token.dt_expired",
	DTCreated: "auth_token.dt_created",
}

// Generated where

type whereHelperint64 struct{ field string }

func (w whereHelperint64) EQ(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint64) NEQ(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint64) LT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint64) LTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint64) GT(x int64) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint64) GTE(x int64) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelperstring struct{ field string }

func (w whereHelperstring) EQ(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperstring) NEQ(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperstring) LT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperstring) LTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperstring) GT(x string) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperstring) GTE(x string) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperstring) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperstring) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpertime_Time struct{ field string }

func (w whereHelpertime_Time) EQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.EQ, x)
}
func (w whereHelpertime_Time) NEQ(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.NEQ, x)
}
func (w whereHelpertime_Time) LT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpertime_Time) LTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpertime_Time) GT(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpertime_Time) GTE(x time.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AuthTokenWhere = struct {
	ID        whereHelperint64
	UserID    whereHelperint64
	Token     whereHelperstring
	Type      whereHelperstring
	DTExpired whereHelpertime_Time
	DTCreated whereHelpertime_Time
}{
	ID:        whereHelperint64{field: "\"goadmin\".\"auth_token\".\"id\""},
	UserID:    whereHelperint64{field: "\"goadmin\".\"auth_token\".\"user_id\""},
	Token:     whereHelperstring{field: "\"goadmin\".\"auth_token\".\"token\""},
	Type:      whereHelperstring{field: "\"goadmin\".\"auth_token\".\"type\""},
	DTExpired: whereHelpertime_Time{field: "\"goadmin\".\"auth_token\".\"dt_expired\""},
	DTCreated: whereHelpertime_Time{field: "\"goadmin\".\"auth_token\".\"dt_created\""},
}

// AuthTokenRels is where relationship names are stored.
var AuthTokenRels = struct {
	User string
}{
	User: "User",
}

// authTokenR is where relationships are stored.
type authTokenR struct {
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*authTokenR) NewStruct() *authTokenR {
	return &authTokenR{}
}

// authTokenL is where Load methods for each relationship are stored.
type authTokenL struct{}

var (
	authTokenAllColumns            = []string{"id", "user_id", "token", "type", "dt_expired", "dt_created"}
	authTokenColumnsWithoutDefault = []string{"user_id", "token", "type", "dt_expired"}
	authTokenColumnsWithDefault    = []string{"id", "dt_created"}
	authTokenPrimaryKeyColumns     = []string{"id"}
)

type (
	// AuthTokenSlice is an alias for a slice of pointers to AuthToken.
	// This should almost always be used instead of []AuthToken.
	AuthTokenSlice []*AuthToken
	// AuthTokenHook is the signature for custom AuthToken hook methods
	AuthTokenHook func(context.Context, boil.ContextExecutor, *AuthToken) error

	authTokenQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	authTokenType                 = reflect.TypeOf(&AuthToken{})
	authTokenMapping              = queries.MakeStructMapping(authTokenType)
	authTokenPrimaryKeyMapping, _ = queries.BindMapping(authTokenType, authTokenMapping, authTokenPrimaryKeyColumns)
	authTokenInsertCacheMut       sync.RWMutex
	authTokenInsertCache          = make(map[string]insertCache)
	authTokenUpdateCacheMut       sync.RWMutex
	authTokenUpdateCache          = make(map[string]updateCache)
	authTokenUpsertCacheMut       sync.RWMutex
	authTokenUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var authTokenBeforeInsertHooks []AuthTokenHook
var authTokenBeforeUpdateHooks []AuthTokenHook
var authTokenBeforeDeleteHooks []AuthTokenHook
var authTokenBeforeUpsertHooks []AuthTokenHook

var authTokenAfterInsertHooks []AuthTokenHook
var authTokenAfterSelectHooks []AuthTokenHook
var authTokenAfterUpdateHooks []AuthTokenHook
var authTokenAfterDeleteHooks []AuthTokenHook
var authTokenAfterUpsertHooks []AuthTokenHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AuthToken) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authTokenBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AuthToken) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authTokenBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AuthToken) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authTokenBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AuthToken) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authTokenBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AuthToken) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authTokenAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AuthToken) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authTokenAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AuthToken) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authTokenAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AuthToken) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authTokenAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AuthToken) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range authTokenAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAuthTokenHook registers your hook function for all future operations.
func AddAuthTokenHook(hookPoint boil.HookPoint, authTokenHook AuthTokenHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		authTokenBeforeInsertHooks = append(authTokenBeforeInsertHooks, authTokenHook)
	case boil.BeforeUpdateHook:
		authTokenBeforeUpdateHooks = append(authTokenBeforeUpdateHooks, authTokenHook)
	case boil.BeforeDeleteHook:
		authTokenBeforeDeleteHooks = append(authTokenBeforeDeleteHooks, authTokenHook)
	case boil.BeforeUpsertHook:
		authTokenBeforeUpsertHooks = append(authTokenBeforeUpsertHooks, authTokenHook)
	case boil.AfterInsertHook:
		authTokenAfterInsertHooks = append(authTokenAfterInsertHooks, authTokenHook)
	case boil.AfterSelectHook:
		authTokenAfterSelectHooks = append(authTokenAfterSelectHooks, authTokenHook)
	case boil.AfterUpdateHook:
		authTokenAfterUpdateHooks = append(authTokenAfterUpdateHooks, authTokenHook)
	case boil.AfterDeleteHook:
		authTokenAfterDeleteHooks = append(authTokenAfterDeleteHooks, authTokenHook)
	case boil.AfterUpsertHook:
		authTokenAfterUpsertHooks = append(authTokenAfterUpsertHooks, authTokenHook)
	}
}

// One returns a single authToken record from the query.
func (q authTokenQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AuthToken, error) {
	o := &AuthToken{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "postgres: failed to execute a one query for auth_token")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AuthToken records from the query.
func (q authTokenQuery) All(ctx context.Context, exec boil.ContextExecutor) (AuthTokenSlice, error) {
	var o []*AuthToken

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "postgres: failed to assign all query results to AuthToken slice")
	}

	if len(authTokenAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AuthToken records in the query.
func (q authTokenQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to count auth_token rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q authTokenQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "postgres: failed to check if auth_token exists")
	}

	return count > 0, nil
}

// User pointed to by the foreign key.
func (o *AuthToken) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"goadmin\".\"user\"")

	return query
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (authTokenL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAuthToken interface{}, mods queries.Applicator) error {
	var slice []*AuthToken
	var object *AuthToken

	if singular {
		object = maybeAuthToken.(*AuthToken)
	} else {
		slice = *maybeAuthToken.(*[]*AuthToken)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &authTokenR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &authTokenR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`goadmin.user`),
		qm.WhereIn(`goadmin.user.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for user")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for user")
	}

	if len(authTokenAfterSelectHooks) != 0 {
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
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.AuthTokens = append(foreign.R.AuthTokens, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.AuthTokens = append(foreign.R.AuthTokens, local)
				break
			}
		}
	}

	return nil
}

// SetUser of the authToken to the related item.
// Sets o.R.User to related.
// Adds o to related.R.AuthTokens.
func (o *AuthToken) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"goadmin\".\"auth_token\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, authTokenPrimaryKeyColumns),
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

	o.UserID = related.ID
	if o.R == nil {
		o.R = &authTokenR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			AuthTokens: AuthTokenSlice{o},
		}
	} else {
		related.R.AuthTokens = append(related.R.AuthTokens, o)
	}

	return nil
}

// AuthTokens retrieves all the records using an executor.
func AuthTokens(mods ...qm.QueryMod) authTokenQuery {
	mods = append(mods, qm.From("\"goadmin\".\"auth_token\""))
	return authTokenQuery{NewQuery(mods...)}
}

// FindAuthToken retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAuthToken(ctx context.Context, exec boil.ContextExecutor, iD int64, selectCols ...string) (*AuthToken, error) {
	authTokenObj := &AuthToken{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"goadmin\".\"auth_token\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, authTokenObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "postgres: unable to select from auth_token")
	}

	if err = authTokenObj.doAfterSelectHooks(ctx, exec); err != nil {
		return authTokenObj, err
	}

	return authTokenObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AuthToken) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("postgres: no auth_token provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authTokenColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	authTokenInsertCacheMut.RLock()
	cache, cached := authTokenInsertCache[key]
	authTokenInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			authTokenAllColumns,
			authTokenColumnsWithDefault,
			authTokenColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(authTokenType, authTokenMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(authTokenType, authTokenMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"goadmin\".\"auth_token\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"goadmin\".\"auth_token\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
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

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "postgres: unable to insert into auth_token")
	}

	if !cached {
		authTokenInsertCacheMut.Lock()
		authTokenInsertCache[key] = cache
		authTokenInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AuthToken.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AuthToken) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	authTokenUpdateCacheMut.RLock()
	cache, cached := authTokenUpdateCache[key]
	authTokenUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			authTokenAllColumns,
			authTokenPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("postgres: unable to update auth_token, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"goadmin\".\"auth_token\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, authTokenPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(authTokenType, authTokenMapping, append(wl, authTokenPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "postgres: unable to update auth_token row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to get rows affected by update for auth_token")
	}

	if !cached {
		authTokenUpdateCacheMut.Lock()
		authTokenUpdateCache[key] = cache
		authTokenUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q authTokenQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to update all for auth_token")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to retrieve rows affected for auth_token")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AuthTokenSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("postgres: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"goadmin\".\"auth_token\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, authTokenPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to update all in authToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to retrieve rows affected all in update all authToken")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AuthToken) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("postgres: no auth_token provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(authTokenColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
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
	key := buf.String()
	strmangle.PutBuffer(buf)

	authTokenUpsertCacheMut.RLock()
	cache, cached := authTokenUpsertCache[key]
	authTokenUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			authTokenAllColumns,
			authTokenColumnsWithDefault,
			authTokenColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			authTokenAllColumns,
			authTokenPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("postgres: unable to upsert auth_token, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(authTokenPrimaryKeyColumns))
			copy(conflict, authTokenPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"goadmin\".\"auth_token\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(authTokenType, authTokenMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(authTokenType, authTokenMapping, ret)
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
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "postgres: unable to upsert auth_token")
	}

	if !cached {
		authTokenUpsertCacheMut.Lock()
		authTokenUpsertCache[key] = cache
		authTokenUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AuthToken record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AuthToken) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("postgres: no AuthToken provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), authTokenPrimaryKeyMapping)
	sql := "DELETE FROM \"goadmin\".\"auth_token\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to delete from auth_token")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to get rows affected by delete for auth_token")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q authTokenQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("postgres: no authTokenQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to delete all from auth_token")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to get rows affected by deleteall for auth_token")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AuthTokenSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(authTokenBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"goadmin\".\"auth_token\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, authTokenPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "postgres: unable to delete all from authToken slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "postgres: failed to get rows affected by deleteall for auth_token")
	}

	if len(authTokenAfterDeleteHooks) != 0 {
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
func (o *AuthToken) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAuthToken(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AuthTokenSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AuthTokenSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), authTokenPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"goadmin\".\"auth_token\".* FROM \"goadmin\".\"auth_token\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, authTokenPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "postgres: unable to reload all in AuthTokenSlice")
	}

	*o = slice

	return nil
}

// AuthTokenExists checks if the AuthToken row exists.
func AuthTokenExists(ctx context.Context, exec boil.ContextExecutor, iD int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"goadmin\".\"auth_token\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "postgres: unable to check if auth_token exists")
	}

	return exists, nil
}
