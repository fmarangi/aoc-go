package day23

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type instruction []string

func (i instruction) exec(regs map[string]int, offset *int) {
	value := func(el string) int {
		if el == "a" || el == "b" {
			return regs[el]
		}
		return utils.ToInt(el)
	}

	switch i[0] {
	case "hlf":
		regs[i[1]] /= 2
	case "tpl":
		regs[i[1]] *= 3
	case "inc":
		regs[i[1]]++
	case "jmp":
		*offset += value(i[1]) - 1
	case "jie":
		if value(i[1])%2 == 0 {
			*offset += value(i[2]) - 1
		}
	case "jio":
		if value(i[1]) == 1 {
			*offset += value(i[2]) - 1
		}
	}

	*offset++
}

func Part1(input string) int {
	instructions := parseInput(input)
	regs := map[string]int{}
	for i, n := 0, len(instructions); i < n; {
		instructions[i].exec(regs, &i)
	}
	return regs["b"]
}

func Part2(input string) int {
	instructions := parseInput(input)
	regs := map[string]int{"a": 1}
	for i, n := 0, len(instructions); i < n; {
		instructions[i].exec(regs, &i)
	}
	return regs["b"]
}

func parseInput(input string) (instructions []instruction) {
	for _, row := range strings.Split(input, "\n") {
		parts := strings.Fields(strings.ReplaceAll(row, ",", ""))
		instructions = append(instructions, parts)
	}
	return
}
