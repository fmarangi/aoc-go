package day09

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func valid(num int, prev []int) bool {
	for i, n := 0, len(prev); i < n; i++ {
		for j := i + 1; j < n; j++ {
			if prev[i]+prev[j] == num {
				return true
			}
		}
	}
	return false
}

func invalid(nums []int, preamble int) int {
	i := preamble
	for valid(nums[i], nums[i-preamble:i]) {
		i++
	}
	return nums[i]
}

func Part1(input string) int {
	return invalid(parseInput(input), 25)
}

func Part2(input string) int {
	nums := parseInput(input)
	target := invalid(nums, 25)

	var weakness []int
	for i, n := 0, len(nums); i < n; i++ {
		sum, j := nums[i], i
		for sum < target {
			j++
			sum += nums[j]
		}

		if sum == target {
			weakness = nums[i : j+1]
			break
		}
	}

	return utils.Min(weakness...) + utils.Max(weakness...)
}

func parseInput(input string) []int {
	return utils.Ints(strings.Fields(input))
}
