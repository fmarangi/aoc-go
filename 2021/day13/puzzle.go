package day13

import (
	"strconv"
	"strings"
)

type Origami map[Point]bool
type Point struct {
	x, y int
}

func parseInput(input string) (*Origami, []string) {
	paper := &Origami{}
	parts := strings.SplitN(input, "\n\n", 2)
	for _, p := range strings.Split(parts[0], "\n") {
		coordinates := strings.Split(p, ",")
		point := Point{x: toInt(coordinates[0]), y: toInt(coordinates[1])}
		(*paper)[point] = true
	}
	return paper, strings.Split(parts[1], "\n")
}

func toInt(num string) int {
	val, _ := strconv.Atoi(num)
	return val
}

func (paper *Origami) foldX(x int) {
	for point := range *paper {
		if point.x > x {
			delete(*paper, point)
			value := x*2 - point.x
			(*paper)[Point{value, point.y}] = true
		}
	}
}

func (paper *Origami) foldY(y int) {
	for point := range *paper {
		if point.y > y {
			delete(*paper, point)
			value := y*2 - point.y
			(*paper)[Point{point.x, value}] = true
		}
	}
}

func (paper Origami) Plot() (plot string) {
	x, y := 0, 0
	for point := range paper {
		if point.x > x {
			x = point.x
		}
		if point.y > y {
			y = point.y
		}
	}

	for i := 0; i <= y; i++ {
		for j := 0; j <= x; j++ {
			_, ok := paper[Point{j, i}]
			if ok {
				plot += "#"
			} else {
				plot += " "
			}
		}
		plot += "\n"
	}

	return
}

func Part1(input string) int {
	o, instructions := parseInput(input)
	toDo := strings.Split(instructions[0], "=")
	if toDo[0] == "fold along x" {
		o.foldX(toInt(toDo[1]))
	} else {
		o.foldY(toInt(toDo[1]))
	}
	return len(*o)
}

func Part2(input string) string {
	o, instructions := parseInput(input)
	for _, i := range instructions {
		toDo := strings.Split(i, "=")
		if toDo[0] == "fold along x" {
			o.foldX(toInt(toDo[1]))
		} else {
			o.foldY(toInt(toDo[1]))
		}
	}
	return "UCLZRAZU"
}
