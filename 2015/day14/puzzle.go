package day14

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Reindeer struct {
	Name              string
	Speed, Time, Rest int
}

func (r Reindeer) Distance(seconds int) int {
	cycle := r.Time + r.Rest
	last := Min(r.Time, seconds%cycle)
	return (seconds/cycle*r.Time + last) * r.Speed
}

func Part1(input string) (max int) {
	for _, r := range parseInput(input) {
		if distance := r.Distance(2503); distance > max {
			max = distance
		}
	}
	return
}

func Part2(input string) int {
	reindeers := parseInput(input)
	scores := make([]int, len(reindeers))
	for i := 1; i <= 2503; i++ {
		max := 0
		for _, r := range reindeers {
			if d := r.Distance(i); d > max {
				max = d
			}
		}

		for c, r := range reindeers {
			if r.Distance(i) == max {
				scores[c]++
			}
		}
	}
	return utils.Max(scores...)
}

func parseInput(input string) (reindeers []Reindeer) {
	for _, r := range strings.Split(input, "\n") {
		reindeer := Reindeer{}
		fmt.Sscanf(
			r, "%s can fly %d km/s for %d seconds, but then must rest for %d seconds.",
			&reindeer.Name, &reindeer.Speed, &reindeer.Time, &reindeer.Rest)
		reindeers = append(reindeers, reindeer)
	}
	return
}
