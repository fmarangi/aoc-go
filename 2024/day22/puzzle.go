package day22

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Sequence [4]int

func mix(a, b int) int  { return a ^ b }
func prune(num int) int { return num % 16777216 }

func Evolve(number int) int {
	result := number
	result = prune(mix(result*64, result))
	result = prune(mix(result/32, result))
	result = prune(mix(result*2048, result))
	return result
}

func priceChanges(curr int) map[Sequence]int {
	var (
		changes          = make(map[Sequence]int)
		next             int
		currSeq, nextSeq Sequence
	)

	for i := 0; i < 2000; i++ {
		next = Evolve(curr)
		nextSeq = Sequence{currSeq[1], currSeq[2], currSeq[3], next%10 - curr%10}
		if _, seen := changes[nextSeq]; !seen && i > 2 {
			changes[nextSeq] = next % 10
		}
		curr, currSeq = next, nextSeq
	}

	return changes
}

func Part1(input string) (sum int) {
	for _, n := range parseInput(input) {
		for i := 0; i < 2000; i++ {
			n = Evolve(n)
		}
		sum += n
	}
	return
}

func Part2(input string) (max int) {
	var changes = make(map[Sequence]int)
	for _, n := range parseInput(input) {
		for seq, bananas := range priceChanges(n) {
			changes[seq] += bananas
			if changes[seq] > max {
				max = changes[seq]
			}
		}
	}
	return
}

func parseInput(input string) []int {
	return utils.Ints(strings.Fields(input))
}
