package filter

import (
	"net/url"
	"strconv"
	"strings"
	"text/template"
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

/*************************************************************
 * built in filters
 *************************************************************/

// Trim string
func Trim(str string, cutSet ...string) string {
	if len(cutSet) > 0 {
		return strings.Trim(str, cutSet[0])
	}

	return strings.TrimSpace(str)
}

// TrimLeft char in the string.
func TrimLeft(s string, cutSet ...string) string {
	if len(cutSet) > 0 {
		return strings.TrimLeft(s, cutSet[0])
	}

	return strings.TrimLeft(s, " ")
}

// TrimRight char in the string.
func TrimRight(s string, cutSet ...string) string {
	if len(cutSet) > 0 {
		return strings.TrimRight(s, cutSet[0])
	}

	return strings.TrimRight(s, " ")
}

// UrlEncode encode url string.
func UrlEncode(s string) string {
	if strings.ContainsRune(s, '?') { // escape query data
		ss := strings.SplitN(s, "?", 2)
		return ss[0] + url.QueryEscape(ss[1])
	}

	return s
}

// Unique value in the given array, slice.
func Unique(val interface{}) interface{} {
	return val // todo
}

// Substr cut string
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	l := pos + length

	if l > len(runes) {
		l = len(runes)
	}

	return string(runes[pos:l])
}

// Email filter, clear invalid chars.
func Email(s string) string {
	s = strings.TrimSpace(s)
	i := strings.LastIndex(s, "@")
	if i == -1 {
		return s
	}

	// According to rfc5321, "The local-part of a mailbox MUST BE treated as case sensitive"
	return s[0:i] + "@" + strings.ToLower(s[i+1:])
}

func stringSplit(str, sep string) (ss []string) {
	str = strings.TrimSpace(str)
	if str == "" {
		return
	}

	for _, val := range strings.Split(str, sep) {
		if val = strings.TrimSpace(val); val != "" {
			ss = append(ss, val)
		}
	}

	return
}

// String definition.
type String string

// CanInt convert.
func (s String) CanInt() bool {
	if s == "" {
		return false
	}

	_, err := strconv.Atoi(s.Trimmed())
	return err == nil
}

// Int convert.
func (s String) Int() (val int) {
	if s == "" {
		return 0
	}

	val, _ = strconv.Atoi(s.Trimmed())
	return
}

// Uint convert.
func (s String) Uint() uint {
	if s == "" {
		return 0
	}

	val, _ := strconv.Atoi(s.Trimmed())
	return uint(val)
}

// Int64 convert.
func (s String) Int64() int64 {
	if s == "" {
		return 0
	}

	val, _ := strconv.ParseInt(s.Trimmed(), 10, 64)
	return val
}

// Bool convert.
func (s String) Bool() bool {
	if s == "" {
		return false
	}

	val, _ := strconv.ParseBool(s.Trimmed())
	return val
}

// Float convert. to float 64
func (s String) Float() float64 {
	if s == "" {
		return 0
	}

	val, _ := strconv.ParseFloat(s.Trimmed(), 0)
	return val
}

// Trimmed string
func (s String) Trimmed() string {
	return strings.TrimSpace(string(s))
}

// Split string to slice
func (s String) Split(sep string) (ss []string) {
	if s == "" {
		return
	}

	return stringSplit(s.String(), sep)
}

// String get
func (s String) String() string {
	return string(s)
}
