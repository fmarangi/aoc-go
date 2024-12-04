package day03

import (
	"regexp"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func multiplications(input string) (result int) {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	for _, v := range r.FindAllStringSubmatch(input, -1) {
		result += utils.ToInt(v[1]) * utils.ToInt(v[2])
	}
	return
}

func Part1(input string) int {
	return multiplications(input)
}

func Part2(input string) (result int) {
	var builder strings.Builder
	for i, n, valid := 0, len(input), true; i < n; i++ {
		if valid {
			if i+7 < n && input[i:i+7] == "don't()" {
				valid, i = false, i+6
				result += multiplications(builder.String())
				builder.Reset()
				continue
			}
			builder.WriteByte(input[i])
		} else if i+4 < n && input[i:i+4] == "do()" {
			valid, i = true, i+3
		}
	}
	return result + multiplications(builder.String())
}
