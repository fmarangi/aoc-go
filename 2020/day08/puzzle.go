package day08

import (
	"errors"
	"fmt"
	"strings"
)

type Instruction struct {
	Op  string
	Arg int
}

func run(code []Instruction) (acc int, e error) {
	var (
		count = len(code)
		seen  = make(map[int]bool, count)
	)

	for i := 0; i < count; i++ {
		if seen[i] {
			e = errors.New("loop detected")
			break
		}

		seen[i] = true
		switch code[i].Op {
		case "jmp":
			i += code[i].Arg - 1
		case "acc":
			acc += code[i].Arg
		}
	}
	return
}

func Part1(input string) (acc int) {
	acc, _ = run(parseInput(input))
	return
}

func Part2(input string) int {
	var (
		code     = parseInput(input)
		count    = len(code)
		repaired = make([]Instruction, count)
	)

	for i := 0; i < count; i++ {
		if code[i].Op != "jmp" {
			continue
		}

		copy(repaired, code)
		repaired[i] = Instruction{Op: "nop"}
		if acc, e := run(repaired); e == nil {
			return acc
		}
	}

	return -1
}

func parseInput(input string) (instructions []Instruction) {
	for _, row := range strings.Split(input, "\n") {
		var i Instruction
		fmt.Sscanf(row, "%s %d", &i.Op, &i.Arg)
		instructions = append(instructions, i)
	}
	return
}
