package day13

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Bus struct{ Id, Delay int }

func Part1(input string) (result int) {
	t, buses := parseInput(input)
	min := 99999
	for _, bus := range buses {
		if wait := bus.Id - t%bus.Id; wait < min {
			min, result = wait, wait*bus.Id
		}
	}
	return
}

func Part2(input string) (timestamp int) {
	_, buses := parseInput(input)
	runningProduct := 1
	for _, bus := range buses {
		for (timestamp+bus.Delay)%bus.Id != 0 {
			timestamp += runningProduct
		}
		runningProduct *= bus.Id
	}
	return
}

func parseInput(input string) (timestamp int, buses []Bus) {
	parts := strings.Fields(input)
	timestamp = utils.ToInt(parts[0])
	for i, b := range strings.Split(parts[1], ",") {
		if b != "x" {
			buses = append(buses, Bus{utils.ToInt(b), i})
		}
	}
	return
}
