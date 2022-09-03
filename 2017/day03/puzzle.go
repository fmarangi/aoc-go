package day03

import "github.com/fmarangi/aoc-go/utils"

const (
	East int = iota
	North
	West
	South
)

type Pos struct{ X, Y int }

func (a Pos) Manhattan(b Pos) int {
	return utils.Abs(a.X-b.X) + utils.Abs(a.Y-b.Y)
}

func (a Pos) Move(b Pos) Pos {
	return Pos{a.X + b.X, a.Y + b.Y}
}

var offsets = map[int]Pos{
	East:  {1, 0},
	North: {0, 1},
	West:  {-1, 0},
	South: {0, -1},
}

func Generate() <-chan Pos {
	var (
		out       = make(chan Pos)
		curr      Pos
		dir, next = East, North
		o, nextO  = offsets[dir], offsets[next]
		seen      = map[Pos]bool{curr: true}
	)

	go func() {
		out <- curr
		curr = curr.Move(o)
		seen[curr] = true
		out <- curr
		for {
			for seen[curr.Move(nextO)] {
				curr = curr.Move(o)
				out <- curr
				seen[curr] = true
			}
			dir, next = (dir+1)%4, (next+1)%4
			o, nextO = offsets[dir], offsets[next]
		}
	}()

	return out
}

func GenerateSums() <-chan int {
	var (
		out  = make(chan int)
		grid = map[Pos]int{{}: 1}
	)

	go func() {
		for next := range Generate() {
			sum := 0
			for i := 0; i < 9; i++ {
				x, y := (i/3)-1, (i%3)-1
				sum += grid[Pos{next.X + x, next.Y + y}]
			}
			grid[next] = sum
			out <- grid[next]
		}
	}()

	return out
}

func Part1(input string) (valid int) {
	cells, pos := Generate(), Pos{}
	for i, n := 0, utils.ToInt(input); i < n; i++ {
		pos = <-cells
	}
	return pos.Manhattan(Pos{0, 0})
}

func Part2(input string) (larger int) {
	sums := GenerateSums()
	target := utils.ToInt(input)
	for larger < target {
		larger = <-sums
	}
	return
}
