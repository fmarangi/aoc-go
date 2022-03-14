package day21

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Game [2]Player
type Player struct {
	Score, Position int
}

func parseInput(input string) (game Game) {
	for _, p := range strings.Split(input, "\n") {
		var i, pos int
		fmt.Sscanf(p, "Player %d starting position: %d", &i, &pos)
		game[i-1] = Player{0, pos}
	}
	return
}

func rolls() map[int]int {
	return map[int]int{
		3: 1, 4: 3,
		5: 6, 6: 7, 7: 6,
		8: 3, 9: 1,
	}
}

func Part1(input string) int {
	game := parseInput(input)
	for i := 0; ; i += 3 {
		player := &game[i%2]
		player.Position = (player.Position + 3*(i+2)) % 10
		player.Score += player.Position
		if player.Position == 0 {
			player.Score += 10
		}

		if player.Score >= 1000 {
			i += 3
			return i * game[i%2].Score
		}
	}
}

func Part2(input string) int {
	scores := map[Game]int{parseInput(input): 1}
	wins := [2]int{}

	for player := 0; len(scores) > 0; player = (player + 1) % 2 {
		newScores := map[Game]int{}
		for game, count := range scores {
			for r, n := range rolls() {
				position := (game[player].Position + r) % 10
				score := game[player].Score + position
				if position == 0 {
					score += 10
				}

				if score >= 21 {
					wins[player] += count * n
					continue
				}

				g := game
				g[player] = Player{score, position}
				newScores[g] += count * n
			}
		}
		scores = newScores
	}

	return utils.Max(wins[:]...)
}
