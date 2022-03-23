package utils

import (
	"io/ioutil"
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
	content, _ := ioutil.ReadFile(path)
	return strings.TrimRightFunc(string(content), unicode.IsSpace)
}
