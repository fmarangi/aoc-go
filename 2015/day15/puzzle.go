package day15

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Ingredient [5]int

func recipes() <-chan []int {
	out := make(chan []int)
	go func() {
		for a := 0; a <= 100; a++ {
			for b := 0; b <= 100-a; b++ {
				for c := 0; c <= 100-a-b; c++ {
					out <- []int{a, b, c, 100 - a - b - c}
				}
			}
		}
		close(out)
	}()
	return out
}

func Score(ingredients []Ingredient, qtys []int) int {
	score := 1
	for i := 0; i < 4; i++ {
		var sum int
		for j := range ingredients {
			sum += ingredients[j][i] * qtys[j]
		}

		if sum <= 0 {
			return 0
		}

		score *= sum
	}
	return score
}

func Calories(ingredients []Ingredient, qtys []int) (calories int) {
	for i := range ingredients {
		calories += ingredients[i][4] * qtys[i]
	}
	return
}

func Part1(input string) (max int) {
	ingredients := parseInput(input)
	for qtys := range recipes() {
		max = utils.Max(max, Score(ingredients, qtys))
	}
	return
}

func Part2(input string) (max int) {
	ingredients := parseInput(input)
	for qtys := range recipes() {
		if Calories(ingredients, qtys) == 500 {
			max = utils.Max(max, Score(ingredients, qtys))
		}
	}
	return
}

func parseInput(input string) (ingredients []Ingredient) {
	for _, row := range strings.Split(input, "\n") {
		data := strings.Split(strings.Split(row, ": ")[1], ",")
		var ingredient Ingredient
		for i, value := range data {
			ingredient[i] = utils.ToInt(strings.Fields(value)[1])
		}
		ingredients = append(ingredients, ingredient)
	}
	return
}
