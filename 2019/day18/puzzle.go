package day18

import (
	"math"
	"sort"

	"github.com/fmarangi/aoc-go/utils"
)

func Part1(input string) (shortest int) {
	maze := Maze(input)
	allKeys := maze.AllKeys()
	shortest = math.MaxInt32
	queue := []Route{{Location{maze.Init(), 0}, 0}}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.Steps >= shortest {
			continue
		}

		opt := map[Location]int{}
		for _, r := range queue {
			if opt[r.Location] == 0 {
				opt[r.Location] = r.Steps
			}
			opt[r.Location] = utils.Min(opt[r.Location], r.Steps)
		}

		for next := range maze.bfs(curr) {
			if next.Location.Keys == allKeys && next.Steps < shortest {
				shortest = next.Steps
				continue
			}

			steps := next.Steps
			if opt[next.Location] == 0 {
				opt[next.Location] = steps
			}
			opt[next.Location] = utils.Min(opt[next.Location], steps)
		}

		queue = make([]Route, 0, len(opt))
		for k, v := range opt {
			queue = append(queue, Route{k, v})
		}

		sort.Slice(queue, func(i, j int) bool {
			return queue[i].Steps < queue[j].Steps
		})
	}
	return
}

func Part2(input string) (shortest int) {
	maze := Maze(input).ToRobots()
	allKeys := maze.AllKeys()
	shortest = math.MaxInt32
	queue := []RobotRoute{maze.InitRobots()}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.Steps >= shortest {
			continue
		}

		opt := map[Robots]int{}
		for _, r := range queue {
			if opt[r.Robots] == 0 {
				opt[r.Robots] = r.Steps
			}
			opt[r.Robots] = utils.Min(opt[r.Robots], r.Steps)
		}

		for next := range maze.robotsBfs(curr) {
			if next.Robots.Keys == allKeys && next.Steps < shortest {
				shortest = next.Steps
				continue
			}

			steps := next.Steps
			if opt[next.Robots] == 0 {
				opt[next.Robots] = steps
			}
			opt[next.Robots] = utils.Min(opt[next.Robots], steps)
		}

		queue = make([]RobotRoute, 0, len(opt))
		for k, v := range opt {
			queue = append(queue, RobotRoute{k, v})
		}

		sort.Slice(queue, func(i, j int) bool {
			return queue[i].Steps < queue[j].Steps
		})
	}
	return
}
