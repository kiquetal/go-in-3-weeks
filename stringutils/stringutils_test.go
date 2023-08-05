package stringutils

import (
	"fmt"
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

func TestWithTable(t *testing.T) {

	type testCase struct {
		lang language
		want string
	}

	var tests = map[string]testCase{
		"english": {lang: "en", want: "hello"},
		"spanish": {lang: "sp", want: "hola"},
		"french":  {lang: "fr", want: "bonjour"},
		"german":  {lang: "de", want: "hello"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := greet(tc.lang)
			if got != tc.want {
				t.Errorf("greet(%q) = %q, want %q", tc.lang, got, tc.want)
			}
		})
	}

}

type language string

var phrasebook = map[language]string{
	"en": "hello",   // English
	"sp": "hola",    // Spanish
	"fr": "bonjour", // French

}

func greet(l language) string {
	greeting, ok := phrasebook[l]
	if !ok {
		return fmt.Sprintf("unsupported language: %q", l)
	}

	return greeting
}
