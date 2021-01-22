package day25

import "strconv"
import "strings"

func Part1(input string) int {
	door, key := parseInput(input)
	return EncryptionKey(key, LoopSize(7, door))
}

func parseInput(input string) (door, key int) {
	keys := strings.Split(input, "\n")
	door, _ = strconv.Atoi(keys[0])
	key, _ = strconv.Atoi(keys[1])
	return door, key
}

func LoopSize(subject, publicKey int) (s int) {
	for n := 1; n != publicKey; s++ {
		n = (n * subject) % 20201227
	}
	return s
}

func EncryptionKey(subject, loopSize int) (key int) {
	key = 1
	for i := 0; i < loopSize; i++ {
		key = (key * subject) % 20201227
	}
	return key
}
