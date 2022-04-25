package day02

import s "strings"

type keypad string

func (k keypad) code(sequence string) string {
	row, current := s.Index(string(k), "\n")+1, s.Index(string(k), "5")
	moves := map[rune]int{'U': -row, 'R': 1, 'D': row, 'L': -1}

	var code s.Builder
	for _, seq := range s.Fields(sequence) {
		for _, m := range seq {
			next := current + moves[m]
			if next >= 0 && next < len(k) && k[next] != '\n' && k[next] != ' ' {
				current = next
			}
		}
		code.WriteRune([]rune(k)[current])
	}
	return code.String()
}

func Part1(input string) string {
	return keypad("123\n456\n789").code(input)
}

func Part2(input string) string {
	return keypad("  1  \n 234 \n56789\n ABC \n  D  ").code(input)
}
