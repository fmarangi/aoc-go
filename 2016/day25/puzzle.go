package day25

import "github.com/fmarangi/aoc-go/2016/assembunny"

func Part1(input string) (a int) {
	for a = 0; ; a++ {
		regs := make(assembunny.Registers, 4)
		regs["a"] = a

		out := make(chan int)
		go regs.Output(input, out)

		if isClock(out) {
			break
		}
	}
	return
}

func isClock(in <-chan int) bool {
	for i, exp := 0, 0; i < 20; i++ {
		if <-in != exp {
			return false
		}
		exp = (exp + 1) % 2
	}
	return true
}
