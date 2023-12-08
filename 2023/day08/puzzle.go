package day08

import (
	"fmt"
	"strings"
)

type Node struct {
	Left, Right string
}

type Network map[string]Node

func (network Network) ToZ(node, instructions string) (steps int) {
	for n := len(instructions); node[2] != 'Z'; steps++ {
		switch instructions[steps%n] {
		case 'L':
			node = network[node].Left
		case 'R':
			node = network[node].Right
		}
	}
	return
}

func Part1(input string) (steps int) {
	var (
		instructions, network = parseInput(input)
		current               = "AAA"
	)

	for n := len(instructions); current != "ZZZ"; steps++ {
		switch instructions[steps%n] {
		case 'L':
			current = network[current].Left
		case 'R':
			current = network[current].Right
		}
	}

	return
}

func Part2(input string) (result int) {
	var (
		instructions, network = parseInput(input)
		steps                 []int
	)

	for n := range network {
		if n[2] == 'A' {
			steps = append(steps, network.ToZ(n, instructions))
		}
	}

	result = steps[0]
	for _, n := range steps[1:] {
		running := result
		for running%n != 0 {
			running += result
		}
		result = running
	}

	return
}

func parseInput(input string) (string, Network) {
	var (
		k       string
		network = make(Network)
		parts   = strings.Split(input, "\n\n")
	)

	for _, row := range strings.Split(parts[1], "\n") {
		var n Node
		fmt.Sscanf(row, "%3s = (%3s, %3s)", &k, &n.Left, &n.Right)
		network[k] = n
	}

	return parts[0], network
}
