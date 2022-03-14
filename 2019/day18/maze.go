package day18

import (
	"fmt"
	"strings"
)

type Location struct{ Pos, Keys int }

func which(keys int) (k string) {
	for i := 0; i < 26; i++ {
		if keys&(1<<i) != 0 {
			k += string(rune('a' + i))
		}
	}
	return
}

func CanOpen(door byte, keys int) bool {
	return 1<<(door-'A')&keys != 0
}

func (loc Location) String() string {
	return fmt.Sprintf("%d [%s]", loc.Pos, which(loc.Keys))
}

type Route struct {
	Location Location
	Steps    int
}

func (r Route) String() string {
	loc := r.Location
	return fmt.Sprintf("{%d steps @ %s}", r.Steps, loc)
}

type Maze string

func (maze Maze) Init() int  { return strings.Index(string(maze), "@") }
func (maze Maze) Width() int { return strings.Index(string(maze), "\n") + 1 }

func (maze Maze) AllKeys() (all int) {
	for i := 'a'; i <= 'z'; i++ {
		if !strings.Contains(string(maze), string(i)) {
			break
		}
		all++
	}
	return 1<<all - 1
}

func (maze Maze) bfs(route Route) chan Route {
	found := make(chan Route)
	go func() {
		keys := route.Location.Keys
		width := maze.Width()
		seen := map[int]bool{route.Location.Pos: true}
		queue := []Route{route}

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]
			for _, off := range [...]int{1, -1, -width, width} {
				next := curr.Location.Pos + off
				if seen[next] {
					continue
				}

				cell := maze[next]
				switch {
				case isKey(cell) && keys&Value(cell) == 0:
					found <- Route{Location{next, keys | Value(cell)}, curr.Steps + 1}
				case cell == '.', cell == '@', isKey(cell), isDoor(cell) && CanOpen(cell, keys):
					seen[next] = true
					queue = append(queue, Route{Location{next, keys}, curr.Steps + 1})
				}
			}
		}
		close(found)
	}()
	return found
}

func isKey(pos byte) bool  { return pos >= 'a' && pos <= 'z' }
func isDoor(pos byte) bool { return pos >= 'A' && pos <= 'Z' }
func Value(key byte) int   { return 1 << (key - 'a') }
