package day13

import "strings"

func Reflection(lines []string) int {
	var reflection bool
	for i, n := 0, len(lines); i < n-1; i++ {
		reflection = true
		for j := 0; i-j >= 0 && i+j+1 < n; j++ {
			if lines[i-j] != lines[i+j+1] {
				reflection = false
				break
			}
		}

		if reflection {
			return i
		}
	}
	return -1
}

func FindSmudge(lines []string) int {
	var values = make([]int, len(lines))
	for i, l := range lines {
		values[i] = Checkum(l)
	}

	for i, n := 0, len(values); i < n-1; i++ {
		var diffs []int
		for j := 0; i-j >= 0 && i+j+1 < n; j++ {
			if d := values[i-j] ^ values[i+j+1]; d != 0 {
				diffs = append(diffs, d)
			}
		}

		if len(diffs) == 1 && countBits(diffs[0]) == 1 {
			return i
		}
	}
	return -1
}

func Columns(pattern string) []string {
	return strings.Fields(pattern)
}

func Rows(pattern string) (rows []string) {
	width := strings.Index(pattern, "\n")
	for i := 0; i < width; i++ {
		var line []byte
		for j, m := 0, len(pattern); j < m; j += width + 1 {
			line = append(line, pattern[i+j])
		}
		rows = append(rows, string(line))
	}
	return
}

func Checkum(pattern string) (value int) {
	for i, c := range pattern {
		if c == '#' {
			value += 1 << i
		}
	}
	return
}

func countBits(num int) (n int) {
	for num > 0 {
		if num&1 == 1 {
			n++
		}
		num = num >> 1
	}
	return
}

func Part1(input string) (summary int) {
	for _, pattern := range strings.Split(input, "\n\n") {
		if r := Reflection(Columns(pattern)); r >= 0 {
			summary += 100 * (r + 1)
		} else if r := Reflection(Rows(pattern)); r >= 0 {
			summary += r + 1
		}
	}
	return
}

func Part2(input string) (summary int) {
	for _, pattern := range strings.Split(input, "\n\n") {
		if r := FindSmudge(Columns(pattern)); r >= 0 {
			summary += 100 * (r + 1)
		} else if r := FindSmudge(Rows(pattern)); r >= 0 {
			summary += r + 1
		}
	}
	return
}
