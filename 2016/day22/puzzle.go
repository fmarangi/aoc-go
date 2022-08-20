package day22

import (
	"fmt"
	"strings"
)

type Pos struct{ X, Y int }
type Node struct{ Used, Avail int }
type Grid map[Pos]Node

type path struct {
	pos   Pos
	steps int
}

func (n Node) Wall() bool {
	return n.Used*100/(n.Avail+n.Used) > 95
}

func (g Grid) Data() (max Pos) {
	for p := range g {
		if p.X > max.X {
			max.X = p.X
		}
	}
	return
}

func (g Grid) Empty() (p Pos) {
	for p = range g {
		if g[p].Used == 0 {
			break
		}
	}
	return
}

func (g Grid) Path(a, b Pos) int {
	seen := map[Pos]bool{a: true}
	for q := []path{{a, 0}}; len(q) > 0; q = q[1:] {
		curr := q[0]
		for _, o := range []Pos{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
			next := Pos{curr.pos.X + o.X, curr.pos.Y + o.Y}
			if next == b {
				return curr.steps + 1
			}

			if n, ok := g[next]; !seen[next] && ok && !n.Wall() {
				q = append(q, path{next, curr.steps + 1})
				seen[next] = true
			}
		}
	}
	return -1
}

func (a Node) Viable(b Node) bool {
	return a.Used > 0 && a.Used < b.Avail
}

func Part1(input string) (viable int) {
	nodes := parseInput(input)
	for a := range nodes {
		for b := range nodes {
			if a != b && nodes[a].Viable(nodes[b]) {
				viable++
			}
		}
	}
	return
}

func Part2(input string) int {
	var (
		grid  = parseInput(input)
		empty = grid.Empty()
		data  = grid.Data()
	)
	return grid.Path(empty, data) + 5*(grid.Path(data, Pos{})-1)
}

func parseInput(input string) (grid Grid) {
	grid = make(Grid)
	for _, row := range strings.Split(input, "\n")[2:] {
		node, pos := Node{}, Pos{}
		parts := strings.Fields(row)
		fmt.Sscanf(parts[0][15:], "x%d-y%d", &pos.X, &pos.Y)
		fmt.Sscanf(parts[2], "%dT", &node.Used)
		fmt.Sscanf(parts[3], "%dT", &node.Avail)
		grid[pos] = node
	}
	return
}
