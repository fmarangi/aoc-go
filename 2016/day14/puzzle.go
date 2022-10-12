package day14

import (
	"crypto/md5"
	"fmt"
)

func Hash(prefix string, num int) string {
	data := []byte(fmt.Sprintf("%s%d", prefix, num))
	return fmt.Sprintf("%x", md5.Sum(data))
}

func StretchedHash(prefix string, num int) (hash string) {
	hash = fmt.Sprintf("%s%d", prefix, num)
	for i := 0; i < 2017; i++ {
		hash = fmt.Sprintf("%x", md5.Sum([]byte(hash)))
	}
	return
}

func has(qty int, hash string) (byte, error) {
	n, found := len(hash), 1
	for i, prev := 1, hash[0]; i < n; i, prev = i+1, hash[i] {
		if hash[i] == prev {
			found++
		} else {
			found = 1
		}

		if found == qty {
			return prev, nil
		}
	}
	return 0, fmt.Errorf("no match found")
}

func Keys(prefix string, fn func(string, int) string) <-chan int {
	var (
		out    = make(chan int)
		hashes [1000]string
		five   = make(map[byte]int, 16)
		hash   string
	)

	go func() {
		// Pre-fill and analyse the first 1000 hashes
		for i := 0; i < 1000; i++ {
			hashes[i] = fn(prefix, i)
			if char, err := has(5, hashes[i]); err == nil {
				five[char] = i
			}
		}

		for i := 0; true; i++ {
			hash = fn(prefix, i+1000)
			if char, err := has(5, hash); err == nil {
				five[char] = i + 1000
			}

			if char, err := has(3, hashes[i%1000]); err == nil && five[char] > i {
				out <- i
			}

			hashes[i%1000] = hash
		}
	}()
	return out
}

func Part1(input string) (key int) {
	hashes := Keys(input, Hash)
	for i := 0; i < 64; i++ {
		key = <-hashes
	}
	return
}

func Part2(input string) (key int) {
	hashes := Keys(input, StretchedHash)
	for i := 0; i < 64; i++ {
		key = <-hashes
	}
	return
}
