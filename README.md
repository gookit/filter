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

## License

**MIT**
