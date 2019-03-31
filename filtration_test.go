package filter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFiltration(t *testing.T) {
	is := assert.New(t)

	fl := New(map[string]interface{}{
		"key0": " abc ",
		"key1": "2",
		"sub":  map[string]string{"k0": "v0"},
		"sub1": map[string]interface{}{"k0": "v0"},
		"sub2": map[interface{}]interface{}{"k0": "v0"},
	})

	is.Equal("strToTime", Name("str2time"))
	is.Equal("some", Name("some"))

	is.Equal("", fl.String("key0"))
	is.Equal(0, fl.Int("key1"))

	val, ok := fl.Get("key1")
	is.True(ok)
	is.Equal("2", val)

	_, ok = fl.Get("sub.not-exist")
	is.False(ok)
	val, ok = fl.Get("sub.k0")
	is.True(ok)
	is.Equal("v0", val)

	val, ok = fl.Safe("key1")
	is.False(ok)
	is.Equal(nil, val)

	_, ok = fl.Raw("sub1.not-exist")
	is.False(ok)
	val, ok = fl.Raw("sub1.k0")
	is.True(ok)
	is.Equal("v0", val)

	val, ok = fl.Raw("sub2.k0")
	is.True(ok)
	is.Equal("v0", val)

	val, ok = fl.Raw("key1")
	is.True(ok)
	is.Equal("2", val)
	val, ok = fl.Raw("not-exist")
	is.False(ok)
	is.Equal(nil, val)

	f := New(map[string]interface{}{
		"key0":     "34",
		"name":     " inhere ",
		"email":    " my@email.com ",
		"ids":      " 1,2, 3",
		"jsCode":   "<script>var a = 23;</script>",
		"htmlCode": "<p>some text</p>",
		"strings":  []string{" a", " b ", "c "},
	})
	f.AddRules(map[string]string{
		"ids":      "strToInts",
		"key0":     "int64",
		"email":    "email",
		"name":     "trim|ucFirst",
		"jsCode":   "escapeJS",
		"htmlCode": "escapeHTML",
		"strings":  "trimStrings",
	})

	is.Nil(f.Sanitize())
	is.NotEmpty(f.CleanData())
	is.Equal(int64(34), f.SafeVal("key0"))
	is.Equal([]int{1, 2, 3}, f.SafeVal("ids"))
	is.Equal([]string{"a", "b", "c"}, f.SafeVal("strings"))
	is.Equal("Inhere", f.String("name"))
	is.Equal("my@email.com", f.String("email"))
	is.Equal(`\x3Cscript\x3Evar a = 23;\x3C/script\x3E`, f.SafeVal("jsCode"))
	is.Equal("&lt;p&gt;some text&lt;/p&gt;", f.SafeVal("htmlCode"))

	// clear all
	f.Clear()
	is.Empty(f.CleanData())
}

func TestFiltration_AddRule(t *testing.T) {
	is := assert.New(t)

	f := New(nil)
	f.LoadData(map[string]interface{}{
		"name": " INHERE ",
		"age":  "50 ",
	})

	is.Panics(func() {
		f.AddRule("", nil)
	})
	is.Panics(func() {
		f.AddRule("name", "")
	})
	is.Panics(func() {
		f.AddRule("name", []int{1})
	})

	f.AddRule("name", func(v interface{}) (interface{}, error) {
		return strings.TrimSpace(v.(string)), nil
	})

	is.NoError(f.Filtering())
	is.Equal("INHERE", f.String("name"))

	is.Len(f.CleanData(), 1)
	is.Contains(f.CleanData(), "name")

	// clear rules and cleanData
	f.ResetRules()
	is.Empty(f.CleanData())
	is.NotEmpty(f.RawData())

	f.AddRule("name", "trim|lower")
	is.NoError(f.Filtering())
	is.Equal("inhere", f.String("name"))

	// clear all data
	is.NotEmpty(f.CleanData())
	f.ResetData(true)
	is.Empty(f.RawData())
	is.Empty(f.CleanData())
	f.LoadData(map[string]interface{}{
		"name": " Inhere0 ",
	})
	is.NoError(f.Filtering())
	is.Equal("inhere0", f.String("name"))
	is.Equal("", f.String("not-exist"))

	f.ResetRules()
	f.AddRule("not-exist", "trim").SetDefaultVal(" def val ")
	is.NoError(f.Filtering())
	is.Equal("def val", f.String("not-exist"))

	// trimStrings error
	f = New(map[string]interface{}{
		"ints": []int{1, 2, 3},
	})
	f.AddRule("ints", "trimStrings")
	is.Error(f.Filtering())
	is.Equal("invalid input parameter", f.Err().Error())

	// stringsToInts error
	f.ResetRules()
	f.AddRule("ints", "stringsToInts")
	is.Error(f.Filtering())
	is.Equal("invalid input parameter", f.Err().Error())
}

