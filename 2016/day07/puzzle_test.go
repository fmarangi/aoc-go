package day07

import (
	"testing"

	"github.com/fmarangi/aoc-go/utils"
)

func TestTls(t *testing.T) {
	utils.Assert(t, true, parseIp("abba[mnop]qrst").Tls())
	utils.Assert(t, false, parseIp("abcd[bddb]xyyx").Tls())
	utils.Assert(t, false, parseIp("aaaa[qwer]tyui").Tls())
	utils.Assert(t, true, parseIp("ioxxoj[asdfgh]zxcvbn").Tls())
}

func TestSsl(t *testing.T) {
	utils.Assert(t, true, parseIp("aba[bab]xyz").Ssl())
	utils.Assert(t, false, parseIp("xyx[xyx]xyx").Ssl())
	utils.Assert(t, true, parseIp("aaa[kek]eke").Ssl())
	utils.Assert(t, true, parseIp("zazbz[bzb]cdb").Ssl())
}

func TestSolvePuzzle(t *testing.T) {
	input := utils.ReadInput("input.txt")
	utils.Assert(t, 105, Part1(input))
	utils.Assert(t, 258, Part2(input))
}
