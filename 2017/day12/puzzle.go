package day12

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Pipes [][]int

func (pipes Pipes) Group(pipe int) (group []int) {
	seen := map[int]bool{pipe: true}
	group = append(group, pipe)
	for queue := []int{pipe}; len(queue) > 0; queue = queue[1:] {
		for _, p := range pipes[queue[0]] {
			if !seen[p] {
				group = append(group, p)
				queue = append(queue, p)
				seen[p] = true
			}
		}
	}
	return
}

func Part1(input string) int {
	return len(parseInput(input).Group(0))
}

func Part2(input string) (groups int) {
	pipes := parseInput(input)
	seen := make([]bool, len(pipes))
	for i := range pipes {
		if seen[i] {
			continue
		}

		groups++
		for _, p := range pipes.Group(i) {
			seen[p] = true
		}
	}
	return
}

func parseInput(input string) (pipes Pipes) {
	for _, p := range strings.Split(input, "\n") {
		parts := strings.Split(p, " <-> ")
		pipes = append(pipes, utils.Ints(strings.Split(parts[1], ", ")))
	}
	return
}
