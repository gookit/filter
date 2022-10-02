package filter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
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
	is.Eq(2, intVal)

	intVal, err = ToInt("-2")
	is.Nil(err)
	is.Eq(-2, intVal)

	is.Eq(-2, MustInt("-2"))
	is.Eq(0, MustInt("2a"))
	is.Eq(0, MustInt(nil))
	for _, in := range tests {
		is.Eq(2, MustInt(in))
	}

	// To uint
	uintVal, err := Uint("2")
	is.Nil(err)
	is.Eq(uint64(2), uintVal)

	_, err = ToUint("-2")
	is.Err(err)

	is.Eq(uint64(0), MustUint("-2"))
	is.Eq(uint64(0), MustUint("2a"))
	is.Eq(uint64(0), MustUint(nil))
	for _, in := range tests {
		is.Eq(uint64(2), MustUint(in))
	}

	// To int64
	i64Val, err := ToInt64("2")
	is.Nil(err)
	is.Eq(int64(2), i64Val)

	i64Val, err = Int64("-2")
	is.Nil(err)
	is.Eq(int64(-2), i64Val)

	is.Eq(int64(0), MustInt64("2a"))
	is.Eq(int64(0), MustInt64(nil))
	for _, in := range tests {
		is.Eq(int64(2), MustInt64(in))
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
		is.Eq("2", MustString(in))
	}

	tests1 := []interface{}{
		float32(2.3), 2.3,
	}
	for _, in := range tests1 {
		is.Eq("2.3", MustString(in))
	}

	str, err := String(2.3)
	is.NoErr(err)
	is.Eq("2.3", str)

	str, err = String(nil)
	is.NoErr(err)
	is.Eq("", str)

	_, err = String([]string{"a"})
	is.Err(err)
}

func TestStrToFloat(t *testing.T) {
	is := assert.New(t)

	is.Eq(123.5, MustFloat("123.5"))
	is.Eq(float64(0), MustFloat("invalid"))

	fltVal, err := ToFloat("123.5")
	is.Nil(err)
	is.Eq(123.5, fltVal)

	fltVal, err = Float("-123.5")
	is.Nil(err)
	is.Eq(-123.5, fltVal)
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
		is.Eq(want, MustBool(str))
	}

	blVal, err := ToBool("1")
	is.Nil(err)
	is.True(blVal)

	blVal, err = Bool("10")
	is.Err(err)
	is.False(blVal)
}

func TestLowerOrUpperFirst(t *testing.T) {
	is := assert.New(t)

	// Uppercase, Lowercase
	is.Eq("ABC", Uppercase("abc"))
	is.Eq("ABC", Upper("abc"))
	is.Eq("abc", Lowercase("ABC"))
	is.Eq("abc", Lower("ABC"))

	tests := []string{
		"Abc-abc",
		"ABC-aBC",
	}

	// UpperFirst, LowerFirst
	for _, sample := range tests {
		ss := strings.Split(sample, "-")
		is.Eq(ss[0], UpperFirst(ss[1]))
		is.Eq(ss[1], LowerFirst(ss[0]))
	}

	is.Eq("", LowerFirst(""))
	is.Eq("", UpperFirst(""))
	is.Eq("abc", LowerFirst("abc"))
	is.Eq("Abc", UpperFirst("Abc"))

	// UpperWord
	is.Eq("", UpperWord(""))
	is.Eq("Hello World!", UpperWord("hello world!"))
}

func TestSnakeCase(t *testing.T) {
	is := assert.New(t)
	tests := map[string]string{
		"RangePrice":  "range_price",
		"rangePrice":  "range_price",
		"range_price": "range_price",
	}

	for sample, want := range tests {
		is.Eq(want, SnakeCase(sample))
	}

	is.Eq("range-price", Snake("rangePrice", "-"))
	is.Eq("range price", SnakeCase("rangePrice", " "))
}

func TestCamelCase(t *testing.T) {
	is := assert.New(t)
	tests := map[string]string{
		"rangePrice":   "rangePrice",
		"range_price":  "rangePrice",
		"_range_price": "RangePrice",
	}

	for sample, want := range tests {
		is.Eq(want, CamelCase(sample))
	}

	is.Eq("rangePrice", Camel("range-price", "-"))
	is.Eq("rangePrice", CamelCase("range price", " "))

	// custom sep char
	is.Eq("rangePrice", CamelCase("range+price", "+"))
	is.Eq("rangePrice", CamelCase("range*price", "*"))
}

func TestStrToInts(t *testing.T) {
	is := assert.New(t)

	ints, err := StrToInts("a,b,c")
	is.Err(err)
	is.Len(ints, 0)

	ints, err = StrToInts("1,2,3")
	is.Nil(err)
	is.Eq([]int{1, 2, 3}, ints)
}

func TestStr2Array(t *testing.T) {
	is := assert.New(t)

	ss := StrToArray("a,b,c", ",")
	is.Len(ss, 3)
	is.Eq(`[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))

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
		is.Eq(`[]string{"a", "b", "c"}`, fmt.Sprintf("%#v", ss))
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
	is.Eq("[]int{1, 2}", fmt.Sprintf("%#v", ints))

	_, err = StringsToInts([]string{"a", "b"})
	is.Err(err)
}

func TestEscape(t *testing.T) {
	tests := struct{ give, want string }{
		"<p>some text</p>",
		"&lt;p&gt;some text&lt;/p&gt;",
	}

	assert.Eq(t, tests.want, EscapeHTML(tests.give))

	tests = struct{ give, want string }{
		"<script>var a = 23;</script>",
		`\x3Cscript\x3Evar a = 23;\x3C/script\x3E`,
	}
	assert.NotEq(t, tests.give, EscapeJS(tests.give))
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
		is.Eq(want, tm.String())
	}

	tm, err := StrToTime("invalid")
	is.Err(err)
	is.True(tm.IsZero())

	tm, err = StrToTime("2018-09-27T15:34", "2018-09-27 15:34:23")
	is.Err(err)
	is.True(tm.IsZero())
}
