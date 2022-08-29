package day22

import "errors"

type Cuboid struct {
	X1, X2, Y1, Y2, Z1, Z2 int
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func (c Cuboid) Volume() int {
	return (c.X2 - c.X1 + 1) * (c.Y2 - c.Y1 + 1) * (c.Z2 - c.Z1 + 1)
}

func (c Cuboid) Intersect(other Cuboid) (Cuboid, error) {
	x1, x2 := Max(c.X1, other.X1), Min(c.X2, other.X2)
	y1, y2 := Max(c.Y1, other.Y1), Min(c.Y2, other.Y2)
	z1, z2 := Max(c.Z1, other.Z1), Min(c.Z2, other.Z2)
	if x1 <= x2 && y1 <= y2 && z1 <= z2 {
		return Cuboid{x1, x2, y1, y2, z1, z2}, nil
	}
	return Cuboid{}, errors.New("no intersection")
}

func (c Cuboid) InRange() bool {
	i, _ := c.Intersect(Cuboid{-50, 50, -50, 50, -50, 50})
	return i == c
}
