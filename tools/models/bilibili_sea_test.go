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

func testBilibiliSeas(t *testing.T) {
	t.Parallel()

	query := BilibiliSeas()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testBilibiliSeasDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
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

	count, err := BilibiliSeas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBilibiliSeasQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := BilibiliSeas().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BilibiliSeas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBilibiliSeasSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BilibiliSeaSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := BilibiliSeas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testBilibiliSeasExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := BilibiliSeaExists(ctx, tx, o.SeasonID)
	if err != nil {
		t.Errorf("Unable to check if BilibiliSea exists: %s", err)
	}
	if !e {
		t.Errorf("Expected BilibiliSeaExists to return true, but got false.")
	}
}

func testBilibiliSeasFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	bilibiliSeaFound, err := FindBilibiliSea(ctx, tx, o.SeasonID)
	if err != nil {
		t.Error(err)
	}

	if bilibiliSeaFound == nil {
		t.Error("want a record, got nil")
	}
}

func testBilibiliSeasBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = BilibiliSeas().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testBilibiliSeasOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := BilibiliSeas().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testBilibiliSeasAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	bilibiliSeaOne := &BilibiliSea{}
	bilibiliSeaTwo := &BilibiliSea{}
	if err = randomize.Struct(seed, bilibiliSeaOne, bilibiliSeaDBTypes, false, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}
	if err = randomize.Struct(seed, bilibiliSeaTwo, bilibiliSeaDBTypes, false, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bilibiliSeaOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bilibiliSeaTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BilibiliSeas().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testBilibiliSeasCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	bilibiliSeaOne := &BilibiliSea{}
	bilibiliSeaTwo := &BilibiliSea{}
	if err = randomize.Struct(seed, bilibiliSeaOne, bilibiliSeaDBTypes, false, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}
	if err = randomize.Struct(seed, bilibiliSeaTwo, bilibiliSeaDBTypes, false, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = bilibiliSeaOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = bilibiliSeaTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BilibiliSeas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func bilibiliSeaBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *BilibiliSea) error {
	*o = BilibiliSea{}
	return nil
}

func bilibiliSeaAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *BilibiliSea) error {
	*o = BilibiliSea{}
	return nil
}

func bilibiliSeaAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *BilibiliSea) error {
	*o = BilibiliSea{}
	return nil
}

func bilibiliSeaBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BilibiliSea) error {
	*o = BilibiliSea{}
	return nil
}

func bilibiliSeaAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *BilibiliSea) error {
	*o = BilibiliSea{}
	return nil
}

func bilibiliSeaBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BilibiliSea) error {
	*o = BilibiliSea{}
	return nil
}

func bilibiliSeaAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *BilibiliSea) error {
	*o = BilibiliSea{}
	return nil
}

func bilibiliSeaBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BilibiliSea) error {
	*o = BilibiliSea{}
	return nil
}

func bilibiliSeaAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *BilibiliSea) error {
	*o = BilibiliSea{}
	return nil
}

func testBilibiliSeasHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &BilibiliSea{}
	o := &BilibiliSea{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, false); err != nil {
		t.Errorf("Unable to randomize BilibiliSea object: %s", err)
	}

	AddBilibiliSeaHook(boil.BeforeInsertHook, bilibiliSeaBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	bilibiliSeaBeforeInsertHooks = []BilibiliSeaHook{}

	AddBilibiliSeaHook(boil.AfterInsertHook, bilibiliSeaAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	bilibiliSeaAfterInsertHooks = []BilibiliSeaHook{}

	AddBilibiliSeaHook(boil.AfterSelectHook, bilibiliSeaAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	bilibiliSeaAfterSelectHooks = []BilibiliSeaHook{}

	AddBilibiliSeaHook(boil.BeforeUpdateHook, bilibiliSeaBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	bilibiliSeaBeforeUpdateHooks = []BilibiliSeaHook{}

	AddBilibiliSeaHook(boil.AfterUpdateHook, bilibiliSeaAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	bilibiliSeaAfterUpdateHooks = []BilibiliSeaHook{}

	AddBilibiliSeaHook(boil.BeforeDeleteHook, bilibiliSeaBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	bilibiliSeaBeforeDeleteHooks = []BilibiliSeaHook{}

	AddBilibiliSeaHook(boil.AfterDeleteHook, bilibiliSeaAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	bilibiliSeaAfterDeleteHooks = []BilibiliSeaHook{}

	AddBilibiliSeaHook(boil.BeforeUpsertHook, bilibiliSeaBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	bilibiliSeaBeforeUpsertHooks = []BilibiliSeaHook{}

	AddBilibiliSeaHook(boil.AfterUpsertHook, bilibiliSeaAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	bilibiliSeaAfterUpsertHooks = []BilibiliSeaHook{}
}

func testBilibiliSeasInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BilibiliSeas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBilibiliSeasInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(bilibiliSeaColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := BilibiliSeas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testBilibiliSeasReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
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

func testBilibiliSeasReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := BilibiliSeaSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testBilibiliSeasSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := BilibiliSeas().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	bilibiliSeaDBTypes = map[string]string{`Cover`: `VARCHAR(128)`, `IndexShow`: `VARCHAR(16)`, `IsFinish`: `BOOLEAN`, `SeasonID`: `INTEGER`, `SeasonType`: `INTEGER`, `Title`: `VARCHAR(128)`, `UpdatedAt`: `DATETIME`, `CreatedAt`: `DATETIME`}
	_                  = bytes.MinRead
)

func testBilibiliSeasUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(bilibiliSeaPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(bilibiliSeaAllColumns) == len(bilibiliSeaPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BilibiliSeas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testBilibiliSeasSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(bilibiliSeaAllColumns) == len(bilibiliSeaPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &BilibiliSea{}
	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := BilibiliSeas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, bilibiliSeaDBTypes, true, bilibiliSeaPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(bilibiliSeaAllColumns, bilibiliSeaPrimaryKeyColumns) {
		fields = bilibiliSeaAllColumns
	} else {
		fields = strmangle.SetComplement(
			bilibiliSeaAllColumns,
			bilibiliSeaPrimaryKeyColumns,
		)
		fields = strmangle.SetComplement(fields, bilibiliSeaGeneratedColumns)
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

	slice := BilibiliSeaSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testBilibiliSeasUpsert(t *testing.T) {
	t.Parallel()
	if len(bilibiliSeaAllColumns) == len(bilibiliSeaPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := BilibiliSea{}
	if err = randomize.Struct(seed, &o, bilibiliSeaDBTypes, true); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BilibiliSea: %s", err)
	}

	count, err := BilibiliSeas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, bilibiliSeaDBTypes, false, bilibiliSeaPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize BilibiliSea struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert BilibiliSea: %s", err)
	}

	count, err = BilibiliSeas().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
