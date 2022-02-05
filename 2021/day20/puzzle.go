package day20

import (
	s "strings"
)

type Algorithm [512]bool
type Image string

func parseInput(input string) (Algorithm, Image) {
	algo := Algorithm{}
	parts := s.Split(input, "\n\n")
	for i := 0; i < 511; i++ {
		algo[i] = parts[0][i] == '#'
	}
	return algo, Image(parts[1])
}

func (img Image) Enlarge(times int) Image {
	empty, next := ".", ""
	line := s.Index(string(img), "\n")

	for i := 0; i < times; i++ {
		next += s.Repeat(empty, line+2*times) + "\n"
	}

	for _, l := range s.Split(string(img), "\n") {
		next += s.Repeat(empty, times) + l + s.Repeat(empty, times) + "\n"
	}

	for i := 0; i < times; i++ {
		next += s.Repeat(empty, line+2*times) + "\n"
	}

	return Image(s.Trim(next, "\n"))
}

func (img Image) ValueAt(pos int) (value int) {
	pixel := img[pos]
	l, max := s.Index(string(img), "\n")+1, len(img)

	for i := 0; i < 9; i++ {
		cell := pos - 1 + i%3 + ((i/3)-1)*l

		// pixels outside of the canvas are assumed to have
		// the value of the current one
		curr := pixel
		if cell >= 0 && cell < max && img[cell] != '\n' {
			curr = img[cell]
		}
		if curr == '#' {
			value += 1 << (8 - i)
		}
	}
	return
}

func (img Image) Enhance(algo Algorithm) Image {
	next := ""
	for p, c := range img {
		switch {
		case c == '\n':
			next += "\n"
		case algo[img.ValueAt(p)]:
			next += "#"
		default:
			next += "."
		}
	}
	return Image(next)
}

func (img Image) Lit() (lit int) {
	for _, p := range img {
		if p == '#' {
			lit++
		}
	}
	return
}

func Part1(input string) int {
	algo, img := parseInput(input)

	img = img.Enlarge(2)
	for i := 0; i < 2; i++ {
		img = img.Enhance(algo)
	}

	return img.Lit()
}

func Part2(input string) int {
	algo, img := parseInput(input)

	img = img.Enlarge(50)
	for i := 0; i < 50; i++ {
		img = img.Enhance(algo)
	}

	return img.Lit()
}
