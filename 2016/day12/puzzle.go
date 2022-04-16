package day13

import "github.com/fmarangi/aoc-go/2016/assembunny"

func Part1(input string) int {
	regs := make(assembunny.Registers)
	regs.Execute(input)
	return regs["a"]
}

func Part2(input string) int {
	regs := make(assembunny.Registers)
	regs["c"] = 1
	regs.Execute(input)
	return regs["a"]
}
