package filter

import (
	"github.com/gookit/goutil/x/assert"
	"testing"
)

func TestURLEncodePreservesFragment(t *testing.T) {
	for _, tc := range []struct{ input, want string }{
		{"https://example.com/?q=a b#part+one", "https://example.com/?q%3Da+b#part+one"},
		{"https://example.com/?q=a#part?b=c", "https://example.com/?q%3Da#part?b=c"},
		{"https://example.com/#part?q=a b", "https://example.com/#part?q=a b"},
		{"https://example.com/?#part", "https://example.com/?#part"},
		{"https://example.com/?q=a b", "https://example.com/?q%3Da+b"},
		{"https://example.com/", "https://example.com/"},
	} {
		t.Run(tc.input, func(t *testing.T) { assert.Eq(t, tc.want, URLEncode(tc.input)) })
	}
}

func TestURLDecodePreservesFragment(t *testing.T) {
	for _, tc := range []struct{ input, want string }{
		{"https://example.com/?q%3Da+b#part+one", "https://example.com/?q=a b#part+one"},
		{"https://example.com/?q%3Da#part%20one", "https://example.com/?q=a#part%20one"},
		{"https://example.com/#part?q%3Da+b", "https://example.com/#part?q%3Da+b"},
		{"https://example.com/?q%3Da#bad%escape", "https://example.com/?q=a#bad%escape"},
		{"https://example.com/?bad%escape#part+one", "https://example.com/?bad%escape#part+one"},
		{"https://example.com/?q%3Da+b", "https://example.com/?q=a b"},
		{"https://example.com/", "https://example.com/"},
	} {
		t.Run(tc.input, func(t *testing.T) { assert.Eq(t, tc.want, URLDecode(tc.input)) })
	}
}
