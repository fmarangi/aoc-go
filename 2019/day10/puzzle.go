package day10

import (
	"math"
	"sort"

	"github.com/fmarangi/aoc-go/utils"
)

type Point struct{ X, Y int }

func (a Point) Angle(b Point) (angle float64) {
	return math.Pi - math.Atan2(float64(b.X-a.X), float64(b.Y-a.Y))
}

func (a Point) Distance(b Point) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y)
}

func bestLocation(asteroids []Point) (best Point, max int) {
	var (
		detected = map[Point]map[float64]bool{}
		n        = len(asteroids)
	)

	for i := 0; i < n; i++ {
		curr := asteroids[i]
		detected[curr] = map[float64]bool{}
		for j := 0; j < n; j++ {
			if i != j {
				detected[curr][curr.Angle(asteroids[j])] = true
			}
		}
	}

	for p, angles := range detected {
		if s := len(angles); s > max {
			max, best = s, p
		}
	}

	return
}

func Part1(input string) (max int) {
	_, max = bestLocation(parseInput(input))
	return
}

func Part2(input string) (result int) {
	var (
		asteroids = parseInput(input)
		best, _   = bestLocation(asteroids)
	)

	sort.Slice(asteroids, func(i, j int) bool {
		a1, a2 := best.Angle(asteroids[i]), best.Angle(asteroids[j])
		if a1 == a2 {
			return best.Distance(asteroids[i]) < best.Distance(asteroids[j])
		}
		return a1 < a2
	})

	prev, count := float64(-1), 0
	for _, a := range asteroids {
		if a == best {
			continue
		}

		if angle := best.Angle(a); angle != prev {
			prev, count, result = angle, count+1, a.X*100+a.Y
		}

		if count == 200 {
			break
		}
	}

	return
}

func parseInput(input string) (asteroids []Point) {
	var x, y int
	for _, p := range input {
		switch p {
		case '\n':
			x, y = 0, y+1
		case '#':
			asteroids = append(asteroids, Point{x, y})
			fallthrough
		default:
			x++
		}
	}
	return
}
