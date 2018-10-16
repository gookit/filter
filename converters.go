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

// some alias methods.
var (
	Lower = strings.ToLower
	Upper = strings.ToUpper
	Title = strings.ToTitle
	// escape string.
	EscapeJS   = template.JSEscapeString
	EscapeHTML = template.HTMLEscapeString
)

// some regex for convert string.
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

// Int convert
func Int(in interface{}) (int, error) {
	return ToInt(in)
}

// MustInt convert
func MustInt(in interface{}) int {
	val, _ := ToInt(in)
	return val
}

// ToInt convert
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

// Uint convert
func Uint(in interface{}) (uint64, error) {
	return ToUint(in)
}

// MustUint convert
func MustUint(in interface{}) uint64 {
	val, _ := ToUint(in)
	return val
}

// ToUint convert
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

// Int64 convert
func Int64(in interface{}) (int64, error) {
	return ToInt64(in)
}

// ToInt64 convert
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

// Float convert
func Float(str string) (float64, error) {
	return ToFloat(str)
}

// ToFloat convert
func ToFloat(str string) (float64, error) {
	return strconv.ParseFloat(Trim(str), 0)
}

// MustFloat convert
func MustFloat(str string) float64 {
	val, _ := strconv.ParseFloat(Trim(str), 0)
	return val
}

// ToBool convert.
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

// SnakeCase convert. eg "RangePrice" -> "range_price"
func SnakeCase(str string, sep ...string) string {
	sepChar := "_"
	if len(sep) > 0 {
		sepChar = sep[0]
	}

	newStr := toSnakeReg.ReplaceAllStringFunc(str, func(s string) string {
		return sepChar + LowerFirst(s)
	})

	return strings.TrimLeft(newStr, sepChar)
}

// CamelCase convert.
// Support:
// 	"range_price" -> "rangePrice"
// 	"range price" -> "rangePrice"
// 	"range-price" -> "rangePrice"
func CamelCase(str string, sep ...string) string {
	sepChar := "_"
	if len(sep) > 0 {
		sepChar = sep[0]
	}

	// not contains sep char
	if !strings.Contains(str, sepChar) {
		return str
	}

	// get regexp instance
	rgx, ok := toCamelRegs[sepChar]
	if !ok {
		rgx = regexp.MustCompile(regexp.QuoteMeta(sepChar) + "+[a-zA-Z]")
	}

	return rgx.ReplaceAllStringFunc(str, func(s string) string {
		s = strings.TrimLeft(s, sepChar)
		return UpperFirst(s)
	})
}

/*************************************************************
 * string to slice, time
 *************************************************************/

// StrToArray alias of the StrToSlice()
func StrToArray(str string, sep ...string) []string {
	return StrToSlice(str, sep...)
}

// StrToSlice split string to array.
func StrToSlice(str string, sep ...string) []string {
	if len(sep) > 0 {
		return stringSplit(str, sep[0])
	}

	return stringSplit(str, ",")
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
		case 16:
			layout = "2006-01-02 15:04"
			if strings.ContainsRune(s, 'T') {
				layout = "2006-01-02T15:04"
			}
		case 19:
			layout = "2006-01-02 15:04:05"
			if strings.ContainsRune(s, 'T') {
				layout = "2006-01-02T15:04:05"
			}
		}
	}

	if layout == "" {
		err = errInvalidParam
		return
	}

	t, err = time.Parse(layout, s)
	// t, err = time.ParseInLocation(layout, s, time.Local)
	return
}
