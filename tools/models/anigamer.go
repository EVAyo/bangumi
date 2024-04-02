// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// Anigamer is an object representing the database table.
type Anigamer struct {
	AcgSN     int64        `boil:"acg_sn" json:"acg_sn" toml:"acg_sn" yaml:"acg_sn"`
	AnimeSN   int64        `boil:"anime_sn" json:"anime_sn" toml:"anime_sn" yaml:"anime_sn"`
	Title     null.String  `boil:"title" json:"title,omitempty" toml:"title" yaml:"title,omitempty"`
	DCC1      int64        `boil:"dc_c1" json:"dc_c1" toml:"dc_c1" yaml:"dc_c1"`
	DCC2      int64        `boil:"dc_c2" json:"dc_c2" toml:"dc_c2" yaml:"dc_c2"`
	Cover     null.String  `boil:"cover" json:"cover,omitempty" toml:"cover" yaml:"cover,omitempty"`
	Popular   null.Int64   `boil:"popular" json:"popular,omitempty" toml:"popular" yaml:"popular,omitempty"`
	Bilingual null.Bool    `boil:"bilingual" json:"bilingual,omitempty" toml:"bilingual" yaml:"bilingual,omitempty"`
	Edition   null.String  `boil:"edition" json:"edition,omitempty" toml:"edition" yaml:"edition,omitempty"`
	VipTime   null.Time    `boil:"vip_time" json:"vip_time,omitempty" toml:"vip_time" yaml:"vip_time,omitempty"`
	Score     null.Float64 `boil:"score" json:"score,omitempty" toml:"score" yaml:"score,omitempty"`
	UpdatedAt time.Time    `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`
	CreatedAt time.Time    `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`

	R *anigamerR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L anigamerL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AnigamerColumns = struct {
	AcgSN     string
	AnimeSN   string
	Title     string
	DCC1      string
	DCC2      string
	Cover     string
	Popular   string
	Bilingual string
	Edition   string
	VipTime   string
	Score     string
	UpdatedAt string
	CreatedAt string
}{
	AcgSN:     "acg_sn",
	AnimeSN:   "anime_sn",
	Title:     "title",
	DCC1:      "dc_c1",
	DCC2:      "dc_c2",
	Cover:     "cover",
	Popular:   "popular",
	Bilingual: "bilingual",
	Edition:   "edition",
	VipTime:   "vip_time",
	Score:     "score",
	UpdatedAt: "updated_at",
	CreatedAt: "created_at",
}

var AnigamerTableColumns = struct {
	AcgSN     string
	AnimeSN   string
	Title     string
	DCC1      string
	DCC2      string
	Cover     string
	Popular   string
	Bilingual string
	Edition   string
	VipTime   string
	Score     string
	UpdatedAt string
	CreatedAt string
}{
	AcgSN:     "anigamer.acg_sn",
	AnimeSN:   "anigamer.anime_sn",
	Title:     "anigamer.title",
	DCC1:      "anigamer.dc_c1",
	DCC2:      "anigamer.dc_c2",
	Cover:     "anigamer.cover",
	Popular:   "anigamer.popular",
	Bilingual: "anigamer.bilingual",
	Edition:   "anigamer.edition",
	VipTime:   "anigamer.vip_time",
	Score:     "anigamer.score",
	UpdatedAt: "anigamer.updated_at",
	CreatedAt: "anigamer.created_at",
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

type whereHelpernull_String struct{ field string }

func (w whereHelpernull_String) EQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_String) NEQ(x null.String) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_String) LT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_String) LTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_String) GT(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_String) GTE(x null.String) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_String) LIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" LIKE ?", x)
}
func (w whereHelpernull_String) NLIKE(x null.String) qm.QueryMod {
	return qm.Where(w.field+" NOT LIKE ?", x)
}
func (w whereHelpernull_String) IN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_String) NIN(slice []string) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_String) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_String) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Int64 struct{ field string }

func (w whereHelpernull_Int64) EQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int64) NEQ(x null.Int64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int64) LT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int64) LTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int64) GT(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int64) GTE(x null.Int64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Int64) IN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Int64) NIN(slice []int64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Int64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Bool struct{ field string }

func (w whereHelpernull_Bool) EQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Bool) NEQ(x null.Bool) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Bool) LT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Bool) LTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Bool) GT(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Bool) GTE(x null.Bool) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

func (w whereHelpernull_Bool) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Bool) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Time struct{ field string }

