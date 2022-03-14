package day18

import (
	"fmt"
)

type Robots struct {
	Pos  [4]int
	Keys int
}

func (bots Robots) String() string {
	return fmt.Sprintf("%d,%d,%d,%d [%s]", bots.Pos[0], bots.Pos[1], bots.Pos[2], bots.Pos[3], which(bots.Keys))
}

type RobotRoute struct {
	Robots Robots
	Steps  int
}

func (maze Maze) ToRobots() Maze {
	start, w, m := maze.Init(), maze.Width(), string(maze)
	robots := m[:start-w-1] + "@#@"
	robots += m[start-w+2:start-1] + "###"
	robots += m[start+2:start+w-1] + "@#@"
	return Maze(robots + m[start+w+2:])
}

func (maze Maze) InitRobots() RobotRoute {
	first, w := maze.Init(), maze.Width()
	bots := Robots{[...]int{first, first + 2, first + 2*w, first + 2*w + 2}, 0}
	return RobotRoute{bots, 0}
}

func (maze Maze) robotsBfs(route RobotRoute) chan RobotRoute {
	found := make(chan RobotRoute)
	go func() {
		keys, w := route.Robots.Keys, maze.Width()

		for r := 0; r < 4; r++ {
			seen := map[int]bool{route.Robots.Pos[r]: true}
			queue := []RobotRoute{route}
			for len(queue) > 0 {
				curr := queue[0]
				queue = queue[1:]
				for _, off := range [...]int{1, -1, -w, w} {
					next := curr.Robots.Pos[r] + off
					if seen[next] {
						continue
					}

					cell := maze[next]
					switch {
					case isKey(cell) && keys&Value(cell) == 0:
						bots := curr.Robots.Pos
						bots[r] = next
						found <- RobotRoute{Robots{bots, keys | Value(cell)}, curr.Steps + 1}
					case cell == '.', cell == '@', isKey(cell), isDoor(cell) && CanOpen(cell, keys):
						bots := curr.Robots.Pos
						bots[r] = next
						queue = append(queue, RobotRoute{Robots{bots, keys}, curr.Steps + 1})
						seen[next] = true
					}
				}
			}
		}

		close(found)
	}()
	return found
}
