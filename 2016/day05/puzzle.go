package day05

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func MD5(prefix string, index int) string {
	value := fmt.Sprintf("%s%d", prefix, index)
	return fmt.Sprintf("%x", md5.Sum([]byte(value)))
}

func password(prefix string) string {
	var pwd strings.Builder
	for c, i := 0, 0; c < 8; i++ {
		hash := MD5(prefix, i)
		if strings.HasPrefix(hash, "00000") {
			pwd.WriteRune(rune(hash[5]))
			c++
		}
	}
	return pwd.String()
}

func passwordV2(prefix string) string {
	pwd := make([]rune, 8)
	for i, c := 0, 0; c < 8; i++ {
		hash := MD5(prefix, i)
		if !strings.HasPrefix(hash, "00000") {
			continue
		}

		if p := int(hash[5] - '0'); p >= 0 && p < 8 && int(pwd[p]) == 0 {
			pwd[p] = rune(hash[6])
			c++
		}
	}
	return string(pwd)
}

func Part1(input string) string {
	return password(input)
}

func Part2(input string) string {
	return passwordV2(input)
}
