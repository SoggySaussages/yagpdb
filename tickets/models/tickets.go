// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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

	"emperror.dev/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// CustomCommands retrieves all the records using an executor.
func CustomCommands(mods ...qm.QueryMod) customCommandQuery {
	mods = append(mods, qm.From("\"custom_commands\""))
	return customCommandQuery{NewQuery(mods...)}
}

// CustomCommandGroup is an object representing the database table.
type CustomCommandGroup struct {
	ID                int64            `boil:"id" json:"id" toml:"id" yaml:"id"`
	GuildID           int64            `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	Name              string           `boil:"name" json:"name" toml:"name" yaml:"name"`
	IgnoreRoles       types.Int64Array `boil:"ignore_roles" json:"ignore_roles,omitempty" toml:"ignore_roles" yaml:"ignore_roles,omitempty"`
	IgnoreChannels    types.Int64Array `boil:"ignore_channels" json:"ignore_channels,omitempty" toml:"ignore_channels" yaml:"ignore_channels,omitempty"`
	WhitelistRoles    types.Int64Array `boil:"whitelist_roles" json:"whitelist_roles,omitempty" toml:"whitelist_roles" yaml:"whitelist_roles,omitempty"`
	WhitelistChannels types.Int64Array `boil:"whitelist_channels" json:"whitelist_channels,omitempty" toml:"whitelist_channels" yaml:"whitelist_channels,omitempty"`

	R *customCommandGroupR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L customCommandGroupL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// CustomCommand is an object representing the database table.
type CustomCommand struct {
	LocalID                   int64             `boil:"local_id" json:"local_id" toml:"local_id" yaml:"local_id"`
	GuildID                   int64             `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	GroupID                   null.Int64        `boil:"group_id" json:"group_id,omitempty" toml:"group_id" yaml:"group_id,omitempty"`
	TriggerType               int               `boil:"trigger_type" json:"trigger_type" toml:"trigger_type" yaml:"trigger_type"`
	TextTrigger               string            `boil:"text_trigger" json:"text_trigger" toml:"text_trigger" yaml:"text_trigger"`
	TextTriggerCaseSensitive  bool              `boil:"text_trigger_case_sensitive" json:"text_trigger_case_sensitive" toml:"text_trigger_case_sensitive" yaml:"text_trigger_case_sensitive"`
	TimeTriggerInterval       int               `boil:"time_trigger_interval" json:"time_trigger_interval" toml:"time_trigger_interval" yaml:"time_trigger_interval"`
	TimeTriggerExcludingDays  types.Int64Array  `boil:"time_trigger_excluding_days" json:"time_trigger_excluding_days" toml:"time_trigger_excluding_days" yaml:"time_trigger_excluding_days"`
	TimeTriggerExcludingHours types.Int64Array  `boil:"time_trigger_excluding_hours" json:"time_trigger_excluding_hours" toml:"time_trigger_excluding_hours" yaml:"time_trigger_excluding_hours"`
	LastRun                   null.Time         `boil:"last_run" json:"last_run,omitempty" toml:"last_run" yaml:"last_run,omitempty"`
	NextRun                   null.Time         `boil:"next_run" json:"next_run,omitempty" toml:"next_run" yaml:"next_run,omitempty"`
	Responses                 types.StringArray `boil:"responses" json:"responses" toml:"responses" yaml:"responses"`
	Channels                  types.Int64Array  `boil:"channels" json:"channels,omitempty" toml:"channels" yaml:"channels,omitempty"`
	ChannelsWhitelistMode     bool              `boil:"channels_whitelist_mode" json:"channels_whitelist_mode" toml:"channels_whitelist_mode" yaml:"channels_whitelist_mode"`
	Roles                     types.Int64Array  `boil:"roles" json:"roles,omitempty" toml:"roles" yaml:"roles,omitempty"`
	RolesWhitelistMode        bool              `boil:"roles_whitelist_mode" json:"roles_whitelist_mode" toml:"roles_whitelist_mode" yaml:"roles_whitelist_mode"`
	ContextChannel            int64             `boil:"context_channel" json:"context_channel" toml:"context_channel" yaml:"context_channel"`
	ReactionTriggerMode       int16             `boil:"reaction_trigger_mode" json:"reaction_trigger_mode" toml:"reaction_trigger_mode" yaml:"reaction_trigger_mode"`
	Disabled                  bool              `boil:"disabled" json:"disabled" toml:"disabled" yaml:"disabled"`
	LastError                 string            `boil:"last_error" json:"last_error" toml:"last_error" yaml:"last_error"`
	LastErrorTime             null.Time         `boil:"last_error_time" json:"last_error_time,omitempty" toml:"last_error_time" yaml:"last_error_time,omitempty"`
	RunCount                  int               `boil:"run_count" json:"run_count" toml:"run_count" yaml:"run_count"`
	ShowErrors                bool              `boil:"show_errors" json:"show_errors" toml:"show_errors" yaml:"show_errors"`

	R *customCommandR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L customCommandL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// customCommandR is where relationships are stored.
type customCommandR struct {
	Group *CustomCommandGroup
}

// NewStruct creates a new relationship struct
func (*customCommandR) NewStruct() *customCommandR {
	return &customCommandR{}
}

// customCommandL is where Load methods for each relationship are stored.
type customCommandL struct{}

type (
	// CustomCommandSlice is an alias for a slice of pointers to CustomCommand.
	// This should generally be used opposed to []CustomCommand.
	CustomCommandSlice []*CustomCommand

	customCommandQuery struct {
		*queries.Query
	}
)

// customCommandGroupR is where relationships are stored.
type customCommandGroupR struct {
	GroupCustomCommands CustomCommandSlice
}

// NewStruct creates a new relationship struct
func (*customCommandGroupR) NewStruct() *customCommandGroupR {
	return &customCommandGroupR{}
}

// customCommandGroupL is where Load methods for each relationship are stored.
type customCommandGroupL struct{}

var (
	customCommandGroupAllColumns            = []string{"id", "guild_id", "name", "ignore_roles", "ignore_channels", "whitelist_roles", "whitelist_channels"}
	customCommandGroupColumnsWithoutDefault = []string{"guild_id", "name", "ignore_roles", "ignore_channels", "whitelist_roles", "whitelist_channels"}
	customCommandGroupColumnsWithDefault    = []string{"id"}
	customCommandGroupPrimaryKeyColumns     = []string{"id"}
)

type (
	// CustomCommandGroupSlice is an alias for a slice of pointers to CustomCommandGroup.
	// This should generally be used opposed to []CustomCommandGroup.
	CustomCommandGroupSlice []*CustomCommandGroup

	customCommandGroupQuery struct {
		*queries.Query
	}
)

// Ticket is an object representing the database table.
type Ticket struct {
	GuildID               int64     `boil:"guild_id" json:"guild_id" toml:"guild_id" yaml:"guild_id"`
	LocalID               int64     `boil:"local_id" json:"local_id" toml:"local_id" yaml:"local_id"`
	ChannelID             int64     `boil:"channel_id" json:"channel_id" toml:"channel_id" yaml:"channel_id"`
	Title                 string    `boil:"title" json:"title" toml:"title" yaml:"title"`
	CreatedAt             time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	ClosedAt              null.Time `boil:"closed_at" json:"closed_at,omitempty" toml:"closed_at" yaml:"closed_at,omitempty"`
	LogsID                int64     `boil:"logs_id" json:"logs_id" toml:"logs_id" yaml:"logs_id"`
	AuthorID              int64     `boil:"author_id" json:"author_id" toml:"author_id" yaml:"author_id"`
	AuthorUsernameDiscrim string    `boil:"author_username_discrim" json:"author_username_discrim" toml:"author_username_discrim" yaml:"author_username_discrim"`

	R *ticketR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L ticketL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TicketColumns = struct {
	GuildID               string
	LocalID               string
	ChannelID             string
	Title                 string
	CreatedAt             string
	ClosedAt              string
	LogsID                string
	AuthorID              string
	AuthorUsernameDiscrim string
}{
	GuildID:               "guild_id",
	LocalID:               "local_id",
	ChannelID:             "channel_id",
	Title:                 "title",
	CreatedAt:             "created_at",
	ClosedAt:              "closed_at",
	LogsID:                "logs_id",
	AuthorID:              "author_id",
	AuthorUsernameDiscrim: "author_username_discrim",
}

// Generated where

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

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Time) LT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Time) LTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Time) GT(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Time) GTE(x null.Time) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var TicketWhere = struct {
	GuildID               whereHelperint64
	LocalID               whereHelperint64
	ChannelID             whereHelperint64
	Title                 whereHelperstring
	CreatedAt             whereHelpertime_Time
	ClosedAt              whereHelpernull_Time
	LogsID                whereHelperint64
	AuthorID              whereHelperint64
	AuthorUsernameDiscrim whereHelperstring
}{
	GuildID:               whereHelperint64{field: "\"tickets\".\"guild_id\""},
	LocalID:               whereHelperint64{field: "\"tickets\".\"local_id\""},
	ChannelID:             whereHelperint64{field: "\"tickets\".\"channel_id\""},
	Title:                 whereHelperstring{field: "\"tickets\".\"title\""},
	CreatedAt:             whereHelpertime_Time{field: "\"tickets\".\"created_at\""},
	ClosedAt:              whereHelpernull_Time{field: "\"tickets\".\"closed_at\""},
	LogsID:                whereHelperint64{field: "\"tickets\".\"logs_id\""},
	AuthorID:              whereHelperint64{field: "\"tickets\".\"author_id\""},
	AuthorUsernameDiscrim: whereHelperstring{field: "\"tickets\".\"author_username_discrim\""},
}

