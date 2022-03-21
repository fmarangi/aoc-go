package day19

import (
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func Part1(input string) int {
	scanners := parseInput(input)
	n, found := len(scanners), map[int]bool{0: true}

	for len(found) != n {
		for i := range found {
			for j := 0; j < n; j++ {
				if j == i || found[j] {
					continue
				}

				offset, rotate := scanners[i].Intersect(scanners[j])
				if rotate != nil {
					scanners[j] = scanners[j].Translate(offset, rotate)
					found[j] = true
				}
			}
		}
	}

	return count(scanners...)
}

func Part2(input string) (max int) {
	scanners := parseInput(input)
	n, found := len(scanners), map[int]Beacon{0: {}}

	for len(found) != n {
		for i := range found {
			for j := 0; j < n; j++ {
				if _, ok := found[j]; ok || j == i {
					continue
				}

				offset, rotate := scanners[i].Intersect(scanners[j])
				if rotate != nil {
					scanners[j] = scanners[j].Translate(offset, rotate)
					found[j] = offset
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if distance := found[i].Distance(found[j]); distance > max {
				max = distance
			}
		}
	}

	return
}

func count(scanners ...Scanner) int {
	all := map[Beacon]bool{}
	for _, s := range scanners {
		for _, b := range s {
			all[b] = true
		}
	}

	return len(all)
}

func parseInput(input string) (scanners []Scanner) {
	for _, s := range strings.Split(input, "\n\n") {
		scanner := Scanner{}
		for _, b := range strings.Split(s, "\n")[1:] {
			vals := utils.Ints(strings.Split(b, ","))
			scanner = append(scanner, Beacon{vals[0], vals[1], vals[2]})
		}
		scanners = append(scanners, scanner)
	}
	return
}
