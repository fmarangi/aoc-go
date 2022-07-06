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

func digits(signal []int, times int) <-chan int {
	var (
		out  = make(chan int)
		n    = len(signal)
		last = signal[n-1]
		curr = make([]int, times+1)
		prev []int
	)

	// the last digit always stays the same
	for i := 0; i <= times; i++ {
		curr[i] = last
	}

	go func() {
		defer close(out)
		out <- curr[times]
		for i := 1; ; i++ {
			curr, prev = make([]int, times+1), curr
			curr[0] = signal[n-(i%n)-1]
			for j := 1; j <= times; j++ {
				curr[j] = (curr[j-1] + prev[j]) % 10
			}
			out <- curr[times]
		}
	}()

	return out
}

func Part2(input string) (first int) {
	var (
		signal = parseInput(input)
		offset = len(signal)*10000 - join(signal[:7]) - 8
		i      = 0
		result = make([]int, 8)
	)

	for k := range digits(signal, 100) {
		if i >= offset {
			j := offset + 7 - i
			if j < 0 {
				break
			}
			result[j] = k
		}
		i++
	}
	return join(result)
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
