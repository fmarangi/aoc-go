package day15

import (
	"fmt"
	"strings"
)

func Generator(init, factor int) chan int {
	out := make(chan int)
	go func() {
		for {
			init = (init * factor) % 2147483647
			out <- init & 65535
		}
	}()
	return out
}

func FilterMultiples(in chan int, multiples int) chan int {
	out := make(chan int)
	go func() {
		for {
			num := <-in
			if num%multiples == 0 {
				out <- num
			}
		}
	}()
	return out
}

func Part1(input string) (pairs int) {
	initA, initB := parseInput(input)
	a, b := Generator(initA, 16807), Generator(initB, 48271)
	for i := 0; i < 40_000_000; i++ {
		if <-a == <-b {
			pairs++
		}
	}
	return
}

func Part2(input string) (pairs int) {
	initA, initB := parseInput(input)
	a := FilterMultiples(Generator(initA, 16807), 4)
	b := FilterMultiples(Generator(initB, 48271), 8)
	for i := 0; i < 5_000_000; i++ {
		if <-a == <-b {
			pairs++
		}
	}
	return
}

func parseInput(input string) (a, b int) {
	row, gen := strings.Split(input, "\n"), ""
	fmt.Sscanf(row[0], "Generator %s starts with %d", &gen, &a)
	fmt.Sscanf(row[1], "Generator %s starts with %d", &gen, &b)
	return
}
