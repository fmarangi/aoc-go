package day22

import (
	"fmt"
	"strings"
)

type Step struct {
	Cuboid Cuboid
	On     bool
}

func parseInput(input string) (steps []Step) {
	var on string
	var x1, x2, y1, y2, z1, z2 int
	for _, row := range strings.Split(input, "\n") {
		fmt.Sscanf(row, "%s x=%d..%d,y=%d..%d,z=%d..%d", &on, &x1, &x2, &y1, &y2, &z1, &z2)
		cuboid := Cuboid{x1, x2, y1, y2, z1, z2}
		steps = append(steps, Step{cuboid, on == "on"})
	}
	return
}

func TotalVolume(steps []Step) (volume int) {
	var added []Step
	for _, step := range steps {
		on, cuboid := step.On, step.Cuboid

		var intersection []Step
		for _, s := range added {
			if i, err := cuboid.Intersect(s.Cuboid); err == nil {
				if on {
					intersection = append(intersection, Step{i, on != s.On})
				} else {
					intersection = append(intersection, Step{i, on == s.On})
				}
			}
		}

		if on {
			added = append(added, step)
		}
		added = append(added, intersection...)
	}

	for _, a := range added {
		if a.On {
			volume += a.Cuboid.Volume()
		} else {
			volume -= a.Cuboid.Volume()
		}
	}

	return
}

func Part1(input string) int {
	var steps []Step
	for _, s := range parseInput(input) {
		if s.Cuboid.InRange() {
			steps = append(steps, s)
		}
	}
	return TotalVolume(steps)
}

func Part2(input string) int {
	return TotalVolume(parseInput(input))
}
