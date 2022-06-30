package day07

import (
	"errors"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Instruction struct {
	Op, Target string
	Values     []string
}

type Wires map[string]int

func (w Wires) Value(v string) (int, error) {
	// If it's numeric
	if first := v[0]; first >= '0' && first <= '9' {
		return utils.ToInt(v), nil
	}

	if res, ok := w[v]; ok {
		return res, nil
	}

	return 0, errors.New(v + " has not been defined yet")
}

func (w Wires) Run(i Instruction) {
	switch i.Op {
	case "VALUE":
		value, err := w.Value(i.Values[0])
		if err == nil {
			w[i.Target] = value
		}

	case "AND", "OR", "LSHIFT", "RSHIFT":
		a, errA := w.Value(i.Values[0])
		b, errB := w.Value(i.Values[1])
		if errA == nil && errB == nil {
			switch i.Op {
			case "AND":
				w[i.Target] = a & b
			case "OR":
				w[i.Target] = a | b
			case "LSHIFT":
				w[i.Target] = a << b
			case "RSHIFT":
				w[i.Target] = a >> b
			}
		}

	case "NOT":
		value, err := w.Value(i.Values[0])
		if err == nil {
			w[i.Target] = ^value
		}
	}
}

func Part1(input string) int {
	var (
		signals      = make(Wires)
		instructions = parseInput(input)
		count        = len(instructions)
	)

	for i := 0; signals["a"] == 0; i++ {
		signals.Run(instructions[i%count])
	}

	return signals["a"]
}

func Part2(input string) int {
	var (
		signals      = Wires{"b": Part1(input)}
		instructions = parseInput(input)
		count        = len(instructions)
		inst         Instruction
	)

	for i := 0; signals["a"] == 0; i++ {
		inst = instructions[i%count]
		if inst.Target != "b" {
			signals.Run(inst)
		}
	}

	return signals["a"]
}

func parseInput(input string) (instructions []Instruction) {
	for _, row := range strings.Split(input, "\n") {
		var (
			i      Instruction
			parts  = strings.Split(row, " -> ")
			values = strings.Fields(parts[0])
		)

		i.Target = parts[1]
		switch {
		case len(values) == 1:
			i.Op, i.Values = "VALUE", values
		case values[0] == "NOT":
			i.Op, i.Values = values[0], values[1:]
		default:
			i.Op = values[1]
			i.Values = append(i.Values, values[0], values[2])
		}

		instructions = append(instructions, i)
	}
	return
}
