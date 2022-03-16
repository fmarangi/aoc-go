package day01

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

const (
	North = iota
	East
	South
	West

	Right = 1
	Left  = -1
)

type Direction struct{ Turn, Steps int }
type Position struct{ X, Y int }

func (p Position) Distance(other Position) int {
	return utils.Abs(p.X-other.X) + utils.Abs(p.Y-other.Y)
}

func Part1(input string) int {
	var pos Position
	for pos = range Move(parseInput(input)) {
	}
	return pos.Distance(Position{})
}

func Part2(input string) int {
	seen := map[Position]bool{}
	for pos := range Move(parseInput(input)) {
		if seen[pos] {
			return pos.Distance(Position{})
		}
		seen[pos] = true
	}
	return -1
}

func Move(dirs []Direction) chan Position {
	out := make(chan Position)
	go func() {
		offset := map[int]Position{
			North: {0, 1},
			East:  {1, 0},
			South: {0, -1},
			West:  {-1, 0},
		}
		curr, dir := Position{}, North
		out <- curr

		for _, d := range dirs {
			dir = (dir + 4 + d.Turn) % 4
			for i := 0; i < d.Steps; i++ {
				x, y := offset[dir].X+curr.X, offset[dir].Y+curr.Y
				curr = Position{x, y}
				out <- curr
			}
		}

		close(out)
	}()
	return out
}

func parseInput(input string) (dirs []Direction) {
	for _, d := range strings.Split(input, ", ") {
		dir, steps := Right, utils.ToInt(d[1:])
		if d[0] == 'L' {
			dir = Left
		}
		dirs = append(dirs, Direction{dir, steps})
	}
	return
}
