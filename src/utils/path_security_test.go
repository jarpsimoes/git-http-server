package utils

import (
	"testing"
)

func TestIsValidPath(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"valid_path123", true},
		{"invalid_path/with_slash", false},
		{"invalid_path/with/slash", false},
		{"invalid/path/with/slash", false},
		{"invalid.path.with.multiple.dots", false},
		{"another.invalid.path.with.multiple.dots", false},
		{"valid_path_with_underscore", true},
		{"valid-path_with-hyphen", true},
		{"valid.path_with_underscore-hyphen", true},
	}

	pathSecurityCheck := PathSecurityCheck{}

	for _, test := range tests {
		t.Run(test.input, func(t *testing.T) {
			result := pathSecurityCheck.IsValidPath(test.input)
			if result != test.expected {
				t.Errorf("Expected IsValidPath(%s) to be %v, but got %v", test.input, test.expected, result)
			}
		})
	}
}
