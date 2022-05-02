package day19

import "github.com/fmarangi/aoc-go/utils"

func pow(n, base int) (result int) {
	for result = 1; result*base <= n; result *= base {
	}
	return
}

func stealPresents(n int) int {
	return 2*(n-pow(n, 2)) + 1
}

func stealPresentsAcross(n int) int {
	p3 := pow(n, 3)
	switch {
	case p3 == n:
		return n
	case n < 2*p3:
		return n - p3
	}
	return 2*n - 3*p3
}

func Part1(input string) int {
	return stealPresents(utils.ToInt(input))
}

func Part2(input string) int {
	return stealPresentsAcross(utils.ToInt(input))
}
