package day01

import (
	"strconv"
	"strings"
)

func Part1(input string) int {
	nums := parseInput(input)
	for i, n := 0, len(nums); i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]+nums[j] == 2020 {
				return nums[i] * nums[j]
			}
		}
	}
	return -1
}

func Part2(input string) int {
	nums := parseInput(input)
	for i, n := 0, len(nums); i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if nums[i]+nums[j]+nums[k] == 2020 {
					return nums[i] * nums[j] * nums[k]
				}
			}
		}
	}
	return -1
}

func parseInput(input string) (nums []int) {
	for _, v := range strings.Split(input, "\n") {
		n, _ := strconv.ParseInt(v, 10, 0)
		nums = append(nums, int(n))
	}
	return nums
}
