package day06

import (
	"io/ioutil"
	"strings"
	"testing"
)

type Solution struct {
	part   func(string) string
	result string
}

func TestSolvePuzzle(t *testing.T) {
	parts := []Solution{
		{Part1, "ygjzvzib"},
		{Part2, "pdesmnoz"},
	}

	content, _ := ioutil.ReadFile("input.txt")
	input := strings.TrimSpace(string(content))
	for _, test := range parts {
		if res := test.part(input); res != test.result {
			t.Errorf("Expected %s, got: %s", test.result, res)
		}
	}
}
