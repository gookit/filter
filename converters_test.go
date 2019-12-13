package filter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValToInt(t *testing.T) {
	is := assert.New(t)

	tests := []interface{}{
		2,
		int8(2), int16(2), int32(2), int64(2),
		uint(2), uint8(2), uint16(2), uint32(2), uint64(2),
		float32(2.2), 2.3,
		"2",
	}

	// To int
	intVal, err := Int("2")
	is.Nil(err)
	is.Equal(2, intVal)

	intVal, err = ToInt("-2")
	is.Nil(err)
	is.Equal(-2, intVal)

	is.Equal(-2, MustInt("-2"))
	is.Equal(0, MustInt("2a"))
	is.Equal(0, MustInt(nil))
	for _, in := range tests {
		is.Equal(2, MustInt(in))
	}

	// To uint
	uintVal, err := Uint("2")
	is.Nil(err)
	is.Equal(uint64(2), uintVal)

	_, err = ToUint("-2")
	is.Error(err)

	is.Equal(uint64(0), MustUint("-2"))
	is.Equal(uint64(0), MustUint("2a"))
	is.Equal(uint64(0), MustUint(nil))
	for _, in := range tests {
		is.Equal(uint64(2), MustUint(in))
	}

	// To int64
	i64Val, err := ToInt64("2")
	is.Nil(err)
	is.Equal(int64(2), i64Val)

	i64Val, err = Int64("-2")
	is.Nil(err)
	is.Equal(int64(-2), i64Val)

	is.Equal(int64(0), MustInt64("2a"))
	is.Equal(int64(0), MustInt64(nil))
	for _, in := range tests {
		is.Equal(int64(2), MustInt64(in))
	}
}

func TestValToStr(t *testing.T) {
	is := assert.New(t)

	tests := []interface{}{
		2,
		int8(2), int16(2), int32(2), int64(2),
		uint(2), uint8(2), uint16(2), uint32(2), uint64(2),
		"2",
	}
	for _, in := range tests {
		is.Equal("2", MustString(in))
	}

	tests1 := []interface{}{
		float32(2.3), 2.3,
	}
	for _, in := range tests1 {
		is.Equal("2.3", MustString(in))
	}

	str, err := String(2.3)
	is.NoError(err)
	is.Equal("2.3", str)

	str, err = String(nil)
	is.NoError(err)
	is.Equal("", str)

	_, err = String([]string{"a"})
	is.Error(err)
}

func TestStrToFloat(t *testing.T) {
	is := assert.New(t)

	is.Equal(123.5, MustFloat("123.5"))
	is.Equal(float64(0), MustFloat("invalid"))

	fltVal, err := ToFloat("123.5")
	is.Nil(err)
	is.Equal(123.5, fltVal)

	fltVal, err = Float("-123.5")
	is.Nil(err)
	is.Equal(-123.5, fltVal)
}

func TestStrToBool(t *testing.T) {
	is := assert.New(t)

	tests1 := map[string]bool{
		"1":     true,
		"on":    true,
		"yes":   true,
		"true":  true,
		"false": false,
		"off":   false,
		"no":    false,
		"0":     false,
	}

	for str, want := range tests1 {
		is.Equal(want, MustBool(str))
	}

	blVal, err := ToBool("1")
	is.Nil(err)
	is.True(blVal)

	blVal, err = Bool("10")
	is.Error(err)
	is.False(blVal)
}

func TestLowerOrUpperFirst(t *testing.T) {
	is := assert.New(t)

	// Uppercase, Lowercase
	is.Equal("ABC", Uppercase("abc"))
	is.Equal("abc", Lowercase("ABC"))

	tests := []string{
		"Abc-abc",
		"ABC-aBC",
	}

	// UpperFirst, LowerFirst
	for _, sample := range tests {
		ss := strings.Split(sample, "-")
		is.Equal(ss[0], UpperFirst(ss[1]))
		is.Equal(ss[1], LowerFirst(ss[0]))
	}

	is.Equal("", LowerFirst(""))
	is.Equal("", UpperFirst(""))
	is.Equal("abc", LowerFirst("abc"))
	is.Equal("Abc", UpperFirst("Abc"))

	// UpperWord
	is.Equal("", UpperWord(""))
	is.Equal("Hello World!", UpperWord("hello world!"))
}

