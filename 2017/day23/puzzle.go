package day23

import (
	"strconv"
	"strings"
)

type Instruction struct {
	Op, X, Y string
}

type Regs map[string]int

func (r Regs) Value(y string) int {
	if v, err := strconv.Atoi(y); err == nil {
		return v
	}
	return r[y]
}

func Part1(input string) (mul int) {
	var (
		program = parseInput(input)
		regs    = make(Regs, 8)
		current Instruction
	)

	for i, n := 0, len(program); i < n; i++ {
		current = program[i]
		switch current.Op {
		case "set":
			regs[current.X] = regs.Value(current.Y)
		case "sub":
			regs[current.X] -= regs.Value(current.Y)
		case "mul":
			regs[current.X] *= regs.Value(current.Y)
			mul++
		case "jnz":
			if regs.Value(current.X) != 0 {
				i += regs.Value(current.Y) - 1
			}
		}
	}

	return
}

func Part2(input string) (composites int) {
	for b := 109300; b <= 126300; b += 17 {
		for d := 2; d*d < b; d++ {
			if b%d == 0 {
				composites++
				break
			}
		}
	}
	return
}

func parseInput(input string) (program []Instruction) {
	for _, row := range strings.Split(input, "\n") {
		parts := strings.Fields(row)
		program = append(program, Instruction{parts[0], parts[1], parts[2]})
	}
	return
}
