package day24

import (
	"regexp"
	"strings"
)

type Tile struct{ X, Y int }

var offsets = map[string]Tile{
	"e":  {2, 0},
	"se": {1, -1},
	"sw": {-1, -1},
	"w":  {-2, 0},
	"nw": {-1, 1},
	"ne": {1, 1},
}

func (t Tile) Neighbours() (n []Tile) {
	for _, o := range offsets {
		n = append(n, Tile{t.X + o.X, t.Y + o.Y})
	}
	return
}

func Floor(paths [][]string) (floor map[Tile]bool) {
	floor = make(map[Tile]bool)
	for _, path := range paths {
		var curr Tile
		for _, step := range path {
			curr.X += offsets[step].X
			curr.Y += offsets[step].Y
		}
		floor[curr] = !floor[curr]
	}

	for tile := range floor {
		if !floor[tile] {
			delete(floor, tile)
		}
	}

	return
}

func Part1(input string) (black int) {
	for range Floor(parseInput(input)) {
		black++
	}

	return
}

func Part2(input string) int {
	floor := Floor(parseInput(input))
	for i := 0; i < 100; i++ {
		// First, load all reached tiles so far, include their neighbours
		for tile := range floor {
			for _, n := range tile.Neighbours() {
				floor[n] = floor[n]
			}
		}

		next := make(map[Tile]bool)
		for tile, black := range floor {
			adj := 0
			for _, n := range tile.Neighbours() {
				if floor[n] {
					adj++
				}
			}

			switch {
			case adj == 2, black && adj == 1:
				next[tile] = true
			}
		}

		floor = next
	}

	return len(floor)
}

func parseInput(input string) (paths [][]string) {
	pattern := regexp.MustCompile(`([ns][ew]|e|w)`)
	for _, row := range strings.Fields(input) {
		paths = append(paths, pattern.FindAllString(row, -1))
	}
	return
}
