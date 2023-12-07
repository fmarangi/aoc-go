package day06

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Race struct {
	Time, Distance int
}

func (r Race) Winning() (options int) {
	for i := 1; i < r.Time; i++ {
		if (r.Time-i)*i > r.Distance {
			options++
		}
	}
	return
}

func Part1(input string) (result int) {
	result = 1
	for _, r := range parseInput(input) {
		if w := r.Winning(); w > 0 {
			result *= w
		}
	}
	return
}

func Part2(input string) int {
	rows := strings.Split(input, "\n")
	race := Race{parseNumber(rows[0]), parseNumber(rows[1])}
	return race.Winning()
}

func parseInput(input string) (races []Race) {
	rows := strings.Split(input, "\n")
	times := utils.Ints(strings.Fields(rows[0])[1:])
	distances := utils.Ints(strings.Fields(rows[1])[1:])
	for i, t := range times {
		races = append(races, Race{t, distances[i]})
	}
	return
}

func parseNumber(x string) int {
	return utils.ToInt(strings.Join(strings.Fields(x)[1:], ""))
}
