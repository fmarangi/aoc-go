package day24

import (
	"fmt"
	"strings"
)

type Levels []Grid

func (levels Levels) AddLevels() Levels {
	next := make(Levels, 1, len(levels)+2)
	next = append(next, levels...)
	return append(next, Grid{})
}

func (levels Levels) NextState() (next Levels) {
	maxDepth := len(levels) - 1
	offsets := [...]struct{ X, Y int }{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}

	for depth, level := range levels {
		nextLevel := Grid{}

		for i := 0; i < 25; i++ {
			bugs, row, col := 0, i/5, i%5
			if row == 2 && col == 2 {
				continue
			}

			for _, o := range offsets {
				r, c := row+o.Y, col+o.X

				if r == 2 && c == 2 && depth < maxDepth {
					switch {
					case row == 1 && col == 2:
						bugs += levels[depth+1].Row(0)
					case row == 3 && col == 2:
						bugs += levels[depth+1].Row(4)
					case row == 2 && col == 1:
						bugs += levels[depth+1].Col(0)
					case row == 2 && col == 3:
						bugs += levels[depth+1].Col(4)
					}
					continue
				}

				if (r == -1 || r == 5 || c == -1 || c == 5) && depth > 0 {
					switch {
					case r == -1 && levels[depth-1][1][2]:
						bugs++
					case r == 5 && levels[depth-1][3][2]:
						bugs++
					case c == -1 && levels[depth-1][2][1]:
						bugs++
					case c == 5 && levels[depth-1][2][3]:
						bugs++
					}
					continue
				}

				if r >= 0 && c >= 0 && r < 5 && c < 5 && level[r][c] {
					bugs++
				}
			}

			if bugs == 1 || (bugs == 2 && !level[row][col]) {
				nextLevel[row][col] = true
			}
		}

		next = append(next, nextLevel)
	}
	return
}

func (levels Levels) Bugs() (bugs int) {
	for _, level := range levels {
		bugs += level.Bugs()
	}
	return
}

func (levels Levels) String() (result string) {
	for i, level := range levels {
		result += fmt.Sprintf("Depth %d:\n%s\n", i-len(levels)/2, level)
	}
	return
}

func Part1(input string) int {
	grid, states := parseInput(input), map[Grid]bool{}
	for !states[grid] {
		states[grid] = true
		grid = grid.NextState()
	}
	return grid.Biodiversity()
}

func Part2(input string) int {
	levels := Levels{parseInput(input)}
	for i := 0; i < 200; i++ {
		if i%2 == 0 {
			levels = levels.AddLevels()
		}
		levels = levels.NextState()
	}
	return levels.Bugs()
}

func parseInput(input string) (grid Grid) {
	for r, row := range strings.Split(input, "\n") {
		for c, value := range row {
			grid[r][c] = value == '#'
		}
	}
	return
}
