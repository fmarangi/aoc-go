package day08

import "strings"

func Part1(input string) (visible int) {
	trees, w := parseInput(input)
	n := len(trees)

	for i, t := range trees {
		row, col := i/w, i%w
		if row == 0 || col == 0 || row+1 == w || col+1 == w {
			visible++
			continue
		}

	outer:
		for _, o := range [4]int{1, -1, w, -w} {
			for j := i + o; j >= 0 && j < n; j += o {
				if (o == 1 || o == -1) && j/w != row {
					break
				}

				if trees[j] >= t {
					continue outer
				}
			}

			visible++
			break
		}
	}
	return
}

func Part2(input string) (best int) {
	trees, w := parseInput(input)
	n := len(trees)

	for i, t := range trees {
		row, col := i/w, i%w
		if row == 0 || col == 0 || row+1 == w || col+1 == w {
			continue
		}

		var seen [4]int
		for x, o := range [4]int{1, -1, w, -w} {
			for j := i + o; j >= 0 && j < n; j += o {
				if (o == 1 || o == -1) && j/w != row {
					break
				}

				seen[x]++
				if trees[j] >= t {
					break
				}
			}
		}

		if score := calcScore(seen); score > best {
			best = score
		}
	}
	return
}

func calcScore(trees [4]int) (score int) {
	score = 1
	for _, t := range trees {
		score *= t
	}
	return
}

func parseInput(input string) (trees []int, width int) {
	width = strings.IndexByte(input, '\n')
	for _, t := range input {
		if t != '\n' {
			trees = append(trees, int(t-'0'))
		}
	}
	return
}
