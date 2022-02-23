package day24

import (
	"testing"
)

func TestBiodiversity(t *testing.T) {
	grid := Grid{}
	grid[3][0] = true
	grid[4][1] = true

	exp := 2129920
	if res := grid.Biodiversity(); res != exp {
		t.Errorf("Expected %d, got: %d", exp, res)
	}
}
