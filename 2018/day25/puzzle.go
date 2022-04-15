package day25

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Point [4]int

func (a Point) Distance(b Point) (distance int) {
	for i := 0; i < 4; i++ {
		distance += utils.Abs(a[i] - b[i])
	}
	return
}

func AllDistances(points []Point) [][]int {
	n := len(points)
	all := make([][]int, n)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if points[i].Distance(points[j]) <= 3 {
				all[i] = append(all[i], j)
				all[j] = append(all[j], i)
			}
		}
	}
	return all
}

func Part1(input string) (constellations int) {
	points := parseInput(input)
	n := len(points)
	distances := AllDistances(points)
	processed := make([]bool, n)

	for i := 0; i < n; i++ {
		if processed[i] {
			continue
		}

		processed[i] = true
		for queue := distances[i][:]; len(queue) > 0; queue = queue[1:] {
			curr := queue[0]
			processed[curr] = true
			for _, next := range distances[curr] {
				if !processed[next] {
					queue = append(queue, next)
				}
			}
		}

		constellations++
	}

	return
}

func parseInput(input string) (points []Point) {
	for _, row := range strings.Fields(input) {
		var p Point
		copy(p[:], utils.Ints(strings.Split(row, ",")))
		points = append(points, p)
	}
	return
}
