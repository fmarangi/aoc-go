package day01

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func Fuel(mass int) int {
	return mass/3 - 2
}

func TotalFuel(mass int) (fuel int) {
	for n := Fuel(mass); n > 0; n = Fuel(n) {
		fuel += n
	}
	return
}

func parseInput(input string) (modules []int) {
	for _, n := range strings.Fields(input) {
		modules = append(modules, utils.ToInt(n))
	}
	return
}

func Part1(input string) (result int) {
	for _, module := range parseInput(input) {
		result += Fuel(module)
	}
	return
}

func Part2(input string) (result int) {
	for _, module := range parseInput(input) {
		result += TotalFuel(module)
	}
	return
}
