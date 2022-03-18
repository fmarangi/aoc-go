package day18

import "strings"

func Part1(input string) int {
	nums := parseInput(input)
	num := nums[0]
	for _, n := range nums[1:] {
		num = Sum(num, n)
	}
	return num.Magnitude()
}

func Part2(input string) (max int) {
	nums := parseInput(input)
	for i, n := 0, len(nums); i < n; i++ {
		for j := i + 1; j < n; j++ {
			if sum := Sum(nums[i], nums[j]).Magnitude(); sum > max {
				max = sum
			}
			if sum := Sum(nums[j], nums[i]).Magnitude(); sum > max {
				max = sum
			}
		}
	}
	return
}

func parseInput(input string) (nums []Number) {
	for _, n := range strings.Fields(input) {
		nums = append(nums, parseNumber(n))
	}
	return
}
