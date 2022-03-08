package day02

import (
	"fmt"
	"strings"
)

type Policy struct {
	from, to int
	letter   string
}

type row struct {
	policy   Policy
	password string
}

func Frequencies(password string) map[string]int {
	freqs := map[string]int{}
	for _, c := range password {
		freqs[string(c)]++
	}
	return freqs
}

func Part1(input string) (valid int) {
	for _, row := range parseInput(input) {
		frequency := Frequencies(row.password)[row.policy.letter]
		if frequency >= row.policy.from && frequency <= row.policy.to {
			valid++
		}
	}
	return
}

func Part2(input string) (valid int) {
	for _, row := range parseInput(input) {
		to, from := row.policy.from-1, row.policy.to-1
		a, b := string(row.password[from]), string(row.password[to])
		if (a == row.policy.letter) != (b == row.policy.letter) {
			valid++
		}
	}
	return
}

func parseInput(input string) (rows []row) {
	for _, r := range strings.Split(input, "\n") {
		pol, pwd := Policy{}, ""
		fmt.Sscanf(r, "%d-%d %1s: %s", &pol.from, &pol.to, &pol.letter, &pwd)
		rows = append(rows, row{pol, pwd})
	}
	return
}
