package day04

import (
	"strconv"
	"strings"
)

const DRAWN = 1000

type Board struct {
	numbers [25]int
}

func (board *Board) markDrawn(drawn int) {
	for i, n := range board.numbers {
		if n == drawn {
			board.numbers[i] = DRAWN
			break
		}
	}
}

func (board Board) score() (score int) {
	for _, n := range board.numbers {
		if n != DRAWN {
			score += n
		}
	}
	return
}

func (board Board) hasWon() bool {
	for i := 0; i < 5; i++ {
		row, col := true, true
		for j := 0; j < 5 && (row || col); j++ {
			row = row && board.numbers[i*5+j] == DRAWN
			col = col && board.numbers[j*5+i] == DRAWN
		}
		if row || col {
			return true
		}
	}
	return false
}

func Part1(input string) int {
	drawn, boards := parseInput(input)
	for _, n := range drawn {
		for _, b := range boards {
			b.markDrawn(n)
			if b.hasWon() {
				return b.score() * n
			}
		}
	}
	return 0
}

func Part2(input string) int {
	drawn, boards := parseInput(input)
	for _, n := range drawn {
		nextRound := []*Board{}
		for _, b := range boards {
			b.markDrawn(n)
			switch {
			case !b.hasWon():
				nextRound = append(nextRound, b)
			case len(boards) == 1:
				return b.score() * n
			}
		}
		boards = nextRound
	}
	return 0
}

func parseInput(input string) (drawn []int, boards []*Board) {
	parts := strings.SplitN(input, "\n\n", 2)
	for _, n := range strings.Split(parts[0], ",") {
		num, _ := strconv.Atoi(n)
		drawn = append(drawn, num)
	}

	for _, b := range strings.Split(parts[1], "\n\n") {
		board := &Board{}
		boards = append(boards, board)
		for i, n := range strings.Fields(b) {
			num, _ := strconv.Atoi(n)
			board.numbers[i] = num
		}
	}

	return
}
