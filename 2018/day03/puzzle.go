package day03

import (
	"fmt"
	"strings"
)

type Pos struct{ X, Y int }

type Area struct{ A, B Pos }

type Claim struct {
	Id   int
	Area Area
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (a Area) Overlap(b Area) (Area, error) {
	x1, x2 := max(a.A.X, b.A.X), min(a.B.X, b.B.X)
	y1, y2 := max(a.A.Y, b.A.Y), min(a.B.Y, b.B.Y)
	if x1 < x2 && y1 < y2 {
		return Area{Pos{x1, y1}, Pos{x2, y2}}, nil
	}
	return Area{}, fmt.Errorf("no overlap")
}

func Part1(input string) (inches int) {
	overlaps := make(map[Pos]int)
	for _, c := range parseInput(input) {
		for x := c.Area.A.X; x < c.Area.B.X; x++ {
			for y := c.Area.A.Y; y < c.Area.B.Y; y++ {
				overlaps[Pos{x, y}]++
			}
		}
	}

	for _, o := range overlaps {
		if o > 1 {
			inches++
		}
	}

	return
}

func Part2(input string) int {
	var (
		claims = parseInput(input)
		count  = len(claims)
	)

	for i := 0; i < count; i++ {
		found := true
		for j := 0; j < count; j++ {
			if i == j {
				continue
			}

			_, e := claims[i].Area.Overlap(claims[j].Area)
			if e == nil {
				found = false
				break
			}
		}

		if found {
			return claims[i].Id
		}
	}

	return -1
}

func parseInput(input string) (claims []Claim) {
	var id int
	var A, B Pos
	for _, row := range strings.Split(input, "\n") {
		fmt.Sscanf(row, "#%d @ %d,%d: %dx%d",
			&id, &A.X, &A.Y, &B.X, &B.Y)

		B.X += A.X
		B.Y += A.Y

		claims = append(claims, Claim{id, Area{A, B}})
	}
	return
}
