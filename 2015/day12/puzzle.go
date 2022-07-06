package day12

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/fmarangi/aoc-go/utils"
)

func Sum(input string) (sum int) {
	var num strings.Builder
	for _, c := range input {
		switch {
		case c == '-', unicode.IsDigit(c):
			num.WriteRune(c)
		case num.Len() != 0:
			sum += utils.ToInt(num.String())
			num.Reset()
		}
	}
	return
}

func Part1(input string) int {
	return Sum(input)
}

func Part2(input string) (sum int) {
	obj := regexp.MustCompile(`\{[^{}]*?\}`)
	for obj.MatchString(input) {
		input = obj.ReplaceAllStringFunc(input, func(s string) string {
			if strings.Contains(s, `:"red"`) {
				return "0"
			}
			return strconv.Itoa(Sum(s))
		})
	}

	return Sum(input)
}