func (w whereHelpernull_Time) EQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Time) NEQ(x null.Time) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
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

func (w whereHelpernull_Time) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Time) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

type whereHelpernull_Float64 struct{ field string }

func (w whereHelpernull_Float64) EQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Float64) NEQ(x null.Float64) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Float64) LT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Float64) LTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Float64) GT(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Float64) GTE(x null.Float64) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}
func (w whereHelpernull_Float64) IN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelpernull_Float64) NIN(slice []float64) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

func (w whereHelpernull_Float64) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Float64) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }

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

var AnigamerWhere = struct {
	AcgSN     whereHelperint64
	AnimeSN   whereHelperint64
	Title     whereHelpernull_String
	DCC1      whereHelperint64
	DCC2      whereHelperint64
	Cover     whereHelpernull_String
	Popular   whereHelpernull_Int64
	Bilingual whereHelpernull_Bool
	Edition   whereHelpernull_String
	VipTime   whereHelpernull_Time
	Score     whereHelpernull_Float64
	UpdatedAt whereHelpertime_Time
	CreatedAt whereHelpertime_Time
}{
	AcgSN:     whereHelperint64{field: "\"anigamer\".\"acg_sn\""},
	AnimeSN:   whereHelperint64{field: "\"anigamer\".\"anime_sn\""},
	Title:     whereHelpernull_String{field: "\"anigamer\".\"title\""},
	DCC1:      whereHelperint64{field: "\"anigamer\".\"dc_c1\""},
	DCC2:      whereHelperint64{field: "\"anigamer\".\"dc_c2\""},
	Cover:     whereHelpernull_String{field: "\"anigamer\".\"cover\""},
	Popular:   whereHelpernull_Int64{field: "\"anigamer\".\"popular\""},
	Bilingual: whereHelpernull_Bool{field: "\"anigamer\".\"bilingual\""},
	Edition:   whereHelpernull_String{field: "\"anigamer\".\"edition\""},
	VipTime:   whereHelpernull_Time{field: "\"anigamer\".\"vip_time\""},
	Score:     whereHelpernull_Float64{field: "\"anigamer\".\"score\""},
	UpdatedAt: whereHelpertime_Time{field: "\"anigamer\".\"updated_at\""},
	CreatedAt: whereHelpertime_Time{field: "\"anigamer\".\"created_at\""},
}

// AnigamerRels is where relationship names are stored.
var AnigamerRels = struct {
}{}

// anigamerR is where relationships are stored.
type anigamerR struct {
}

// NewStruct creates a new relationship struct
func (*anigamerR) NewStruct() *anigamerR {
	return &anigamerR{}
}

// anigamerL is where Load methods for each relationship are stored.
type anigamerL struct{}

var (
	anigamerAllColumns            = []string{"acg_sn", "anime_sn", "title", "dc_c1", "dc_c2", "cover", "popular", "bilingual", "edition", "vip_time", "score", "updated_at", "created_at"}
	anigamerColumnsWithoutDefault = []string{"acg_sn", "dc_c1", "dc_c2", "updated_at", "created_at"}
	anigamerColumnsWithDefault    = []string{"anime_sn", "title", "cover", "popular", "bilingual", "edition", "vip_time", "score"}
	anigamerPrimaryKeyColumns     = []string{"anime_sn"}
	anigamerGeneratedColumns      = []string{"anime_sn"}
)

