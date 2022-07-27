package day14

import (
	"fmt"
	"strconv"
	"strings"
)

func toInt(num string) int {
	n, _ := strconv.ParseInt(num, 2, 64)
	return int(n)
}

type Mask string

func (m Mask) Apply(num int) int {
	one := toInt(strings.ReplaceAll(string(m), "X", "1"))
	zero := toInt(strings.ReplaceAll(string(m), "X", "0"))
	return (num | zero) & one
}

func (m Mask) Addresses(num int) (addresses []int) {
	var (
		one  = toInt(strings.ReplaceAll(string(m), "X", "1"))
		zero = strings.ReplaceAll(string(m), "0", "1")
	)

	for q := []string{zero}; len(q) > 0; q = q[1:] {
		if strings.Contains(q[0], "X") {
			q = append(q,
				strings.Replace(q[0], "X", "1", 1),
				strings.Replace(q[0], "X", "0", 1),
			)
			continue
		}
		addresses = append(addresses, (num|one)&toInt(q[0]))
	}

	return
}

type Program struct {
	Mask   Mask
	Values [][2]int
}

func Part1(input string) (sum int) {
	mem := make(map[int]int)
	for _, p := range parseInput(input) {
		for _, v := range p.Values {
			mem[v[0]] = p.Mask.Apply(v[1])
		}
	}

	for _, val := range mem {
		sum += val
	}

	return
}

func Part2(input string) (sum int) {
	mem := make(map[int]int)
	for _, p := range parseInput(input) {
		for _, v := range p.Values {
			for _, a := range p.Mask.Addresses(v[0]) {
				mem[a] = v[1]
			}
		}
	}

	for _, val := range mem {
		sum += val
	}

	return
}

func parseInput(input string) (programs []Program) {
	var program Program
	for _, row := range strings.Split(input, "\n") {
		switch {
		case strings.HasPrefix(row, "mem"):
			var value [2]int
			fmt.Sscanf(row, "mem[%d] = %d", &value[0], &value[1])
			program.Values = append(program.Values, value)
		case program.Mask != "":
			programs = append(programs, program)
			fallthrough
		default:
			program = Program{}
			program.Mask = Mask(row[7:])
		}
	}
	programs = append(programs, program)
	return
}
