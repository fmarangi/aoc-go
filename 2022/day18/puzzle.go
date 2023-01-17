package day18

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Cube struct{ X, Y, Z int }

func (a Cube) Adjacent(b Cube) bool {
	return utils.Abs(a.X-b.X)+utils.Abs(a.Y-b.Y)+utils.Abs(a.Z-b.Z) == 1
}

func (a Cube) Neighbours() [6]Cube {
	return [6]Cube{
		{a.X + 1, a.Y, a.Z}, {a.X, a.Y + 1, a.Z}, {a.X, a.Y, a.Z + 1},
		{a.X - 1, a.Y, a.Z}, {a.X, a.Y - 1, a.Z}, {a.X, a.Y, a.Z - 1},
	}
}

func (a Cube) InRange(min, max Cube) bool {
	return a.X >= min.X && a.Y >= min.Y && a.Z >= min.Z &&
		a.X <= max.X && a.Y <= max.Y && a.Z <= max.Z
}

func Part1(input string) (area int) {
	cubes := parseInput(input)
	for i, n := 0, len(cubes); i < n; i++ {
		area += 6
		for j := i + 1; j < n; j++ {
			if cubes[i].Adjacent(cubes[j]) {
				area -= 2
			}
		}
	}
	return
}

func Part2(input string) (area int) {
	var (
		cubes     = parseInput(input)
		available = make(map[Cube]bool, len(cubes))
		min, max  Cube
		x, y, z   []int
	)

	for _, c := range cubes {
		available[c] = true
		x, y, z = append(x, c.X), append(y, c.Y), append(z, c.Z)
	}

	min.X, max.X = utils.Min(x...)-1, utils.Max(x...)+1
	min.Y, max.Y = utils.Min(y...)-1, utils.Max(y...)+1
	min.Z, max.Z = utils.Min(z...)-1, utils.Max(z...)+1

	seen := map[Cube]bool{min: true}
	for q := []Cube{min}; len(q) > 0; q = q[1:] {
		current := q[0]
		for _, n := range current.Neighbours() {
			switch {
			case available[n]:
				area++
			case n.InRange(min, max) && !seen[n]:
				q = append(q, n)
				seen[n] = true
			}
		}
	}

	return
}

func parseInput(input string) (cubes []Cube) {
	for _, c := range strings.Fields(input) {
		var cube Cube
		fmt.Sscanf(c, "%d,%d,%d", &cube.X, &cube.Y, &cube.Z)
		cubes = append(cubes, cube)
	}
	return
}
