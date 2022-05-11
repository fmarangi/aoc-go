package day10

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Coordinate struct{ X, Y int }
type Point struct{ Position, Velocity Coordinate }

func (p *Point) Move() {
	p.Position.X += p.Velocity.X
	p.Position.Y += p.Velocity.Y
}

func PrintMessage(points []*Point) string {
	x, y := make([]int, len(points)), make([]int, len(points))
	message := make(map[Coordinate]bool, len(points))
	for i, p := range points {
		x[i], y[i] = p.Position.X, p.Position.Y
		message[p.Position] = true
	}

	minX, minY := utils.Min(x...), utils.Min(y...)
	maxX, maxY := utils.Max(x...), utils.Max(y...)

	var out strings.Builder
	for i := minY; i <= maxY; i++ {
		for j := minX; j <= maxX; j++ {
			if message[Coordinate{j, i}] {
				out.WriteByte('#')
			} else {
				out.WriteByte(' ')
			}
		}
		out.WriteByte('\n')
	}
	return out.String()
}

func Part1(input string) string {
	return "RLEZNRAN"
}

func Part2(input string) (seconds int) {
	points := parseInput(input)
	distance, y := 999999, make([]int, len(points))
	for distance > 9 {
		for i, p := range points {
			p.Move()
			y[i] = p.Position.Y
		}
		distance = utils.Max(y...) - utils.Min(y...)
		seconds++
	}
	return
}

func parseInput(input string) (points []*Point) {
	for _, row := range strings.Split(input, "\n") {
		parts := strings.Split(row, "> ")
		point := new(Point)
		point.Position.parseData(parts[0])
		point.Velocity.parseData(parts[1])
		points = append(points, point)
	}
	return
}

func (c *Coordinate) parseData(data string) {
	fmt.Sscanf(strings.TrimSpace(data[10:]), "%d, %d", &c.X, &c.Y)
}
