package utils

import "testing"

func TestToInt(t *testing.T) {
	Assert(t, 42, ToInt("42"))
}

func TestAbs(t *testing.T) {
	Assert(t, 42, Abs(42))
	Assert(t, 42, Abs(-42))
}

func TestSum(t *testing.T) {
	Assert(t, 10, Sum([]int{1, 2, 3, 4}))
	Assert(t, 127, Sum([]int{1, 2, 4, 8, 16, 32, 64}))
}

func TestMax(t *testing.T) {
	Assert(t, 4, Max(1, 2, 3, 4))
	Assert(t, 64, Max(1, 4, 64, 2, 8, 32, 16))
	Assert(t, -1, Max(-1, -2, -3, -4))
}

func TestMin(t *testing.T) {
	Assert(t, 1, Min(1, 2, 3, 4))
	Assert(t, 4, Min(64, 8, 32, 4, 16))
	Assert(t, -4, Min(-1, -2, -4, -3))
}
