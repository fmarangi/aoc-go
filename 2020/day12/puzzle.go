package day12

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

const (
	East = iota
	South
	West
	North
)

var offsets = [...]struct{ n, e int }{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}

type Ship struct{ N, E, dir int }

func (ship *Ship) Forward(value int) {
	ship.Move(ship.dir, value)
}

func (ship *Ship) Move(dir, value int) {
	ship.N += offsets[dir].n * value
	ship.E += offsets[dir].e * value
}

func (ship *Ship) Turn(dir string, value int) {
	switch dir {
	case "R":
		ship.dir = (ship.dir + value/90) % 4
	case "L":
		ship.dir = (ship.dir + 3*value/90) % 4
	}
}

func (ship *Ship) Sail(wp Waypoint, value int) {
	ship.N += wp.N * value
	ship.E += wp.E * value
}

func (ship Ship) Distance() int {
	return utils.Abs(ship.N) + utils.Abs(ship.E)
}

type Waypoint struct{ N, E int }

func (wp *Waypoint) Turn(dir string, value int) {
	n := value / 90
	if dir == "L" {
		n = (3 * n) % 4
	}

	for i := 0; i < n; i++ {
		wp.N, wp.E = -wp.E, wp.N
	}
}

func (wp *Waypoint) Move(dir, value int) {
	wp.N += offsets[dir].n * value
	wp.E += offsets[dir].e * value
}

type Instruction struct {
	Action string
	Value  int
}

func Part1(input string) int {
	ship := new(Ship)
	for _, i := range parseInput(input) {
		switch i.Action {
		case "F":
			ship.Forward(i.Value)
		case "L", "R":
			ship.Turn(i.Action, i.Value)
		case "E":
			ship.Move(East, i.Value)
		case "S":
			ship.Move(South, i.Value)
		case "W":
			ship.Move(West, i.Value)
		case "N":
			ship.Move(North, i.Value)
		}
	}
	return ship.Distance()
}

func Part2(input string) int {
	ship, waypoint := new(Ship), &Waypoint{1, 10}
	for _, i := range parseInput(input) {
		switch i.Action {
		case "F":
			ship.Sail(*waypoint, i.Value)
		case "L", "R":
			waypoint.Turn(i.Action, i.Value)
		case "E":
			waypoint.Move(East, i.Value)
		case "S":
			waypoint.Move(South, i.Value)
		case "W":
			waypoint.Move(West, i.Value)
		case "N":
			waypoint.Move(North, i.Value)
		}
	}
	return ship.Distance()
}

func parseInput(input string) (instructions []Instruction) {
	for _, row := range strings.Fields(input) {
		instructions = append(instructions, Instruction{row[0:1], utils.ToInt(row[1:])})
	}
	return
}
