package filter

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
)

// Some alias methods.
var (
	Lower = strings.ToLower
	Upper = strings.ToUpper
	Title = strings.ToTitle
	// escape string.
	EscapeJS   = template.JSEscapeString
	EscapeHTML = template.HTMLEscapeString
)

// Some regex for convert string.
var (
	toSnakeReg  = regexp.MustCompile("[A-Z][a-z]")
	toCamelRegs = map[string]*regexp.Regexp{
		" ": regexp.MustCompile(" +[a-zA-Z]"),
		"-": regexp.MustCompile("-+[a-zA-Z]"),
		"_": regexp.MustCompile("_+[a-zA-Z]"),
	}
	errConvertFail  = errors.New("convert data is failure")
	errInvalidParam = errors.New("invalid input parameter")
)

/*************************************************************
 * string to int,bool,float
 *************************************************************/

// Int convert string to int
func Int(in interface{}) (int, error) {
	return ToInt(in)
}

// MustInt convert string to int
func MustInt(in interface{}) int {
	val, _ := ToInt(in)
	return val
}

// ToInt convert string to int
func ToInt(in interface{}) (iVal int, err error) {
	switch tVal := in.(type) {
	case int:
		iVal = tVal
	case int8:
		iVal = int(tVal)
	case int16:
		iVal = int(tVal)
	case int32:
		iVal = int(tVal)
	case int64:
		iVal = int(tVal)
	case uint:
		iVal = int(tVal)
	case uint8:
		iVal = int(tVal)
	case uint16:
		iVal = int(tVal)
	case uint32:
		iVal = int(tVal)
	case uint64:
		iVal = int(tVal)
	case float32:
		iVal = int(tVal)
	case float64:
		iVal = int(tVal)
	case string:
		iVal, err = strconv.Atoi(Trim(tVal))
	default:
		err = errConvertFail
	}

	return
}

// Uint convert string to uint
func Uint(in interface{}) (uint64, error) {
	return ToUint(in)
}

// MustUint convert string to uint
func MustUint(in interface{}) uint64 {
	val, _ := ToUint(in)
	return val
}

// ToUint convert string to uint
func ToUint(in interface{}) (u64 uint64, err error) {
	switch tVal := in.(type) {
	case int:
		u64 = uint64(tVal)
	case int8:
		u64 = uint64(tVal)
	case int16:
		u64 = uint64(tVal)
	case int32:
		u64 = uint64(tVal)
	case int64:
		u64 = uint64(tVal)
	case uint:
		u64 = uint64(tVal)
	case uint8:
		u64 = uint64(tVal)
	case uint16:
		u64 = uint64(tVal)
	case uint32:
		u64 = uint64(tVal)
	case uint64:
		u64 = tVal
	case float32:
		u64 = uint64(tVal)
	case float64:
		u64 = uint64(tVal)
	case string:
		u64, err = strconv.ParseUint(Trim(tVal), 10, 0)
	default:
		err = errConvertFail
	}

	return
}

// Int64 convert string to int64
func Int64(in interface{}) (int64, error) {
	return ToInt64(in)
}

// ToInt64 convert string to int64
func ToInt64(in interface{}) (i64 int64, err error) {
	switch tVal := in.(type) {
	case string:
		i64, err = strconv.ParseInt(Trim(tVal), 10, 0)
	case int:
		i64 = int64(tVal)
	case int8:
		i64 = int64(tVal)
	case int16:
		i64 = int64(tVal)
	case int32:
		i64 = int64(tVal)
	case int64:
		i64 = tVal
	case uint:
		i64 = int64(tVal)
	case uint8:
		i64 = int64(tVal)
	case uint16:
		i64 = int64(tVal)
	case uint32:
		i64 = int64(tVal)
	case uint64:
		i64 = int64(tVal)
	case float32:
		i64 = int64(tVal)
	case float64:
		i64 = int64(tVal)
	default:
		err = errConvertFail
	}

	return
}

// MustInt64 convert
func MustInt64(in interface{}) int64 {
	i64, _ := ToInt64(in)
	return i64
}

// Float convert string to float
func Float(s string) (float64, error) {
	return ToFloat(s)
}

// ToFloat convert string to float
func ToFloat(s string) (float64, error) {
	return strconv.ParseFloat(Trim(s), 0)
}

// MustFloat convert string to float
func MustFloat(s string) float64 {
	val, _ := strconv.ParseFloat(Trim(s), 0)
	return val
}

// ToBool convert string to bool
func ToBool(s string) (bool, error) {
	return Bool(s)
}

// Bool parse string to bool
func Bool(s string) (bool, error) {
	// return strconv.ParseBool(Trim(s))
	lower := strings.ToLower(s)
	switch lower {
	case "1", "on", "yes", "true":
		return true, nil
	case "0", "off", "no", "false":
		return false, nil
	}

	return false, fmt.Errorf("'%s' cannot convert to bool", s)
}

