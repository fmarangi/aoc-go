package day11

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Point struct{ X, Y int }

func (a Point) Distance(b Point) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y)
}

func Part1(input string) (sum int) {
	galaxies := parseInput(input, 2)
	for i, n := 0, len(galaxies); i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sum += galaxies[i].Distance(galaxies[j])
		}
	}
	return
}

func Part2(input string) (sum int) {
	galaxies := parseInput(input, 1000000)
	for i, n := 0, len(galaxies); i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			sum += galaxies[i].Distance(galaxies[j])
		}
	}
	return
}

func parseInput(input string, factor int) (galaxies []Point) {
	var (
		w    = strings.Index(input, "\n") + 1
		m    = len(input)
		cols = emptyCols(input)
		rows = emptyRows(input)
		c, r int
	)

	for y := 0; y*w < m; y++ {
		if rows[y] {
			r++
			continue
		}

		c = 0
		for x := 0; x < w-1; x++ {
			if cols[x] {
				c++
			}

			if input[x+y*w] == '#' {
				galaxies = append(galaxies, Point{x + c*(factor-1), y + r*(factor-1)})
			}
		}
	}

	return
}

func emptyRows(chart string) map[int]bool {
	rows := make(map[int]bool)
	for i, row := range strings.Fields(chart) {
		if !strings.Contains(row, "#") {
			rows[i] = true
		}
	}
	return rows
}

func emptyCols(chart string) map[int]bool {
	var (
		cols  = make(map[int]bool)
		w     = strings.Index(chart, "\n") + 1
		m     = len(chart)
		empty bool
	)

	for i := 0; i < w-1; i++ {
		empty = true
		for j := 0; j*w+i < m; j++ {
			if chart[j*w+i] == '#' {
				empty = false
				break
			}
		}

		if empty {
			cols[i] = true
		}
	}

	return cols
}
