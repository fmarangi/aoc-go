package day17

import (
	"crypto/md5"
	"fmt"
)

var (
	doors   = [4]string{"U", "D", "L", "R"}
	offsets = [4]room{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}
	target  = room{3, 3}
)

type room struct{ x, y int }

type Path struct {
	Steps string
	Pos   room
}

func (p Path) Next(passcode string) (next []Path) {
	hash := md5.Sum([]byte(passcode + p.Steps))
	for i, d := range fmt.Sprintf("%x", hash)[:4] {
		if d < 'b' || d > 'f' {
			continue
		}

		pos := room{p.Pos.x + offsets[i].x, p.Pos.y + offsets[i].y}
		if pos.x >= 0 && pos.y >= 0 && pos.x < 4 && pos.y < 4 {
			next = append(next, Path{p.Steps + doors[i], pos})
		}
	}
	return
}

func shortest(passcode string) string {
	for q := []Path{{}}; len(q) > 0; q = q[1:] {
		for _, path := range q[0].Next(passcode) {
			if path.Pos == target {
				return path.Steps
			}
			q = append(q, path)
		}
	}
	return ""
}

func longest(passcode string) (steps int) {
	for q := []Path{{}}; len(q) > 0; q = q[1:] {
		for _, path := range q[0].Next(passcode) {
			if path.Pos == target {
				steps = len(path.Steps)
				continue
			}
			q = append(q, path)
		}
	}
	return
}

func Part1(input string) string {
	return shortest(input)
}

func Part2(input string) (time int) {
	return longest(input)
}
