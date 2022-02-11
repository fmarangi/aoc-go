package day24

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type ALU [4]int
type ModelNumber [14]int
type Monad []Instruction
type Instruction struct {
	verb string
	a    int
	b    string
}

func toVar(val string) int {
	return int(val[0] - 'w')
}

func parseInput(input string) (i Monad) {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)
		switch parts[0] {
		case "inp":
			i = append(i, Instruction{verb: parts[0], a: toVar(parts[1])})
		default:
			i = append(i, Instruction{parts[0], toVar(parts[1]), parts[2]})
		}
	}
	return
}

func Process(input Monad, data ModelNumber) (alu ALU) {
	w := 0
	for _, i := range input {
		if i.verb == "inp" {
			alu[i.a] = data[w]
			w++
			continue
		}

		var b int
		switch i.b {
		case "w", "x", "y", "z":
			b = alu[toVar(i.b)]
		default:
			b = utils.ToInt(i.b)
		}

		switch i.verb {
		case "add":
			alu[i.a] += b
		case "mul":
			alu[i.a] *= b
		case "div":
			alu[i.a] /= b
		case "mod":
			alu[i.a] = alu[i.a] % b
		case "eql":
			if alu[i.a] == b {
				alu[i.a] = 1
			} else {
				alu[i.a] = 0
			}
		}
	}
	return
}

func initModelNumber(num int) (number ModelNumber) {
	for i := 0; i < 14; i++ {
		number[i] = num
	}
	return
}

func (number ModelNumber) toInt() (result int) {
	for _, n := range number {
		result = result*10 + n
	}
	return
}

func Part1(input string) (result int) {
	number := initModelNumber(9)
	number[4] = number[5] - 7
	number[7] = number[6] - 7
	number[8] = number[3] - 2
	number[10] = number[9] - 3
	number[11] = number[2] - 8
	number[1] = number[12] - 5
	number[0] = number[13]

	if Process(parseInput(input), number)[3] == 0 {
		return number.toInt()
	}
	return
}

func Part2(input string) (result int) {
	number := initModelNumber(1)
	number[5] = number[4] + 7
	number[6] = number[7] + 7
	number[3] = number[8] + 2
	number[9] = number[10] + 3
	number[2] = number[11] + 8
	number[12] = number[1] + 5
	number[0] = number[13]

	if Process(parseInput(input), number)[3] == 0 {
		return number.toInt()
	}
	return
}
