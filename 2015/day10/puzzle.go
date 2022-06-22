package day10

import (
	"strconv"
	"strings"
)

func Sequence(init string) string {
	var (
		out   strings.Builder
		curr  rune
		count int
	)

	for i, c := range init {
		switch {
		case i == 0:
			curr, count = c, 1
		case c == curr:
			count++
		default:
			out.WriteString(strconv.Itoa(count))
			out.WriteRune(curr)
			curr, count = c, 1
		}
	}

	out.WriteString(strconv.Itoa(count))
	out.WriteRune(curr)
	return out.String()
}

func Part1(input string) int {
	for i := 0; i < 40; i++ {
		input = Sequence(input)
	}
	return len(input)
}

func Part2(input string) int {
	for i := 0; i < 50; i++ {
		input = Sequence(input)
	}
	return len(input)
}
