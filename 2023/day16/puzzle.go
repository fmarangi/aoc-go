package day16

import "strings"

const (
	Right int = iota
	Down
	Left
	Up
)

func Energize(from Beam, input string) int {
	var (
		m = len(input)
		w = strings.Index(input, "\n") + 1
		q = []Beam{from}
		s = map[Beam]bool{from: true}
		o = map[int]int{
			Right: 1,
			Down:  w,
			Left:  -1,
			Up:    -w,
		}
		curr Beam
	)

	for len(q) > 0 {
		curr, q = q[0], q[1:]
		for _, next := range curr.Move(input[curr.P]) {
			b := Beam{curr.P + o[next], next}
			if b.P >= 0 && b.P < m && input[b.P] != '\n' && !s[b] {
				q = append(q, b)
				s[b] = true
			}
		}
	}

	e := map[int]bool{}
	for c := range s {
		e[c.P] = true
	}

	return len(e)
}

func Part1(input string) int {
	return Energize(Beam{0, Right}, input)
}

func Part2(input string) (max int) {
	w, m := strings.Index(input, "\n"), len(input)

	for i := 0; i < w; i++ {
		if e := Energize(Beam{i, Down}, input); e > max {
			max = e
		}
	}

	for i := 0; i < m; i += w + 1 {
		if e := Energize(Beam{i, Right}, input); e > max {
			max = e
		}

		if e := Energize(Beam{i + w - 1, Left}, input); e > max {
			max = e
		}
	}

	for i := 0; i < w; i++ {
		if e := Energize(Beam{m - i - 1, Up}, input); e > max {
			max = e
		}
	}

	return
}
