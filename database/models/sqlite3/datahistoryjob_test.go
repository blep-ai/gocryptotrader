// Code generated by SQLBoiler 3.5.0-gct (https://github.com/thrasher-corp/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlite3

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/thrasher-corp/sqlboiler/boil"
	"github.com/thrasher-corp/sqlboiler/queries"
	"github.com/thrasher-corp/sqlboiler/randomize"
	"github.com/thrasher-corp/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testDatahistoryjobs(t *testing.T) {
	t.Parallel()

	query := Datahistoryjobs()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testDatahistoryjobsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
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

	count, err := Datahistoryjobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDatahistoryjobsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Datahistoryjobs().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Datahistoryjobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDatahistoryjobsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DatahistoryjobSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Datahistoryjobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testDatahistoryjobsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := DatahistoryjobExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Datahistoryjob exists: %s", err)
	}
	if !e {
		t.Errorf("Expected DatahistoryjobExists to return true, but got false.")
	}
}

func testDatahistoryjobsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	datahistoryjobFound, err := FindDatahistoryjob(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if datahistoryjobFound == nil {
		t.Error("want a record, got nil")
	}
}

func testDatahistoryjobsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Datahistoryjobs().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testDatahistoryjobsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Datahistoryjobs().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testDatahistoryjobsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	datahistoryjobOne := &Datahistoryjob{}
	datahistoryjobTwo := &Datahistoryjob{}
	if err = randomize.Struct(seed, datahistoryjobOne, datahistoryjobDBTypes, false, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}
	if err = randomize.Struct(seed, datahistoryjobTwo, datahistoryjobDBTypes, false, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = datahistoryjobOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = datahistoryjobTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Datahistoryjobs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testDatahistoryjobsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	datahistoryjobOne := &Datahistoryjob{}
	datahistoryjobTwo := &Datahistoryjob{}
	if err = randomize.Struct(seed, datahistoryjobOne, datahistoryjobDBTypes, false, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}
	if err = randomize.Struct(seed, datahistoryjobTwo, datahistoryjobDBTypes, false, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = datahistoryjobOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = datahistoryjobTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Datahistoryjobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func datahistoryjobBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Datahistoryjob) error {
	*o = Datahistoryjob{}
	return nil
}

func datahistoryjobAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Datahistoryjob) error {
	*o = Datahistoryjob{}
	return nil
}

func datahistoryjobAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Datahistoryjob) error {
	*o = Datahistoryjob{}
	return nil
}

func datahistoryjobBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Datahistoryjob) error {
	*o = Datahistoryjob{}
	return nil
}

func datahistoryjobAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Datahistoryjob) error {
	*o = Datahistoryjob{}
	return nil
}

func datahistoryjobBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Datahistoryjob) error {
	*o = Datahistoryjob{}
	return nil
}

func datahistoryjobAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Datahistoryjob) error {
	*o = Datahistoryjob{}
	return nil
}

func datahistoryjobBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Datahistoryjob) error {
	*o = Datahistoryjob{}
	return nil
}

func datahistoryjobAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Datahistoryjob) error {
	*o = Datahistoryjob{}
	return nil
}

func testDatahistoryjobsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Datahistoryjob{}
	o := &Datahistoryjob{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob object: %s", err)
	}

	AddDatahistoryjobHook(boil.BeforeInsertHook, datahistoryjobBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	datahistoryjobBeforeInsertHooks = []DatahistoryjobHook{}

	AddDatahistoryjobHook(boil.AfterInsertHook, datahistoryjobAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	datahistoryjobAfterInsertHooks = []DatahistoryjobHook{}

	AddDatahistoryjobHook(boil.AfterSelectHook, datahistoryjobAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	datahistoryjobAfterSelectHooks = []DatahistoryjobHook{}

	AddDatahistoryjobHook(boil.BeforeUpdateHook, datahistoryjobBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	datahistoryjobBeforeUpdateHooks = []DatahistoryjobHook{}

	AddDatahistoryjobHook(boil.AfterUpdateHook, datahistoryjobAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	datahistoryjobAfterUpdateHooks = []DatahistoryjobHook{}

	AddDatahistoryjobHook(boil.BeforeDeleteHook, datahistoryjobBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	datahistoryjobBeforeDeleteHooks = []DatahistoryjobHook{}

	AddDatahistoryjobHook(boil.AfterDeleteHook, datahistoryjobAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	datahistoryjobAfterDeleteHooks = []DatahistoryjobHook{}

	AddDatahistoryjobHook(boil.BeforeUpsertHook, datahistoryjobBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	datahistoryjobBeforeUpsertHooks = []DatahistoryjobHook{}

	AddDatahistoryjobHook(boil.AfterUpsertHook, datahistoryjobAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	datahistoryjobAfterUpsertHooks = []DatahistoryjobHook{}
}

func testDatahistoryjobsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Datahistoryjobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDatahistoryjobsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(datahistoryjobColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Datahistoryjobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testDatahistoryjobToManyJobDatahistoryjobresults(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Datahistoryjob
	var b, c Datahistoryjobresult

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, datahistoryjobresultDBTypes, false, datahistoryjobresultColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, datahistoryjobresultDBTypes, false, datahistoryjobresultColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.JobID = a.ID
	c.JobID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.JobDatahistoryjobresults().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.JobID == b.JobID {
			bFound = true
		}
		if v.JobID == c.JobID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := DatahistoryjobSlice{&a}
	if err = a.L.LoadJobDatahistoryjobresults(ctx, tx, false, (*[]*Datahistoryjob)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.JobDatahistoryjobresults); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.JobDatahistoryjobresults = nil
	if err = a.L.LoadJobDatahistoryjobresults(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.JobDatahistoryjobresults); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testDatahistoryjobToManyAddOpJobDatahistoryjobresults(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Datahistoryjob
	var b, c, d, e Datahistoryjobresult

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, datahistoryjobDBTypes, false, strmangle.SetComplement(datahistoryjobPrimaryKeyColumns, datahistoryjobColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Datahistoryjobresult{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, datahistoryjobresultDBTypes, false, strmangle.SetComplement(datahistoryjobresultPrimaryKeyColumns, datahistoryjobresultColumnsWithoutDefault)...); err != nil {
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

	foreignersSplitByInsertion := [][]*Datahistoryjobresult{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddJobDatahistoryjobresults(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.JobID {
			t.Error("foreign key was wrong value", a.ID, first.JobID)
		}
		if a.ID != second.JobID {
			t.Error("foreign key was wrong value", a.ID, second.JobID)
		}

		if first.R.Job != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Job != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.JobDatahistoryjobresults[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.JobDatahistoryjobresults[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.JobDatahistoryjobresults().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testDatahistoryjobToOneExchangeUsingExchangeName(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Datahistoryjob
	var foreign Exchange

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, datahistoryjobDBTypes, false, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, exchangeDBTypes, false, exchangeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Exchange struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ExchangeNameID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.ExchangeName().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := DatahistoryjobSlice{&local}
	if err = local.L.LoadExchangeName(ctx, tx, false, (*[]*Datahistoryjob)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ExchangeName == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.ExchangeName = nil
	if err = local.L.LoadExchangeName(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.ExchangeName == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testDatahistoryjobToOneSetOpExchangeUsingExchangeName(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Datahistoryjob
	var b, c Exchange

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, datahistoryjobDBTypes, false, strmangle.SetComplement(datahistoryjobPrimaryKeyColumns, datahistoryjobColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, exchangeDBTypes, false, strmangle.SetComplement(exchangePrimaryKeyColumns, exchangeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, exchangeDBTypes, false, strmangle.SetComplement(exchangePrimaryKeyColumns, exchangeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Exchange{&b, &c} {
		err = a.SetExchangeName(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.ExchangeName != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.ExchangeNameDatahistoryjobs[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ExchangeNameID != x.ID {
			t.Error("foreign key was wrong value", a.ExchangeNameID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ExchangeNameID))
		reflect.Indirect(reflect.ValueOf(&a.ExchangeNameID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ExchangeNameID != x.ID {
			t.Error("foreign key was wrong value", a.ExchangeNameID, x.ID)
		}
	}
}

func testDatahistoryjobsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
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

func testDatahistoryjobsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := DatahistoryjobSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testDatahistoryjobsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Datahistoryjobs().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	datahistoryjobDBTypes = map[string]string{`ID`: `TEXT`, `Nickname`: `TEXT`, `ExchangeNameID`: `TEXT`, `Asset`: `TEXT`, `Base`: `TEXT`, `Quote`: `TEXT`, `StartTime`: `TIMESTAMP`, `EndTime`: `TIMESTAMP`, `Interval`: `REAL`, `DataType`: `REAL`, `RequestSize`: `REAL`, `MaxRetries`: `REAL`, `BatchCount`: `REAL`, `Status`: `REAL`, `Created`: `TIMESTAMP`}
	_                     = bytes.MinRead
)

func testDatahistoryjobsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(datahistoryjobPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(datahistoryjobAllColumns) == len(datahistoryjobPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Datahistoryjobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testDatahistoryjobsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(datahistoryjobAllColumns) == len(datahistoryjobPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Datahistoryjob{}
	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Datahistoryjobs().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, datahistoryjobDBTypes, true, datahistoryjobPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Datahistoryjob struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(datahistoryjobAllColumns, datahistoryjobPrimaryKeyColumns) {
		fields = datahistoryjobAllColumns
	} else {
		fields = strmangle.SetComplement(
			datahistoryjobAllColumns,
			datahistoryjobPrimaryKeyColumns,
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

	slice := DatahistoryjobSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}
