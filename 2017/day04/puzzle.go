package day04

import "strings"

type Letters [26]int

func hasDuplicates(passphrase string) bool {
	words := map[string]bool{}
	for _, w := range strings.Fields(passphrase) {
		if words[w] {
			return true
		}
		words[w] = true
	}
	return false
}

func wordLetters(word string) (letters Letters) {
	for _, c := range word {
		letters[int(c-'a')]++
	}
	return
}

func hasAnagrams(passphrase string) bool {
	words := map[Letters]bool{}
	for _, w := range strings.Fields(passphrase) {
		letters := wordLetters(w)
		if words[letters] {
			return true
		}
		words[letters] = true
	}
	return false
}

func Part1(input string) (valid int) {
	for _, passphrase := range parseInput(input) {
		if !hasDuplicates(passphrase) {
			valid++
		}
	}
	return
}

func Part2(input string) (valid int) {
	for _, passphrase := range parseInput(input) {
		if !hasAnagrams(passphrase) {
			valid++
		}
	}
	return
}

func parseInput(input string) []string {
	return strings.Split(input, "\n")
}
