package filter_test

import (
	"testing"

	"github.com/gookit/filter"
	"github.com/gookit/goutil/x/assert"
)

func TestApply(t *testing.T) {
	str := "  abc  "
	ret, err := filter.Apply("trim", str, nil)
	assert.NoErr(t, err)
	assert.Equal(t, "abc", ret)

	// test pointer string
	ps := &str
	ret, err = filter.Apply("trim", ps, nil)
	assert.NoErr(t, err)
	assert.Equal(t, "abc", ret)
}

func TestApplyTrimStringsCutSet(t *testing.T) {
	ret, err := filter.Apply("trimStrings", []string{"_first_", "__second__", " third "}, []string{"_"})
	assert.NoErr(t, err)
	assert.Equal(t, []string{"first", "second", " third "}, ret)
}

func TestFiltrationTrimStringsCutSet(t *testing.T) {
	f := filter.New(map[string]any{
		"tags": []string{"_first_", "__second__", " third "},
	})
	f.AddRule("tags", "trimStrings:_")
	assert.NoErr(t, f.Filtering())
	assert.Equal(t, []string{"first", "second", " third "}, f.SafeVal("tags"))
}
