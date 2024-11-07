package day18

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Point struct{ X, Y int }

func (a Point) Add(b Point, times int) Point {
	return Point{a.X + b.X*times, a.Y + b.Y*times}
}

func (a Point) Distance(b Point) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y)
}

func (a Point) CrossProduct(b Point) int {
	return a.X*b.Y - a.Y*b.X
}

type Lagoon []Point

func (l Lagoon) Perimeter() (p int) {
	for i, n := 0, len(l); i < n; i++ {
		p += l[i].Distance(l[(i+1)%n])
	}
	return
}

func (l Lagoon) Area() (area int) {
	for i, n := 0, len(l); i < n; i++ {
		area += l[i].CrossProduct(l[(i+1)%n])
	}
	return area/2 + l.Perimeter()/2 + 1
}

var offsets = map[string]Point{
	"R": {1, 0},
	"D": {0, 1},
	"L": {-1, 0},
	"U": {0, -1},

	"0": {1, 0},
	"1": {0, 1},
	"2": {-1, 0},
	"3": {0, -1},
}

func Part1(input string) int {
	return parseInput(input).Area()
}

func Part2(input string) (power int) {
	return parseHex(input).Area()
}

func parseInput(input string) (l Lagoon) {
	var (
		curr Point
		n    int
		d    string
	)

	for _, row := range strings.Split(input, "\n") {
		fmt.Sscanf(row, "%s %d", &d, &n)
		curr = curr.Add(offsets[d], n)
		l = append(l, curr)
	}
	return
}

func parseHex(input string) (l Lagoon) {
	var (
		curr Point
		n    int
		d    string
	)

	for _, row := range strings.Split(input, "#")[1:] {
		fmt.Sscanf(row, "%5x%1s", &n, &d)
		curr = curr.Add(offsets[d], n)
		l = append(l, curr)
	}

	return
}
