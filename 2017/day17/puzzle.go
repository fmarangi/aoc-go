package day17

import "github.com/fmarangi/aoc-go/utils"

type Node struct {
	Value int
	Next  *Node
}

func (n *Node) Insert(other *Node) {
	n.Next, other.Next = other, n.Next
}

func Spinlock(nodes, steps int) int {
	spinlock := &Node{Value: 0}
	spinlock.Next = spinlock

	for i := 1; i <= nodes; i++ {
		for j := 0; j < steps%i; j++ {
			spinlock = spinlock.Next
		}
		spinlock.Insert(&Node{Value: i})
		spinlock = spinlock.Next
	}

	return spinlock.Next.Value
}

func Part1(input string) int {
	return Spinlock(2017, utils.ToInt(input))
}

func Part2(input string) (next int) {
	curr, inc := 0, utils.ToInt(input)
	for i := 1; i <= 50000000; i++ {
		curr = (curr + inc + 1) % i
		if curr == 0 {
			next = i
		}
	}
	return
}
