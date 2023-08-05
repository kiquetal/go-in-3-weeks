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
