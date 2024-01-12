package day23

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

const (
	N uint8 = 1 << iota
	NE
	E
	SE
	S
	SW
	W
	NW
)

var offsets = map[uint8]Position{
	N:  {0, -1},
	NE: {1, -1},
	E:  {1, 0},
	SE: {1, 1},
	S:  {0, 1},
	SW: {-1, 1},
	W:  {-1, 0},
	NW: {-1, -1},
}

var constraints = map[uint8]uint8{
	N: NW | N | NE,
	E: NE | E | SE,
	S: SW | S | SE,
	W: NW | W | SW,
}

type Grove map[Position]bool

func (grove Grove) Neighbours(pos Position) (n uint8) {
	for o, p := range offsets {
		if grove[pos.Move(p)] {
			n |= o
		}
	}
	return
}

func (grove Grove) GetMove(elf Position, round int) Position {
	var (
		neighbours = grove.Neighbours(elf)
		sequence   = [4]uint8{N, S, W, E}
	)

	for i := 0; neighbours != 0 && i < 4; i++ {
		var direction = sequence[(i+round)%4]
		if neighbours&constraints[direction] == 0 {
			return elf.Move(offsets[direction])
		}
	}

	return elf
}

func (grove Grove) Next(round int) (Grove, bool) {
	var (
		next  = make(Grove)
		moves = make(map[Position][]Position)
		move  Position
		moved = false
	)

	for elf := range grove {
		move = grove.GetMove(elf, round)
		moves[move] = append(moves[move], elf)
	}

	for move, elves := range moves {
		if len(elves) == 1 {
			next[move] = true
			if move != elves[0] {
				moved = true
			}
		} else {
			for _, elf := range elves {
				next[elf] = true
			}
		}
	}

	return next, moved
}

func (grove Grove) Edges() (Position, Position) {
	var tl, br = Position{99999, 99999}, Position{-99999, -99999}
	for elf := range grove {
		tl.X, tl.Y = min(tl.X, elf.X), min(tl.Y, elf.Y)
		br.X, br.Y = max(br.X, elf.X), max(br.Y, elf.Y)
	}
	return tl, br
}

func (grove Grove) EmptyTiles() int {
	var (
		tl, br = grove.Edges()
		w      = utils.Abs(br.X-tl.X) + 1
		h      = utils.Abs(br.Y-tl.Y) + 1
	)
	return h*w - len(grove)
}

func (grove Grove) String() string {
	var (
		out    strings.Builder
		tl, br = grove.Edges()
	)

	for y := tl.Y; y <= br.Y; y++ {
		for x := tl.X; x <= br.X; x++ {
			if grove[Position{x, y}] {
				out.WriteByte('#')
			} else {
				out.WriteByte('.')
			}
		}
		out.WriteByte('\n')
	}

	return out.String()
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
