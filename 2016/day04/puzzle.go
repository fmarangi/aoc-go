package day04

import (
	"regexp"
	"sort"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Room struct {
	Name, Checksum string
	SectorId       int
}

func (room Room) IsValid() bool {
	letters := map[rune]int{}
	for _, c := range room.Name {
		if c != '-' {
			letters[c]++
		}
	}

	pairs := []struct {
		letter rune
		count  int
	}{}
	for k, v := range letters {
		pairs = append(pairs, struct {
			letter rune
			count  int
		}{k, v})
	}

	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].count == pairs[j].count {
			return pairs[i].letter < pairs[j].letter
		}
		return pairs[i].count > pairs[j].count
	})

	var checkum strings.Builder
	for _, p := range pairs[:5] {
		checkum.WriteRune(p.letter)
	}

	return checkum.String() == room.Checksum
}

func (room Room) Decrypt() string {
	var builder strings.Builder
	for _, c := range room.Name {
		if c == '-' {
			builder.WriteString(" ")
			continue
		}
		builder.WriteRune(rune((int(c-'a')+room.SectorId)%26 + 'a'))
	}
	return builder.String()
}

func Part1(input string) (sum int) {
	for _, room := range parseInput(input) {
		if room.IsValid() {
			sum += room.SectorId
		}
	}
	return
}

func Part2(input string) int {
	for _, room := range parseInput(input) {
		if room.Decrypt() == "northpole object storage" {
			return room.SectorId
		}
	}
	return -1
}

func parseInput(input string) (rooms []Room) {
	var pattern = regexp.MustCompile(`^(.+?)-(\d+)\[([a-z]{5})\]$`)
	for _, row := range strings.Fields(input) {
		data := pattern.FindStringSubmatch(row)
		if len(data) == 4 {
			rooms = append(rooms, Room{data[1], data[3], utils.ToInt(data[2])})
		}
	}
	return
}
