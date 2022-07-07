package day24

import "github.com/fmarangi/aoc-go/utils"

func Part1(input string) (min int) {
	var (
		maze      = Maze(input)
		distances = maze.Distances()
		poi       []string
	)

	for p := range maze.POI() {
		if p != "0" {
			poi = append(poi, p)
		}
	}

	min = 999999
	for _, route := range utils.Permutations(poi) {
		route = append(route, "0")
		if l := distances.Length(route); l < min {
			min = l
		}
	}

	return
}

func Part2(input string) (min int) {
	var (
		maze      = Maze(input)
		distances = maze.Distances()
		poi       []string
	)

	for p := range maze.POI() {
		if p != "0" {
			poi = append(poi, p)
		}
	}

	min = 999999
	for _, route := range utils.Permutations(poi) {
		route = append([]string{"0"}, route...)
		route = append(route, "0")
		if l := distances.Length(route); l < min {
			min = l
		}
	}

	return
}
