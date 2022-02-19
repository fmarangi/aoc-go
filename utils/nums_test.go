package utils

import "testing"

func TestToInt(t *testing.T) {
	tests := []struct {
		input string
		exp   int
	}{{"42", 42}}

	for _, test := range tests {
		if result := ToInt(test.input); result != test.exp {
			t.Errorf("Expected %d, got: %d", test.exp, result)
		}
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		input int
		exp   int
	}{{42, 42}, {-42, 42}}

	for _, test := range tests {
		if result := Abs(test.input); result != test.exp {
			t.Errorf("Expected %d, got: %d", test.exp, result)
		}
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		input []int
		exp   int
	}{
		{[]int{1, 2, 3, 4}, 10},
		{[]int{1, 2, 4, 8, 16, 32, 64}, 127},
	}

	for _, test := range tests {
		if result := Sum(test.input); result != test.exp {
			t.Errorf("Expected %d, got: %d", test.exp, result)
		}
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		input []int
		exp   int
	}{
		{[]int{1, 2, 3, 4}, 4},
		{[]int{1, 4, 64, 2, 8, 32, 16}, 64},
		{[]int{-1, -2, -3, -4}, -1},
	}

	for _, test := range tests {
		if result := Max(test.input); result != test.exp {
			t.Errorf("Expected %d, got: %d", test.exp, result)
		}
	}
}

func TestMin(t *testing.T) {
	tests := []struct {
		input []int
		exp   int
	}{
		{[]int{1, 2, 3, 4}, 1},
		{[]int{64, 8, 32, 4, 16}, 4},
		{[]int{-1, -2, -4, -3}, -4},
	}

	for _, test := range tests {
		if result := Min(test.input); result != test.exp {
			t.Errorf("Expected %d, got: %d", test.exp, result)
		}
	}
}
