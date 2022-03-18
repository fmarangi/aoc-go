package day18

import "github.com/fmarangi/aoc-go/utils"

type Element struct{ Value, Level int }
type Number []Element

func parseNumber(input string) (num Number) {
	for i, n, level := 0, len(input), -1; i < n; i++ {
		switch input[i] {
		case '[':
			level++
		case ']':
			level--
		case ',':
			continue
		default:
			j := i
			if input[i+1] >= '0' && input[i+1] <= '9' {
				i++
			}
			num = append(num, Element{utils.ToInt(input[j : i+1]), level})
		}
	}
	return
}

func (num Number) Magnitude() int {
	// Make sure the number is fully reduced
	num = num.Reduce()

	// The highest level should be 3
	for l := 3; l >= 0; l-- {
		next := Number{}
		for i, n := 0, len(num); i < n; i++ {
			if num[i].Level == l && num[i].Level == num[i+1].Level {
				value := num[i].Value*3 + num[i+1].Value*2
				next = append(next, Element{value, l - 1})
				i++
				continue
			}
			next = append(next, num[i])
		}
		num = next
	}

	return num[0].Value
}

func Sum(a, b Number) (c Number) {
	for _, p := range [...]Number{a, b} {
		for _, e := range p {
			c = append(c, Element{e.Value, e.Level + 1})
		}
	}
	return c.Reduce()
}

func (num Number) Reduce() Number {
	reduced, split := false, -1
	for !reduced {
		reduced, split = true, -1
		for i, n := 0, len(num); i < n; i++ {
			// Look for a pair to explode
			if num[i].Level > 3 {
				num = num.explode(i)
				reduced, split = false, -1
				break
			}

			if split < 0 && num[i].Value > 9 {
				split = i
			}
		}

		// No explode, but maybe we can split...
		if split >= 0 {
			num = num.split(split)
			reduced = false
		}
	}
	return num
}

func (num Number) explode(pos int) Number {
	next := make(Number, 0, len(num)-1)
	next = append(next, num[:pos]...)
	next = append(next, Element{Level: num[pos].Level - 1})
	next = append(next, num[pos+2:]...)
	if pos >= 1 {
		next[pos-1].Value += num[pos].Value
	}
	if pos+1 < len(next) {
		next[pos+1].Value += num[pos+1].Value
	}
	return next
}

func (num Number) split(pos int) Number {
	a, b, l := num[pos].Value/2, (num[pos].Value+1)/2, num[pos].Level+1
	next := make(Number, 0, len(num)+1)
	next = append(next, num[:pos]...)
	next = append(next, Element{a, l}, Element{b, l})
	next = append(next, num[pos+1:]...)
	return next
}
