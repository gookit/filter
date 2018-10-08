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

	// TrimLeft
	is.Equal("abc ", TrimLeft(" abc "))
	is.Equal("abc ,", TrimLeft(", abc ,", " ,"))
	is.Equal("abc ,", TrimLeft(", abc ,", ", "))

	// TrimRight
	is.Equal(" abc", TrimRight(" abc "))
	is.Equal(", abc", TrimRight(", abc ,", ", "))
}

func TestEmail(t *testing.T) {
	is := assert.New(t)
	is.Equal("THE@inhere.com", Email("   THE@INHere.com  "))
	is.Equal("inhere.xyz", Email("   inhere.xyz  "))
}

func TestStrOperate(t *testing.T) {
	is := assert.New(t)

	// Substr
	is.Equal("DEF", Substr("abcDEF", 3, 3))
	is.Equal("DEF", Substr("abcDEF", 3, 5))
	is.Equal("", Substr("abcDEF", 23, 5))
}

func TestUrlEncode(t *testing.T) {
	is := assert.New(t)

	is.Equal("a.com/?name%3D%E4%BD%A0%E5%A5%BD", UrlEncode("a.com/?name=你好"))
	is.Equal("a.com/?name=你好", UrlDecode("a.com/?name%3D%E4%BD%A0%E5%A5%BD"))
}

func TestFiltration(t *testing.T) {
	is := assert.New(t)

	fl := New(map[string]interface{}{
		"key0": " abc ",
		"key1": "2",
	})

	is.Equal("", fl.Trimmed("not-exist"))
	is.Equal("abc", fl.Trimmed("key0"))
	is.Equal(2, fl.Int("key1"))
	is.Equal(0, fl.Int("not-exist"))
}
