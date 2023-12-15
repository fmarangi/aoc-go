package day14

import (
	"sort"
)

const (
	Empty int = iota
	Wall
	Rock
)

const (
	North int = iota
	West
	South
	East
)

type Position struct{ X, Y int }

type Platform map[Position]int

var offsets = map[int]Position{
	North: {0, -1},
	West:  {-1, 0},
	South: {0, 1},
	East:  {1, 0},
}

func (platform Platform) Tilt(direction int) Platform {
	var (
		next   = make(Platform)
		o      = offsets[direction]
		rocks  []Position
		maxCol int
		maxRow int
		step   Position
	)

	for pos, p := range platform {
		switch p {
		case Wall:
			next[pos] = Wall
		case Rock:
			rocks = append(rocks, pos)
		}

		maxCol = max(maxCol, pos.X)
		maxRow = max(maxRow, pos.Y)
	}

	sort.Slice(rocks, func(i, j int) bool {
		a, b := rocks[i], rocks[j]
		switch direction {
		case South:
			return a.Y > b.Y
		case West:
			return a.X < b.X
		case East:
			return a.X > b.X
		}
		return a.Y < b.Y
	})

	for _, rock := range rocks {
		for step = rock; next[step] == Empty && step.X >= 0 && step.Y >= 0 && step.X <= maxCol && step.Y <= maxRow; step.X, step.Y = step.X+o.X, step.Y+o.Y {
			rock = step
		}
		next[rock] = Rock
	}

	return next
}

func (platform Platform) Cycle() Platform {
	for i := 0; i < 4; i++ {
		platform = platform.Tilt(i)
	}
	return platform
}

func (platform Platform) Load() (load int) {
	var lastRow int
	for pos := range platform {
		lastRow = max(lastRow, pos.Y)
	}

	for pos, p := range platform {
		if p == Rock {
			load += lastRow - pos.Y + 1
		}
	}

	return
}

func (platform Platform) String() (output string) {
	var maxCol, maxRow int
	for pos := range platform {
		maxCol = max(maxCol, pos.X)
		maxRow = max(maxRow, pos.Y)
	}

	for i := 0; i <= maxRow; i++ {
		for j := 0; j <= maxCol; j++ {
			switch platform[Position{j, i}] {
			case Empty:
				output += "."
			case Wall:
				output += "#"
			case Rock:
				output += "O"
			}
		}
		output += "\n"
	}

	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
