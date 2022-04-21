package day16

import "github.com/fmarangi/aoc-go/utils"

var pattern = [4]int{0, 1, 0, -1}

func mul(row, col int) int {
	return pattern[(row+1)/(col+1)%4]
}

func phase(signal []int) []int {
	n := len(signal)
	next := make([]int, n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			next[j] += signal[i] * mul(i, j)
		}
	}

	for i := 0; i < n; i++ {
		next[i] = utils.Abs(next[i]) % 10
	}

	return next
}

func FFT(signal []int, phases int) []int {
	for i := 0; i < phases; i++ {
		signal = phase(signal)
	}
	return signal
}

func Part1(input string) (first int) {
	signal := parseInput(input)
	return join(FFT(signal, 100)[:8])
}

func parseInput(input string) []int {
	signal := make([]int, len(input))
	for i, c := range input {
		signal[i] = int(c - '0')
	}
	return signal
}

func join(nums []int) (result int) {
	for _, n := range nums {
		result = result*10 + n
	}
	return
}
