package day04

import "strings"

type Point struct{ X, Y int }

func (a Point) Add(b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}

type Board map[Point]rune

func (b Board) Xmas(p Point) (xmas int) {
	if b[p] == 'X' {
		var offsets = []Point{
			{1, 0}, {0, 1}, {-1, 0}, {0, -1},
			{1, 1}, {1, -1}, {-1, -1}, {-1, 1},
		}
		for _, o := range offsets {
			var pos = p
			var value strings.Builder
			for i := 0; i < 3; i++ {
				pos = pos.Add(o)
				value.WriteRune(b[pos])
			}

			if value.String() == "MAS" {
				xmas++
			}
		}
	}
	return
}

func (b Board) Mas(p Point) (xmas int) {
	if b[p] == 'A' {
		var offsets = []Point{
			{-1, -1}, {1, -1},
			{-1, 1}, {1, 1},
		}

		var value strings.Builder
		for _, o := range offsets {
			value.WriteRune(b[p.Add(o)])
		}

		switch value.String() {
		case "MSMS", "SMSM", "MMSS", "SSMM":
			xmas++
		}
	}
	return
}

func Part1(input string) (xmas int) {
	var board = parseInput(input)
	for c := range board {
		xmas += board.Xmas(c)
	}
	return
}

func Part2(input string) (xmas int) {
	var board = parseInput(input)
	for c := range board {
		xmas += board.Mas(c)
	}
	return

}

func parseInput(input string) (board Board) {
	board = make(Board)

	var x, y int
	for _, c := range input {
		switch c {
		case 'X', 'M', 'A', 'S':
			board[Point{x, y}] = c
			x++
		case '\n':
			x, y = 0, y+1
		}
	}

	return
}
