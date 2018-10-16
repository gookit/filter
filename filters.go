package filter

import (
	"net/url"
	"strings"
)

var dontLimitType = map[string]int{
	"int":   1,
	"uint":  1,
	"int64": 1,
	//
	"stringsToInts": 1,
}

var filterAliases = map[string]string{
	"toInt":   "int",
	"toUint":  "uint",
	"toInt64": "int64",
	"toBool":  "bool",
	"camel":   "camelCase",
	"snake":   "snakeCase",
	//
	"lcFirst":    "lowerFirst",
	"ucFirst":    "upperFirst",
	"ucWord":     "upperWord",
	"trimSpace":  "trim",
	"uppercase":  "upper",
	"lowercase":  "lower",
	"escapeJs":   "escapeJS",
	"escapeHtml": "escapeHTML",
	//
	"str2arr":   "strToArray",
	"str2array": "strToArray",
	"strToArr":  "strToArray",
	"str2time":  "strToTime",
	// strings2ints
	"strings2ints": "stringsToInts",
}

// Name get real filter name.
func Name(name string) string {
	if rName, ok := filterAliases[name]; ok {
		return rName
	}

	return name
}

/*************************************************************
 * built in filters
 *************************************************************/

// Trim string
func Trim(s string, cutSet ...string) string {
	if len(cutSet) > 0 && cutSet[0] != "" {
		return strings.Trim(s, cutSet[0])
	}

	return strings.TrimSpace(s)
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

// TrimStrings trim string slice item.
func TrimStrings(ss []string, cutSet ...string) (ns []string) {
	hasCutSet := len(cutSet) > 0 && cutSet[0] != ""

	for _, str := range ss {
		if hasCutSet {
			ns = append(ns, strings.Trim(str, cutSet[0]))
		} else {
			ns = append(ns, strings.TrimSpace(str))
		}
	}

	return
}

// URLEncode encode url string.
func URLEncode(s string) string {
	if pos := strings.IndexRune(s, '?'); pos > -1 { // escape query data
		return s[0:pos+1] + url.QueryEscape(s[pos+1:])
	}

	return s
}

// URLDecode decode url string.
func URLDecode(s string) string {
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
