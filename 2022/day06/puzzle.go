package day06

func marker(input string) bool {
	els := make(map[rune]bool)
	for _, e := range input {
		if els[e] {
			return false
		}
		els[e] = true
	}
	return true
}

func findMarker(input string, length int) int {
	for i, n := 0, len(input)-length; i < n; i++ {
		if marker(input[i : i+length]) {
			return i + length
		}
	}
	return -1
}

func Part1(input string) int {
	return findMarker(input, 4)
}

func Part2(input string) int {
	return findMarker(input, 14)
}
