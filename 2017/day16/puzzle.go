package day16

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

const Initial = "abcdefghijklmnop"

func Spin(programs string, n int) string {
	o := len(programs) - n
	return programs[o:] + programs[:o]
}

func Exchange(programs string, a, b int) string {
	c := []byte(programs)
	c[a], c[b] = c[b], c[a]
	return string(c)
}

func Partner(programs string, a, b byte) string {
	x, y := strings.IndexByte(programs, a), strings.IndexByte(programs, b)
	return Exchange(programs, x, y)
}

func Apply(programs string, moves []string) string {
	for _, move := range moves {
		switch move[0] {
		case 's':
			programs = Spin(programs, utils.ToInt(move[1:]))
		case 'x':
			ab := strings.Split(move[1:], "/")
			programs = Exchange(programs, utils.ToInt(ab[0]), utils.ToInt(ab[1]))
		case 'p':
			programs = Partner(programs, move[1], move[3])
		}
	}
	return programs
}

func Part1(input string) string {
	return Apply(Initial, strings.Split(input, ","))
}

func Part2(input string) string {
	var (
		moves    = strings.Split(input, ",")
		programs = Apply(Initial, moves)
		cycle    = 1
	)

	for programs != Initial {
		programs = Apply(programs, moves)
		cycle++
	}

	programs = Initial
	for i, n := 0, 1000000000%cycle; i < n; i++ {
		programs = Apply(programs, moves)
	}

	return programs
}
