package day04

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func Hash(secret string, n int) string {
	data := []byte(fmt.Sprintf("%s%d", secret, n))
	return fmt.Sprintf("%x", md5.Sum(data))
}

func Part1(input string) (i int) {
	for ; !strings.HasPrefix(Hash(input, i), "00000"); i++ {
	}
	return i
}

func Part2(input string) (i int) {
	for ; !strings.HasPrefix(Hash(input, i), "000000"); i++ {
	}
	return i
}
