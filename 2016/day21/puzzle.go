package day21

import (
	"fmt"
	"strings"
)

func position(haystack, letter string) int {
	for i, n := 0, len(haystack); i < n; i++ {
		if haystack[i] == letter[0] {
			return i
		}
	}
	return -1
}

func Swap(password string, x, y int) string {
	chars := []byte(password)
	chars[x], chars[y] = chars[y], chars[x]
	return string(chars)
}

func SwapLetter(password, x, y string) string {
	return Swap(password, position(password, x), position(password, y))
}

func Rotate(password string, steps int) string {
	n := len(password)
	steps = (n*2 - steps) % n
	return password[steps:] + password[:steps]
}

func RotateLetter(password, letter string) string {
	pos := position(password, letter)
	if pos >= 4 {
		pos += 1
	}
	return Rotate(password, pos+1)
}

func Reverse(password string, x, y int) string {
	chars := []byte(password)
	for i, j := x, y; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}
	return string(chars)
}

func Move(password string, x, y int) string {
	letter := password[x]
	password = password[:x] + password[x+1:]
	return password[:y] + string(letter) + password[y:]
}

func Part1(input string) (password string) {
	password = "abcdefgh"
	for _, op := range strings.Split(input, "\n") {
		switch {
		case strings.HasPrefix(op, "swap position"):
			var x, y int
			fmt.Sscanf(op, "swap position %d with position %d", &x, &y)
			password = Swap(password, x, y)

		case strings.HasPrefix(op, "swap letter"):
			var x, y string
			fmt.Sscanf(op, "swap letter %s with letter %s", &x, &y)
			password = SwapLetter(password, x, y)

		case strings.HasPrefix(op, "rotate based"):
			var x string
			fmt.Sscanf(op, "rotate based on position of letter %s", &x)
			password = RotateLetter(password, x)

		case strings.HasPrefix(op, "rotate left"):
			var x int
			fmt.Sscanf(op, "rotate left %d step", &x)
			password = Rotate(password, -x)

		case strings.HasPrefix(op, "rotate right"):
			var x int
			fmt.Sscanf(op, "rotate right %d step", &x)
			password = Rotate(password, x)

		case strings.HasPrefix(op, "reverse"):
			var x, y int
			fmt.Sscanf(op, "reverse positions %d through %d", &x, &y)
			password = Reverse(password, x, y)

		case strings.HasPrefix(op, "move position"):
			var x, y int
			fmt.Sscanf(op, "move position %d to position %d", &x, &y)
			password = Move(password, x, y)
		}
	}
	return
}

func Part2(input string) (password string) {
	var (
		ops     = strings.Split(input, "\n")
		offsets = []int{7, -1, 2, -2, 1, -3, 0, -4}
	)

	password = "fbgdceah"
	for i := len(ops) - 1; i >= 0; i-- {
		op := ops[i]
		switch {
		case strings.HasPrefix(op, "swap position"):
			var x, y int
			fmt.Sscanf(op, "swap position %d with position %d", &x, &y)
			password = Swap(password, x, y)

		case strings.HasPrefix(op, "swap letter"):
			var x, y string
			fmt.Sscanf(op, "swap letter %s with letter %s", &x, &y)
			password = SwapLetter(password, x, y)

		case strings.HasPrefix(op, "rotate based"):
			var x string
			fmt.Sscanf(op, "rotate based on position of letter %s", &x)
			password = Rotate(password, offsets[position(password, x)])

		case strings.HasPrefix(op, "rotate left"):
			var x int
			fmt.Sscanf(op, "rotate left %d step", &x)
			password = Rotate(password, x)

		case strings.HasPrefix(op, "rotate right"):
			var x int
			fmt.Sscanf(op, "rotate right %d step", &x)
			password = Rotate(password, -x)

		case strings.HasPrefix(op, "reverse"):
			var x, y int
			fmt.Sscanf(op, "reverse positions %d through %d", &x, &y)
			password = Reverse(password, x, y)

		case strings.HasPrefix(op, "move position"):
			var x, y int
			fmt.Sscanf(op, "move position %d to position %d", &x, &y)
			password = Move(password, y, x)
		}
	}
	return
}
