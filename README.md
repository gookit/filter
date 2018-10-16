# Filter

[![GoDoc](https://godoc.org/github.com/gookit/filter?status.svg)](https://godoc.org/github.com/gookit/filter)
[![Build Status](https://travis-ci.org/gookit/filter.svg?branch=master)](https://travis-ci.org/gookit/filter)
[![Coverage Status](https://coveralls.io/repos/github/gookit/filter/badge.svg?branch=master)](https://coveralls.io/github/gookit/filter?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/filter)](https://goreportcard.com/report/github.com/gookit/filter)

package filter provide filter, sanitize, convert golang data.

## GoDoc

- [godoc for gopkg](https://godoc.org/gopkg.in/gookit/filter.v1)
- [godoc for github](https://godoc.org/github.com/gookit/filter)

## Usage

```go
intVal, err := filter.Int("20") // int(20)
strArr := filter.Str2Array("a,b, c", ",") // []string{"a", "b", "c"}
```

## Filters & Converters

- `ToBool/Bool(s string) (bool, error)`
- `ToFloat/Float(s string) (float64, error)`
- `ToInt/Int(s string) (int, error)`
- `ToUint/Uint(s string) (uint64, error)`
- `ToInt64/Int64(s string) (int64, error)`
- `MustBool(s string) bool`
- `MustFloat(s string) float64`
- `MustInt(s string) int`
- `MustInt64(s string) int64`
- `MustUint(s string) uint64`
- `Trim(s string, cutSet ...string) string`
- `TrimLeft(s string, cutSet ...string) string`
- `TrimRight(s string, cutSet ...string) string`
- `TrimStrings(ss []string, cutSet ...string) (ns []string)`
- `Substr(s string, pos, length int) string`
- `Lowercase(s string) string`
- `Uppercase(s string) string`
- `LowerFirst(s string) string`
- `UpperFirst(s string) string`
- `UpperWord(s string) string`
- `CamelCase(s string, sep ...string) string`
- `SnakeCase(s string, sep ...string) string`
- `Email(s string) string`
- `URLDecode(s string) string`
- `URLEncode(s string) string`
- `EscapeJS(s string) string`
- `EscapeHTML(s string) string`
- `Unique(val interface{}) interface{}`
- `StrToArray(s string, sep ...string) []string`
- `StrToTime(s string, layouts ...string) (t time.Time, err error)`
- `StringsToInts(ss []string) (ints []int, err error)`

## License

**[MIT](LICENSE)**
