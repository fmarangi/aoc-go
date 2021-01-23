package day05

import (
	"regexp"
	"strings"
)

var threeVowels = regexp.MustCompile("[aeiou]")
var hasForbidden = regexp.MustCompile("(ab|cd|pq|xy)")

func Part1(input string) (nice int) {
	for _, s := range strings.Split(input, "\n") {
		if IsNice(s) {
			nice += 1
		}
	}
	return nice
}

func Part2(input string) (nice int) {
	for _, s := range strings.Split(input, "\n") {
		if hasPair(s) && hasRepeating(s) {
			nice += 1
		}
	}
	return nice
}

func IsNice(s string) bool {
	return len(threeVowels.FindAllString(s, -1)) >= 3 && twoConsecutive(s) && len(hasForbidden.FindString(s)) == 0
}

func twoConsecutive(s string) bool {
	for i, n := 0, len(s)-1; i < n; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

func hasPair(s string) bool {
	runes := []rune(s)
	pairs := map[string]bool{}
	var curr string
	for i, n := 0, len(s)-1; i < n; i++ {
		pair := string(runes[i : i+2])
		if pairs[pair] && pair != curr {
			return true
		}

		pairs[pair] = true
		if curr == pair {
			curr = ""
		} else {
			curr = pair
		}
	}
	return false
}

func hasRepeating(s string) bool {
	for i, n := 0, len(s)-2; i < n; i++ {
		if s[i] == s[i+2] {
			return true
		}
	}
	return false
}
