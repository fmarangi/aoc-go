package day02

import (
	"sort"
	"strconv"
	"strings"
)

type Present struct {
	l, h, w int
}

func (p *Present) WrappingPaper() int {
	dims := []int{p.l, p.h, p.w}
	sort.Ints(dims)
	return (p.l*p.h+p.h*p.w+p.w*p.l)*2 + dims[0]*dims[1]
}

func (p *Present) Ribbon() int {
	dims := []int{p.l, p.h, p.w}
	sort.Ints(dims)
	return 2*(dims[0]+dims[1]) + p.l*p.h*p.w
}

func Part1(input string) int {
	sum := 0
	for _, p := range parseInput(input) {
		sum += p.WrappingPaper()
	}
	return sum
}

func Part2(input string) int {
	sum := 0
	for _, p := range parseInput(input) {
		sum += p.Ribbon()
	}
	return sum
}

func parseInput(input string) []Present {
	presents := []Present{}
	for _, p := range strings.Split(input, "\n") {
		dims := strings.Split(p, "x")
		l, _ := strconv.ParseInt(dims[0], 10, 0)
		h, _ := strconv.ParseInt(dims[1], 10, 0)
		w, _ := strconv.ParseInt(dims[2], 10, 0)
		presents = append(presents, Present{int(l), int(h), int(w)})
	}
	return presents
}
