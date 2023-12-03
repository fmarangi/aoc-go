package day03

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func isDigit(x byte) bool {
	return x >= '0' && x <= '9'
}

func partNumbers(input string) func(int, int) int {
	w, m := strings.Index(input, "\n")+1, len(input)
	ok := func(n int) bool {
		return n >= 0 && n < m && input[n] != '.' && input[n] != '\n'
	}

	return func(from, n int) int {
		for i := -1; i <= 1; i++ {
			for j := -1; j < n+1; j++ {
				if p := from + i*w + j; (p < from || p >= from+n) && ok(p) {
					return p
				}
			}
		}
		return -1
	}
}

func Part1(input string) (sum int) {
	pn := partNumbers(input)
	for i, m := 0, len(input); i < m; i++ {
		if !isDigit(input[i]) {
			continue
		}

		j := 0
		for isDigit(input[i+j]) {
			j++
		}

		if pn(i, j) > -1 {
			sum += utils.ToInt(input[i : i+j])
		}

		i += j
	}
	return
}

func Part2(input string) (sum int) {
	pn := partNumbers(input)
	gears := make(map[int][]int)

	for i, m := 0, len(input); i < m; i++ {
		if !isDigit(input[i]) {
			continue
		}

		j := 0
		for isDigit(input[i+j]) {
			j++
		}

		if gear := pn(i, j); gear > -1 && input[gear] == '*' {
			gears[gear] = append(gears[gear], utils.ToInt(input[i:i+j]))
		}

		i += j
	}

	for _, values := range gears {
		if len(values) == 2 {
			sum += values[0] * values[1]
		}
	}

	return
}
