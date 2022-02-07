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
