package day08

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Op struct {
	reg, op, ref, cmp string
	delta, value      int
}

type Registrers map[string]int

func (r Registrers) valid(op Op) bool {
	switch op.cmp {
	case "==":
		return r[op.ref] == op.value
	case "!=":
		return r[op.ref] != op.value
	case ">":
		return r[op.ref] > op.value
	case ">=":
		return r[op.ref] >= op.value
	case "<":
		return r[op.ref] < op.value
	case "<=":
		return r[op.ref] <= op.value
	}
	return false
}

func Part1(input string) int {
	regs := Registrers{}
	for _, op := range parseInput(input) {
		if !regs.valid(op) {
			continue
		}

		switch op.op {
		case "inc":
			regs[op.reg] += op.delta
		case "dec":
			regs[op.reg] -= op.delta
		}
	}

	values := make([]int, 0, len(regs))
	for _, v := range regs {
		values = append(values, v)
	}
	return utils.Max(values...)
}

func Part2(input string) (highest int) {
	regs := Registrers{}
	for _, op := range parseInput(input) {
		if !regs.valid(op) {
			continue
		}

		switch op.op {
		case "inc":
			regs[op.reg] += op.delta
		case "dec":
			regs[op.reg] -= op.delta
		}

		if regs[op.reg] > highest {
			highest = regs[op.reg]
		}
	}
	return
}

func parseInput(input string) (ops []Op) {
	for _, row := range strings.Split(input, "\n") {
		var op Op
		fmt.Sscanf(
			row, "%s %s %d if %s %s %d",
			&op.reg, &op.op, &op.delta, &op.ref, &op.cmp, &op.value)
		ops = append(ops, op)
	}
	return
}
