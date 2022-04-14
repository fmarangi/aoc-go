package day20

import (
	"math"
	"regexp"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Coordinate struct{ X, Y, Z int }
type Particle struct{ P, V, A Coordinate }

func (a Coordinate) Distance(b Coordinate) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y) + utils.Abs(a.Z-b.Z)
}

func (p *Particle) Update() {
	p.V.X += p.A.X
	p.V.Y += p.A.Y
	p.V.Z += p.A.Z
	p.P.X += p.V.X
	p.P.Y += p.V.Y
	p.P.Z += p.V.Z
}

func Closest(particles []Particle, other Coordinate) (closest int) {
	distance := math.MaxInt32
	for i, p := range particles {
		if d := p.P.Distance(other); d < distance {
			distance, closest = d, i
		}
	}
	return
}

func RemoveCollisions(particles []Particle) (filtered []Particle) {
	positions := make(map[Coordinate]int)
	for _, p := range particles {
		positions[p.P]++
	}

	for _, p := range particles {
		if positions[p.P] == 1 {
			filtered = append(filtered, p)
		}
	}
	return
}

func Part1(input string) int {
	particles := parseInput(input)

	n := len(particles)
	for i := 0; i < 500*n; i++ {
		(&particles[i%n]).Update()
	}

	return Closest(particles, Coordinate{})
}

func Part2(input string) int {
	particles := parseInput(input)

	for i := 0; i < 50; i++ {
		for j, n := 0, len(particles); j < n; j++ {
			(&particles[j]).Update()
		}
		particles = RemoveCollisions(particles)
	}

	return len(particles)
}

func parseInput(input string) (particles []Particle) {
	pattern := regexp.MustCompile("<(.+?)>")
	for _, row := range strings.Split(input, "\n") {
		pav := pattern.FindAllStringSubmatch(row, 3)
		particles = append(particles, Particle{
			parseCoordinates(pav[0][1]),
			parseCoordinates(pav[1][1]),
			parseCoordinates(pav[2][1]),
		})
	}
	return
}

func parseCoordinates(input string) Coordinate {
	values := utils.Ints(strings.Split(input, ","))
	return Coordinate{values[0], values[1], values[2]}
}
