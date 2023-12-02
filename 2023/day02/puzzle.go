package day02

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Hand struct {
	Blue, Green, Red int
}

func (h Hand) Possible() bool {
	return h.Blue <= 14 && h.Green <= 13 && h.Red <= 12
}

type Game struct {
	Id    int
	Hands []Hand
}

func (g Game) Possible() bool {
	for _, hand := range g.Hands {
		if !hand.Possible() {
			return false
		}
	}
	return true
}

func (g Game) Power() int {
	var blue, green, red int
	for _, hand := range g.Hands {
		blue = utils.Max(blue, hand.Blue)
		green = utils.Max(green, hand.Green)
		red = utils.Max(red, hand.Red)
	}
	return blue * green * red
}

func Part1(input string) (possible int) {
	for _, game := range parseInput(input) {
		if game.Possible() {
			possible += game.Id
		}
	}
	return
}

func Part2(input string) (sum int) {
	for _, game := range parseInput(input) {
		sum += game.Power()
	}
	return
}

func parseInput(input string) (games []Game) {
	var n int
	var t string

	for _, line := range strings.Split(input, "\n") {
		var game Game
		parts := strings.Split(line[5:], ": ")
		game.Id = utils.ToInt(parts[0])

		for _, h := range strings.Split(parts[1], "; ") {
			var hand Hand
			for _, cubes := range strings.Split(h, ", ") {
				fmt.Sscanf(cubes, "%d %s", &n, &t)
				switch t {
				case "blue":
					hand.Blue = n
				case "green":
					hand.Green = n
				case "red":
					hand.Red = n
				}
			}
			game.Hands = append(game.Hands, hand)
		}

		games = append(games, game)
	}
	return
}
