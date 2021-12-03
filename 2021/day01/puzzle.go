package day01

import (
	"strconv"
	"strings"
)

func Part1(input string) (larger int) {
	var prev int
	for i, num := range parseInput(input) {
		if i > 0 && prev < num {
			larger++
		}
		prev = num
	}
	return
}

func Part2(input string) (larger int) {
	prev, nums := 0, parseInput(input)
	for i, n := 0, len(nums); i < n-2; i++ {
		curr := nums[i] + nums[i+1] + nums[i+2]
		if i > 0 && prev < curr {
			larger++
		}
		prev = curr
	}
	return
}

func parseInput(input string) (nums []int) {
	for _, line := range strings.Split(input, "\n") {
		num, _ := strconv.Atoi(line)
		nums = append(nums, num)
	}
	return
}
