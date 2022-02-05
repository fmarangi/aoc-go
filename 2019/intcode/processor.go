package intcode

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

const (
	ADD = iota + 1
	MULTIPLY
	INPUT
	OUTPUT
	JUMP_IF_TRUE
	JUMP_IF_FALSE
	LESS_THAN
	EQUALS
	EXIT = 99
)

const (
	MODE_POSITION = iota
	MODE_IMMEDIATE
	MODE_RELATIVE
)

type Program []int

func ParseProgram(program string) (intcode Program) {
	for _, address := range strings.Split(program, ",") {
		intcode = append(intcode, utils.ToInt(address))
	}
	return
}

func (p Program) value(at, mode int) int {
	if mode == MODE_POSITION {
		at = p[at]
	}
	return p[at]
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

func (p Program) Run(in chan int) chan int {
	out := make(chan int)

	go func() {
		pos := 0
		for {
			op := p[pos]
			switch op % 100 {
			case EXIT:
				return

			case INPUT:
				p[p[pos+1]] = <-in
				pos += 2

			case OUTPUT:
				out <- p.value(pos+1, op/100)
				pos += 2

			case ADD:
				p[p[pos+3]] = p.value(pos+1, (op/100)%10) + p.value(pos+2, (op/1000)%10)
				pos += 4

			case MULTIPLY:
				p[p[pos+3]] = p.value(pos+1, (op/100)%10) * p.value(pos+2, (op/1000)%10)
				pos += 4

			case JUMP_IF_TRUE:
				if p.value(pos+1, (op/100)%10) != 0 {
					pos = p.value(pos+2, (op/1000)%10)
					continue
				}
				pos += 3

			case JUMP_IF_FALSE:
				if p.value(pos+1, (op/100)%10) == 0 {
					pos = p.value(pos+2, (op/1000)%10)
					continue
				}
				pos += 3

			case LESS_THAN:
				p[p[pos+3]] = 0
				if p.value(pos+1, (op/100)%10) < p.value(pos+2, (op/1000)%10) {
					p[p[pos+3]] = 1
				}
				pos += 4

			case EQUALS:
				p[p[pos+3]] = 0
				if p.value(pos+1, (op/100)%10) == p.value(pos+2, (op/1000)%10) {
					p[p[pos+3]] = 1
				}
				pos += 4

			default:
				panic(fmt.Sprintf("Unknown operation: %d", op))
			}
		}
	}()

	return out
}
