// Package filter provide data filter, sanitize, convert
package filter

import (
	"fmt"
	"strings"
)

// Filtration definition. Sanitization Sanitizing Sanitize
type Filtration struct {
	err error
	// raw data
	data map[string]interface{}
	// mark has apply filters
	filtered bool
	// filter rules
	filterRules []*FilterRule
	// filtered data
	filteredData map[string]interface{}
}

// New a Filtration
func New(data map[string]interface{}) *Filtration {
	return &Filtration{
		data: data,
		// init
		filteredData: make(map[string]interface{}),
	}
}

// Raw get raw value by key
func (f *Filtration) Raw(key string) (interface{}, bool) {
	return GetByPath(key, f.data)
}

// Get value by key
func (f *Filtration) Get(key string) (interface{}, bool) {
	val, ok := GetByPath(key, f.filteredData)
	if !ok {
		val, ok = GetByPath(key, f.data)
	}

	return val, ok
}

// MustGet value by key
func (f *Filtration) MustGet(key string) interface{} {
	val, _ := f.Get(key)
	return val
}

// Trimmed get trimmed string value
func (f *Filtration) Trimmed(key string) string {
	val, ok := f.Raw(key)
	if !ok {
		return ""
	}

	// string.
	if str, ok := val.(string); ok {
		return strings.TrimSpace(str)
	}

	return fmt.Sprint(val)
}

// Int value get
func (f *Filtration) Int(key string) int {
	val, ok := f.Raw(key)
	if !ok {
		return 0
	}

	return MustInt(val)
}

// AddRule add filter(s) rule.
// Usage:
// 	f.AddRule("name", "trim")
// 	f.AddRule("age", "int")
// 	f.AddRule("age", "trim|int")
func (f *Filtration) AddRule(fields string, rule string) *Filtration {
	rule = strings.TrimSpace(rule)
	rules := stringSplit(strings.Trim(rule, "|:"), "|")

	fieldList := stringSplit(fields, ",")
	if len(fieldList) > 0 {
		r := newFilterRule(fieldList)
		r.AddFilters(rules...)
		f.filterRules = append(f.filterRules, r)
	}

	return f
}

// Filtering todo
func (f *Filtration) Filtering() error {
	if f.filtered || f.err != nil {
		return f.err
	}

	// apply rule to validate data.
	for _, rule := range f.filterRules {
		if err := rule.Apply(f); err != nil { // has error
			f.err = err
			break
		}
	}

	f.filtered = true
	return f.err
}

// IsOK of the apply filters
func (f *Filtration) IsOK() bool {
	return f.err == nil
}

/*************************************************************
 * filtering rule
 *************************************************************/

// FilterRule definition
type FilterRule struct {
	// fields to filter
	fields []string
	// filter list, can with args. eg. "int" "str2arr:,"
	filters map[string]string
}

func newFilterRule(fields []string) *FilterRule {
	return &FilterRule{
		fields:  fields,
		filters: make(map[string]string),
	}
}

// AddFilters add filter(s).
// Usage:
// 	r.AddFilters("int", "str2arr:,")
func (r *FilterRule) AddFilters(filters ...string) *FilterRule {
	for _, filterName := range filters {
		pos := strings.IndexRune(filterName, ':')

		// has filter args
		if pos > 0 {
			name := filterName[:pos]
			r.filters[name] = filterName[pos+1:]
		} else {
			r.filters[filterName] = ""
		}
	}

	return r
}

// Apply rule for the rule fields
func (r *FilterRule) Apply(f *Filtration) (err error) {
	// validate field
	for _, field := range r.Fields() {
		// get field value.
		val, has := f.Get(field)
		if !has { // no field
			continue
		}

		// call filters
		for name, argStr := range r.filters {
			args := parseArgString(argStr)

			val, err = callFilter(name, val, args)
			if err != nil {
				return err
			}
		}

		// save filtered value.
		f.filteredData[field] = val
	}

	return
}

// Fields name get
func (r *FilterRule) Fields() []string {
	return r.fields
}

func callFilter(name string, val interface{}, args []string) (interface{}, error) {
	var err error
	realName := Name(name)
	str, isStr := val.(string)

	switch realName {
	case "int":
		val, err = ToInt(val)
	case "uint":
		val, err = ToUint(val)
	case "int64":
		val, err = ToInt64(val)
	case "bool":
		if !isStr {
			return nil, errInvalidParam
		}
		val, err = ToBool(str)
	case "float":
		if !isStr {
			return nil, errInvalidParam
		}
		val, err = ToFloat(str)
	case "trim":
		if !isStr {
			return nil, errInvalidParam
		}
		val = Trim(str, args...)
	case "ltrim":
		if !isStr {
			return nil, errInvalidParam
		}
		val = TrimLeft(str, args...)
	case "rtrim":
		if !isStr {
			return nil, errInvalidParam
		}
		val = TrimRight(str, args...)
	case "lower":
		if !isStr {
			return nil, errInvalidParam
		}
		val = Lowercase(str)
	case "upper":
		if !isStr {
			return nil, errInvalidParam
		}
		val = Uppercase(str)
	case "lowerFirst":
		if !isStr {
			return nil, errInvalidParam
		}
		val = LowerFirst(str)
	case "upperFirst":
		if !isStr {
			return nil, errInvalidParam
		}
		val = UpperFirst(str)
	case "upperWord":
		if !isStr {
			return nil, errInvalidParam
		}
		val = UpperWord(str)
	case "snakeCase":
		if !isStr {
			return nil, errInvalidParam
		}
		val = SnakeCase(str, args...)
	case "camelCase":
		if !isStr {
			return nil, errInvalidParam
		}
		val = CamelCase(str, args...)
	case "strToArray":
		if !isStr {
			return nil, errInvalidParam
		}
		val = StrToArray(str, args...)
	case "strToTime":
		if !isStr {
			return nil, errInvalidParam
		}
		val, err = StrToTime(str)
	case "stringsToInts":
		if ss, ok := val.([]string); ok {
			val, err = StringsToInts(ss)
		} else {
			return nil, errInvalidParam
		}
	default:
		err = errInvalidParam
	}

	return val, err
}
