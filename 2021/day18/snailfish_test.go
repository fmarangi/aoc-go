package day18

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestMagnitude(t *testing.T) {
	nums := map[string]int{
		"[[1,2],[[3,4],5]]":                                             143,
		"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]":                             1384,
		"[[[[1,1],[2,2]],[3,3]],[4,4]]":                                 445,
		"[[[[3,0],[5,3]],[4,4]],[5,5]]":                                 791,
		"[[[[5,0],[7,4]],[5,5]],[6,6]]":                                 1137,
		"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]":         3488,
		"[[[[6,6],[7,6]],[[7,7],[7,0]]],[[[7,7],[7,7]],[[7,8],[9,9]]]]": 4140,
	}

	for k, v := range nums {
		utils.Assert(t, parseNumber(k).Magnitude(), v)
	}
}

func TestExplode(t *testing.T) {
	nums := []struct {
		from, to string
		pos      int
	}{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]", 0},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]", 4},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]", 3},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", 3},
	}

	for _, row := range nums {
		a, b := parseNumber(row.from), parseNumber(row.to)
		utils.Assert(t, a.explode(row.pos).Magnitude(), b.Magnitude())
	}
}

func TestSplit(t *testing.T) {
	nums := []struct {
		from, to string
		pos      int
	}{
		{"[[[[0,7],4],[15,[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", 3},
		{"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]", 6},
	}

	for _, row := range nums {
		a, b := parseNumber(row.from), parseNumber(row.to)
		utils.Assert(t, a.split(row.pos).Magnitude(), b.Magnitude())
	}
}
