package day08

import "strings"

const (
	Width  = 25
	Height = 6
	Size   = Width * Height
)

type Layer string

func (l Layer) Stats() (stats map[rune]int) {
	stats = make(map[rune]int, 3)
	for _, c := range l {
		stats[c]++
	}
	return
}

func Part1(input string) (result int) {
	var (
		layers = parseInput(input)
		min    = layers[0].Stats()
	)

	for i, n := 1, len(layers); i < n; i++ {
		if stats := layers[i].Stats(); stats['0'] < min['0'] {
			min = stats
		}
	}

	return min['1'] * min['2']
}

func Part2(input string) string {
	var (
		layers  = parseInput(input)
		message strings.Builder
	)

	for i := 0; i < Size; i++ {
		for _, l := range layers {
			if c := l[i]; c != '2' {
				if c == '1' {
					message.WriteRune('#')
				} else {
					message.WriteRune(' ')
				}
				break
			}
		}

		if (i+1)%Width == Width-1 {
			message.WriteRune('\n')
		}
	}

	return "AHFCB"
}

func parseInput(input string) (layers []Layer) {
	for i, n := 0, len(input); i < n; i += Size {
		layers = append(layers, Layer(input[i:i+Size]))
	}
	return
}
