package day12

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Pos struct{ X, Y, Z int }
type Moon struct{ Pos, Vel Pos }

func (m Moon) Pot() int {
	return utils.Abs(m.Pos.X) + utils.Abs(m.Pos.Y) + utils.Abs(m.Pos.Z)
}

func (m Moon) Kin() int {
	return utils.Abs(m.Vel.X) + utils.Abs(m.Vel.Y) + utils.Abs(m.Vel.Z)
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func compare(a, b int) int {
	switch {
	case a < b:
		return 1
	case a > b:
		return -1
	}
	return 0
}

func Step(moons []Moon) []Moon {
	var n = len(moons)

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			moons[i].Vel.X += compare(moons[i].Pos.X, moons[j].Pos.X)
			moons[i].Vel.Y += compare(moons[i].Pos.Y, moons[j].Pos.Y)
			moons[i].Vel.Z += compare(moons[i].Pos.Z, moons[j].Pos.Z)

			moons[j].Vel.X += compare(moons[j].Pos.X, moons[i].Pos.X)
			moons[j].Vel.Y += compare(moons[j].Pos.Y, moons[i].Pos.Y)
			moons[j].Vel.Z += compare(moons[j].Pos.Z, moons[i].Pos.Z)
		}
	}

	for i := 0; i < n; i++ {
		moons[i].Pos.X += moons[i].Vel.X
		moons[i].Pos.Y += moons[i].Vel.Y
		moons[i].Pos.Z += moons[i].Vel.Z
	}

	return moons
}

func FindLoop(values []int) (steps int) {
	var (
		n    = len(values)
		vel  = make([]int, len(values))
		seen = make(map[string]bool)
	)

	for !seen[fmt.Sprintf("%v;%v", values, vel)] {
		seen[fmt.Sprintf("%v;%v", values, vel)] = true
		for i := 0; i < n; i++ {
			for j := i; j < n; j++ {
				v := compare(values[i], values[j])
				vel[i] += +v
				vel[j] += -v
			}
		}

		for i := 0; i < n; i++ {
			values[i] += vel[i]
		}

		steps++
	}

	return
}

func Part1(input string) (energy int) {
	moons := parseInput(input)
	for i := 0; i < 1000; i++ {
		moons = Step(moons)
	}

	for _, moon := range moons {
		energy += moon.Pot() * moon.Kin()
	}

	return
}

func Part2(input string) (loop int) {
	axes := [3][]int{}
	for _, m := range parseInput(input) {
		axes[0] = append(axes[0], m.Pos.X)
		axes[1] = append(axes[1], m.Pos.Y)
		axes[2] = append(axes[2], m.Pos.Z)
	}

	loop = 1
	for i := 0; i < 3; i++ {
		l := FindLoop(axes[i])
		loop = l * loop / GCD(l, loop)
	}

	return
}

func parseInput(input string) (moons []Moon) {
	for _, row := range strings.Split(input, "\n") {
		var moon Moon
		fmt.Sscanf(row, "<x=%d, y=%d, z=%d>", &moon.Pos.X, &moon.Pos.Y, &moon.Pos.Z)
		moons = append(moons, moon)
	}
	return
}
