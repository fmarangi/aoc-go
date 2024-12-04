package day02

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Report []int

func (r Report) IsSafe() bool {
	direction := 1
	if r[0] > r[1] {
		direction = -1
	}
	for i, n, incr := 1, len(r), 0; i < n; i++ {
		incr = direction * (r[i] - r[i-1])
		if incr < 1 || incr > 3 {
			return false
		}
	}
	return true
}

func (r Report) Variants() (variants []Report) {
	for i, n := 0, len(r); i < n; i++ {
		var variant Report
		variant = append(variant, r[:i]...)
		variant = append(variant, r[i+1:]...)
		variants = append(variants, variant)
	}
	return
}

func (r Report) ProblemDampener() bool {
	for _, v := range r.Variants() {
		if v.IsSafe() {
			return true
		}
	}
	return false
}

func Part1(input string) (safe int) {
	for _, report := range parseInput(input) {
		if report.IsSafe() {
			safe++
		}
	}
	return
}

func Part2(input string) (safe int) {
	for _, report := range parseInput(input) {
		if report.IsSafe() || report.ProblemDampener() {
			safe++
		}
	}
	return
}

func parseInput(input string) (reports []Report) {
	for _, row := range strings.Split(input, "\n") {
		levels := utils.Ints(strings.Fields(row))
		reports = append(reports, levels)
	}
	return
}
