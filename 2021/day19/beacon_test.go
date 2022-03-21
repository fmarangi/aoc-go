package day19

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestDistance(t *testing.T) {
	utils.Assert(t, Beacon{1, 2, 3}.Distance(Beacon{}), 6)
	utils.Assert(t, Beacon{-1, 2, -3}.Distance(Beacon{}), 6)
}

func TestOffset(t *testing.T) {
	utils.Assert(t, Beacon{1, 2, 3}.Offset(Beacon{}), Beacon{-1, -2, -3})
}

func TestMove(t *testing.T) {
	utils.Assert(t, Beacon{1, 2, 3}.Move(Beacon{-1, -2, -3}), Beacon{})
}
