package day06

import "strings"

type Orbits map[string][]string

type route struct {
	obj   string
	steps int
}

func (o Orbits) Path(from, to string) (steps int) {
	var (
		curr  route
		seen  = map[string]bool{from: true}
		queue = []route{{from, 0}}
	)

	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		for _, next := range o[curr.obj] {
			switch {
			case next == to:
				return curr.steps + 1
			case !seen[next]:
				seen[next] = true
				queue = append(queue, route{next, curr.steps + 1})
			}
		}
	}
	return
}

func Part1(input string) (checksum int) {
	orbits := parseInput(input, false)
	for obj := range orbits {
		checksum += orbits.Path(obj, "COM")
	}
	return
}

func Part2(input string) int {
	return parseInput(input, true).Path("YOU", "SAN") - 2
}

func parseInput(input string, full bool) (orbits Orbits) {
	orbits = Orbits{}
	for _, row := range strings.Fields(input) {
		o := strings.Split(row, ")")
		orbits[o[1]] = append(orbits[o[1]], o[0])
		if full {
			orbits[o[0]] = append(orbits[o[0]], o[1])
		}
	}
	return
}
