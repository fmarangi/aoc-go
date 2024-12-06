package day05

import (
	"sort"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Pair [2]int

func Part1(input string) (result int) {
	updates, cmp := parseInput(input)
	for _, update := range updates {
		if sort.SliceIsSorted(update, cmp(update)) {
			result += update[len(update)/2]
		}
	}
	return
}

func Part2(input string) (result int) {
	updates, cmpFunc := parseInput(input)
	for _, update := range updates {
		cmp := cmpFunc(update)
		if !sort.SliceIsSorted(update, cmp) {
			sort.Slice(update, cmp)
			result += update[len(update)/2]
		}
	}
	return
}

func parseInput(input string) (updates [][]int, cmp func([]int) func(int, int) bool) {
	pairs := make(map[Pair]bool)

	var parts = strings.Split(input, "\n\n")
	for _, p := range strings.Fields(parts[0]) {
		pages := utils.Ints(strings.Split(p, "|"))
		pairs[Pair{pages[0], pages[1]}] = true
	}

	for _, u := range strings.Fields(parts[1]) {
		updates = append(updates, utils.Ints(strings.Split(u, ",")))
	}

	cmp = func(u []int) func(int, int) bool {
		return func(i, j int) bool {
			return pairs[Pair{u[i], u[j]}]
		}
	}

	return
}