// MustBool convert.
func MustBool(s string) bool {
	val, _ := Bool(Trim(s))
	return val
}

/*************************************************************
 * change string case
 *************************************************************/

// Lowercase alias of the strings.ToLower()
func Lowercase(s string) string {
	return strings.ToLower(s)
}

// Uppercase alias of the strings.ToUpper()
func Uppercase(s string) string {
	return strings.ToUpper(s)
}

// UpperWord Change the first character of each word to uppercase
func UpperWord(s string) string {
	if len(s) == 0 {
		return s
	}

	ss := strings.Split(s, " ")
	ns := make([]string, len(ss))
	for i, word := range ss {
		ns[i] = UpperFirst(word)
	}
	return strings.Join(ns, " ")
}

// LowerFirst lower first char
func LowerFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	f := s[0]
	if f >= 'A' && f <= 'Z' {
		return strings.ToLower(string(f)) + string(s[1:])
	}
	return s
}

// UpperFirst upper first char
func UpperFirst(s string) string {
	if len(s) == 0 {
		return s
	}

	f := s[0]
	if f >= 'a' && f <= 'z' {
		return strings.ToUpper(string(f)) + string(s[1:])
	}
	return s
}

// Snake alias of the SnakeCase
func Snake(s string, sep ...string) string {
	return SnakeCase(s, sep...)
}

// SnakeCase convert. eg "RangePrice" -> "range_price"
func SnakeCase(s string, sep ...string) string {
	sepChar := "_"
	if len(sep) > 0 {
		sepChar = sep[0]
	}

	newStr := toSnakeReg.ReplaceAllStringFunc(s, func(s string) string {
		return sepChar + LowerFirst(s)
	})

	return strings.TrimLeft(newStr, sepChar)
}

// Camel alias of the CamelCase
func Camel(s string, sep ...string) string {
	return CamelCase(s, sep...)
}

// CamelCase convert string to camel case.
// Support:
// 	"range_price" -> "rangePrice"
// 	"range price" -> "rangePrice"
// 	"range-price" -> "rangePrice"
func CamelCase(s string, sep ...string) string {
	sepChar := "_"
	if len(sep) > 0 {
		sepChar = sep[0]
	}

	// Not contains sep char
	if !strings.Contains(s, sepChar) {
		return s
	}

	// Get regexp instance
	rgx, ok := toCamelRegs[sepChar]
	if !ok {
		rgx = regexp.MustCompile(regexp.QuoteMeta(sepChar) + "+[a-zA-Z]")
	}

	return rgx.ReplaceAllStringFunc(s, func(s string) string {
		s = strings.TrimLeft(s, sepChar)
		return UpperFirst(s)
	})
}

/*************************************************************
 * string to slice, time
 *************************************************************/

// StrToInts split string to slice and convert item to int.
func StrToInts(s string, sep ...string) (ints []int, err error) {
	ss := StrToSlice(s, sep...)
	for _, item := range ss {
		iVal, err := ToInt(item)
		if err != nil {
			return []int{}, err
		}

		ints = append(ints, iVal)
	}

	return
}

// StrToArray alias of the StrToSlice()
func StrToArray(s string, sep ...string) []string {
	return StrToSlice(s, sep...)
}

// StrToSlice split string to array.
func StrToSlice(s string, sep ...string) []string {
	if len(sep) > 0 {
		return stringSplit(s, sep[0])
	}

	return stringSplit(s, ",")
}

// StringsToInts string slice to int slice
func StringsToInts(ss []string) (ints []int, err error) {
	for _, str := range ss {
		iVal, err := strconv.Atoi(str)
		if err != nil {
			return []int{}, err
		}

		ints = append(ints, iVal)
	}

	return
}

// StrToTime convert date string to time.Time
func StrToTime(s string, layouts ...string) (t time.Time, err error) {
	var layout string
	if len(layouts) > 0 { // custom layout
		layout = layouts[0]
	} else { // auto match layout.
		switch len(s) {
		case 8:
			layout = "20060102"
		case 10:
			layout = "2006-01-02"
		case 13:
			layout = "2006-01-02 15"
		case 16:
			layout = "2006-01-02 15:04"
		case 19:
			layout = "2006-01-02 15:04:05"
		case 20: // time.RFC3339
			layout = "2006-01-02T15:04:05Z07:00"
		}
	}

	if layout == "" {
		err = errInvalidParam
		return
	}

	// has 'T' eg.2006-01-02T15:04:05
	if strings.ContainsRune(s, 'T') {
		layout = strings.Replace(layout, " ", "T", -1)
	}

	// eg: 2006/01/02 15:04:05
	if strings.ContainsRune(s, '/') {
		layout = strings.Replace(layout, "-", "/", -1)
	}

	t, err = time.Parse(layout, s)
	// t, err = time.ParseInLocation(layout, s, time.Local)
	return
}
