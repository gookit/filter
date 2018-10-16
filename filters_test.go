package filter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrim(t *testing.T) {
	is := assert.New(t)

	// Trim
	tests := map[string]string{
		"abc ":  "",
		" abc":  "",
		" abc ": "",
		"abc,,": ",",
		"abc,.": ",.",
	}
	for sample, cutSet := range tests {
		is.Equal("abc", Trim(sample, cutSet))
	}

	is.Equal("abc", Trim("abc,.", ".,"))
	// is.Equal("", Trim(nil))

	// TrimLeft
	is.Equal("abc ", TrimLeft(" abc "))
	is.Equal("abc ,", TrimLeft(", abc ,", " ,"))
	is.Equal("abc ,", TrimLeft(", abc ,", ", "))

	// TrimRight
	is.Equal(" abc", TrimRight(" abc "))
	is.Equal(", abc", TrimRight(", abc ,", ", "))

	// TrimStrings
	ss := TrimStrings([]string{" a", "b ", " c "})
	is.Equal("[a b c]", fmt.Sprint(ss))
	ss = TrimStrings([]string{",a", "b.", ",.c,"}, ",.")
	is.Equal("[a b c]", fmt.Sprint(ss))
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
	is.Equal("a.com", UrlEncode("a.com"))
	is.Equal("a.com", UrlDecode("a.com"))
}
