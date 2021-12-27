package day15

import (
	"fmt"
	"sort"
	"strings"
)

type Cell struct {
	pos, risk int
}

type Path struct {
	pos, risk int
}

func (p Path) MoveTo(cell Cell) Path {
	return Path{cell.pos, p.risk + cell.risk}
}

func toInt(cell byte) int {
	return int(cell) - 48 // int('0')
}

func Explore(cave string) func(int) []Cell {
	line := strings.Index(cave, "\n") + 1
	goal := len(cave) - 1
	offsets := [...]int{1, line, -1, -line}
	return func(pos int) (cells []Cell) {
		for _, o := range offsets {
			k := o + pos
			if k >= 0 && k <= goal && cave[k] != '\n' {
				cells = append(cells, Cell{k, toInt(cave[k])})
			}
		}
		return
	}
}

func FullMap(input string) (full string) {
	lines := strings.Split(input, "\n")
	for i := 0; i < 5; i++ {
		for _, line := range lines {
			for j := 0; j < 5; j++ {
				for _, pos := range line {
					value := (toInt(byte(pos))-1+i+j)%9 + 1
					full += fmt.Sprint(value)
				}
			}
			full += "\n"
		}
	}
	return strings.TrimSpace(full)
}

func Part1(input string) int {
	explore := Explore(input)
	goal := len(input) - 1
	seen := map[int]bool{0: true}

	q := []Path{{}}
	for len(q) > 0 {
		current := q[0]
		for _, p := range explore(current.pos) {
			next := current.MoveTo(p)
			if seen[next.pos] {
				continue
			}

			seen[next.pos] = true
			if next.pos == goal {
				return next.risk
			}

			q = append(q, next)
		}

		q = q[1:]
		sort.Slice(q, func(i, j int) bool {
			return q[i].risk < q[j].risk
		})
	}
	return -1
}

func Part2(input string) int {
	return Part1(FullMap(input))
}
