package day11

import "regexp"

var forbidden = regexp.MustCompile("[ilo]")

func straight(pwd string) bool {
	for i := len(pwd) - 1; i > 1; i-- {
		if pwd[i]-1 == pwd[i-1] && pwd[i]-2 == pwd[i-2] {
			return true
		}
	}
	return false
}

func pairs(pwd string) bool {
	for i, n, pair := 0, len(pwd)-1, ""; i < n; i++ {
		switch {
		case pwd[i] != pwd[i+1]:
			continue
		case pair == "":
			pair = pwd[i : i+2]
		case pair != pwd[i:i+2]:
			return true
		}

		i++
	}
	return false
}

func Valid(pwd string) bool {
	return !forbidden.MatchString(pwd) && straight(pwd) && pairs(pwd)
}

func Next(password string) (next string) {
	for i, d := len(password)-1, 1; i >= 0; i-- {
		p := int(password[i]-'a') + d
		d, next = p/26, string(rune(p%26+'a'))+next
	}
	return next
}

func Part1(input string) (pwd string) {
	pwd = Next(input)
	for !Valid(pwd) {
		pwd = Next(pwd)
	}
	return
}

func Part2(input string) (pwd string) {
	pwd = Next(Part1(input))
	for !Valid(pwd) {
		pwd = Next(pwd)
	}
	return
}
