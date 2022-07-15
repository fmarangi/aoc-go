package day20

import (
	"sort"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Range struct{ From, To int }

func Part1(input string) (min int) {
	for _, r := range parseInput(input) {
		if min < r.From {
			break
		}
		min = utils.Max(min, r.To+1)
	}
	return
}

func Part2(input string) (valid int) {
	var x int
	for _, r := range parseInput(input) {
		if x < r.From {
			valid += r.From - x
		}
		x = utils.Max(x, r.To+1)
	}
	return
}

func parseInput(input string) (ranges []Range) {
	for _, row := range strings.Fields(input) {
		r := utils.Ints(strings.Split(row, "-"))
		ranges = append(ranges, Range{r[0], r[1]})
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].From < ranges[j].From
	})

	return
}
