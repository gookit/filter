package filter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFiltration(t *testing.T) {
	is := assert.New(t)

	fl := New(map[string]interface{}{
		"key0": " abc ",
		"key1": "2",
		"sub":  map[string]string{"k0": "v0"},
	})

	is.Equal("strToTime", Name("str2time"))
	is.Equal("some", Name("some"))

	is.Equal("", fl.Trimmed("not-exist"))
	is.Equal("abc", fl.Trimmed("key0"))
	is.Equal(2, fl.Int("key1"))
	is.Equal(0, fl.Int("not-exist"))

	val, ok := fl.Get("key1")
	is.True(ok)
	is.Equal("2", val)

	val, ok = fl.Get("sub.k0")
	is.True(ok)
	is.Equal("v0", val)

	val, ok = fl.Raw("key1")
	is.True(ok)
	is.Equal("2", val)
	val, ok = fl.Raw("not-exist")
	is.False(ok)
	is.Equal(nil, val)
}

func TestFiltration_Filtering(t *testing.T) {
	is := assert.New(t)

	f := New(map[string]interface{}{
		"name":     "inhere",
		"age":      "50",
		"money":    "50.34",
		"remember": "yes",
		"sub":      map[string]string{"k0": "v0"},
		"sub1":     []string{"1", "2"},
		"tags":     "go;lib",
	})
	f.AddRule("name", "upper")
	f.AddRule("age", "int")
	f.AddRule("money", "float")
	f.AddRule("remember", "bool")
	f.AddRule("sub1", "strings2ints")
	f.AddRule("tags", "str2arr:;")

	is.Nil(f.Filtering())
	is.Equal(50, f.MustGet("age"))
	is.Equal(50.34, f.MustGet("money"))
	is.Equal([]int{1, 2}, f.MustGet("sub1"))
	is.Equal([]string{"go", "lib"}, f.MustGet("tags"))
}
