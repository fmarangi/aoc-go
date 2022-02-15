package day02

import (
	"strings"
)

type Box string

func (box Box) frequencies() (freqs map[rune]int) {
	freqs = map[rune]int{}
	for _, c := range box {
		freqs[c]++
	}
	return
}

func (box Box) check() (two, three bool) {
	for _, c := range box.frequencies() {
		switch c {
		case 2:
			two = true
		case 3:
			three = true
		}
	}
	return
}

func Part1(input string) int {
	two, three := 0, 0
	for _, box := range strings.Fields(input) {
		c2, c3 := Box(box).check()
		if c2 {
			two++
		}
		if c3 {
			three++
		}
	}
	return two * three
}

func Part2(input string) string {
	boxes := strings.Fields(input)
	count := len(boxes)
	for i := 0; i < count; i++ {
		chars := len(boxes[i])
		for j := i; j < count; j++ {
			diff, same := 0, ""
			for c := 0; c < chars && diff < 2; c++ {
				if boxes[i][c] != boxes[j][c] {
					diff++
					continue
				}
				same = same + string(boxes[i][c])
			}

			if diff == 1 {
				return same
			}
		}
	}
	return ""
}
