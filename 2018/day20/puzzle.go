package day20

import "strings"

type Room struct{ X, Y int }

type Path struct {
	Room  Room
	Steps int
}

func (r Room) Next(dir byte) Room {
	o := map[byte]Room{
		'N': {0, 1},
		'S': {0, -1},
		'W': {-1, 0},
		'E': {1, 0},
	}
	return Room{r.X + o[dir].X, r.Y + o[dir].Y}
}

func Explore(maze string) map[Room]int {
	var (
		reader = strings.NewReader(maze)
		curr   Path
		queue  = []Path{}
		seen   = map[Room]int{curr.Room: 0}
	)

	for {
		dir, err := reader.ReadByte()
		if err != nil {
			break
		}

		switch dir {
		case 'N', 'S', 'W', 'E':
			curr = Path{curr.Room.Next(dir), curr.Steps + 1}
		case '(':
			queue = append(queue, curr)
		case '|':
			curr = queue[len(queue)-1]
		case ')':
			curr = queue[len(queue)-1]
			queue = queue[:len(queue)-1]
		}

		if _, ok := seen[curr.Room]; !ok {
			seen[curr.Room] = curr.Steps
		}
	}

	return seen
}

func Part1(input string) (max int) {
	for _, steps := range Explore(input) {
		if steps > max {
			max = steps
		}
	}
	return
}

func Part2(input string) (rooms int) {
	for _, steps := range Explore(input) {
		if steps >= 1000 {
			rooms++
		}
	}
	return
}