// TicketRels is where relationship names are stored.
var TicketRels = struct {
}{}

// ticketR is where relationships are stored.
type ticketR struct {
}

// NewStruct creates a new relationship struct
func (*ticketR) NewStruct() *ticketR {
	return &ticketR{}
}

// ticketL is where Load methods for each relationship are stored.
type ticketL struct{}

var (
	ticketAllColumns            = []string{"guild_id", "local_id", "channel_id", "title", "created_at", "closed_at", "logs_id", "author_id", "author_username_discrim"}
	ticketColumnsWithoutDefault = []string{"guild_id", "local_id", "channel_id", "title", "created_at", "closed_at", "logs_id", "author_id", "author_username_discrim"}
	ticketColumnsWithDefault    = []string{}
	ticketPrimaryKeyColumns     = []string{"guild_id", "local_id"}
)

type (
	// TicketSlice is an alias for a slice of pointers to Ticket.
	// This should generally be used opposed to []Ticket.
	TicketSlice []*Ticket

	ticketQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	ticketType                 = reflect.TypeOf(&Ticket{})
	ticketMapping              = queries.MakeStructMapping(ticketType)
	ticketPrimaryKeyMapping, _ = queries.BindMapping(ticketType, ticketMapping, ticketPrimaryKeyColumns)
	ticketInsertCacheMut       sync.RWMutex
	ticketInsertCache          = make(map[string]insertCache)
	ticketUpdateCacheMut       sync.RWMutex
	ticketUpdateCache          = make(map[string]updateCache)
	ticketUpsertCacheMut       sync.RWMutex
	ticketUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// OneG returns a single ticket record from the query using the global executor.
func (q ticketQuery) OneG(ctx context.Context) (*Ticket, error) {
	return q.One(ctx, boil.GetContextDB())
}

// OneG returns a single customCommand record from the query using the global executor.
func (q customCommandQuery) OneG(ctx context.Context) (*CustomCommand, error) {
	return q.One(ctx, boil.GetContextDB())
}

// One returns a single ticket record from the query.
func (q ticketQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Ticket, error) {
	o := &Ticket{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.WrapIf(err, "models: failed to execute a one query for tickets")
	}

	return o, nil
}

// One returns a single customCommand record from the query.
func (q customCommandQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CustomCommand, error) {
	o := &CustomCommand{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for custom_commands")
	}

	return o, nil
}

// AllG returns all Ticket records from the query using the global executor.
func (q ticketQuery) AllG(ctx context.Context) (TicketSlice, error) {
	return q.All(ctx, boil.GetContextDB())
}

// All returns all Ticket records from the query.
func (q ticketQuery) All(ctx context.Context, exec boil.ContextExecutor) (TicketSlice, error) {
	var o []*Ticket

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.WrapIf(err, "models: failed to assign all query results to Ticket slice")
	}

	return o, nil
}