type (
	// AnigamerSlice is an alias for a slice of pointers to Anigamer.
	// This should almost always be used instead of []Anigamer.
	AnigamerSlice []*Anigamer
	// AnigamerHook is the signature for custom Anigamer hook methods
	AnigamerHook func(context.Context, boil.ContextExecutor, *Anigamer) error

	anigamerQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	anigamerType                 = reflect.TypeOf(&Anigamer{})
	anigamerMapping              = queries.MakeStructMapping(anigamerType)
	anigamerPrimaryKeyMapping, _ = queries.BindMapping(anigamerType, anigamerMapping, anigamerPrimaryKeyColumns)
	anigamerInsertCacheMut       sync.RWMutex
	anigamerInsertCache          = make(map[string]insertCache)
	anigamerUpdateCacheMut       sync.RWMutex
	anigamerUpdateCache          = make(map[string]updateCache)
	anigamerUpsertCacheMut       sync.RWMutex
	anigamerUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var anigamerAfterSelectMu sync.Mutex
var anigamerAfterSelectHooks []AnigamerHook

var anigamerBeforeInsertMu sync.Mutex
var anigamerBeforeInsertHooks []AnigamerHook
var anigamerAfterInsertMu sync.Mutex
var anigamerAfterInsertHooks []AnigamerHook

var anigamerBeforeUpdateMu sync.Mutex
var anigamerBeforeUpdateHooks []AnigamerHook
var anigamerAfterUpdateMu sync.Mutex
var anigamerAfterUpdateHooks []AnigamerHook

var anigamerBeforeDeleteMu sync.Mutex
var anigamerBeforeDeleteHooks []AnigamerHook
var anigamerAfterDeleteMu sync.Mutex
var anigamerAfterDeleteHooks []AnigamerHook

var anigamerBeforeUpsertMu sync.Mutex
var anigamerBeforeUpsertHooks []AnigamerHook
var anigamerAfterUpsertMu sync.Mutex
var anigamerAfterUpsertHooks []AnigamerHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Anigamer) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range anigamerAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Anigamer) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range anigamerBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Anigamer) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range anigamerAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Anigamer) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range anigamerBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Anigamer) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range anigamerAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Anigamer) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range anigamerBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Anigamer) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range anigamerAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Anigamer) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range anigamerBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Anigamer) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range anigamerAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAnigamerHook registers your hook function for all future operations.
func AddAnigamerHook(hookPoint boil.HookPoint, anigamerHook AnigamerHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		anigamerAfterSelectMu.Lock()
		anigamerAfterSelectHooks = append(anigamerAfterSelectHooks, anigamerHook)
		anigamerAfterSelectMu.Unlock()
	case boil.BeforeInsertHook:
		anigamerBeforeInsertMu.Lock()
		anigamerBeforeInsertHooks = append(anigamerBeforeInsertHooks, anigamerHook)
		anigamerBeforeInsertMu.Unlock()
	case boil.AfterInsertHook:
		anigamerAfterInsertMu.Lock()
		anigamerAfterInsertHooks = append(anigamerAfterInsertHooks, anigamerHook)
		anigamerAfterInsertMu.Unlock()
	case boil.BeforeUpdateHook:
		anigamerBeforeUpdateMu.Lock()
		anigamerBeforeUpdateHooks = append(anigamerBeforeUpdateHooks, anigamerHook)
		anigamerBeforeUpdateMu.Unlock()
	case boil.AfterUpdateHook:
		anigamerAfterUpdateMu.Lock()
		anigamerAfterUpdateHooks = append(anigamerAfterUpdateHooks, anigamerHook)
		anigamerAfterUpdateMu.Unlock()
	case boil.BeforeDeleteHook:
		anigamerBeforeDeleteMu.Lock()
		anigamerBeforeDeleteHooks = append(anigamerBeforeDeleteHooks, anigamerHook)
		anigamerBeforeDeleteMu.Unlock()
	case boil.AfterDeleteHook:
		anigamerAfterDeleteMu.Lock()
		anigamerAfterDeleteHooks = append(anigamerAfterDeleteHooks, anigamerHook)
		anigamerAfterDeleteMu.Unlock()
	case boil.BeforeUpsertHook:
		anigamerBeforeUpsertMu.Lock()
		anigamerBeforeUpsertHooks = append(anigamerBeforeUpsertHooks, anigamerHook)
		anigamerBeforeUpsertMu.Unlock()
	case boil.AfterUpsertHook:
		anigamerAfterUpsertMu.Lock()
		anigamerAfterUpsertHooks = append(anigamerAfterUpsertHooks, anigamerHook)
		anigamerAfterUpsertMu.Unlock()
	}
}

