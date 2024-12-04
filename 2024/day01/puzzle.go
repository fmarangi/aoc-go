package day01

import (
	"sort"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func Part1(input string) (distance int) {
	a, b := parseInput(input)
	sort.Ints(a)
	sort.Ints(b)
	for i, n := range a {
		distance += utils.Abs(n - b[i])
	}
	return
}

func Part2(input string) (similarity int) {
	a, b := parseInput(input)
	times := make(map[int]int)
	for _, n := range b {
		times[n]++
	}
	for _, i := range a {
		similarity += i * times[i]
	}
	return
}

func parseInput(input string) (a, b []int) {
	for i, val := range strings.Fields(input) {
		if i%2 == 0 {
			a = append(a, utils.ToInt(val))
		} else {
			b = append(b, utils.ToInt(val))
		}
	}
	return
}
