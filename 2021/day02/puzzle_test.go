package day02

import (
	"io/ioutil"
	"strings"
	"testing"
)

type Solution struct {
	part   func(string) int
	result int
}

func TestSolvePuzzle(t *testing.T) {
	parts := []Solution{
		{Part1, 1936494},
		{Part2, 1997106066},
	}

	content, _ := ioutil.ReadFile("input.txt")
	input := strings.TrimSpace(string(content))
	for _, test := range parts {
		if res := test.part(input); res != test.result {
			t.Errorf("Expected %d, got: %d", test.result, res)
		}
	}
}
