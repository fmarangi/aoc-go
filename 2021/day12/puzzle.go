package day12

import (
	"strings"
)

type Caves map[string][]string

type Path struct {
	visited map[string]int
	last    string
	twice   bool
}

func (p Path) CurrentStep() string {
	if p.last != "" {
		return p.last
	}
	return "start"
}

func (p Path) Has(cave string) bool {
	return strings.ToLower(cave) == cave && p.visited[cave] > 0
}

func (p Path) HasVisitedTwice(cave string) bool {
	switch cave {
	case strings.ToLower(cave):
		visited := p.visited[cave]
		if p.twice {
			return visited > 0
		}
		return visited > 1
	}
	return false
}

func (p Path) MoveTo(cave string) Path {
	next := Path{map[string]int{}, cave, p.twice}
	for k, v := range p.visited {
		next.visited[k] = v
	}
	if strings.ToLower(cave) == cave {
		next.visited[cave]++
		next.twice = next.twice || next.visited[cave] > 1
	}
	return next
}

func parseInput(input string) (caves Caves) {
	caves = Caves{}
	for _, line := range strings.Split(input, "\n") {
		c := strings.SplitN(line, "-", 2)
		caves[c[0]] = append(caves[c[0]], c[1])
		caves[c[1]] = append(caves[c[1]], c[0])
	}
	return
}

func Part1(input string) (found int) {
	caves := parseInput(input)
	for q := []Path{{}}; len(q) > 0; q = q[1:] {
		current := q[0]
		for _, c := range caves[current.CurrentStep()] {
			switch {
			case c == "start":
			case current.Has(c):
				continue
			case c == "end":
				found++
			default:
				q = append(q, current.MoveTo(c))
			}
		}
	}
	return
}

func Part2(input string) (found int) {
	caves := parseInput(input)
	for q := []Path{{}}; len(q) > 0; q = q[1:] {
		current := q[0]
		for _, c := range caves[current.CurrentStep()] {
			switch {
			case c == "start":
			case current.HasVisitedTwice(c):
				continue
			case c == "end":
				found++
			default:
				q = append(q, current.MoveTo(c))
			}
		}
	}
	return
}
