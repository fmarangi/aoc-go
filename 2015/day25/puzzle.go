package day25

import (
	"fmt"
	"strings"
)

func Part1(input string) int {
	var (
		row, col = parseInput(input)
		times    = (row+col-1)*(row+col-2)/2 + col
		code     = 20151125
	)

	for i := 1; i < times; i++ {
		code = code * 252533 % 33554393
	}

	return code
}

func parseInput(input string) (row, col int) {
	parts := strings.Split(input, "row ")
	fmt.Sscanf(parts[1], "%d, column %d.", &row, &col)
	return
}
