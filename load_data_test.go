package filter_test

import (
	"strings"
	"testing"

	"github.com/gookit/filter"
)

func TestLoadDataReappliesRules(t *testing.T) {
	calls := 0
	f := filter.New(map[string]any{"value": " first "})
	f.AddRule("value", func(value any) (any, error) {
		calls++
		return strings.TrimSpace(value.(string)), nil
	})
	if err := f.Filtering(); err != nil {
		t.Fatal(err)
	}
	previous := f.CleanData()
	if err := f.Filtering(); err != nil || calls != 1 {
		t.Fatalf("unchanged input was processed again: calls=%d, err=%v", calls, err)
	}

	data := map[string]any{"value": " second "}
	f.LoadData(data)
	if len(f.CleanData()) != 0 {
		t.Fatal("LoadData retained results from the previous input")
	}
	if err := f.Filtering(); err != nil {
		t.Fatal(err)
	}
	if f.String("value") != "second" || calls != 2 {
		t.Fatalf("new input was not filtered: value=%q, calls=%d", f.String("value"), calls)
	}
	if previous["value"] != "first" || data["value"] != " second " {
		t.Fatal("loading or filtering mutated a previous result or the caller's input")
	}
	if err := f.Filtering(); err != nil || calls != 2 {
		t.Fatalf("reloaded input was processed again: calls=%d, err=%v", calls, err)
	}
}

func TestLoadDataClearsPreviousError(t *testing.T) {
	f := filter.New(map[string]any{"age": "bad"})
	f.AddRule("age", "int")
	if err := f.Filtering(); err == nil {
		t.Fatal("invalid integer was accepted")
	}
	f.LoadData(map[string]any{"age": "42"})
	if f.Err() != nil || !f.IsOK() {
		t.Fatal("LoadData retained the previous error")
	}
	if err := f.Filtering(); err != nil {
		t.Fatal(err)
	}
	if f.Int("age") != 42 {
		t.Fatalf("new integer was not filtered: %d", f.Int("age"))
	}
	f.LoadData(map[string]any{"age": "invalid"})
	if err := f.Filtering(); err == nil {
		t.Fatal("the previous success hid an error in the new input")
	}
	if _, ok := f.Safe("age"); ok {
		t.Fatal("the previous successful integer survived a failed reload")
	}
}

func TestLoadDataRefreshesMissingAndDefaultValues(t *testing.T) {
	f := filter.New(map[string]any{"removed": " old "})
	f.AddRule("removed", "trim")
	f.AddRule("added", "trim")
	f.AddRule("fallback", "trim").SetDefaultVal(" default ")
	if err := f.Filtering(); err != nil {
		t.Fatal(err)
	}
	f.LoadData(map[string]any{"added": " new ", "fallback": ""})
	if err := f.Filtering(); err != nil {
		t.Fatal(err)
	}
	if _, ok := f.Safe("removed"); ok {
		t.Fatal("a removed field retained its previous clean value")
	}
	if f.String("added") != "new" || f.String("fallback") != "" {
		t.Fatal("a newly present or explicitly empty field was not filtered")
	}
	f.LoadData(nil)
	if err := f.Filtering(); err != nil {
		t.Fatal(err)
	}
	if len(f.CleanData()) != 1 || f.String("fallback") != "default" {
		t.Fatalf("nil input did not reapply the retained default rule: %v", f.CleanData())
	}
}
