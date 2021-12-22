package day14

import (
	"sort"
	"strings"
)

type Polymer map[string]int

func countElements(p Polymer) []int {
	count := map[string]int{}
	for k, v := range p {
		count[k[0:1]] += v
		count[k[1:]] += v
	}

	values := []int{}
	for _, v := range count {
		values = append(values, (v+1)/2)
	}

	sort.Ints(values)
	return values
}

func parseInput(input string) (Polymer, func(Polymer) Polymer) {
	parts := strings.Split(input, "\n\n")
	polymer := Polymer{}
	rules := map[string]string{}

	for i, n := 1, len(parts[0]); i < n; i++ {
		polymer[parts[0][i-1:i+1]] += 1
	}

	for _, rule := range strings.Split(parts[1], "\n") {
		fromto := strings.Split(rule, " -> ")
		rules[fromto[0]] = fromto[1]
	}

	process := func(p Polymer) Polymer {
		next := Polymer{}
		for from, n := range p {
			to := rules[from]
			next[from[0:1]+to] += n
			next[to+from[1:]] += n
		}
		return next
	}

	return polymer, process
}

func Part1(input string) int {
	polymer, process := parseInput(input)
	for i := 0; i < 10; i++ {
		polymer = process(polymer)
	}
	c := countElements(polymer)
	return c[len(c)-1] - c[0]
}

func Part2(input string) int {
	polymer, process := parseInput(input)
	for i := 0; i < 40; i++ {
		polymer = process(polymer)
	}
	c := countElements(polymer)
	return c[len(c)-1] - c[0]
}
