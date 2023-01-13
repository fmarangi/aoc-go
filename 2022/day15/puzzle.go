package day15

import (
	"fmt"
	"sort"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Point struct{ X, Y int }
type Sensor struct{ Position, Beacon Point }

func (a Point) Distance(b Point) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y)
}

func (s Sensor) Range() int {
	return s.Position.Distance(s.Beacon)
}

func (s Sensor) Coverage(row int) ([2]int, error) {
	x := s.Position.X
	if d := s.Range() - s.Position.Distance(Point{x, row}); d >= 0 {
		return [2]int{x - d, x + d}, nil
	}
	return [2]int{}, fmt.Errorf("cannot detect on row %d", row)
}

func Gap(sensors []Sensor, row int) int {
	var v [][2]int
	for _, s := range sensors {
		if c, err := s.Coverage(row); err == nil {
			v = append(v, c)
		}
	}

	sort.Slice(v, func(i, j int) bool {
		if v[i][0] == v[j][0] {
			return v[i][1] < v[j][1]
		}
		return v[i][0] < v[j][0]
	})

	var m = v[0][1]
	for i, n := 1, len(v); i < n && m < 4000000; i++ {
		switch {
		case v[i][0] > m+1:
			return m + 1
		case v[i][1] > m:
			m = v[i][1]
		}
	}

	return -1

}

func Part1(input string) int {
	var min, max []int
	for _, s := range parseInput(input) {
		if c, err := s.Coverage(2000000); err == nil {
			min, max = append(min, c[0]), append(max, c[1])
		}
	}
	return utils.Max(max...) - utils.Min(min...)
}

func Part2(input string) int {
	var sensors = parseInput(input)
	for i := 0; i <= 4000000; i++ {
		if g := Gap(sensors, i); g > 0 {
			return g*4000000 + i
		}
	}
	return -1
}

func parseInput(input string) (sensors []Sensor) {
	for _, row := range strings.Split(strings.TrimSpace(input), "\n") {
		var s Sensor
		fmt.Sscanf(row, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d",
			&s.Position.X, &s.Position.Y, &s.Beacon.X, &s.Beacon.Y)
		sensors = append(sensors, s)
	}
	return
}