func TestFiltration_Filtering(t *testing.T) {
	is := assert.New(t)

	data := map[string]interface{}{
		"name":     "inhere",
		"age":      "50",
		"money":    "50.34",
		"remember": "yes",
		//
		"sub":  map[string]string{"k0": "v0"},
		"sub1": []string{"1", "2"},
		"tags": "go;lib",
		"str1": " word ",
		"str2": "HELLO",
		"ids":  []int{1, 2, 2, 1},
	}
	f := New(data)
	f.AddRule("name", "upper")
	f.AddRule("age", "int")
	f.AddRule("money", "float")
	f.AddRule("remember", "bool")
	f.AddRule("sub1", "strings2ints")
	f.AddRule("tags", "str2arr:;")
	f.AddRule("ids", "unique")
	f.AddRule("str1", "ltrim|rtrim")
	f.AddRule("not-exist", "unique")
	f.AddRule("str2", "lower")

	is.Nil(f.Filtering())
	is.Nil(f.Filtering())
	is.True(f.IsOK())

	// get value
	is.True(f.Bool("remember"))
	is.False(f.Bool("not-exist"))
	is.Equal(50, f.Int("age"))
	is.Equal(0, f.Int("not-exist"))
	is.Equal(50, f.MustGet("age"))
	is.Equal(int64(50), f.Int64("age"))
	is.Equal(int64(0), f.Int64("not-exist"))
	is.Equal(50.34, f.MustGet("money"))
	is.Equal([]int{1, 2}, f.MustGet("sub1"))
	is.Len(f.MustGet("ids"), 2)
	is.Equal([]string{"go", "lib"}, f.MustGet("tags"))
	is.Equal("INHERE", f.CleanData()["name"])
	is.Equal("word", f.String("str1"))
	is.Equal("hello", f.String("str2"))

	f = New(data)
	f.AddRule("name", "int")
	is.Error(f.Sanitize())

	data["name"] = " inhere "
	data["sDate"] = "2018-10-16 12:34"
	data["msg"] = " hello world "
	data["msg1"] = "helloWorld"
	data["msg2"] = "hello_world"
	f = New(data)
	f.AddRules(map[string]string{
		"age":   "uint",
		"money": "float",
		"name":  "trim|ucFirst",
		"str1":  "trim|upper",
		"sDate": "str2time",
		"msg":   "trim|ucWord",
		"msg1":  "snake",
		"msg2":  "camel",
		"str2":  "lowerFirst",
	})
	is.Nil(f.Sanitize())
	is.Equal("Inhere", f.String("name"))
	is.Equal("WORD", f.String("str1"))
	is.Equal("Hello World", f.String("msg"))
	is.Equal("hello_world", f.String("msg1"))
	is.Equal("helloWorld", f.String("msg2"))
	is.Equal("hELLO", f.String("str2"))

	sTime, ok := f.Safe("sDate")
	is.True(ok)
	is.Equal("2018-10-16 12:34:00 +0000 UTC", fmt.Sprintf("%v", sTime))

	data["url"] = "a.com?p=1"
	f = New(data)
	f.AddRule("url", "urlEncode")
	f.AddRule("msg1", "substr:0,2")
	is.Nil(f.Sanitize())
	is.Equal("he", f.String("msg1"))
	is.Equal("a.com?p%3D1", f.String("url"))

	// bind
	f = New(map[string]interface{}{
		"name": " inhere ",
		"age":  " 89 ",
	})
	f.AddRules(map[string]string{
		"age":  "trim|int",
		"name": "trim|ucFirst",
	})
	is.Nil(f.Filtering())
	user := &struct {
		Age  int
		Name string
	}{}
	err := f.BindStruct(user)
	is.Nil(err)
	is.Equal(89, user.Age)
	is.Equal("Inhere", user.Name)
}
