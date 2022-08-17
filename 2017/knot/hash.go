package knot

import "encoding/hex"

type Circle [256]byte

func Init() (c Circle) {
	for i := range c {
		c[i] = byte(i)
	}
	return
}

func (c Circle) String() string {
	return hex.EncodeToString(c.toDenseHash())
}

func (c *Circle) Reverse(from, size int) {
	for i := 0; i < size/2; i++ {
		a, b := (from+i)%256, (from+size-i-1)%256
		c[a], c[b] = c[b], c[a]
	}
}

func (c Circle) toDenseHash() []byte {
	result := make([]byte, 16)
	for i := range c {
		result[i/16] ^= c[i]
	}
	return result
}

func Hash(input string) string {
	circle, lengths := Init(), toLengths(input)
	for i, skip, pos := 0, 0, 0; i < 64; i++ {
		for _, l := range lengths {
			circle.Reverse(pos, l)
			pos += l + skip
			skip++
		}
	}
	return circle.String()
}

func toLengths(input string) []int {
	result := make([]int, len(input), len(input)+5)
	for i, c := range input {
		result[i] = int(c)
	}
	return append(result, 17, 31, 73, 47, 23)
}
