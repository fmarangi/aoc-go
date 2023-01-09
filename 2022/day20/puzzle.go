package day20

import (
	"container/ring"
	"strings"

	"github.com/fmarangi/aoc-go/utils"
)

func Ring(n int) (r *ring.Ring) {
	r = ring.New(n)
	for i := 0; i < n; i++ {
		r.Value, r = i, r.Next()
	}
	return
}

func Shuffle(nums []int, times int) (sum int) {
	var (
		qty  = len(nums)
		ring = Ring(qty)
	)

	for j := 0; j < qty*times; j++ {
		i, n := j%qty, nums[j%qty]
		for ring.Value != i {
			ring = ring.Next()
		}

		ring = ring.Prev()
		s := ring.Unlink(1)
		ring.Move(n % (qty - 1)).Link(s)
	}

	for nums[ring.Value.(int)] != 0 {
		ring = ring.Next()
	}

	for i := 0; i < 3; i++ {
		ring = ring.Move(1000)
		sum += nums[ring.Value.(int)]
	}

	return
}

func Part1(input string) int {
	return Shuffle(parseInput(input), 1)
}

func Part2(input string) int {
	nums := parseInput(input)
	for i := range nums {
		nums[i] *= 811589153
	}
	return Shuffle(nums, 10)
}

func parseInput(input string) []int {
	return utils.Ints(strings.Fields(input))
}
