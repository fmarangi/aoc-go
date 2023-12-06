package day05

import (
	"math"
	"sort"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

type Range struct {
	Start, Length int
}

func (r Range) End() int {
	return r.Start + r.Length
}

func (r Range) Empty() bool {
	return r.Length == 0
}

func (r Range) Contains(value int) bool {
	return value >= r.Start && value < r.End()
}

func (a Range) Overlap(b Range) int {
	if a.Contains(b.Start) {
		if a.End() < b.End() {
			return a.End() - b.Start
		}
		return b.End() - b.Start
	}
	return 0
}

type Conversion struct {
	From, To Range
}

type Section []Conversion

func (s Section) Convert(value int) int {
	for _, c := range s {
		if c.From.Contains(value) {
			return value - c.From.Start + c.To.Start
		}
	}
	return value
}

func (s Section) ConvertRange(r Range) (Range, Range) {
	for i, n := 0, len(s); i < n; i++ {
		if c := s[i]; c.From.Contains(r.Start) {
			var converted, remainder Range
			converted.Start = r.Start - c.From.Start + c.To.Start
			converted.Length = c.From.Overlap(r)

			if converted.Length < r.Length {
				remainder.Start = c.From.End()
				remainder.Length = r.Length - converted.Length
			}

			return converted, remainder
		}
	}
	return r, Range{}
}

func Part1(input string) (lowest int) {
	seeds, sections := parseInput(input)

	lowest = math.MaxInt32
	for _, seed := range seeds {
		for _, section := range sections {
			seed = section.Convert(seed)
		}

		if seed < lowest {
			lowest = seed
		}
	}

	return
}

func Part2(input string) (lowest int) {
	var (
		seeds, sections = parseInput(input)
		ranges          = toRanges(seeds)
	)

	for _, s := range sections {
		var next []Range

		for len(ranges) > 0 {
			c, r := s.ConvertRange(ranges[0])

			ranges = ranges[1:]
			next = append(next, c)

			if !r.Empty() {
				ranges = append(ranges, r)
			}
		}

		ranges = next
	}

	lowest = math.MaxInt32
	for _, r := range ranges {
		if r.Start < lowest {
			lowest = r.Start
		}
	}

	return
}

func parseInput(input string) (seeds []int, sections []Section) {
	parts := strings.Split(input, "\n\n")

	seeds = utils.Ints(strings.Fields(parts[0][7:]))
	for _, group := range parts[1:] {
		var s Section

		ranges := strings.Split(group, "\n")
		for _, r := range ranges[1:] {
			values := utils.Ints(strings.Fields(r))
			s = append(s, Conversion{
				Range{values[1], values[2]},
				Range{values[0], values[2]},
			})
		}

		sort.Slice(s, func(i, j int) bool {
			return s[i].From.Start < s[j].To.Start
		})

		sections = append(sections, s)
	}

	return
}

func toRanges(seeds []int) (ranges []Range) {
	for i, n := 0, len(seeds); i < n; i += 2 {
		ranges = append(ranges, Range{seeds[i], seeds[i+1]})
	}
	return
}
