package day18

import "strings"

const (
	OpenGround = '.'
	Tree       = '|'
	Lumberyard = '#'
)

func Change(state string) string {
	var (
		width  = strings.Index(state, "\n") + 1
		length = len(state)
		next   strings.Builder
		x      int
	)

	for i, acre := range state {
		if acre == '\n' {
			next.WriteRune(acre)
			continue
		}

		acres := make(map[byte]int, 3)
		for j := 0; j < 9; j++ {
			x = i + (j/3-1)*width + j%3 - 1
			if x != i && x >= 0 && x < length && state[x] != '\n' {
				acres[state[x]]++
			}
		}

		switch {
		case acre == OpenGround && acres[Tree] >= 3:
			next.WriteRune(Tree)
		case acre == Tree && acres[Lumberyard] >= 3:
			next.WriteRune(Lumberyard)
		case acre == Lumberyard && (acres[Lumberyard] < 1 || acres[Tree] < 1):
			next.WriteRune(OpenGround)
		default:
			next.WriteRune(acre)
		}
	}

	return next.String()
}

func Resources(state string) int {
	acres := make(map[rune]int, 2)
	for _, acre := range state {
		if acre == Tree || acre == Lumberyard {
			acres[acre]++
		}
	}
	return acres[Tree] * acres[Lumberyard]
}

func Part1(input string) int {
	for i := 0; i < 10; i++ {
		input = Change(input)
	}
	return Resources(input)
}

func Part2(input string) int {
	var (
		minutes     = 1000000000
		states      = map[string]int{input: 0}
		minutesLeft int
	)

	for i := 1; minutesLeft == 0; i++ {
		input = Change(input)
		if at, seen := states[input]; seen {
			// Loop detected, calculate remainder
			minutesLeft = (minutes - at) % (i - at)
		}
		states[input] = i
	}

	// Run missing minutes
	for i := 0; i < minutesLeft; i++ {
		input = Change(input)
	}

	return Resources(input)
}
