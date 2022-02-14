package day17

import (
	"fmt"

	"github.com/fmarangi/aoc-go/utils"
)

type Position struct {
	X, Y int
}

type Area struct {
	A, B Position
}

func parseInput(input string) Area {
	var a, b, c, d int
	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &a, &b, &c, &d)
	return Area{A: Position{a, d}, B: Position{b, c}}
}

func (p Position) Within(area Area) bool {
	return p.X >= area.A.X && p.X <= area.B.X && p.Y <= area.A.Y && p.Y >= area.B.Y
}

func (p Position) Past(area Area) bool {
	return p.Y < area.B.Y || p.X > area.B.X
}

func (area Area) checkVelocity(x, y int) bool {
	p := Position{x, y}
	for !p.Past(area) {
		if p.Within(area) {
			return true
		}
		y--
		if x > 0 {
			x--
		}
		p = Position{p.X + x, p.Y + y}
	}
	return false
}

func xCandidates(from, to int) (candidates []int) {
	for i := 1; i <= to; i++ {
		for p, j := 0, i; j > 0 && p <= to; p, j = p+j, j-1 {
			if p >= from {
				candidates = append(candidates, i)
				break
			}
		}
	}
	return
}

func (a Area) AllVelocities() (all []Position) {
	for _, x := range xCandidates(a.A.X, a.B.X) {
		for y := utils.Abs(a.B.Y) - 1; y >= a.B.Y; y-- {
			if a.checkVelocity(x, y) {
				all = append(all, Position{x, y})
			}
		}
	}
	return
}

func Part1(input string) int {
	y := parseInput(input).B.Y
	return y * (y + 1) / 2
}

func Part2(input string) int {
	return len(parseInput(input).AllVelocities())
}
