package day09

import (
	"fmt"
	"strings"
)

func Decompress(data string) (length int) {
	var n, t, start, stop int
	for {
		start = strings.Index(data, "(")
		if start == -1 {
			break
		}

		stop = strings.Index(data, ")")
		marker := data[start+1 : stop]
		fmt.Sscanf(marker, "%dx%d", &n, &t)
		length += start + n*t
		data = data[stop+n+1:]
	}
	return length + len(data)
}

func DecompressV2(data string) (length int) {
	var n, t, start, stop int
	for {
		start = strings.Index(data, "(")
		if start == -1 {
			break
		}

		stop = strings.Index(data, ")")
		marker := data[start+1 : stop]
		fmt.Sscanf(marker, "%dx%d", &n, &t)
		length += start + t*DecompressV2(data[stop+1:stop+n+1])
		data = data[stop+n+1:]
	}
	return length + len(data)
}

func Part1(input string) int {
	return Decompress(input)
}

func Part2(input string) int {
	return DecompressV2(input)
}
