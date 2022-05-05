package day14

import "github.com/fmarangi/aoc-go/utils"

func generator() <-chan int {
	out := make(chan int)
	go func() {
		first, second := 0, 1
		recipes := make([]int, 0)
		push := func(n int) {
			recipes = append(recipes, n)
			out <- n
		}

		push(3)
		push(7)
		for {
			recipe := recipes[first] + recipes[second]
			if recipe > 9 {
				push(recipe / 10)
			}
			push(recipe % 10)
			n := len(recipes)
			first, second = (first+recipes[first]+1)%n, (second+recipes[second]+1)%n
		}
	}()
	return out
}

func Scores(after int) (result int) {
	chocolate := generator()
	for i := 0; i < after; i++ {
		<-chocolate
	}

	for i := 0; i < 10; i++ {
		result = result*10 + <-chocolate
	}
	return
}

func SequenceAfter(scores int) (after int) {
	chocolate := generator()
	for sequence := 0; sequence != scores; after++ {
		sequence = (sequence*10 + <-chocolate) % 1000000
	}
	return after - 6
}

func Part1(input string) int {
	return Scores(utils.ToInt(input))
}

func Part2(input string) int {
	return SequenceAfter(utils.ToInt(input))
}
