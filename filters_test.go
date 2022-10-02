package filter

import (
	"fmt"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
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
		is.Eq("abc", Trim(sample, cutSet))
	}

	is.Eq("abc", Trim("abc,.", ".,"))
	// is.Eq("", Trim(nil))

	// TrimLeft
	is.Eq("abc ", TrimLeft(" abc "))
	is.Eq("abc ,", TrimLeft(", abc ,", " ,"))
	is.Eq("abc ,", TrimLeft(", abc ,", ", "))

	// TrimRight
	is.Eq(" abc", TrimRight(" abc "))
	is.Eq(", abc", TrimRight(", abc ,", ", "))

	// TrimStrings
	ss := TrimStrings([]string{" a", "b ", " c "})
	is.Eq("[a b c]", fmt.Sprint(ss))
	ss = TrimStrings([]string{",a", "b.", ",.c,"}, ",.")
	is.Eq("[a b c]", fmt.Sprint(ss))
}

func TestEmail(t *testing.T) {
	is := assert.New(t)
	is.Eq("THE@inhere.com", Email("   THE@INHere.com  "))
	is.Eq("inhere.xyz", Email("   inhere.xyz  "))
}

func TestStrOperate(t *testing.T) {
	is := assert.New(t)

	// Substr
	is.Eq("DEF", Substr("abcDEF", 3, 3))
	is.Eq("DEF", Substr("abcDEF", 3, 5))
	is.Eq("", Substr("abcDEF", 23, 5))
}

func TestURLEnDecode(t *testing.T) {
	is := assert.New(t)

	is.Eq("a.com/?name%3D%E4%BD%A0%E5%A5%BD", URLEncode("a.com/?name=你好"))
	is.Eq("a.com/?name=你好", URLDecode("a.com/?name%3D%E4%BD%A0%E5%A5%BD"))
	is.Eq("a.com", URLEncode("a.com"))
	is.Eq("a.com", URLDecode("a.com"))
}

func TestUnique(t *testing.T) {
	is := assert.New(t)

	is.Len(Unique([]int{1, 2}), 2)
	is.Len(Unique([]int{1, 2, 2, 1}), 2)
	is.Len(Unique([]int64{1, 2}), 2)
	is.Len(Unique([]int64{1, 2, 2, 1}), 2)
	is.Len(Unique([]string{"a", "b"}), 2)
	is.Len(Unique([]string{"a", "b", "b"}), 2)
	is.Eq("invalid", Unique("invalid"))
}
