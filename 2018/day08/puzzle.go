package day08

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Node struct {
	Meta     []int
	Children []Node
}

func (node Node) Sum() (sum int) {
	for _, child := range node.Children {
		sum += child.Sum()
	}
	return sum + utils.Sum(node.Meta)
}

func (node Node) Value() (value int) {
	if len(node.Children) == 0 {
		return utils.Sum(node.Meta)
	}

	for _, m := range node.Meta {
		if m <= len(node.Children) {
			value += node.Children[m-1].Value()
		}
	}
	return
}

func Part1(input string) int {
	return parseInput(input).Sum()
}

func Part2(input string) int {
	return parseInput(input).Value()
}

func parseInput(input string) Node {
	data := utils.Ints(strings.Fields(input))
	node, _ := parseNode(data)
	return node
}

func parseNode(data []int) (Node, []int) {
	node := Node{}
	children, meta, data := data[0], data[1], data[2:]
	for i := 0; i < children; i++ {
		child, rest := parseNode(data)
		node.Children = append(node.Children, child)
		data = rest
	}
	node.Meta = data[:meta]
	return node, data[meta:]
}