// One returns a single anigamer record from the query.
func (q anigamerQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Anigamer, error) {
	o := &Anigamer{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for anigamer")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Anigamer records from the query.
func (q anigamerQuery) All(ctx context.Context, exec boil.ContextExecutor) (AnigamerSlice, error) {
	var o []*Anigamer

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Anigamer slice")
	}

	if len(anigamerAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Anigamer records in the query.
func (q anigamerQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count anigamer rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q anigamerQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if anigamer exists")
	}

	return count > 0, nil
}

// Anigamers retrieves all the records using an executor.
func Anigamers(mods ...qm.QueryMod) anigamerQuery {
	mods = append(mods, qm.From("\"anigamer\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"anigamer\".*"})
	}

	return anigamerQuery{q}
}

// FindAnigamer retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAnigamer(ctx context.Context, exec boil.ContextExecutor, animeSN int64, selectCols ...string) (*Anigamer, error) {
	anigamerObj := &Anigamer{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"anigamer\" where \"anime_sn\"=?", sel,
	)

	q := queries.Raw(query, animeSN)

	err := q.Bind(ctx, exec, anigamerObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from anigamer")
	}

	if err = anigamerObj.doAfterSelectHooks(ctx, exec); err != nil {
		return anigamerObj, err
	}

	return anigamerObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Anigamer) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no anigamer provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(anigamerColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	anigamerInsertCacheMut.RLock()
	cache, cached := anigamerInsertCache[key]
	anigamerInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			anigamerAllColumns,
			anigamerColumnsWithDefault,
			anigamerColumnsWithoutDefault,
			nzDefaults,
		)
		wl = strmangle.SetComplement(wl, anigamerGeneratedColumns)

		cache.valueMapping, err = queries.BindMapping(anigamerType, anigamerMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(anigamerType, anigamerMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"anigamer\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"anigamer\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "models: unable to insert into anigamer")
	}

	if !cached {
		anigamerInsertCacheMut.Lock()
		anigamerInsertCache[key] = cache
		anigamerInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Anigamer.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Anigamer) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	anigamerUpdateCacheMut.RLock()
	cache, cached := anigamerUpdateCache[key]
	anigamerUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			anigamerAllColumns,
			anigamerPrimaryKeyColumns,
		)
		wl = strmangle.SetComplement(wl, anigamerGeneratedColumns)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update anigamer, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"anigamer\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 0, wl),
			strmangle.WhereClause("\"", "\"", 0, anigamerPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(anigamerType, anigamerMapping, append(wl, anigamerPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update anigamer row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for anigamer")
	}

	if !cached {
		anigamerUpdateCacheMut.Lock()
		anigamerUpdateCache[key] = cache
		anigamerUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q anigamerQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for anigamer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for anigamer")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AnigamerSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), anigamerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"anigamer\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, anigamerPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in anigamer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all anigamer")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Anigamer) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no anigamer provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(anigamerColumnsWithDefault, o)

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

	anigamerUpsertCacheMut.RLock()
	cache, cached := anigamerUpsertCache[key]
	anigamerUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, _ := insertColumns.InsertColumnSet(
			anigamerAllColumns,
			anigamerColumnsWithDefault,
			anigamerColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			anigamerAllColumns,
			anigamerPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert anigamer, could not build update column list")
		}

		ret := strmangle.SetComplement(anigamerAllColumns, strmangle.SetIntersect(insert, update))

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(anigamerPrimaryKeyColumns))
			copy(conflict, anigamerPrimaryKeyColumns)
		}
		cache.query = buildUpsertQuerySQLite(dialect, "\"anigamer\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(anigamerType, anigamerMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(anigamerType, anigamerMapping, ret)
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
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert anigamer")
	}

	if !cached {
		anigamerUpsertCacheMut.Lock()
		anigamerUpsertCache[key] = cache
		anigamerUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Anigamer record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Anigamer) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Anigamer provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), anigamerPrimaryKeyMapping)
	sql := "DELETE FROM \"anigamer\" WHERE \"anime_sn\"=?"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from anigamer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for anigamer")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q anigamerQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no anigamerQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from anigamer")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for anigamer")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AnigamerSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(anigamerBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), anigamerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"anigamer\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, anigamerPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from anigamer slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for anigamer")
	}

	if len(anigamerAfterDeleteHooks) != 0 {
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
func (o *Anigamer) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAnigamer(ctx, exec, o.AnimeSN)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AnigamerSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AnigamerSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), anigamerPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"anigamer\".* FROM \"anigamer\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, anigamerPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AnigamerSlice")
	}

	*o = slice

	return nil
}

// AnigamerExists checks if the Anigamer row exists.
func AnigamerExists(ctx context.Context, exec boil.ContextExecutor, animeSN int64) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"anigamer\" where \"anime_sn\"=? limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, animeSN)
	}
	row := exec.QueryRowContext(ctx, sql, animeSN)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if anigamer exists")
	}

	return exists, nil
}

// Exists checks if the Anigamer row exists.
func (o *Anigamer) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return AnigamerExists(ctx, exec, o.AnimeSN)
}
