package day16

import "strings"

func reverse(input string) string {
	var output strings.Builder
	for i := len(input) - 1; i >= 0; i-- {
		output.WriteByte((input[i]-'0'+1)%2 + '0')
	}
	return output.String()
}

func Checksum(input string) string {
	var checksum strings.Builder
	checksum.WriteString(input)
	n := checksum.Len()
	for n%2 == 0 {
		curr := checksum.String()
		checksum.Reset()
		checksum.Grow(n / 2)
		for i := 0; i < n; i += 2 {
			if curr[i] == curr[i+1] {
				checksum.WriteByte('1')
			} else {
				checksum.WriteByte('0')
			}
		}
		n = checksum.Len()
	}
	return checksum.String()
}

func Fill(initial string, length int) (output string) {
	output = initial
	for len(output) < length {
		output += "0" + reverse(output)
	}
	return output[:length]
}

func Part1(input string) string {
	return Checksum(Fill(input, 272))
}

func Part2(input string) string {
	return Checksum(Fill(input, 35651584))
}
