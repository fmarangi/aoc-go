package day16

type Beam struct{ P, dir int }

func (b Beam) Move(cell byte) (next []int) {
	switch cell {
	case '.':
		next = append(next, b.dir)

	case '/':
		move := map[int]int{
			Right: Up,
			Down:  Left,
			Left:  Down,
			Up:    Right,
		}
		next = append(next, move[b.dir])

	case '\\':
		move := map[int]int{
			Right: Down,
			Down:  Right,
			Left:  Up,
			Up:    Left,
		}
		next = append(next, move[b.dir])

	case '-':
		if b.dir == Right || b.dir == Left {
			next = append(next, b.dir)
		} else {
			next = append(next, Right)
			next = append(next, Left)
		}

	case '|':
		if b.dir == Up || b.dir == Down {
			next = append(next, b.dir)
		} else {
			next = append(next, Up)
			next = append(next, Down)
		}
	}

	return
}