func TestSnakeCase(t *testing.T) {
	is := assert.New(t)
	tests := map[string]string{
		"RangePrice":  "range_price",
		"rangePrice":  "range_price",
		"range_price": "range_price",
	}

	for sample, want := range tests {
		is.Equal(want, SnakeCase(sample))
	}

	is.Equal("range-price", Snake("rangePrice", "-"))
	is.Equal("range price", SnakeCase("rangePrice", " "))
}

func TestCamelCase(t *testing.T) {
	is := assert.New(t)
	tests := map[string]string{
		"rangePrice":   "rangePrice",
		"range_price":  "rangePrice",
		"_range_price": "RangePrice",
	}

	for sample, want := range tests {
		is.Equal(want, CamelCase(sample))
	}

	is.Equal("rangePrice", Camel("range-price", "-"))
	is.Equal("rangePrice", CamelCase("range price", " "))

	// custom sep char
	is.Equal("rangePrice", CamelCase("range+price", "+"))
	is.Equal("rangePrice", CamelCase("range*price", "*"))
}

func TestStrToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := StrToInts("a,b,c")
	is.Error(err)
	is.Len(ints, 0)

	ints, err = StrToInts("1,2,3")
	is.Nil(err)
	is.Equal([]int{1, 2, 3}, ints)
}

func TestStr2Array(t *testing.T) {
	is := assert.New(t)

	ss := StrToArray("a,b,c", ",")
	is.Len(ss, 3)
	is.Equal(`[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

	tests := []string{
		// sample
		"a,b,c",
		"a,b,c,",
		",a,b,c",
		"a, b,c",
		"a,,b,c",
		"a, , b,c",
	}

	for _, sample := range tests {
		ss = StrToArray(sample)
		is.Equal(`[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))
	}

	ss = StrToSlice("", ",")
	is.Len(ss, 0)

	ss = StrToArray(", , ", ",")
	is.Len(ss, 0)
}

func TestStringsToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := StringsToInts([]string{"1", "2"})
	is.Nil(err)
	is.Equal("[]int{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = StringsToInts([]string{"a", "b"})
	is.Error(err)
}

func TestEscape(t *testing.T) {
	tests := struct{ give, want string }{
		"<p>some text</p>",
		"&lt;p&gt;some text&lt;/p&gt;",
	}

	assert.Equal(t, tests.want, EscapeHTML(tests.give))

	tests = struct{ give, want string }{
		"<script>var a = 23;</script>",
		`\x3Cscript\x3Evar a = 23;\x3C/script\x3E`,
	}
	assert.Equal(t, tests.want, EscapeJS(tests.give))
}

func TestStrToTime(t *testing.T) {
	is := assert.New(t)
	tests := map[string]string{
		"20180927":             "2018-09-27 00:00:00 +0000 UTC",
		"2018-09-27":           "2018-09-27 00:00:00 +0000 UTC",
		"2018-09-27 12":        "2018-09-27 12:00:00 +0000 UTC",
		"2018-09-27T12":        "2018-09-27 12:00:00 +0000 UTC",
		"2018-09-27 12:34":     "2018-09-27 12:34:00 +0000 UTC",
		"2018-09-27T12:34":     "2018-09-27 12:34:00 +0000 UTC",
		"2018-09-27 12:34:45":  "2018-09-27 12:34:45 +0000 UTC",
		"2018-09-27T12:34:45":  "2018-09-27 12:34:45 +0000 UTC",
		"2018/09/27 12:34:45":  "2018-09-27 12:34:45 +0000 UTC",
		"2018/09/27T12:34:45Z": "2018-09-27 12:34:45 +0000 UTC",
	}

	for sample, want := range tests {
		tm, err := StrToTime(sample)
		is.Nil(err)
		is.Equal(want, tm.String())
	}

	tm, err := StrToTime("invalid")
	is.Error(err)
	is.True(tm.IsZero())

	tm, err = StrToTime("2018-09-27T15:34", "2018-09-27 15:34:23")
	is.Error(err)
	is.True(tm.IsZero())
}
