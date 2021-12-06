package day05

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func increment(a, b int) int {
	if a < b {
		return 1
	}
	return -1
}

func overlap(diagram map[string]int) (count int) {
	for _, val := range diagram {
		if val > 1 {
			count++
		}
	}
	return
}

func Part1(input string) int {
	diagram := map[string]int{}
	for _, line := range parseInput(input) {
		switch {
		case line[0] == line[2]:
			inc := increment(line[1], line[3])
			for i := line[1]; i != line[3]+inc; i += inc {
				p := fmt.Sprintf("%d,%d", line[0], i)
				diagram[p] = diagram[p] + 1
			}
		case line[1] == line[3]:
			inc := increment(line[0], line[2])
			for i := line[0]; i != line[2]+inc; i += inc {
				p := fmt.Sprintf("%d,%d", i, line[1])
				diagram[p] = diagram[p] + 1
			}
		}
	}
	return overlap(diagram)
}

func Part2(input string) int {
	diagram := map[string]int{}
	for _, line := range parseInput(input) {
		switch {
		case line[0] == line[2]:
			inc := increment(line[1], line[3])
			for i := line[1]; i != line[3]+inc; i += inc {
				p := fmt.Sprintf("%d,%d", line[0], i)
				diagram[p] = diagram[p] + 1
			}
		case line[1] == line[3]:
			inc := increment(line[0], line[2])
			for i := line[0]; i != line[2]+inc; i += inc {
				p := fmt.Sprintf("%d,%d", i, line[1])
				diagram[p] = diagram[p] + 1
			}
		default:
			xInc, yInc := increment(line[0], line[2]), increment(line[1], line[3])
			for i, j := line[0], line[1]; i != line[2]+xInc; i, j = i+xInc, j+yInc {
				p := fmt.Sprintf("%d,%d", i, j)
				diagram[p] += 1
			}
		}
	}
	return overlap(diagram)
}

func parseInput(input string) (lines [][]int) {
	pattern := regexp.MustCompile(`\d+`)
	for _, p := range strings.Split(input, "\n") {
		line := make([]int, 4)
		lines = append(lines, line)
		for i, n := range pattern.FindAllString(p, 4) {
			num, _ := strconv.Atoi(n)
			line[i] = num
		}
	}
	return
}
