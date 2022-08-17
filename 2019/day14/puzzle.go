package day14

import (
	"fmt"
	"strings"
)

func ceil(x, y int) int {
	return 1 + (x-1)/y
}

type Component struct {
	Element  string
	Quantity int
}

type Requirements []Component

type Reaction struct {
	Quantity     int
	Requirements Requirements
}

type Reactions map[string]Reaction

func (rx Reactions) OreAmount(fuel int) (ore int) {
	stock := make(map[string]int)
	for q := []Component{{"FUEL", fuel}}; len(q) > 0; q = q[1:] {
		req := q[0]
		if req.Quantity > stock[req.Element] {
			var (
				component = rx[req.Element]
				needed    = req.Quantity - stock[req.Element]
				times     = ceil(needed, component.Quantity)
			)

			stock[req.Element] += times * component.Quantity
			for _, c := range rx[req.Element].Requirements {
				if c.Element == "ORE" {
					ore += c.Quantity * times
				} else {
					q = append(q, Component{c.Element, c.Quantity * times})
				}
			}
		}

		stock[req.Element] -= req.Quantity
	}
	return
}

func Part1(input string) (fuel int) {
	return parseInput(input).OreAmount(1)
}

func Part2(input string) int {
	var (
		ore       = 1000000000000
		reactions = parseInput(input)
		min, max  = 1, ore
		mid       int
	)

	for max-min > 1 {
		mid = (max + min) / 2
		if reactions.OreAmount(mid) > ore {
			max = mid
		} else {
			min = mid
		}
	}

	return min
}

func parseInput(input string) (reactions Reactions) {
	reactions = make(Reactions)
	for _, row := range strings.Split(input, "\n") {
		var (
			parts    = strings.Split(row, " => ")
			element  = parseItem(parts[1])
			reaction = Reaction{Quantity: element.Quantity}
		)

		for _, el := range strings.Split(parts[0], ", ") {
			reaction.Requirements = append(reaction.Requirements, parseItem(el))
		}

		reactions[element.Element] = reaction
	}
	return
}

func parseItem(input string) (c Component) {
	fmt.Sscanf(input, "%d %s", &c.Quantity, &c.Element)
	return
}
