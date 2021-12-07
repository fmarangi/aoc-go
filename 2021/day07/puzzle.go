package day07

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

func parseInput(input string) (nums []int) {
	for _, n := range strings.Split(input, ",") {
		num, _ := strconv.Atoi(n)
		nums = append(nums, num)
	}
	return
}

func abs(num int) int {
	if num > 0 {
		return num
	}
	return -num
}

func avg(nums []int) float64 {
	average := 0
	for _, n := range nums {
		average += n
	}
	return float64(average) / float64(len(nums))
}

func totalFuel(distance int) int {
	return distance * (distance + 1) / 2
}

func Part1(input string) (fuel int) {
	nums := parseInput(input)
	sort.Ints(nums)
	median := nums[len(nums)/2]
	for _, n := range nums {
		fuel += abs(n - median)
	}
	return
}

func Part2(input string) int {
	nums := parseInput(input)
	average := avg(nums)
	a, b := int(average), int(math.Ceil(average))
	fuelA, fuelB := 0, 0
	for _, n := range nums {
		fuelA += totalFuel(abs(n - a))
		fuelB += totalFuel(abs(n - b))
	}
	if fuelA < fuelB {
		return fuelA
	}
	return fuelB
}
