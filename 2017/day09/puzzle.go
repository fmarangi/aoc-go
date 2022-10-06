package day09

func Score(stream string) (score int) {
	var (
		level     = 0
		isGarbage = false
		curr      byte
	)

	for i, n := 0, len(stream); i < n; i++ {
		curr = stream[i]
		switch {
		case curr == '!':
			i++
		case curr == '<':
			isGarbage = true
		case curr == '>':
			isGarbage = false
		case isGarbage:
			continue
		case curr == '{':
			level++
		case curr == '}':
			score += level
			level--
		}
	}

	return
}

func Garbage(stream string) (garbage int) {
	var (
		isGarbage bool
		curr      byte
	)

	for i, n := 0, len(stream); i < n; i++ {
		curr = stream[i]
		switch {
		case curr == '!':
			i++
		case curr == '<' && !isGarbage:
			isGarbage = true
		case curr == '>':
			isGarbage = false
		case isGarbage:
			garbage++
		}
	}

	return
}

func Part1(input string) int {
	return Score(input)
}

func Part2(input string) int {
	return Garbage(input)
}
