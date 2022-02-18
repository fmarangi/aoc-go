package day06

import "strings"

type Frequency map[rune]int

func (f Frequency) Max() (max rune) {
	for k, v := range f {
		if v > f[max] {
			max = k
		}
	}
	return
}

func (f Frequency) Min() (min rune) {
	for k, v := range f {
		if min == 0 || v < f[min] {
			min = k
		}
	}
	return
}

func collectFrequencies(messages []string) (freqs []Frequency) {
	for i, c, m := 0, len(messages[0]), len(messages); i < c; i++ {
		freq := Frequency{}
		for j := 0; j < m; j++ {
			freq[rune(messages[j][i])]++
		}
		freqs = append(freqs, freq)
	}
	return
}

func Part1(input string) (result string) {
	freqs := collectFrequencies(strings.Fields(input))
	for _, f := range freqs {
		result = result + string(f.Max())
	}
	return
}

func Part2(input string) (result string) {
	freqs := collectFrequencies(strings.Fields(input))
	for _, f := range freqs {
		result = result + string(f.Min())
	}
	return
}
