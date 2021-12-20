package day16

import (
	"io/ioutil"
	"strings"
	"testing"
)

type Solution struct {
	part   func(string) int
	result int
}

func TestPacketValue(t *testing.T) {
	packets := map[string]int{
		"C200B40A82":                 3,
		"04005AC33890":               54,
		"880086C3E88112":             7,
		"CE00C43D881120":             9,
		"D8005AC2A8F0":               1,
		"F600BC2D8F":                 0,
		"9C005AC2F8F0":               0,
		"9C0141080250320F1802104A08": 1,
	}

	for p, r := range packets {
		if packet, _ := readPacket(ToBin(p)); packet.Value() != r {
			t.Errorf("Expected %d, got: %d", r, packet.Value())
		}
	}
}

func TestSolvePuzzle(t *testing.T) {
	parts := []Solution{
		{Part1, 953},
		{Part2, 246225449979},
	}

	content, _ := ioutil.ReadFile("input.txt")
	input := strings.TrimSpace(string(content))
	for _, test := range parts {
		if res := test.part(input); res != test.result {
			t.Errorf("Expected %d, got: %d", test.result, res)
		}
	}
}
