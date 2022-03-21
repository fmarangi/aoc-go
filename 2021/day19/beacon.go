package day19

import (
	"errors"

	"github.com/fmarangi/aoc-go/utils"
)

type Beacon struct{ X, Y, Z int }
type Rotate func(Beacon) Beacon

func (a Beacon) Distance(b Beacon) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y) + utils.Abs(a.Z-b.Z)
}

func (a Beacon) Offset(b Beacon) Beacon {
	return Beacon{b.X - a.X, b.Y - a.Y, b.Z - a.Z}
}

func (a Beacon) Move(b Beacon) Beacon {
	return Beacon{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

func findRotation(from, to Beacon) (Rotate, error) {
	for _, r := range Rotations() {
		if r(to) == from {
			return r, nil
		}
	}
	return nil, errors.New("no suitable rotation found")
}

func Rotations() []Rotate {
	return []Rotate{
		func(b Beacon) Beacon { return b },
		func(b Beacon) Beacon { return Beacon{-b.Y, b.X, b.Z} },
		func(b Beacon) Beacon { return Beacon{-b.X, -b.Y, b.Z} },
		func(b Beacon) Beacon { return Beacon{b.Y, -b.X, b.Z} },
		func(b Beacon) Beacon { return Beacon{b.Z, b.Y, -b.X} },
		func(b Beacon) Beacon { return Beacon{-b.Y, b.Z, -b.X} },
		func(b Beacon) Beacon { return Beacon{-b.Z, -b.Y, -b.X} },
		func(b Beacon) Beacon { return Beacon{b.Y, -b.Z, -b.X} },
		func(b Beacon) Beacon { return Beacon{-b.Z, b.Y, b.X} },
		func(b Beacon) Beacon { return Beacon{-b.Y, -b.Z, b.X} },
		func(b Beacon) Beacon { return Beacon{b.Z, -b.Y, b.X} },
		func(b Beacon) Beacon { return Beacon{b.Y, b.Z, b.X} },
		func(b Beacon) Beacon { return Beacon{b.X, -b.Z, b.Y} },
		func(b Beacon) Beacon { return Beacon{b.Z, b.X, b.Y} },
		func(b Beacon) Beacon { return Beacon{-b.X, b.Z, b.Y} },
		func(b Beacon) Beacon { return Beacon{-b.Z, -b.X, b.Y} },
		func(b Beacon) Beacon { return Beacon{b.X, -b.Y, -b.Z} },
		func(b Beacon) Beacon { return Beacon{b.Y, b.X, -b.Z} },
		func(b Beacon) Beacon { return Beacon{-b.X, b.Y, -b.Z} },
		func(b Beacon) Beacon { return Beacon{-b.Y, -b.X, -b.Z} },
		func(b Beacon) Beacon { return Beacon{b.X, b.Z, -b.Y} },
		func(b Beacon) Beacon { return Beacon{-b.Z, b.X, -b.Y} },
		func(b Beacon) Beacon { return Beacon{-b.X, -b.Z, -b.Y} },
		func(b Beacon) Beacon { return Beacon{b.Z, -b.X, -b.Y} },
	}
}
