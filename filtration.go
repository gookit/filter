// Package filter provide data filter, sanitize, convert
package filter

// Filtration definition. Sanitization Sanitizing Sanitize
type Filtration struct {
	// filtered data
	filtered map[string]interface{}
}

// Get value by key
func (f *Filtration) Get(key string) (interface{}, bool) {
	return GetByPath(key, f.filtered)
}

// Set value by key
func (f *Filtration) Set(field string, val interface{}) error {
	panic("implement me")
}
