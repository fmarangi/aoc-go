package day06

import (
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	fish := make([]int, 9)
	for _, t := range strings.Split(input, ",") {
		timer, _ := strconv.Atoi(t)
		fish[timer]++
	}
	return fish
}

func simulate(fish []int, days int) (count int) {
	for i := 0; i < days; i++ {
		zero := fish[0]
		copy(fish, fish[1:])
		fish[6] += zero
		fish[8] = zero
	}
	for _, c := range fish {
		count += c
	}
	return
}

func Part1(input string) int {
	return simulate(parseInput(input), 80)
}

func Part2(input string) int {
	return simulate(parseInput(input), 256)
}
