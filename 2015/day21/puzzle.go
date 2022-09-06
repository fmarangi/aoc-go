package day21

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func Part1(input string) (gold int) {
	gold = 99999
	boss := parseInput(input)
	for battle := range Battles() {
		if battle.Fighter().Wins(boss) {
			gold = min(gold, battle.Cost())
		}
	}
	return
}

func Part2(input string) (gold int) {
	boss := parseInput(input)
	for battle := range Battles() {
		if !battle.Fighter().Wins(boss) {
			gold = max(gold, battle.Cost())
		}
	}
	return
}

func parseInput(input string) (boss Fighter) {
	for _, row := range strings.Split(input, "\n") {
		var (
			parts = strings.Split(row, ": ")
			value = utils.ToInt(parts[1])
		)

		switch parts[0] {
		case "Hit Points":
			boss.HitPoints = value
		case "Damage":
			boss.Damage = value
		case "Armor":
			boss.Armor = value
		}
	}
	return
}
