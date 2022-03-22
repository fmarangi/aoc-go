package day09

import "regexp"

func RemoveGarbage(stream string) string {
	stream = regexp.MustCompile("!.").ReplaceAllString(stream, "")
	return regexp.MustCompile("<[^>]*?>").ReplaceAllString(stream, "")
}

func Score(stream string) (score int) {
	level := 0
	for _, c := range stream {
		switch c {
		case '{':
			level++
		case '}':
			score += level
			level--
		}
	}
	return
}

func Part1(input string) int {
	return Score(RemoveGarbage(input))
}

func Part2(input string) int {
	stream := regexp.MustCompile("!.").ReplaceAllString(input, "")
	garbage := regexp.MustCompile("<[^>]*?>")
	return len(stream) - len(garbage.ReplaceAllString(stream, "<>"))
}
