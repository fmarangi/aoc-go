package day05

import "strings"

func opposite(a, b byte) bool {
	return a+32 == b || a-32 == b
}

func filter(polymer string, a, b rune) string {
	var builder strings.Builder
	for _, c := range polymer {
		if c != a && c != b {
			builder.WriteRune(c)
		}
	}
	return builder.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Reduce(polymer string) string {
	var builder strings.Builder

	for n, f := len(polymer), true; f; polymer, n = builder.String(), builder.Len() {
		builder.Reset()
		builder.Grow(n)

		f = false
		for i := 0; i < n; i++ {
			if i+1 != n && opposite(polymer[i], polymer[i+1]) {
				i, f = i+1, true
				continue
			}
			builder.WriteByte(polymer[i])
		}
	}

	return polymer
}

func Part1(input string) int {
	return len(Reduce(input))
}

func Part2(input string) (shortest int) {
	shortest = len(input)
	for i := 0; i < 26; i++ {
		reduced := Reduce(filter(input, rune('a'+i), rune('A'+i)))
		shortest = min(shortest, len(reduced))
	}
	return
}
