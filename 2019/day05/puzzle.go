package day05

import (
	"github.com/fmarangi/aoc-go/2019/intcode"
)

func Part1(input string) (result int) {
	in := make(chan int, 1)
	in <- 1

	out := intcode.ParseProgram(input).Run(in)
	for result == 0 {
		result = <-out
	}
	return
}

func Part2(input string) (result int) {
	in := make(chan int, 1)
	in <- 5

	out := intcode.ParseProgram(input).Run(in)
	for result == 0 {
		result = <-out
	}
	return
}
