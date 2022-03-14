package day18

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestLocation(t *testing.T) {
	loc := Location{42, 5}
	utils.Assert(t, "42 [ac]", loc.String())
	utils.Assert(t, true, CanOpen('A', 5))
	utils.Assert(t, false, CanOpen('B', 5))
	utils.Assert(t, true, CanOpen('C', 5))
}
