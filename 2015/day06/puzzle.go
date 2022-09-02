package day06

import (
	"fmt"
	"strings"
)

const (
	On int = iota
	Off
	Toggle
)

type Pos struct{ X, Y int }

type Instruction struct {
	Type     int
	From, To Pos
}

type Lights [1000][1000]uint8

func (i Instruction) Lights() <-chan Pos {
	out := make(chan Pos)
	go func() {
		defer close(out)
		for x := i.From.X; x <= i.To.X; x++ {
			for y := i.From.Y; y <= i.To.Y; y++ {
				out <- Pos{x, y}
			}
		}
	}()
	return out
}

func (lights Lights) Sum() (sum int) {
	for _, y := range lights {
		for _, x := range y {
			sum += int(x)
		}
	}
	return
}

func Part1(input string) int {
	var lights Lights
	for _, i := range parseInput(input) {
		for p := range i.Lights() {
			switch i.Type {
			case On:
				lights[p.Y][p.X] = 1
			case Off:
				lights[p.Y][p.X] = 0
			case Toggle:
				lights[p.Y][p.X] = (lights[p.Y][p.X] + 1) % 2
			}
		}
	}
	return lights.Sum()
}

func Part2(input string) int {
	var lights Lights
	for _, i := range parseInput(input) {
		for p := range i.Lights() {
			switch i.Type {
			case On:
				lights[p.Y][p.X] += 1
			case Off:
				if lights[p.Y][p.X] > 0 {
					lights[p.Y][p.X] -= 1
				}
			case Toggle:
				lights[p.Y][p.X] += 2
			}
		}
	}
	return lights.Sum()
}

func parseInput(input string) (instructions []Instruction) {
	for _, row := range strings.Split(input, "\n") {
		var i Instruction
		switch {
		case strings.HasPrefix(row, "turn on "):
			i.Type, row = On, row[8:]
		case strings.HasPrefix(row, "turn off "):
			i.Type, row = Off, row[9:]
		case strings.HasPrefix(row, "toggle "):
			i.Type, row = Toggle, row[7:]
		}

		fmt.Sscanf(row, "%d,%d through %d,%d",
			&i.From.X, &i.From.Y, &i.To.X, &i.To.Y)

		instructions = append(instructions, i)
	}
	return
}
