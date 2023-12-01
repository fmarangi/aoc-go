package day01

import "strings"

var digits = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func isDigit(c byte) bool {
	return c > '0' && c <= '9'
}

func Part1(input string) (sum int) {
	for _, line := range strings.Fields(input) {
		first, last := -1, -1
		for _, c := range []byte(line) {
			if isDigit(c) {
				last = int(c - '0')
				if first == -1 {
					first = int(c - '0')
				}
			}
		}
		sum += first*10 + last
	}
	return
}

func Part2(input string) (sum int) {
	for _, line := range strings.Fields(input) {
		first, last, n := -1, -1, len(line)
		for i, c := range []byte(line) {
			if isDigit(c) {
				last = int(c - '0')
				if first == -1 {
					first = int(c - '0')
				}
				continue
			}

			for d, v := range digits {
				if i+len(d) <= n && line[i:i+len(d)] == d {
					last = v
					if first == -1 {
						first = v
					}
					break
				}
			}
		}
		sum += first*10 + last
	}
	return
}
