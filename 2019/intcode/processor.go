package intcode

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

const (
	ADD = iota + 1
	MULTIPLY
	EXIT = 99
)

type Program []int

func ParseProgram(program string) (intcode Program) {
	for _, address := range strings.Split(program, ",") {
		intcode = append(intcode, utils.ToInt(address))
	}
	return
}

func (p Program) Process() {
	pos := 0
	for p[pos] != EXIT {
		op, a, b, c := p[pos], p[pos+1], p[pos+2], p[pos+3]
		switch op {
		case ADD:
			p[c] = p[a] + p[b]
		case MULTIPLY:
			p[c] = p[a] * p[b]
		}
		pos += 4
	}
}
