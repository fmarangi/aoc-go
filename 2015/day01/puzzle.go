package day01

var move = map[rune]int{
	'(': +1,
	')': -1,
}

func Part1(input string) int {
	at := 0
	for _, c := range input {
		at += move[c]
	}
	return at
}

func Part2(input string) int {
	at := 0
	for p, c := range input {
		at += move[c]
		if at < 0 {
			return p + 1
		}
	}
	return -1
}
