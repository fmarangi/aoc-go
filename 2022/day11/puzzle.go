package day11

import (
	"sort"
	"strings"
)

type MonkeyBusiness []int

func (mb *MonkeyBusiness) Round(monkeys []*Monkey) {
	for n, m := range monkeys {
		(*mb)[n] += len(m.Items)
		for _, i := range m.Items {
			worryLevel := m.Operation(i) / 3
			t := m.Target0
			if worryLevel%m.DivisibleBy == 0 {
				t = m.Target1
			}
			monkeys[t].Items = append(monkeys[t].Items, worryLevel)
		}
		m.Items = make([]int, 0)
	}
}

func (mb *MonkeyBusiness) RoundWithLcm(monkeys []*Monkey, lcm int) {
	for n, m := range monkeys {
		(*mb)[n] += len(m.Items)
		for _, i := range m.Items {
			worryLevel := m.Operation(i) % lcm
			t := m.Target0
			if worryLevel%m.DivisibleBy == 0 {
				t = m.Target1
			}
			monkeys[t].Items = append(monkeys[t].Items, worryLevel)
		}
		m.Items = make([]int, 0)
	}
}

func Part1(input string) int {
	monkeys := parseInput(input)
	mb := make(MonkeyBusiness, len(monkeys))
	for i := 0; i < 20; i++ {
		mb.Round(monkeys)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(mb)))
	return mb[0] * mb[1]
}

func Part2(input string) int {
	monkeys := parseInput(input)
	lcm := 1
	for _, m := range monkeys {
		lcm *= m.DivisibleBy
	}

	mb := make(MonkeyBusiness, len(monkeys))
	for i := 0; i < 10000; i++ {
		mb.RoundWithLcm(monkeys, lcm)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(mb)))
	return mb[0] * mb[1]
}

func parseInput(input string) (monkeys []*Monkey) {
	for _, m := range strings.Split(input, "\n\n") {
		monkeys = append(monkeys, parseMonkey(m))
	}
	return
}
