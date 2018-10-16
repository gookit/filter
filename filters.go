package filter

import (
	"net/url"
	"strings"
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
	if pos := strings.IndexRune(s, '?'); pos > -1 { // escape query data
		return s[0:pos+1] + url.QueryEscape(s[pos+1:])
	}

	return s
}

// UrlDecode decode url string.
func UrlDecode(s string) string {
	if pos := strings.IndexRune(s, '?'); pos > -1 { // un-escape query data
		qy, err := url.QueryUnescape(s[pos+1:])

		if err == nil {
			return s[0:pos+1] + qy
		}
	}

	return s
}

// Unique value in the given array, slice.
func Unique(val interface{}) interface{} {
	// switch tv := val.(type) {
	// case []string:
	// 	cp := make([]string, len(tv))
	// 	copy(cp, tv)
	// }

	return val // todo
}

// Substr cut string
func Substr(s string, pos, length int) string {
	runes := []rune(s)
	strLen := len(runes)

	if pos >= strLen {
		return ""
	}

	l := pos + length
	if l > strLen {
		l = strLen
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
