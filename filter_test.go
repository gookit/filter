package filter_test

import (
	"testing"

	"github.com/gookit/filter"
	"github.com/gookit/goutil/testutil/assert"
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
