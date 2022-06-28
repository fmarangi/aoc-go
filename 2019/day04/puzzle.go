package day04

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func digits(num int) (result [6]int) {
	for i := 5; i >= 0; i-- {
		num, result[i] = num/10, num%10
	}
	return
}

func valid(num int) bool {
	d := digits(num)
	prev, pair := d[0], false
	for _, curr := range d[1:] {
		switch {
		case curr < prev:
			return false
		case curr == prev:
			pair = true
		}
		prev = curr
	}
	return pair
}

func validPt2(num int) bool {
	d := digits(num)
	prev, pair, count := d[0], false, 1
	for _, curr := range d[1:] {
		switch {
		case curr < prev:
			return false
		case curr == prev:
			count++
		default:
			if count == 2 {
				pair = true
			}
			count = 1
		}
		prev = curr
	}
	return pair || count == 2
}

func Part1(input string) (passwords int) {
	from, to := parseInput(input)
	for n := from; n <= to; n++ {
		if valid(n) {
			passwords++
		}
	}
	return
}

func Part2(input string) (passwords int) {
	from, to := parseInput(input)
	for n := from; n <= to; n++ {
		if validPt2(n) {
			passwords++
		}
	}
	return
}

func parseInput(input string) (int, int) {
	parts := strings.Split(input, "-")
	return utils.ToInt(parts[0]), utils.ToInt(parts[1])
}
