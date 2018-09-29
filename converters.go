package filter

import (
	"errors"
	"regexp"
	"strings"
	"time"
)

// some regex for convert string.
var (
	toSnakeReg  = regexp.MustCompile("[A-Z][a-z]")
	toCamelRegs = map[string]*regexp.Regexp{
		" ": regexp.MustCompile(" +[a-zA-Z]"),
		"-": regexp.MustCompile("-+[a-zA-Z]"),
		"_": regexp.MustCompile("_+[a-zA-Z]"),
	}
	errInvalidParam = errors.New("invalid input parameter")
)

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

// StrToArray split string to array.
func StrToArray(str string, sep ...string) []string {
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
