package day02

import (
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) (checksum int) {
	for _, row := range parseInput(input) {
		checksum += row[len(row)-1] - row[0]
	}
	return
}

func Part2(input string) (checksum int) {
	for _, row := range parseInput(input) {
		value, found := FindDivisible(row)
		if found {
			checksum += value
		}
	}
	return
}

func parseInput(input string) (rows [][]int) {
	for _, line := range strings.Split(input, "\n") {
		row := []int{}
		for _, n := range strings.Fields(line) {
			v, _ := strconv.Atoi(n)
			row = append(row, v)
		}
		sort.Ints(row)
		rows = append(rows, row)
	}
	return
}

func FindDivisible(nums []int) (result int, ok bool) {
	for i, n := 0, len(nums); i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j]%nums[i] == 0 {
				return nums[j] / nums[i], true
			}
		}
	}
	return -1, false
}
