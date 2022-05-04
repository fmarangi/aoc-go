package day11

import (
	"fmt"

	"github.com/fmarangi/aoc-go/utils"
)

type Cell struct{ X, Y int }

func (c Cell) Power(serial int) int {
	rackId := c.X + 10
	return ((rackId*c.Y + serial) * rackId / 100 % 10) - 5
}

func (c Cell) String() string {
	return fmt.Sprintf("%d,%d", c.X, c.Y)
}

type Grid map[Cell]int

func (g Grid) TotalPower(cell Cell, size int) (power int) {
	for _, o := range [...]Cell{{0, 0}, {size, size}} {
		power += g[Cell{cell.X + o.X - 1, cell.Y + o.Y - 1}]
	}
	for _, o := range [...]Cell{{0, size}, {size, 0}} {
		power -= g[Cell{cell.X + o.X - 1, cell.Y + o.Y - 1}]
	}
	return
}

func Part1(input string) string {
	grid := makeGrid(utils.ToInt(input), 300)
	highest, max := Cell{}, 0
	for x := 1; x <= 300-2; x++ {
		for y := 1; y <= 300-2; y++ {
			cell := Cell{x, y}
			if pow := grid.TotalPower(cell, 3); pow > max {
				highest, max = cell, pow
			}
		}
	}
	return highest.String()
}

func Part2(input string) string {
	grid := makeGrid(utils.ToInt(input), 300)
	highest, max, size := Cell{}, 0, 0
	for s := 1; s <= 300; s++ {
		for x := 1; x <= 301-s; x++ {
			for y := 1; y <= 301-s; y++ {
				cell := Cell{x, y}
				if pow := grid.TotalPower(cell, s); pow > max {
					highest, max, size = cell, pow, s
				}
			}
		}
	}
	return fmt.Sprintf("%s,%d", highest, size)
}

// Summed-area table of the power levels
func makeGrid(serial int, size int) (grid Grid) {
	grid = make(map[Cell]int, size*size)
	for x := 1; x <= size; x++ {
		for y := 1; y <= size; y++ {
			cell := Cell{x, y}
			grid[cell] = cell.Power(serial) - grid[Cell{x - 1, y - 1}]
			grid[cell] += grid[Cell{x - 1, y}] + grid[Cell{x, y - 1}]
		}
	}
	return
}
