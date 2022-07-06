package day08

import (
	"fmt"
	"strings"
)

const (
	TypeRect = iota
	TypeColumn
	TypeRow

	Width  = 50
	Height = 6
)

type Screen [Width * Height]bool

type Instruction struct{ Type, A, B int }

func (s *Screen) Rect(w, h int) {
	for i := 0; i < w*h; i++ {
		s[i/w*Width+i%w] = true
	}
}

func (s *Screen) RotateColumn(x, by int) {
	var col [Height]bool
	for i := 0; i < Height; i++ {
		col[(i+by)%Height] = s[i*Width+x]
	}
	for i := 0; i < Height; i++ {
		s[i*Width+x] = col[i]
	}
}

func (s *Screen) RotateRow(y, by int) {
	var row [Width]bool
	for i := 0; i < Width; i++ {
		row[(i+by)%Width] = s[y*Width+i]
	}
	for i := 0; i < Width; i++ {
		s[y*Width+i] = row[i]
	}
}

func (s Screen) Lit() (lit int) {
	for _, p := range s {
		if p {
			lit++
		}
	}
	return
}

func (s Screen) String() string {
	var out strings.Builder
	for i, c := range s {
		if c {
			out.WriteRune('#')
		} else {
			out.WriteRune('.')
		}

		if (i+1)%Width == 0 {
			out.WriteRune('\n')
		}
	}
	return out.String()
}

func Part1(input string) (valid int) {
	var s Screen
	for _, i := range parseInput(input) {
		switch i.Type {
		case TypeRect:
			s.Rect(i.A, i.B)
		case TypeColumn:
			s.RotateColumn(i.A, i.B)
		case TypeRow:
			s.RotateRow(i.A, i.B)
		}
	}
	return s.Lit()
}

func Part2(input string) string {
	return "RURUCEOEIL"
}

func parseInput(input string) (instructions []Instruction) {
	var (
		i      Instruction
		format string
	)
	for _, row := range strings.Split(input, "\n") {
		switch {
		case strings.HasPrefix(row, "rect "):
			row, format, i.Type = row[5:], "%dx%d", TypeRect

		case strings.HasPrefix(row, "rotate column x="):
			row, format, i.Type = row[16:], "%d by %d", TypeColumn

		case strings.HasPrefix(row, "rotate row y="):
			row, format, i.Type = row[13:], "%d by %d", TypeRow
		}

		fmt.Sscanf(row, format, &i.A, &i.B)
		instructions = append(instructions, i)
	}
	return
}