// CountG returns the count of all Ticket records in the query, and panics on error.
func (q ticketQuery) CountG(ctx context.Context) (int64, error) {
	return q.Count(ctx, boil.GetContextDB())
}

// Count returns the count of all Ticket records in the query.
func (q ticketQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to count tickets rows")
	}

	return count, nil
}

// ExistsG checks if the row exists in the table, and panics on error.
func (q ticketQuery) ExistsG(ctx context.Context) (bool, error) {
	return q.Exists(ctx, boil.GetContextDB())
}

// Exists checks if the row exists in the table.
func (q ticketQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.WrapIf(err, "models: failed to check if tickets exists")
	}

	return count > 0, nil
}

// Tickets retrieves all the records using an executor.
func Tickets(mods ...qm.QueryMod) ticketQuery {
	mods = append(mods, qm.From("\"tickets\""))
	return ticketQuery{NewQuery(mods...)}
}

// FindTicketG retrieves a single record by ID.
func FindTicketG(ctx context.Context, guildID int64, localID int64, selectCols ...string) (*Ticket, error) {
	return FindTicket(ctx, boil.GetContextDB(), guildID, localID, selectCols...)
}

// FindTicket retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTicket(ctx context.Context, exec boil.ContextExecutor, guildID int64, localID int64, selectCols ...string) (*Ticket, error) {
	ticketObj := &Ticket{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"tickets\" where \"guild_id\"=$1 AND \"local_id\"=$2", sel,
	)

	q := queries.Raw(query, guildID, localID)

	err := q.Bind(ctx, exec, ticketObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.WrapIf(err, "models: unable to select from tickets")
	}

	return ticketObj, nil
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Ticket) InsertG(ctx context.Context, columns boil.Columns) error {
	return o.Insert(ctx, boil.GetContextDB(), columns)
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Ticket) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tickets provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(ticketColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	ticketInsertCacheMut.RLock()
	cache, cached := ticketInsertCache[key]
	ticketInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			ticketAllColumns,
			ticketColumnsWithDefault,
			ticketColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(ticketType, ticketMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(ticketType, ticketMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"tickets\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"tickets\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.WrapIf(err, "models: unable to insert into tickets")
	}

	if !cached {
		ticketInsertCacheMut.Lock()
		ticketInsertCache[key] = cache
		ticketInsertCacheMut.Unlock()
	}

	return nil
}

// UpdateG a single Ticket record using the global executor.
// See Update for more documentation.
func (o *Ticket) UpdateG(ctx context.Context, columns boil.Columns) (int64, error) {
	return o.Update(ctx, boil.GetContextDB(), columns)
}

// Update uses an executor to update the Ticket.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Ticket) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	ticketUpdateCacheMut.RLock()
	cache, cached := ticketUpdateCache[key]
	ticketUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			ticketAllColumns,
			ticketPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update tickets, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"tickets\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, ticketPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(ticketType, ticketMapping, append(wl, ticketPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update tickets row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by update for tickets")
	}

	if !cached {
		ticketUpdateCacheMut.Lock()
		ticketUpdateCache[key] = cache
		ticketUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (q ticketQuery) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return q.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values.
func (q ticketQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update all for tickets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to retrieve rows affected for tickets")
	}

	return rowsAff, nil
}

