package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrim(t *testing.T) {
	is := assert.New(t)

	is.Equal("abc", Trim("abc "))
	is.Equal("abc", Trim(" abc"))
	is.Equal("abc", Trim(" abc "))

	is.Equal("abc", Trim("abc,,", ","))
	is.Equal("abc", Trim("abc,.", ",."))
	is.Equal("abc", Trim("abc,.", ".,"))

	is.Equal("abc ", TrimLeft(" abc "))
	is.Equal(" abc", TrimRight(" abc "))
}

func TestEmail(t *testing.T) {
	is := assert.New(t)
	is.Equal("THE@inhere.com", Email("   THE@INHere.com  "))
	is.Equal("inhere.xyz", Email("   inhere.xyz  "))
}
