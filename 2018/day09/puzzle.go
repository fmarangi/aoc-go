package day09

import (
	"container/ring"
	"fmt"

	"github.com/fmarangi/aoc-go/utils"
)

func NewMarble(value int) *ring.Ring {
	marble := ring.New(1)
	marble.Value = value
	return marble
}

func Game(rounds, players int) int {
	marbles, scores := NewMarble(0), make([]int, players)

	for i := 1; i <= rounds; i++ {
		if i%23 == 0 {
			marbles = marbles.Move(-8)
			scores[(i/23-1)%players] += i + marbles.Unlink(1).Value.(int)
			marbles = marbles.Next()
			continue
		}

		marble := NewMarble(i)
		marbles.Next().Link(marble)
		marbles = marble
	}

	return utils.Max(scores...)
}

func Part1(input string) int {
	rounds, players := parseInput(input)
	return Game(rounds, players)
}

func Part2(input string) int {
	rounds, players := parseInput(input)
	return Game(rounds*100, players)
}

func parseInput(input string) (rounds, players int) {
	format := "%d players; last marble is worth %d points"
	fmt.Sscanf(input, format, &players, &rounds)
	return
}
