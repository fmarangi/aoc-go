package day16

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Rule func(int) bool
type Ticket []int

func (t Ticket) Invalid(rules []Rule) (values []int) {
	for _, v := range t {
		valid := false
		for _, rule := range rules {
			if rule(v) {
				valid = true
				break
			}
		}

		if !valid {
			values = append(values, v)
		}
	}
	return
}

func matching(tickets []Ticket, rules []Rule) [][]int {
	n := len(rules)
	options := make([][]int, n)

	for i, rule := range rules {
		for p := 0; p < n; p++ {
			ok := true
			for _, t := range tickets {
				if !rule(t[p]) {
					ok = false
					break
				}
			}

			if ok {
				options[i] = append(options[i], p)
			}
		}
	}

	return options
}

func Part1(input string) int {
	invalid := []int{}
	rules, _, tickets := parseInput(input)
	for _, t := range tickets {
		invalid = append(invalid, t.Invalid(rules)...)
	}
	return utils.Sum(invalid)
}

func Part2(input string) int {
	rules, mine, tickets := parseInput(input)

	valid := []Ticket{}
	for _, t := range tickets {
		if len(t.Invalid(rules)) == 0 {
			valid = append(valid, t)
		}
	}

	n := len(rules)
	candidates := matching(valid, rules)
	mapping, assigned := make(map[int]int, n), make(map[int]bool, n)
	for len(mapping) < n {
		for rule, possibilities := range candidates {
			pos := []int{}
			for _, p := range possibilities {
				if !assigned[p] {
					pos = append(pos, p)
				}
			}

			if len(pos) == 1 {
				mapping[rule] = pos[0]
				assigned[pos[0]] = true
			}
		}
	}

	result := 1
	for i := 0; i < 6; i++ {
		result *= mine[mapping[i]]
	}

	return result
}

func parseInput(input string) (rules []Rule, mine Ticket, tickets []Ticket) {
	parts := strings.Split(input, "\n\n")

	for _, r := range strings.Split(parts[0], "\n") {
		var a, b, c, d int
		fmt.Sscanf(strings.Split(r, ": ")[1], "%d-%d or %d-%d", &a, &b, &c, &d)
		rules = append(rules, func(i int) bool {
			return a <= i && i <= b || c <= i && i <= d
		})
	}

	mine = utils.Ints(strings.Split(strings.Split(parts[1], "\n")[1], ","))

	for _, t := range strings.Split(parts[2], "\n")[1:] {
		tickets = append(tickets, utils.Ints(strings.Split(t, ",")))
	}

	return
}
