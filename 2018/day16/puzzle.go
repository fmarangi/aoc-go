package day16

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Sample struct{ before, after, values [4]int }

func matchingOpcodes(sample Sample) (matching int) {
	var test [4]int
	for _, op := range opcodes {
		copy(test[:], sample.before[:])
		op(&test, sample.values)
		if test == sample.after {
			matching++
		}
	}
	return
}

func Part1(input string) (matching int) {
	samples, _ := parseInput(input)
	for _, sample := range samples {
		if matchingOpcodes(sample) >= 3 {
			matching++
		}
	}
	return
}

func Part2(input string) int {
	var (
		samples, instructions = parseInput(input)
		numbers               = OpcodeNumbers(samples)
		registers             [4]int
	)

	for _, i := range instructions {
		opcodes[numbers[i[0]]](&registers, i)
	}

	return registers[0]
}

func parseInput(input string) (samples []Sample, codes [][4]int) {
	parts := strings.Split(input, "\n\n\n\n")
	for _, s := range strings.Split(parts[0], "\n\n") {
		samples = append(samples, parseSample(s))
	}

	for _, c := range strings.Split(parts[1], "\n") {
		var code [4]int
		copy(code[:], utils.Ints(strings.Fields(c)))
		codes = append(codes, code)
	}
	return
}

func parseSample(input string) (sameple Sample) {
	parts := strings.Split(input, "\n")
	copy(sameple.before[:], utils.Ints(strings.Split(parts[0][9:19], ", ")))
	copy(sameple.values[:], utils.Ints(strings.Fields(parts[1])))
	copy(sameple.after[:], utils.Ints(strings.Split(parts[2][9:19], ", ")))
	return
}
