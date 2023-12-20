package day15

import (
	"fmt"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Lens struct {
	Label string
	Value int
	next  *Lens
}

type Box struct{ head *Lens }

func (box *Box) Insert(label string, value int) {
	if box.head == nil {
		box.head = &Lens{label, value, nil}
		return
	}

	var curr *Lens
	for curr = box.head; ; curr = curr.next {
		if curr.Label == label {
			curr.Value = value
			return
		}

		if curr.next == nil {
			break
		}
	}

	curr.next = &Lens{label, value, nil}
}

func (box *Box) Remove(label string) {
	if box.head != nil && box.head.Label == label {
		box.head = box.head.next
		return
	}

	for curr := box.head; curr != nil; curr = curr.next {
		if curr.next != nil && curr.next.Label == label {
			curr.next = curr.next.next
			break
		}
	}
}

func (box Box) Power() (pow int) {
	for i, curr := 1, box.head; curr != nil; i, curr = i+1, curr.next {
		pow += i * curr.Value
	}

	return
}

func (box Box) String() string {
	var output strings.Builder
	for curr := box.head; curr != nil; curr = curr.next {
		output.WriteString(fmt.Sprintf("[%s %d]", curr.Label, curr.Value))
	}
	return output.String()
}

func Hash(text string) (result int) {
	for _, c := range text {
		result = 17 * (result + int(c)) % 256
	}
	return
}

func Part1(input string) (sum int) {
	for _, step := range strings.Split(input, ",") {
		sum += Hash(step)
	}
	return
}

func Part2(input string) (power int) {
	var (
		boxes [256]Box
		label string
		value int
	)

	for _, i := range strings.Split(input, ",") {
		switch {
		case strings.HasSuffix(i, "-"):
			label = strings.Trim(i, "-")
			boxes[Hash(label)].Remove(label)
		default:
			parts := strings.Split(i, "=")
			label, value = parts[0], utils.ToInt(parts[1])
			boxes[Hash(label)].Insert(label, value)
		}
	}

	for n, box := range boxes {
		power += (n + 1) * box.Power()
	}

	return
}
