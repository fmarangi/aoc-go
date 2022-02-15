package day02

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestSolvePuzzle(t *testing.T) {
	content, _ := ioutil.ReadFile("input.txt")
	input := strings.TrimSpace(string(content))

	if res := Part1(input); res != 8118 {
		t.Errorf("Expected %d, got: %d", 8118, res)
	}

	if res := Part2(input); res != "jbbenqtlaxhivmwyscjukztdp" {
		t.Errorf("Expected %s, got: %s", "jbbenqtlaxhivmwyscjukztdp", res)
	}
}
