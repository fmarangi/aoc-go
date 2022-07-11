package day13

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Seat struct{ Guest, Next string }
type Seats map[Seat]int

func (s Seats) Happiness(arrangement []string) (h int) {
	for i, n := 0, len(arrangement); i < n; i++ {
		a, b := arrangement[i], arrangement[(i+1)%n]
		h += s[Seat{a, b}]
		h += s[Seat{b, a}]
	}
	return
}

func allGuests(seats map[Seat]int) (guests []string) {
	g := map[string]bool{}
	for s := range seats {
		g[s.Guest] = true
	}

	for guest := range g {
		guests = append(guests, guest)
	}

	return
}

func Part1(input string) (max int) {
	var (
		seats        = parseInput(input)
		guests       = allGuests(seats)
		arrangements = utils.Permutations(guests)
	)

	for _, arrangement := range arrangements {
		if h := seats.Happiness(arrangement); h > max {
			max = h
		}
	}

	return
}

func Part2(input string) (max int) {
	var (
		seats        = parseInput(input)
		guests       = append([]string{"Me"}, allGuests(seats)...)
		arrangements = utils.Permutations(guests)
	)

	for _, arrangement := range arrangements {
		if h := seats.Happiness(arrangement); h > max {
			max = h
		}
	}

	return
}

func parseInput(input string) Seats {
	var (
		seats     = make(Seats)
		format    = "%s would %s %d happiness units by sitting next to %s."
		seat      Seat
		happiness int
		gain      string
	)

	for _, row := range strings.Split(input, "\n") {
		fmt.Sscanf(row[:len(row)-1], format,
			&seat.Guest, &gain, &happiness, &seat.Next)
		seats[seat] = happiness
		if gain == "lose" {
			seats[seat] *= -1
		}
	}

	return seats
}
