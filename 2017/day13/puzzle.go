package day13

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Scanner struct {
	Range, Depth int
}

func (s Scanner) Caught(delay int) bool {
	r := s.Range + delay
	return r == 0 || r%((s.Depth-1)*2) == 0
}

func (s Scanner) Severity() int {
	return s.Range * s.Depth
}

func Escape(scanners []Scanner, delay int) bool {
	for _, s := range scanners {
		if s.Caught(delay) {
			return false
		}
	}
	return true
}

func Part1(input string) (severity int) {
	for _, scanner := range parseInput(input) {
		if scanner.Caught(0) {
			severity += scanner.Severity()
		}
	}
	return
}

func Part2(input string) (delay int) {
	scanners := parseInput(input)
	for !Escape(scanners, delay) {
		// There is a scanner with Depth=2
		// odd delays will always get caught
		delay += 2
	}
	return
}

func parseInput(input string) (scanners []Scanner) {
	for _, row := range strings.Split(input, "\n") {
		values := utils.Ints(strings.Split(row, ": "))
		scanners = append(scanners, Scanner{values[0], values[1]})
	}
	return
}
