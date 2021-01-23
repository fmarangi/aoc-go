package day03

import (
	"sort"
	"strconv"
	"strings"
)

func Part1(input string) (triangles int) {
	for _, sides := range parseInput(input) {
		if IsTriangle(sides) {
			triangles += 1
		}
	}
	return triangles
}

func Part2(input string) (triangles int) {
	data := parseInput(input)
	for i, n := 0, len(data); i < n; i += 3 {
		for j := 0; j < 3; j++ {
			if IsTriangle([]int{data[i][j], data[i+1][j], data[i+2][j]}) {
				triangles += 1
			}
		}
	}
	return triangles
}

func parseInput(input string) (triangles [][]int) {
	for _, line := range strings.Split(input, "\n") {
		triangle := make([]int, 3)
		for i, side := range strings.Fields(line) {
			triangle[i], _ = strconv.Atoi(side)
		}
		triangles = append(triangles, triangle)
	}
	return triangles
}

func IsTriangle(sides []int) bool {
	sort.Ints(sides)
	return (sides[0] + sides[1]) > sides[2]
}
