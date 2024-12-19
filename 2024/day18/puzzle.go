package day18

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Point struct{ X, Y int }

func (p Point) String() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (p Point) Neighbours() [4]Point {
	return [4]Point{
		{p.X - 1, p.Y}, {p.X + 1, p.Y},
		{p.X, p.Y - 1}, {p.X, p.Y + 1},
	}
}

type Maze map[Point]bool

func (maze Maze) Solve(goal Point) int {
	type path struct {
		Pos   Point
		Steps int
	}

	var (
		start Point
		queue = []path{{start, 0}}
		seen  = map[Point]bool{start: true}
		curr  path
	)

	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		for _, p := range curr.Pos.Neighbours() {
			if p == goal {
				return curr.Steps + 1
			}

			if maze[p] {
				continue
			}

			if p.X >= 0 && p.Y >= 0 && p.X <= goal.X && p.Y <= goal.Y && !seen[p] {
				seen[p] = true
				queue = append(queue, path{p, curr.Steps + 1})
			}
		}
	}

	return -1
}

func Part1(input string) int {
	var corrupted = parseInput(input)
	var maze = make(Maze)

	for _, c := range corrupted[:1024] {
		maze[c] = true
	}

	return maze.Solve(Point{70, 70})
}

func Part2(input string) string {
	var corrupted = parseInput(input)
	var goal = Point{70, 70}
	for i, n := 1025, len(corrupted); i < n; i++ {
		var maze = make(Maze)
		for _, c := range corrupted[:i] {
			maze[c] = true
		}

		if maze.Solve(goal) < 0 {
			return corrupted[i-1].String()
		}
	}

	return "no results"
}

func parseInput(input string) (corrupted []Point) {
	for _, row := range strings.Fields(input) {
		values := utils.Ints(strings.Split(row, ","))
		corrupted = append(corrupted, Point{values[0], values[1]})
	}
	return
}
