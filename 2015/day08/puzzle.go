package day08

import (
	"encoding/hex"
	"regexp"
	"strings"
)

var (
	decode = regexp.MustCompile(`(\\x[0-9a-f]{2}|\\\\|\\")`)
	encode = regexp.MustCompile(`(\\|\")`)
)

func Decode(str string) string {
	str = str[1 : len(str)-1]
	return decode.ReplaceAllStringFunc(str, func(s string) string {
		switch s {
		case `\\`, `\"`:
			return s[1:]
		default:
			str, _ := hex.DecodeString(s[2:])
			return string(str)
		}
	})
}

func Encode(str string) string {
	return `"` + encode.ReplaceAllString(str, `\$1`) + `"`
}

func Part1(input string) (characters int) {
	for _, str := range parseInput(input) {
		characters += len(str) - len(Decode(str))
	}
	return
}

func Part2(input string) (characters int) {
	for _, str := range parseInput(input) {
		characters += len(Encode(str)) - len(str)
	}
	return
}

func parseInput(input string) []string {
	return strings.Fields(input)
}
