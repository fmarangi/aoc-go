package day03

import (
	"strconv"
	"strings"
)

func bitsAt(nums []int, pos int) (n int) {
	for _, k := range nums {
		if 1<<pos&k != 0 {
			n++
		}
	}
	return
}

func Part1(input string) int {
	nums, bits := parseInput(input), strings.Index(input, "\n")
	gamma, epsilon, n := 0, 0, len(nums)
	for i := 0; i < bits; i++ {
		c := bitsAt(nums, i)
		if c > n-c {
			gamma += 1 << i
		} else {
			epsilon += 1 << i
		}
	}
	return gamma * epsilon
}

func Part2(input string) int {
	nums := parseInput(input)

	generator := nums
	for i := strings.Index(input, "\n"); len(generator) > 1; i-- {
		generator = findRating(generator, i-1, 1)
	}

	scrubber := nums
	for i := strings.Index(input, "\n"); len(scrubber) > 1; i-- {
		scrubber = findRating(scrubber, i-1, 0)
	}

	return generator[0] * scrubber[0]
}

func findRating(nums []int, at, keep int) (rating []int) {
	c := bitsAt(nums, at)
	if c < len(nums)-c {
		keep = (keep + 1) % 2
	}
	for _, num := range nums {
		if num>>at&1 == keep {
			rating = append(rating, num)
		}
	}
	return
}

func parseInput(input string) (nums []int) {
	for _, line := range strings.Split(input, "\n") {
		n, _ := strconv.ParseInt(line, 2, 64)
		nums = append(nums, int(n))
	}
	return
}
