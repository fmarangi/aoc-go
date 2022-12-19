package day11

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Monkey struct {
	Items            []int
	Operation        func(int) int
	DivisibleBy      int
	Target1, Target0 int
}

func parseMonkey(data string) *Monkey {
	m := new(Monkey)
	for _, row := range strings.Split(data, "\n")[1:] {
		switch {
		case strings.HasPrefix(row, "  Starting items: "):
			m.Items = utils.Ints(strings.Split(row[18:], ", "))
		case strings.HasPrefix(row, "  Operation: new = "):
			m.Operation = parseOperation(row[19:])
		case strings.HasPrefix(row, "  Test: divisible by "):
			m.DivisibleBy = utils.ToInt(row[21:])
		case strings.HasPrefix(row, "    If true: throw to monkey "):
			m.Target1 = utils.ToInt(row[29:])
		case strings.HasPrefix(row, "    If false: throw to monkey "):
			m.Target0 = utils.ToInt(row[30:])
		}
	}
	return m
}

func parseOperation(data string) func(int) int {
	parts := strings.Fields(data)
	if parts[2] != "old" {
		y := utils.ToInt(parts[2])
		switch parts[1] {
		case "+":
			return func(x int) int { return x + y }
		case "*":
			return func(x int) int { return x * y }
		}
	}
	return func(x int) int { return x * x }
}
