package filter

import (
	"fmt"
	"html/template"
	"net/url"
	"strconv"
	"strings"
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

// UrlEncode
func UrlEncode(s string) string {
	u, err := url.Parse(s)
	if err != nil {
		return s
	}

	return u.EscapedPath()
}

// Int convert
func Int(str string) (int, error) {
	return ToInt(str)
}

// ToInt convert
func ToInt(str string) (int, error) {
	return strconv.Atoi(Trim(str))
}

// MustInt convert
func MustInt(str string) int {
	val, _ := strconv.Atoi(Trim(str))
	return val
}

// Uint convert
func Uint(str string) (uint64, error) {
	return ToUint(str)
}

// ToUint convert
func ToUint(str string) (uint64, error) {
	return strconv.ParseUint(Trim(str), 10, 0)
}

// MustUint convert
func MustUint(str string) uint64 {
	val, _ := strconv.ParseUint(Trim(str), 10, 0)
	return val
}

// Int64 convert
func Int64(str string) (int64, error) {
	return ToInt64(str)
}

// ToInt64 convert
func ToInt64(str string) (int64, error) {
	return strconv.ParseInt(Trim(str), 10, 0)
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

// StrToArray split string to array.
func StrToArray(str string, sep ...string) []string {
	if len(sep) > 0 {
		return stringSplit(str, sep[0])
	}

	return stringSplit(str, ",")
}

// StrToTime convert date string to time.Time
func StrToTime(s string) (t time.Time, err error) {
	var layout string
	switch len(s) {
	case 10: // 2006-01-02
		layout = "2006-01-02"
	case 19: // "2006-01-02 12:24:36" "2006-01-02T15:04:05"
		layout = "2006-01-02 15:04:05"
		if strings.ContainsRune(s, 'T') {
			layout = "2006-01-02T15:04:05"
		}
	default:
		return
	}

	t, err = time.Parse(layout, s)
	// t, err = time.ParseInLocation(layout, s, time.Local)
	return
}

// Unique value in the given array, slice.
func Unique(val interface{}) (interface{}) {
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

// Email filter, clear invalid chars.
func Email(s string) string {
	// According to rfc5321, "The local-part of a mailbox MUST BE treated as case sensitive"
	return emailLocalPart(s) + "@" + strings.ToLower(emailDomainPart(s))
}

// a valid email will only have one "@", but let's treat the last "@" as the domain part separator
func emailLocalPart(s string) string {
	i := strings.LastIndex(s, "@")
	if i == -1 {
		return s
	}

	return s[0:i]
}

func emailDomainPart(s string) string {
	i := strings.LastIndex(s, "@")
	if i == -1 {
		return ""
	}

	return s[i+1:]
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

func panicf(format string, args ...interface{}) {
	panic("filter: " + fmt.Sprintf(format, args...))
}
