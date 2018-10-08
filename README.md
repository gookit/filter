# Filter

[![GoDoc](https://godoc.org/github.com/gookit/filter?status.svg)](https://godoc.org/github.com/gookit/filter)
[![Go Report Card](https://goreportcard.com/badge/github.com/gookit/filter)](https://goreportcard.com/report/github.com/gookit/filter)

package filter provide filter, sanitize, convert golang data.

## GoDoc

- [godoc for gopkg](https://godoc.org/gopkg.in/gookit/filter.v1)
- [godoc for github](https://godoc.org/github.com/gookit/filter)

## Usage

```go
intVal, err := filter.Int("20")
strArr := filter.Str2Array("a,b, c", ",")
```

## Filters & Converters

- `ToBool/Bool(s string) (bool, error)`
- `ToFloat/Float(str string) (float64, error)`
- `ToInt/Int(str string) (int, error)`
- `ToUint/Uint(str string) (uint64, error)`
- `ToInt64/Int64(str string) (int64, error)`
- `MustBool(s string) bool`
- `MustFloat(str string) float64`
- `MustInt(str string) int`
- `MustInt64(str string) int64`
- `MustUint(str string) uint64`
- `CamelCase(str string, sep ...string) string`
- `SnakeCase(str string, sep ...string) string`
- `Email(s string) string`
- `UrlEncode(s string) string`
- `Substr(s string, pos, length int) string`
- `Trim(str string, cutSet ...string) string`
- `TrimLeft(s string, cutSet ...string) string`
- `TrimRight(s string, cutSet ...string) string`
- `LowerFirst(s string) string`
- `UpperFirst(s string) string`
- `UpperWord(s string) string`
- `Unique(val interface{}) interface{}`
- `StrToArray(str string, sep ...string) []string`
- `StrToTime(s string, layouts ...string) (t time.Time, err error)`

## License

**MIT**
