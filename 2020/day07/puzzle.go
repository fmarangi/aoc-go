package day07

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Bags map[string]*Bag

type Bag struct {
	Contents  map[string]int
	Contained []string
}

func (bags Bags) Inside(bagType string) (count int) {
	for t, qty := range bags.Get(bagType).Contents {
		count += qty * (bags.Inside(t) + 1)
	}
	return
}

func (bags *Bags) Get(bagType string) *Bag {
	if _, ok := (*bags)[bagType]; !ok {
		(*bags)[bagType] = &Bag{Contents: make(map[string]int)}
	}
	return (*bags)[bagType]
}

func Part1(input string) int {
	bags := parseInput(input)
	shinyGold := bags.Get("shiny gold")
	queue := make([]string, len(shinyGold.Contained))
	copy(queue, shinyGold.Contained)

	colors := make(map[string]bool)
	for len(queue) > 0 {
		colors[queue[0]] = true
		queue = append(queue[1:], bags[queue[0]].Contained...)
	}
	return len(colors)
}

func Part2(input string) int {
	return parseInput(input).Inside("shiny gold")
}

func parseInput(input string) (bags Bags) {
	bags = make(Bags)
	for _, row := range strings.Split(input, "\n") {
		data := strings.Split(row, " bags contain ")
		if strings.HasPrefix(data[1], "no other") {
			continue
		}

		bag := bags.Get(data[0])
		for _, c := range strings.Split(data[1], ", ") {
			name, qty := parseContent(c)
			bag.Contents[name] = qty

			content := bags.Get(name)
			content.Contained = append(content.Contained, data[0])
		}
	}
	return
}

func parseContent(content string) (string, int) {
	end := strings.LastIndex(content, " ")
	parts := strings.SplitN(content[:end], " ", 2)
	return parts[1], utils.ToInt(parts[0])
}