// UpdateAllG updates all rows with the specified column values.
func (o TicketSlice) UpdateAllG(ctx context.Context, cols M) (int64, error) {
	return o.UpdateAll(ctx, boil.GetContextDB(), cols)
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TicketSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"tickets\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, ticketPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to update all in ticket slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to retrieve rows affected all in update all ticket")
	}
	return rowsAff, nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Ticket) UpsertG(ctx context.Context, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	return o.Upsert(ctx, boil.GetContextDB(), updateOnConflict, conflictColumns, updateColumns, insertColumns)
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Ticket) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no tickets provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	nzDefaults := queries.NonZeroDefaultSet(ticketColumnsWithDefault, o)

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

	ticketUpsertCacheMut.RLock()
	cache, cached := ticketUpsertCache[key]
	ticketUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			ticketAllColumns,
			ticketColumnsWithDefault,
			ticketColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			ticketAllColumns,
			ticketPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert tickets, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(ticketPrimaryKeyColumns))
			copy(conflict, ticketPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"tickets\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(ticketType, ticketMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(ticketType, ticketMapping, ret)
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

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
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
		return errors.WrapIf(err, "models: unable to upsert tickets")
	}

	if !cached {
		ticketUpsertCacheMut.Lock()
		ticketUpsertCache[key] = cache
		ticketUpsertCacheMut.Unlock()
	}

	return nil
}

