package day22

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestVolume(t *testing.T) {
	utils.Assert(t, 27, Cuboid{10, 12, 10, 12, 10, 12}.Volume())
	utils.Assert(t, 27, Cuboid{-3, -1, -3, -1, -3, -1}.Volume())
	utils.Assert(t, 27, Cuboid{1, 3, 1, 3, 1, 3}.Volume())
}
