package day03

import "strings"

type Contents map[rune]bool

func content(comparment string) Contents {
	items := make(Contents)
	for _, i := range comparment {
		items[i] = true
	}
	return items
}

func (a Contents) intersect(b Contents) Contents {
	var c = make(Contents)
	for i := range a {
		if b[i] {
			c[i] = true
		}
	}
	return c
}

func priority(item rune) int {
	if item > 'Z' {
		return int(item-'a') + 1
	}
	return int(item-'A') + 27
}

func Part1(input string) (sum int) {
	for _, rucksack := range strings.Fields(input) {
		n := len(rucksack) / 2
		first := content(rucksack[:n])
		for _, i := range rucksack[n:] {
			if first[i] {
				sum += priority(i)
				break
			}
		}
	}
	return
}

func Part2(input string) (sum int) {
	rows := strings.Fields(input)
	for i, q := 0, len(rows); i < q; i += 3 {
		first := content(rows[i])
		candidates := first.intersect(content(rows[i+1]))
		for _, i := range rows[i+2] {
			if candidates[i] {
				sum += priority(i)
				break
			}
		}
	}
	return
}
