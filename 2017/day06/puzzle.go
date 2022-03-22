package day06

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type MemoryBanks [16]int

func (banks MemoryBanks) Most() (pos, blocks int) {
	for i, b := range banks {
		if b > blocks {
			pos, blocks = i, b
		}
	}
	return
}

func Part1(input string) (cycles int) {
	banks := parseInput(input)
	seen := map[MemoryBanks]bool{}

	for ; !seen[banks]; cycles++ {
		seen[banks] = true

		most, blocks := banks.Most()
		banks[most] = 0
		for pos := most + 1; blocks > 0; blocks, pos = blocks-1, pos+1 {
			banks[pos%16]++
		}
	}

	return
}

func Part2(input string) (cycles int) {
	banks := parseInput(input)
	seen := map[MemoryBanks]int{}

	for ; seen[banks] == 0; cycles++ {
		seen[banks] = cycles

		most, blocks := banks.Most()
		banks[most] = 0
		for pos := most + 1; blocks > 0; blocks, pos = blocks-1, pos+1 {
			banks[pos%16]++
		}
	}

	return cycles - seen[banks]
}

func parseInput(input string) (banks MemoryBanks) {
	for i, b := range utils.Ints(strings.Fields(input)) {
		banks[i] = b
	}
	return
}
