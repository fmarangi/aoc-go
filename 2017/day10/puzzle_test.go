package day10

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestSolvePuzzle(t *testing.T) {
	content, _ := ioutil.ReadFile("input.txt")
	input := strings.TrimSpace(string(content))

	if res := Part1(input); res != 8536 {
		t.Errorf("Expected %d, got: %d", 8536, res)
	}

	exp := "aff593797989d665349efe11bb4fd99b"
	if res := Part2(input); res != exp {
		t.Errorf("Expected %s, got: %s", exp, res)
	}
}
