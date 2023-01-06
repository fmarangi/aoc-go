package day25

import (
	"strconv"
	"strings"
)

func FromSnafu(num string) (value int) {
	for _, c := range num {
		value *= 5
		switch c {
		case '-':
			value -= 1
		case '=':
			value -= 2
		default:
			value += int(c - '0')
		}
	}
	return
}

func ToSnafu(num int) (result string) {
	for ; num > 0; num /= 5 {
		switch r := num % 5; r {
		default:
			result = strconv.Itoa(r) + result
		case 3:
			result = "=" + result
			num += 5
		case 4:
			result = "-" + result
			num += 5
		}
	}
	return
}

func Part1(input string) string {
	var sum int
	for _, n := range strings.Fields(input) {
		sum += FromSnafu(n)
	}
	return ToSnafu(sum)
}
