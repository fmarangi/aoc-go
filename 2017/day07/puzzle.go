package day07

import (
	"fmt"
	"strings"
)

type Tower map[string]Program
type Program struct {
	Weight   int
	Children []string
}

func (tower Tower) Bottom() (bottom string) {
	children := make(map[string]bool, len(tower)-1)
	for _, p := range tower {
		for _, c := range p.Children {
			children[c] = true
		}
	}

	for program := range tower {
		if !children[program] {
			bottom = program
			break
		}
	}

	return
}

func (tower Tower) Weight(program string) (weight int) {
	for _, child := range tower[program].Children {
		weight += tower.Weight(child)
	}
	return weight + tower[program].Weight
}

func (tower Tower) WrongWeight(program string) (name string, diff int) {
	weights, programs := map[int]int{}, map[int]string{}
	for _, child := range tower[program].Children {
		weight := tower.Weight(child)
		weights[weight]++
		programs[weight] = child
	}

	var wrong, others int
	for weight, count := range weights {
		if count == 1 {
			wrong, name = weight, programs[weight]
		} else {
			others = weight
		}
	}

	if wrong != 0 {
		diff = wrong - others
	}

	return
}

func Part1(input string) (root string) {
	return parseInput(input).Bottom()
}

func Part2(input string) (weight int) {
	tower := parseInput(input)
	for w, d := tower.Bottom(), -1; d != 0; w, d = tower.WrongWeight(w) {
		weight = tower[w].Weight - d
	}
	return
}

func parseInput(input string) (tower Tower) {
	rows, name, weight := strings.Split(input, "\n"), "", 0
	tower = make(Tower, len(rows))
	for _, r := range rows {
		parts := strings.Split(r, " -> ")
		switch len(parts) {
		case 1:
			fmt.Sscanf(parts[0], "%s (%d)", &name, &weight)
			tower[name] = Program{Weight: weight}
		case 2:
			fmt.Sscanf(parts[0], "%s (%d)", &name, &weight)
			tower[name] = Program{weight, strings.Split(parts[1], ", ")}
		}
	}
	return
}
