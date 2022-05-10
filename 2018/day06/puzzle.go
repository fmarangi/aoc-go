package day06

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Point struct{ X, Y int }

func (a Point) Distance(b Point) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y)
}

func (from Point) Closest(points []Point) (closest Point) {
	min, counts := 999999, make(map[int]int)
	for _, p := range points {
		dist := from.Distance(p)
		if dist < min {
			min, closest = dist, p
		}
		counts[dist]++
	}

	if counts[min] != 1 {
		return Point{}
	}

	return
}

func (from Point) TotalDistance(points []Point) (tot int) {
	for _, p := range points {
		tot += from.Distance(p)
	}
	return
}

type Area struct{ TopLeft, BottomRight Point }

func (a Area) Points() <-chan Point {
	out := make(chan Point)
	go func() {
		defer close(out)
		for y := a.TopLeft.Y; y <= a.BottomRight.Y; y++ {
			for x := a.TopLeft.X; x <= a.BottomRight.X; x++ {
				out <- Point{x, y}
			}
		}
	}()
	return out
}

func Part1(input string) (largest int) {
	area, points := parseInput(input)
	closest := make(map[Point]int, len(points))

	for p := range area.Points() {
		closest[p.Closest(points)]++
	}

	for _, p := range points {
		largest = utils.Max(closest[p], largest)
	}

	return
}

func Part2(input string) (region int) {
	area, points := parseInput(input)
	for p := range area.Points() {
		if p.TotalDistance(points) < 10000 {
			region++
		}
	}
	return
}

func parseInput(input string) (area Area, points []Point) {
	area.TopLeft.X, area.TopLeft.Y = 99999, 99999
	for _, row := range strings.Split(input, "\n") {
		var point Point
		fmt.Sscanf(row, "%d, %d", &point.X, &point.Y)
		points = append(points, point)

		area.TopLeft.X = utils.Min(area.TopLeft.X, point.X)
		area.TopLeft.Y = utils.Min(area.TopLeft.Y, point.Y)
		area.BottomRight.X = utils.Max(area.BottomRight.X, point.X)
		area.BottomRight.Y = utils.Max(area.BottomRight.Y, point.Y)
	}
	return
}
