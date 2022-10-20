package day24

import (
	"fmt"
	"strings"
)

type Component [2]int
type Bridge []Component
type mapping map[int]map[Component]bool

func (c Component) Value() int {
	if c[0] < c[1] {
		return c[0]*100 + c[1]
	}
	return c[1]*100 + c[0]
}

func (c Component) Flip() Component {
	return Component{c[1], c[0]}
}

func (b Bridge) Has(c Component) bool {
	for _, el := range b {
		if c.Value() == el.Value() {
			return true
		}
	}
	return false
}

func (b Bridge) Tail() int {
	if n := len(b); n > 0 {
		return b[len(b)-1][1]
	}
	return 0
}

func (b Bridge) Strength() (s int) {
	for _, c := range b {
		s += c[0] + c[1]
	}
	return
}

func mappedComponents(components []Component) (m mapping) {
	m = make(mapping)
	for _, c := range components {
		if _, ok := m[c[0]]; !ok {
			m[c[0]] = make(map[Component]bool)
		}
		if _, ok := m[c[1]]; !ok {
			m[c[1]] = make(map[Component]bool)
		}
		m[c[0]][c] = true
		m[c[1]][c] = true
	}
	return
}

func Part1(input string) (strength int) {
	var (
		components = parseInput(input)
		mapped     = mappedComponents(components)
		queue      = []Bridge{{}}
		curr, b    Bridge
	)

	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		for c := range mapped[curr.Tail()] {
			if curr.Has(c) {
				continue
			}

			b = make(Bridge, len(curr), len(curr)+1)
			copy(b, curr)
			if c[0] == curr.Tail() {
				b = append(b, c)
			} else {
				b = append(b, c.Flip())
			}

			queue = append(queue, b)
		}

		if s := curr.Strength(); s > strength {
			strength = s
		}
	}

	return
}

func Part2(input string) (strength int) {
	var (
		components = parseInput(input)
		mapped     = mappedComponents(components)
		queue      = []Bridge{{}}
		curr, b    Bridge
		longest, n int
	)

	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		for c := range mapped[curr.Tail()] {
			if curr.Has(c) {
				continue
			}

			b = make(Bridge, len(curr), len(curr)+1)
			copy(b, curr)
			if c[0] == curr.Tail() {
				b = append(b, c)
			} else {
				b = append(b, c.Flip())
			}

			queue = append(queue, b)
		}

		n = len(curr)
		switch {
		case n > longest:
			longest, strength = n, curr.Strength()
		case n == longest:
			if s := curr.Strength(); s > strength {
				strength = s
			}
		}
	}

	return
}

func parseInput(input string) (components []Component) {
	for _, row := range strings.Fields(input) {
		var k Component
		fmt.Sscanf(row, "%d/%d", &k[0], &k[1])
		components = append(components, k)
	}
	return
}
