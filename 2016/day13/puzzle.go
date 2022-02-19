package day13

import "strconv"

type Maze struct {
	number int
	points map[Point]bool
}

type Point struct {
	X, Y int
}

type Route struct {
	position Point
	steps    int
}

func Part1(input string) int {
	number, _ := strconv.Atoi(input)
	return Bfs(&Maze{number, map[Point]bool{}}, Point{1, 1}, Point{31, 39})
}

func Part2(input string) int {
	number, _ := strconv.Atoi(input)
	return Within50(&Maze{number, map[Point]bool{}}, Point{1, 1})
}

func Bfs(maze *Maze, from, to Point) int {
	var route Route
	queue := []Route{{from, 0}}
	seen := map[Point]bool{from: true}
	for len(queue) > 0 {
		route, queue = queue[0], queue[1:]
		for _, p := range Neighbours(route.position) {
			switch {
			case p == to:
				return route.steps + 1
			case !maze.IsWall(p) && !seen[p]:
				seen[p] = true
				queue = append(queue, Route{p, route.steps + 1})
			}
		}
	}
	return -1
}

func Within50(maze *Maze, from Point) int {
	var route Route
	queue := []Route{{from, 0}}
	seen := map[Point]bool{from: true}
	for len(queue) > 0 {
		route, queue = queue[0], queue[1:]
		for _, p := range Neighbours(route.position) {
			if !maze.IsWall(p) && !seen[p] && route.steps < 50 {
				queue = append(queue, Route{p, route.steps + 1})
				seen[p] = true
			}
		}
	}
	return len(seen)
}

func (maze *Maze) IsWall(p Point) bool {
	_, ok := maze.points[p]
	if !ok {
		x, y := p.X, p.Y
		c, score := 0, maze.number+x*x+3*x+2*x*y+y+y*y
		for ; score > 0; score >>= 1 {
			c += score & 1
		}
		maze.points[p] = c%2 == 1
	}
	return maze.points[p]
}

func Neighbours(p Point) (n []Point) {
	for _, i := range [][]int{{0, -1}, {1, 0}, {0, 1}, {-1, 0}} {
		if p.X+i[0] >= 0 && p.Y+i[1] >= 0 {
			n = append(n, Point{p.X + i[0], p.Y + i[1]})
		}
	}
	return
}
