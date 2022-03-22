package day14

import (
	"fmt"
	"strconv"

	"github.com/fmarangi/aoc-go/2017/knot"
)

type Disk [128][128]bool
type Square struct{ X, Y int }

func (s Square) Neighbours() (neighbours []Square) {
	offsets := [4]Square{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	for _, o := range offsets {
		x, y := s.X+o.X, s.Y+o.Y
		if x >= 0 && x < 128 && y >= 0 && y < 128 {
			neighbours = append(neighbours, Square{x, y})
		}
	}
	return
}

func toInt(hex string) int {
	num, _ := strconv.ParseInt(hex, 16, 32)
	return int(num)
}

func row(key string, num int) (row [16]int) {
	hash := knot.Hash(fmt.Sprintf("%s-%d", key, num))
	for i := 0; i < 16; i++ {
		row[i] = toInt(hash[i*2 : i*2+2])
	}
	return
}

func GetDisk(key string) (disk Disk) {
	for i := 0; i < 128; i++ {
		for j, r := 0, row(key, i); j < 128; j++ {
			if r[j/8]&(1<<(7-j%8)) != 0 {
				disk[i][j] = true
			}
		}
	}
	return
}

func Part1(input string) (used int) {
	disk := GetDisk(input)
	for i := 0; i < 1<<14; i++ {
		if disk[i/128][i%128] {
			used++
		}
	}
	return
}

func Part2(input string) (areas int) {
	disk := GetDisk(input)
	seen := map[Square]bool{}
	for i := 0; i < 1<<14; i++ {
		square := Square{i / 128, i % 128}
		if seen[square] || !disk[square.X][square.Y] {
			continue
		}

		areas++
		for q := []Square{square}; len(q) > 0; q = q[1:] {
			current := q[0]
			for _, n := range current.Neighbours() {
				if !seen[n] && disk[n.X][n.Y] {
					seen[n] = true
					q = append(q, n)
				}
			}
		}
	}
	return
}
