package day07

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Equation struct {
	Result  int
	Numbers []int
}

func (eq Equation) TestValues() bool {
	for i, n := 0, len(eq.Numbers); i < utils.IntPow(2, n-1); i++ {
		result := eq.Numbers[0]
		for j, k := 1, i; j < n; j, k = j+1, k/2 {
			switch k % 2 {
			case 0:
				result += eq.Numbers[j]
			case 1:
				result *= eq.Numbers[j]
			}
		}

		if result == eq.Result {
			return true
		}
	}
	return false
}

func (eq Equation) TestValuesWithConcat() bool {
	for i, n := 0, len(eq.Numbers); i < utils.IntPow(3, n-1); i++ {
		result := eq.Numbers[0]
		for j, k := 1, i; j < n; j, k = j+1, k/3 {
			switch k % 3 {
			case 0:
				result += eq.Numbers[j]
			case 1:
				result *= eq.Numbers[j]
			case 2:
				result = result * utils.IntPow(10, log10(eq.Numbers[j]))
				result += eq.Numbers[j]
			}
		}

		if result == eq.Result {
			return true
		}
	}
	return false
}

func Part1(input string) (result int) {
	for _, eq := range parseInput(input) {
		if eq.TestValues() {
			result += eq.Result
		}
	}
	return
}

func Part2(input string) (result int) {
	for _, eq := range parseInput(input) {
		if eq.TestValuesWithConcat() {
			result += eq.Result
		}
	}
	return
}

func parseInput(input string) (equations []Equation) {
	for _, row := range strings.Split(input, "\n") {
		parts := strings.Split(row, ": ")
		equations = append(equations, Equation{
			Result:  utils.ToInt(parts[0]),
			Numbers: utils.Ints(strings.Fields(parts[1])),
		})
	}
	return
}

func log10(i int) (n int) {
	for x := 10; x <= i; n++ {
		x *= 10
	}
	return n + 1
}
