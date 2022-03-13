package day15

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func MemoryGame(init []int) chan int {
	out := make(chan int)
	go func() {
		spoken, i := map[int]int{}, 1
		for _, n := range init[:len(init)-1] {
			out <- n
			spoken[n] = i
			i++
		}

		last := init[len(init)-1]
		for {
			out <- last
			last, spoken[last] = spoken[last], i
			if last > 0 {
				last = i - last
			}
			i++
		}
	}()
	return out
}

func Part1(input string) (result int) {
	numbers := MemoryGame(parseInput(input))
	for i := 0; i < 2020; i++ {
		result = <-numbers
	}
	return
}

func Part2(input string) (result int) {
	numbers := MemoryGame(parseInput(input))
	for i := 0; i < 30000000; i++ {
		result = <-numbers
	}
	return
}

func parseInput(input string) []int {
	return utils.Ints(strings.Split(input, ","))
}
