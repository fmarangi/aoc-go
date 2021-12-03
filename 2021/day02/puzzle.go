package day02

import (
	"strconv"
	"strings"
)

func Part1(input string) int {
	h, d := 0, 0
	for _, s := range strings.Split(input, "\n") {
		line := strings.Fields(s)
		qty, _ := strconv.Atoi(line[1])
		switch line[0] {
		case "up":
			d -= qty
		case "down":
			d += qty
		case "forward":
			h += qty
		}
	}
	return h * d
}

func Part2(input string) int {
	h, d, aim := 0, 0, 0
	for _, s := range strings.Split(input, "\n") {
		line := strings.Fields(s)
		qty, _ := strconv.Atoi(line[1])
		switch line[0] {
		case "up":
			aim -= qty
		case "down":
			aim += qty
		case "forward":
			h += qty
			d += qty * aim
		}
	}
	return h * d
}
