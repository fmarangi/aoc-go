package day05

import (
	"sort"
	"strings"
)

type Seat struct {
	row, column int
}

func (seat Seat) Id() int {
	return seat.row*8 + seat.column
}

func Part1(input string) (highest int) {
	for _, s := range strings.Fields(input) {
		id := parseSeat(s).Id()
		if id > highest {
			highest = id
		}
	}
	return
}

func Part2(input string) int {
	seats := []int{}
	for _, s := range strings.Fields(input) {
		seats = append(seats, parseSeat(s).Id())
	}

	sort.Ints(seats)
	current := seats[0] - 1
	for _, s := range seats {
		if s-1 != current {
			return current + 1
		}
		current = s
	}
	return -1
}

func parseSeat(data string) (seat Seat) {
	for i := 0; i < 7; i++ {
		if data[6-i] == 'B' {
			seat.row += 1 << i
		}
	}

	for i := 0; i < 3; i++ {
		if data[9-i] == 'R' {
			seat.column += 1 << i
		}
	}

	return
}
