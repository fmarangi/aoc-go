package day18

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

var pattern = regexp.MustCompile(`\(([^())]+)\)`)
var addition = regexp.MustCompile(`\d+ \+ \d+`)

func Calculate(expression string) (result int) {
	for strings.Contains(expression, "(") {
		expression = pattern.ReplaceAllStringFunc(expression, func(s string) string {
			return strconv.Itoa(Calculate(s[1 : len(s)-1]))
		})
	}

	items := strings.Fields(expression)
	result = utils.ToInt(items[0])
	for i, n := 1, len(items); i < n; i += 2 {
		switch items[i] {
		case "+":
			result += utils.ToInt(items[i+1])
		case "*":
			result *= utils.ToInt(items[i+1])
		}
	}
	return
}

func CalculateWithPrecedence(expression string) (result int) {
	for strings.Contains(expression, "(") {
		expression = pattern.ReplaceAllStringFunc(expression, func(s string) string {
			return strconv.Itoa(CalculateWithPrecedence(s[1 : len(s)-1]))
		})
	}

	for strings.Contains(expression, "+") {
		expression = addition.ReplaceAllStringFunc(expression, func(s string) string {
			return strconv.Itoa(Calculate(s))
		})
	}

	items := strings.Fields(expression)
	result = utils.ToInt(items[0])
	for i, n := 1, len(items); i < n; i += 2 {
		result *= utils.ToInt(items[i+1])
	}
	return
}

func Part1(input string) (sum int) {
	for _, e := range strings.Split(input, "\n") {
		sum += Calculate(e)
	}
	return
}

func Part2(input string) (sum int) {
	for _, e := range strings.Split(input, "\n") {
		sum += CalculateWithPrecedence(e)
	}
	return
}
