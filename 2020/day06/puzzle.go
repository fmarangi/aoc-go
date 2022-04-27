package day06

import "strings"

type group []string

func (g group) questions() int {
	q := make(map[rune]bool)
	for _, p := range g {
		for _, a := range p {
			q[a] = true
		}
	}
	return len(q)
}

func (g group) allAgree() (all int) {
	q, n := make(map[rune]int), len(g)
	for _, p := range g {
		for _, a := range p {
			q[a]++
		}
	}

	for _, a := range q {
		if a == n {
			all++
		}
	}

	return
}

func Part1(input string) (answers int) {
	for _, g := range parseInput(input) {
		answers += g.questions()
	}
	return
}

func Part2(input string) (answers int) {
	for _, g := range parseInput(input) {
		answers += g.allAgree()
	}
	return
}

func parseInput(input string) (groups []group) {
	for _, g := range strings.Split(input, "\n\n") {
		groups = append(groups, group(strings.Fields(g)))
	}
	return
}
