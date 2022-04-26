package day03

import "strings"

func slope(grid string, right, down int) (trees int) {
	w := strings.Index(grid, "\n") + 1
	for i, n := 0, len(grid)/w; i*down <= n; i++ {
		p := down*w*i + (right*i)%(w-1)
		if grid[p] == '#' {
			trees++
		}
	}
	return
}

func Part1(input string) int {
	return slope(input, 3, 1)
}

func Part2(input string) (trees int) {
	slopes := []struct{ r, d int }{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	trees = 1
	for _, s := range slopes {
		trees *= slope(input, s.r, s.d)
	}

	return
}
