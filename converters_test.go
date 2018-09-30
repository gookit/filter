package filter

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestStrToInt(t *testing.T) {
	is := assert.New(t)

	intVal, err := Int("2")
	is.Nil(err)
	is.Equal(2, intVal)

	intVal, err = ToInt("-2")
	is.Nil(err)
	is.Equal(-2, intVal)

	is.Equal(2, MustInt("2"))
	is.Equal(-2, MustInt("-2"))
	is.Equal(0, MustInt("2a"))

	uintVal, err := Uint("2")
	is.Nil(err)
	is.Equal(uint64(2), uintVal)

	_, err = ToUint("-2")
	is.Error(err)

	is.Equal(uint64(0), MustUint("-2"))
	is.Equal(uint64(0), MustUint("2a"))

	i64Val, err := ToInt64("2")
	is.Nil(err)
	is.Equal(int64(2), i64Val)

	i64Val, err = Int64("-2")
	is.Nil(err)
	is.Equal(int64(-2), i64Val)

	is.Equal(int64(2), MustInt64("2"))
	is.Equal(int64(0), MustInt64("2a"))
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
	tests := []string{
		"Abc-abc",
		"ABC-aBC",
	}

	for _, sample := range tests {
		ss := strings.Split(sample, "-")
		is.Equal(ss[0], UpperFirst(ss[1]))
		is.Equal(ss[1], LowerFirst(ss[0]))
	}

	is.Equal("", LowerFirst(""))
	is.Equal("", UpperFirst(""))
	is.Equal("abc", LowerFirst("abc"))
	is.Equal("Abc", UpperFirst("Abc"))
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

	is.Equal("range-price", SnakeCase("rangePrice", "-"))
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

	is.Equal("rangePrice", CamelCase("range-price", "-"))
	is.Equal("rangePrice", CamelCase("range price", " "))

	// custom sep char
	is.Equal("rangePrice", CamelCase("range+price", "+"))
	is.Equal("rangePrice", CamelCase("range*price", "*"))
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

func TestStrToTime(t *testing.T) {
	is := assert.New(t)

	tm, err := StrToTime("20180927")
	is.Nil(err)
	is.Equal("2018-09-27 00:00:00 +0000 UTC", tm.String())

	tm, err = StrToTime("2018-09-27")
	is.Nil(err)
	is.Equal("2018-09-27 00:00:00 +0000 UTC", tm.String())

	tm, err = StrToTime("2018-09-27 15:34")
	is.Nil(err)
	is.Equal("2018-09-27 15:34:00 +0000 UTC", tm.String())

	tm, err = StrToTime("2018-09-27T15:34")
	is.Nil(err)
	is.Equal("2018-09-27 15:34:00 +0000 UTC", tm.String())

	tm, err = StrToTime("2018-09-27 15:34:23")
	is.Nil(err)
	is.Equal("2018-09-27 15:34:23 +0000 UTC", tm.String())

	tm, err = StrToTime("2018-09-27T15:34:23")
	is.Nil(err)
	is.False(tm.IsZero())
	is.Equal("2018-09-27 15:34:23 +0000 UTC", tm.String())

	tm, err = StrToTime("invalid")
	is.Error(err)
	is.True(tm.IsZero())

	tm, err = StrToTime("2018-09-27T15:34", "2018-09-27 15:34:23")
	is.Error(err)
	is.True(tm.IsZero())
}
