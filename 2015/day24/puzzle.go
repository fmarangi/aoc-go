package day24

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func QE(nums []int, target int) int {
	var (
		qty, qe int
		open    [][]int
	)

	for _, n := range nums {
		for _, o := range open {
			if qty != 0 && len(o) > qty {
				continue
			}

			sum := utils.Sum(o) + n
			switch {
			case sum < target:
				open = append(open, append([]int{n}, o...))
			case sum == target:
				prod := Product(append([]int{n}, o...))
				if qe == 0 || len(o)+1 < qty || prod < qe {
					qty, qe = len(o)+1, prod
				}
			}
		}

		open = append(open, []int{n})
	}

	return qe
}

func Product(packages []int) (prod int) {
	prod = 1
	for _, p := range packages {
		prod *= p
	}
	return
}

func Part1(input string) int {
	nums := parseInput(input)
	return QE(nums, utils.Sum(nums)/3)
}

func Part2(input string) int {
	nums := parseInput(input)
	return QE(nums, utils.Sum(nums)/4)
}

func parseInput(input string) []int {
	return utils.Ints(strings.Fields(input))
}
