package day10

import (
	"fmt"
	"strings"
)

type Bot struct{ A, B int }

func (b Bot) HasBoth() bool {
	return b.A != 0 && b.B != 0
}

func (b *Bot) Give() (low, high int) {
	low, high = b.A, b.B
	if low > high {
		high, low = b.A, b.B
	}
	b.A, b.B = 0, 0
	return
}

func (b *Bot) Receive(value int) {
	switch {
	case b.A == 0:
		b.A = value
	case b.B == 0:
		b.B = value
	}
}

type Bots map[int]*Bot

func (bots Bots) Get(id int) *Bot {
	if _, ok := bots[id]; !ok {
		bots[id] = &Bot{}
	}
	return bots[id]
}

type Instruction struct {
	Bot, A, B    int
	WhatA, WhatB string
}

func Part1(input string) (target int) {
	output := map[int]int{}
	bots, instructions := parseInput(input)
	for i, n := 0, len(instructions); ; i++ {
		instruction := instructions[i%n]
		if bot := bots.Get(instruction.Bot); bot.HasBoth() {
			a, b := bot.Give()

			if a == 17 && b == 61 {
				target = instruction.Bot
				break
			}

			switch instruction.WhatA {
			case "output":
				output[instruction.A] = a
			default:
				bots.Get(instruction.A).Receive(a)
			}

			switch instruction.WhatB {
			case "output":
				output[instruction.B] = b
			default:
				bots.Get(instruction.B).Receive(b)
			}
		}
	}
	return
}

func Part2(input string) (target int) {
	output := map[int]int{}
	bots, instructions := parseInput(input)
	for i, n := 0, len(instructions); ; i++ {
		instruction := instructions[i%n]
		if bot := bots.Get(instruction.Bot); bot.HasBoth() {
			a, b := bot.Give()

			switch instruction.WhatA {
			case "output":
				output[instruction.A] = a
			default:
				bots.Get(instruction.A).Receive(a)
			}

			switch instruction.WhatB {
			case "output":
				output[instruction.B] = b
			default:
				bots.Get(instruction.B).Receive(b)
			}

			if output[0] != 0 && output[1] != 0 && output[2] != 0 {
				target = output[0] * output[1] * output[2]
				break
			}
		}
	}
	return
}

func parseInput(input string) (bots Bots, instructions []Instruction) {
	bots = make(Bots)
	for _, row := range strings.Split(input, "\n") {
		if strings.HasPrefix(row, "value") {
			var value, who int
			fmt.Sscanf(row, "value %d goes to bot %d", &value, &who)
			bots.Get(who).Receive(value)
			continue
		}

		i := Instruction{}
		fmt.Sscanf(row, "bot %d gives low to %s %d and high to %s %d", &i.Bot, &i.WhatA, &i.A, &i.WhatB, &i.B)
		instructions = append(instructions, i)
	}
	return
}
