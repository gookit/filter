package filter

import (
	"fmt"
	"strings"
)

// Apply a filter by name. for filter value.
func Apply(name string, val interface{}, args []string) (interface{}, error) {
	var err error
	realName := Name(name)

	// don't Limit value type
	if _, ok := dontLimitType[realName]; ok {
		switch realName {
		case "int":
			val, err = ToInt(val)
		case "uint":
			val, err = ToUint(val)
		case "int64":
			val, err = ToInt64(val)
		case "unique":
			val = Unique(val)
		case "trimStrings":
			if ss, ok := val.([]string); ok {
				val = TrimStrings(ss)
			} else {
				err = errInvalidParam
			}
		case "stringsToInts":
			if ss, ok := val.([]string); ok {
				val, err = StringsToInts(ss)
			} else {
				err = errInvalidParam
			}
		}
		return val, err
	}

	str, isString := val.(string)
	if !isString {
		return nil, fmt.Errorf("filter: '%s' only use for string type value", name)
	}

	// val is must be string.
	switch realName {
	case "bool":
		val, err = ToBool(str)
	case "float":
		val, err = ToFloat(str)
	case "trim":
		val = Trim(str, args...)
	case "trimLeft":
		val = TrimLeft(str, args...)
	case "trimRight":
		val = TrimRight(str, args...)
	case "title":
		val = Title(str)
	case "email":
		val = Email(str)
	case "substr":
		val = Substr(str, MustInt(args[0]), MustInt(args[1]))
	case "lower":
		val = Lowercase(str)
	case "upper":
		val = Uppercase(str)
	case "lowerFirst":
		val = LowerFirst(str)
	case "upperFirst":
		val = UpperFirst(str)
	case "upperWord":
		val = UpperWord(str)
	case "snakeCase":
		val = SnakeCase(str, args...)
	case "camelCase":
		val = CamelCase(str, args...)
	case "URLEncode":
		val = URLEncode(str)
	case "URLDecode":
		val = URLDecode(str)
	case "escapeJS":
		val = EscapeJS(str)
	case "escapeHTML":
		val = EscapeHTML(str)
	case "strToInts":
		val, err = StrToInts(str, args...)
	case "strToSlice":
		val = StrToSlice(str, args...)
	case "strToTime":
		val, err = StrToTime(str)
	}

	return val, err
}

// GetByPath get value from a map[string]interface{}. eg "top" "top.sub"
func GetByPath(key string, mp map[string]interface{}) (val interface{}, ok bool) {
	if val, ok := mp[key]; ok {
		return val, true
	}

	// has sub key? eg. "top.sub"
	if !strings.ContainsRune(key, '.') {
		return nil, false
	}

	keys := strings.Split(key, ".")
	topK := keys[0]

	// find top item data based on top key
	var item interface{}
	if item, ok = mp[topK]; !ok {
		return
	}

	for _, k := range keys[1:] {
		switch tData := item.(type) {
		case map[string]string: // is simple map
			item, ok = tData[k]
			if !ok {
				return
			}
		case map[string]interface{}: // is map(decode from toml/json)
			if item, ok = tData[k]; !ok {
				return
			}
		case map[interface{}]interface{}: // is map(decode from yaml)
			if item, ok = tData[k]; !ok {
				return
			}
		default: // error
			ok = false
			return
		}
	}

	return item, true
}

func parseArgString(argStr string) (ss []string) {
	if argStr == "" { // no arg
		return
	}

	if len(argStr) == 1 { // one char
		return []string{argStr}
	}

	return stringSplit(argStr, ",")
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
