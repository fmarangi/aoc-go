package day17

import (
	"sort"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func Combinations(containers []int, liters int) <-chan []int {
	sort.Sort(sort.Reverse(sort.IntSlice(containers)))

	var (
		out   = make(chan []int)
		queue = make([][]int, 0)
	)

	go func() {
		defer close(out)
		for _, c := range containers {
			for _, open := range queue {
				group := append([]int{c}, open...)
				sum := utils.Sum(group)
				switch {
				case sum == liters:
					out <- group
				case sum < liters:
					queue = append(queue, group)
				}
			}
			queue = append(queue, []int{c})
		}
	}()

	return out
}

func Part1(input string) (count int) {
	for range Combinations(parseInput(input), 150) {
		count++
	}
	return
}

func Part2(input string) (count int) {
	lengths := make(map[int]int)
	for c := range Combinations(parseInput(input), 150) {
		lengths[len(c)]++
	}

	min := 999
	for containers, combos := range lengths {
		if containers < min {
			min, count = containers, combos
		}
	}

	return
}

func parseInput(input string) []int {
	return utils.Ints(strings.Fields(input))
}
