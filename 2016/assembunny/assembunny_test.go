package assembunny

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

var example string = `cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a`

func TestExecute(t *testing.T) {
	regs := make(Registers)
	regs.Execute(example)
	utils.Assert(t, 42, regs["a"])
}
