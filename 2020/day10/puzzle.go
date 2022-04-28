package day10

import (
	"sort"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func diffs(adapters []int) (d []int) {
	for i, n := 1, len(adapters); i < n; i++ {
		d = append(d, adapters[i]-adapters[i-1])
	}
	return
}

func Part1(input string) int {
	count := make(map[int]int, 2)
	for _, d := range diffs(parseInput(input)) {
		count[d]++
	}
	return count[1] * count[3]
}

func Part2(input string) (n int) {
	var (
		count  int
		groups []int
		values = map[int]int{2: 2, 3: 4, 4: 7}
	)

	for _, d := range diffs(parseInput(input)) {
		switch d {
		case 1:
			count++
		case 3:
			if count > 1 {
				groups = append(groups, count)
			}
			count = 0
		}
	}

	n = 1
	for _, c := range groups {
		n *= values[c]
	}

	return
}

func parseInput(input string) []int {
	adapters := append(utils.Ints(strings.Fields(input)), 0)
	sort.Ints(adapters)
	return append(adapters, adapters[len(adapters)-1]+3)
}
