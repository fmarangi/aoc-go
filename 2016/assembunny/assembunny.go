package assembunny

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Registers map[string]int
type instruction []string

func (regs Registers) value(val string) int {
	switch val {
	case "a", "b", "c", "d":
		return regs[val]
	}
	return utils.ToInt(val)
}

func (regs Registers) Execute(input string) {
	instructions := parse(input)
	for i, n := 0, len(instructions); i < n; i++ {
		c := instructions[i]
		switch c[0] {
		case "cpy":
			regs[c[2]] = regs.value(c[1])
		case "inc":
			regs[c[1]]++
		case "dec":
			regs[c[1]]--
		case "jnz":
			if regs.value(c[1]) != 0 {
				i += regs.value(c[2]) - 1
			}
		}
	}
}

func parse(code string) (instructions []instruction) {
	for _, row := range strings.Split(code, "\n") {
		instructions = append(instructions, strings.Fields(row))
	}
	return
}
