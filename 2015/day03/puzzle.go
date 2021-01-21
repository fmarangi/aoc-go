package day03

var move = map[rune]int{
	'^': -1000,
	'>': +1,
	'v': +1000,
	'<': -1,
}

func Part1(input string) int {
	curr := 0
	seen := map[int]bool{curr: true}
	for _, p := range input {
		curr += move[p]
		seen[curr] = true
	}
	return len(seen)
}

func Part2(input string) int {
	curr := map[int]int{0: 0, 1: 0}
	seen := map[int]bool{0: true}
	for i, p := range input {
		who := i % 2
		curr[who] += move[p]
		seen[curr[who]] = true
	}
	return len(seen)
}
