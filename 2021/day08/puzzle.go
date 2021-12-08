package day08

import (
	"math"
	"strings"
)

type Entry struct {
	mappings, digits []string
}

func decode(mappings []string) (nums map[uint8]int) {
	m := make([]uint8, 10)
	for _, d := range mappings {
		switch len(d) {
		case 2:
			m[1] = toBin(d)
		case 4:
			m[4] = toBin(d)
		case 3:
			m[7] = toBin(d)
		case 7:
			m[8] = toBin(d)
		}
	}

	one, four := m[1], m[4]
	for _, d := range mappings {
		if len(d) != 6 {
			continue
		}

		value := toBin(d)
		switch {
		case value&one != one:
			m[6] = value
		case value&four != four:
			m[0] = value
		default:
			m[9] = value
		}
	}

	six := m[6]
	for _, d := range mappings {
		if len(d) != 5 {
			continue
		}

		value := toBin(d)
		switch {
		case value&one == one:
			m[3] = value
		case six&value == value:
			m[5] = value
		default:
			m[2] = value
		}
	}

	nums = make(map[uint8]int)
	for i, n := range m {
		nums[n] = i
	}

	return
}

func parseInput(input string) (entries []Entry) {
	for _, n := range strings.Split(input, "\n") {
		digits := strings.SplitN(n, " | ", 2)
		entries = append(entries, Entry{
			strings.Fields(digits[0]),
			strings.Fields(digits[1]),
		})
	}
	return
}

func toBin(digit string) (value uint8) {
	for _, d := range digit {
		value += 1 << (int(d) - int('a'))
	}
	return
}

func Part1(input string) (count int) {
	for _, entry := range parseInput(input) {
		for _, d := range entry.digits {
			switch len(d) {
			case 2, 3, 4, 7:
				count++
			}
		}
	}
	return
}

func Part2(input string) (sum int) {
	for _, entry := range parseInput(input) {
		values := decode(entry.mappings)
		l := len(entry.digits)
		for i, d := range entry.digits {
			sum += values[toBin(d)] * int(math.Pow10(l-i-1))
		}
	}
	return
}
