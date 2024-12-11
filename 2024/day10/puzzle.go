package day10

type Point struct{ X, Y int }

func (a Point) Add(b Point) Point {
	return Point{a.X + b.X, a.Y + b.Y}
}

type Area map[Point]int

func (area Area) Trails(p Point) int {
	var (
		queue      = []Point{p}
		curr, next Point
		found      = make(map[Point]bool)
		offsets    = []Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	)

	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		for _, o := range offsets {
			next = curr.Add(o)
			if v := area[next]; v == area[curr]+1 {
				if v == 9 {
					found[next] = true
				} else {
					queue = append(queue, next)
				}
			}
		}
	}

	return len(found)
}

func (area Area) Ratings(p Point) (ratings int) {
	var (
		queue      = []Point{p}
		curr, next Point
		offsets    = []Point{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	)

	for len(queue) > 0 {
		curr, queue = queue[0], queue[1:]
		for _, o := range offsets {
			next = curr.Add(o)
			if v := area[next]; v == area[curr]+1 {
				if v == 9 {
					ratings++
				} else {
					queue = append(queue, next)
				}
			}
		}
	}

	return
}

func Part1(input string) (trails int) {
	var area = parseInput(input)
	for p, v := range area {
		if v == 0 {
			trails += area.Trails(p)
		}
	}
	return
}

func Part2(input string) (ratings int) {
	var area = parseInput(input)
	for p, v := range area {
		if v == 0 {
			ratings += area.Ratings(p)
		}
	}
	return
}

func parseInput(input string) (area Area) {
	area = make(Area)

	var x, y int
	for _, c := range input {
		if c == '\n' {
			x, y = 0, y+1
		} else {
			area[Point{x, y}] = int(c - '0')
			x++
		}
	}

	return
}
