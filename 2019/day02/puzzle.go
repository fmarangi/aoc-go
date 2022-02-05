package day02

import (
	"github.com/fmarangi/aoc-go/2019/intcode"
)

func run(program intcode.Program, noun, verb int) int {
	program[1], program[2] = noun, verb
	program.Process()
	return program[0]
}

func Part1(input string) int {
	return run(intcode.ParseProgram(input), 12, 2)
}

func Part2(input string) (result int) {
	program := intcode.ParseProgram(input)
	for i := 0; i < 9999; i++ {
		copia := make([]int, len(program))
		copy(copia, program)
		if run(copia, i/100, i%100) == 19690720 {
			return i
		}
	}
	return -1
}
