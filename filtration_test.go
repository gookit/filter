package filter

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gookit/goutil/testutil/assert"
)

func TestFiltration(t *testing.T) {
	is := assert.New(t)

	fl := New(map[string]any{
		"key0": " abc ",
		"key1": "2",
		"sub":  map[string]string{"k0": "v0"},
		"sub1": map[string]any{"k0": "v0"},
		"sub2": map[any]any{"k0": "v0"},
	})

	is.Eq("strToTime", Name("str2time"))
	is.Eq("some", Name("some"))

	is.Eq("", fl.String("key0"))
	is.Eq(0, fl.Int("key1"))

	val, ok := fl.Get("key1")
	is.True(ok)
	is.Eq("2", val)

	_, ok = fl.Get("sub.not-exist")
	is.False(ok)
	val, ok = fl.Get("sub.k0")
	is.True(ok)
	is.Eq("v0", val)

	val, ok = fl.Safe("key1")
	is.False(ok)
	is.Eq(nil, val)

	_, ok = fl.Raw("sub1.not-exist")
	is.False(ok)
	val, ok = fl.Raw("sub1.k0")
	is.True(ok)
	is.Eq("v0", val)

	val, ok = fl.Raw("sub2.k0")
	is.True(ok)
	is.Eq("v0", val)

	val, ok = fl.Raw("key1")
	is.True(ok)
	is.Eq("2", val)
	val, ok = fl.Raw("not-exist")
	is.False(ok)
	is.Eq(nil, val)

	f := New(map[string]any{
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
	is.Eq(int64(34), f.SafeVal("key0"))
	is.Eq([]int{1, 2, 3}, f.SafeVal("ids"))
	is.Eq([]string{"a", "b", "c"}, f.SafeVal("strings"))
	is.Eq("Inhere", f.String("name"))
	is.Eq("my@email.com", f.String("email"))
	is.Eq("&lt;p&gt;some text&lt;/p&gt;", f.SafeVal("htmlCode"))
	// < 1.15 \x3Cscript\x3Evar a = 23;\x3C/script\x3E
	// >= 1.15 \u003Cscript\u003Evar a \u003D 23;\u003C/script\u003E
	// is.Eq(`\x3Cscript\x3Evar a = 23;\x3C/script\x3E`, f.SafeVal("jsCode"))
	is.NotEq("<script>var a = 23;</script>", f.SafeVal("jsCode"))

	// clear all
	f.Clear()
	is.Empty(f.CleanData())
}

func TestFiltration_AddRule(t *testing.T) {
	is := assert.New(t)

	f := New(nil)
	f.LoadData(map[string]any{
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

	f.AddRule("name", func(v any) (any, error) {
		return strings.TrimSpace(v.(string)), nil
	})

	is.NoErr(f.Filtering())
	is.Eq("INHERE", f.String("name"))

	is.Len(f.CleanData(), 1)
	is.Contains(f.CleanData(), "name")

	// clear rules and cleanData
	f.ResetRules()
	is.Empty(f.CleanData())
	is.NotEmpty(f.RawData())

	f.AddRule("name", "trim|lower")
	is.NoErr(f.Filtering())
	is.Eq("inhere", f.String("name"))

	// clear all data
	is.NotEmpty(f.CleanData())
	f.ResetData(true)
	is.Empty(f.RawData())
	is.Empty(f.CleanData())
	f.LoadData(map[string]any{
		"name": " Inhere0 ",
	})
	is.NoErr(f.Filtering())
	is.Eq("inhere0", f.String("name"))
	is.Eq("", f.String("not-exist"))

	f.ResetRules()
	f.AddRule("not-exist", "trim").SetDefaultVal(" def val ")
	is.NoErr(f.Filtering())
	is.Eq("def val", f.String("not-exist"))

	// trimStrings error
	f = New(map[string]any{
		"ints": []int{1, 2, 3},
	})
	f.AddRule("ints", "trimStrings")
	is.Err(f.Filtering())
	is.Eq("invalid input parameter", f.Err().Error())

	// stringsToInts error
	f.ResetRules()
	f.AddRule("ints", "stringsToInts")
	is.Err(f.Filtering())
	is.Eq("invalid input parameter", f.Err().Error())
}

func TestFiltration_Filtering(t *testing.T) {
	is := assert.New(t)

	data := map[string]any{
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
	is.Eq(50, f.Int("age"))
	is.Eq(0, f.Int("not-exist"))
	is.Eq(50, f.MustGet("age"))
	is.Eq(int64(50), f.Int64("age"))
	is.Eq(int64(0), f.Int64("not-exist"))
	is.Eq(50.34, f.MustGet("money"))
	is.Eq([]int{1, 2}, f.MustGet("sub1"))
	is.Len(f.MustGet("ids"), 2)
	is.Eq([]string{"go", "lib"}, f.MustGet("tags"))
	is.Eq("INHERE", f.CleanData()["name"])
	is.Eq("word", f.String("str1"))
	is.Eq("hello", f.String("str2"))

	f = New(data)
	f.AddRule("name", "int")
	is.Err(f.Sanitize())

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
	is.Eq("Inhere", f.String("name"))
	is.Eq("WORD", f.String("str1"))
	is.Eq("Hello World", f.String("msg"))
	is.Eq("hello_world", f.String("msg1"))
	is.Eq("helloWorld", f.String("msg2"))
	is.Eq("hELLO", f.String("str2"))

	sTime, ok := f.Safe("sDate")
	is.True(ok)
	is.Eq("2018-10-16 12:34:00 +0000 UTC", fmt.Sprintf("%v", sTime))

	data["url"] = "a.com?p=1"
	f = New(data)
	f.AddRule("url", "urlEncode")
	f.AddRule("msg1", "substr:0,2")
	is.Nil(f.Sanitize())
	is.Eq("he", f.String("msg1"))
	is.Eq("a.com?p%3D1", f.String("url"))

	// bind
	f = New(map[string]any{
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
	is.Eq(89, user.Age)
	is.Eq("Inhere", user.Name)
}
