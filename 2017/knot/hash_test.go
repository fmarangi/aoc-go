package knot

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestHash(t *testing.T) {
	tests := []struct{ input, expected string }{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}

	for _, test := range tests {
		utils.Assert(t, test.expected, Hash(test.input))
	}
}
