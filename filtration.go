// Package filter provide data filter, sanitize, convert
package filter

// Filtration definition. Sanitization Sanitizing Sanitize
type Filtration struct {
	data map[string]interface{}
	// filtered data
	filtered map[string]interface{}
}

// Get value by key
func (f *Filtration) Get(key string) (interface{}, bool) {
	return GetByPath(key, f.filtered)
}

// Filtering
func (f *Filtration) Filtering() error {
	return nil
}
