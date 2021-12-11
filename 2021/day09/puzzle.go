package day09

import (
	"sort"
	"strings"
)

func toInt(point byte) int {
	return int(point) - 48 // int('0')
}

func next(input string) func(int) []int {
	line, max := strings.Index(input, "\n")+1, len(input)
	return func(p int) (offsets []int) {
		for _, i := range [...]int{1, -1, line, -line} {
			offset := p + i
			if offset >= 0 && offset < max && input[offset] != '\n' {
				offsets = append(offsets, offset)
			}
		}
		return
	}
}

func Part1(input string) (risk int) {
	adjacient := next(input)
	lowPoints := []int{}
	for i, c := range input {
		if c == '\n' {
			continue
		}

		a := adjacient(i)
		low, val := len(a) > 0, toInt(byte(c))
		for _, n := range a {
			low = low && (val < toInt(input[n]))
		}

		if low {
			lowPoints = append(lowPoints, val)
		}
	}

	for _, l := range lowPoints {
		risk += l + 1
	}

	return
}

func Part2(input string) int {
	adjacient := next(input)
	basins := []int{}
	seen := map[int]bool{}
	for p, c := range input {
		_, ok := seen[p]
		if c == '\n' || c == '9' || ok {
			continue
		}

		basin := 0
		seen[p] = true
		for queue := []int{p}; len(queue) > 0; queue = queue[1:] {
			curr := queue[0]
			basin++
			for _, a := range adjacient(curr) {
				if _, visited := seen[a]; !visited && input[a] != '9' {
					queue = append(queue, a)
					seen[a] = true
				}
			}
		}

		basins = append(basins, basin)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(basins)))
	return basins[0] * basins[1] * basins[2]
}