// DeleteG deletes a single Ticket record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Ticket) DeleteG(ctx context.Context) (int64, error) {
	return o.Delete(ctx, boil.GetContextDB())
}

// Delete deletes a single Ticket record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Ticket) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Ticket provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), ticketPrimaryKeyMapping)
	sql := "DELETE FROM \"tickets\" WHERE \"guild_id\"=$1 AND \"local_id\"=$2"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete from tickets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by delete for tickets")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q ticketQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no ticketQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete all from tickets")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by deleteall for tickets")
	}

	return rowsAff, nil
}

// DeleteAllG deletes all rows in the slice.
func (o TicketSlice) DeleteAllG(ctx context.Context) (int64, error) {
	return o.DeleteAll(ctx, boil.GetContextDB())
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TicketSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"tickets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ticketPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.WrapIf(err, "models: unable to delete all from ticket slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.WrapIf(err, "models: failed to get rows affected by deleteall for tickets")
	}

	return rowsAff, nil
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Ticket) ReloadG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: no Ticket provided for reload")
	}

	return o.Reload(ctx, boil.GetContextDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Ticket) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTicket(ctx, exec, o.GuildID, o.LocalID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TicketSlice) ReloadAllG(ctx context.Context) error {
	if o == nil {
		return errors.New("models: empty TicketSlice provided for reload all")
	}

	return o.ReloadAll(ctx, boil.GetContextDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TicketSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TicketSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), ticketPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"tickets\".* FROM \"tickets\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, ticketPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.WrapIf(err, "models: unable to reload all in TicketSlice")
	}

	*o = slice

	return nil
}

// TicketExistsG checks if the Ticket row exists.
func TicketExistsG(ctx context.Context, guildID int64, localID int64) (bool, error) {
	return TicketExists(ctx, boil.GetContextDB(), guildID, localID)
}

// TicketExists checks if the Ticket row exists.
func TicketExists(ctx context.Context, exec boil.ContextExecutor, guildID int64, localID int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"tickets\" where \"guild_id\"=$1 AND \"local_id\"=$2 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, guildID, localID)
	}

	row := exec.QueryRowContext(ctx, sql, guildID, localID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.WrapIf(err, "models: unable to check if tickets exists")
	}

	return exists, nil
}
