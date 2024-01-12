package utils

import (
	"os"
	"strings"
	"testing"
	"unicode"
)

func Assert(t *testing.T, a, b interface{}) {
	if a != b {
		t.Errorf("Expected %v, got: %v", a, b)
	}
}

func ReadInput(path string) string {
	content, _ := os.ReadFile(path)
	return strings.TrimRightFunc(string(content), unicode.IsSpace)
}
