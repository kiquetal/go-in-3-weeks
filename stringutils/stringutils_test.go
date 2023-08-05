package stringutils

import (
	"testing"
)

func TestUpper(t *testing.T) {
	upper := Upper("test")
	if upper != "TEST" {
		t.Error("Expected TEST, got ", upper)
	}
}

func TestLower(t *testing.T) {

	tests := map[string]struct {
		input    string
		expected string
	}{
		"basic": {input: "TEST", expected: "test"},
		"empty": {input: "", expected: ""},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			lower := Lower(tc.input)
			if lower != tc.expected {
				t.Errorf("Expected %s, got %s", tc.expected, lower)
			}
		})
	}
}
