package day04

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Sleep struct{ From, To int }
type Guard []Sleep

func (guard Guard) Asleep() (minutes int) {
	for _, s := range guard {
		minutes += s.To - s.From
	}
	return
}

func (guard Guard) BestMinute() (best, max int) {
	minutes := [60]int{}
	for _, sleep := range guard {
		for i := sleep.From; i < sleep.To; i++ {
			minutes[i]++
		}
	}

	for minute, count := range minutes {
		if count > max {
			best, max = minute, count
		}
	}
	return
}

func Part1(input string) int {
	max, best, guards := 0, -1, parseInput(input)
	for id, guard := range guards {
		if asleep := guard.Asleep(); asleep > max {
			max, best = asleep, id
		}
	}
	minute, _ := guards[best].BestMinute()
	return best * minute
}

func Part2(input string) int {
	var best, max, minute int
	for id, guard := range parseInput(input) {
		min, count := guard.BestMinute()
		if count > max {
			best, max, minute = id, count, min
		}
	}
	return best * minute
}

func parseInput(input string) (guards map[int]Guard) {
	guards = map[int]Guard{}
	records := strings.Split(input, "\n")
	sort.Strings(records)

	var id int
	var sleep Sleep
	for _, record := range records {
		parts := strings.Split(record[15:], "] ")
		switch parts[1][:5] {
		case "Guard":
			fmt.Sscanf(parts[1], "Guard #%d begins shift", &id)
		case "falls":
			sleep = Sleep{From: utils.ToInt(parts[0])}
		case "wakes":
			sleep.To = utils.ToInt(parts[0])
			if _, ok := guards[id]; !ok {
				guards[id] = []Sleep{}
			}
			guards[id] = append(guards[id], sleep)
		}
	}
	return
}
