// Package filter provide data filter, sanitize, convert
package filter

import (
	"fmt"
	"strings"
)

// Filtration definition. Sanitization Sanitizing Sanitize
type Filtration struct {
	data map[string]interface{}
	// filtered data
	filtered map[string]interface{}
}

// New a Filtration
func New(data map[string]interface{}) *Filtration {
	return &Filtration{data: data}
}

// Raw get raw value by key
func (f *Filtration) Raw(key string) (interface{}, bool) {
	return GetByPath(key, f.data)
}

// Get value by key
func (f *Filtration) Get(key string) (interface{}, bool) {
	return GetByPath(key, f.filtered)
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

// AddRule
func (f *Filtration) AddRule(field string, rule string) *Filtration {

	return f
}

// Filtering
func (f *Filtration) Filtering() error {

	return nil
}
