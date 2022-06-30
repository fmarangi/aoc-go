package day09

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type segment struct{ from, to string }

type Distances map[segment]int

func (d Distances) AllLocations() (loc []string) {
	set := make(map[string]bool)
	for k := range d {
		set[k.from], set[k.to] = true, true
	}
	for k := range set {
		loc = append(loc, k)
	}
	return
}

func (d Distances) CalcDistance(route []string) (distance int) {
	for i, n := 1, len(route); i < n; i++ {
		distance += d[segment{route[i-1], route[i]}]
	}
	return
}

func Part1(input string) int {
	var (
		distances = parseInput(input)
		locations = distances.AllLocations()
		routes    = []int{}
	)
	for _, route := range utils.Permutations(locations) {
		routes = append(routes, distances.CalcDistance(route))
	}
	return utils.Min(routes...)
}

func Part2(input string) int {
	var (
		distances = parseInput(input)
		locations = distances.AllLocations()
		routes    = []int{}
	)
	for _, route := range utils.Permutations(locations) {
		routes = append(routes, distances.CalcDistance(route))
	}
	return utils.Max(routes...)
}

func parseInput(input string) Distances {
	var (
		distances = make(Distances)
		from, to  string
		distance  int
	)

	for _, row := range strings.Split(input, "\n") {
		fmt.Sscanf(row, "%s to %s = %d", &from, &to, &distance)
		distances[segment{from, to}] = distance
		distances[segment{to, from}] = distance
	}

	return distances
}
