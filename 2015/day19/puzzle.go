package day19

import (
	"regexp"
	"sort"
	"strings"
)

type Rule struct{ from, to string }

func Part1(input string) int {
	molecule, rules := parseInput(input)

	molecules := make(map[string]bool)
	for _, rule := range rules {
		pattern := regexp.MustCompile(rule.from)
		for _, m := range pattern.FindAllStringIndex(molecule, -1) {
			molecules[molecule[:m[0]]+rule.to+molecule[m[1]:]] = true
		}
	}

	return len(molecules)
}

func Part2(input string) (steps int) {
	molecule, rules := parseInput(input)

	// First, transform everything between Rn and Ar
	pattern := regexp.MustCompile(`Rn([^Rn]+?)Ar`)
	molecule = pattern.ReplaceAllStringFunc(molecule, func(s string) string {
		res, moves := transform(rules, s)
		steps += moves
		return res
	})

	// Now we can safely reduce the molecule
	if result, s := transform(rules, molecule); result == "e" {
		return steps + s
	}

	return -1
}

func parseInput(input string) (molecule string, rules []Rule) {
	parts := strings.Split(input, "\n\n")
	molecule = parts[1]
	for _, row := range strings.Split(parts[0], "\n") {
		rule := strings.Split(row, " => ")
		rules = append(rules, Rule{rule[0], rule[1]})
	}

	sort.Slice(rules, func(i, j int) bool {
		return len(rules[i].to) > len(rules[j].to)
	})

	return
}
func transform(rules []Rule, initial string) (molecule string, steps int) {
	var before string
	for molecule = initial; before != molecule; steps++ {
		before = molecule
		for _, rule := range rules {
			if i := strings.Index(molecule, rule.to); i >= 0 {
				molecule = molecule[:i] + rule.from + molecule[i+len(rule.to):]
				break
			}
		}
	}
	return molecule, steps - 1
}
