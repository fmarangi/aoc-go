package day13

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func compare(a, b interface{}) int {
	valueA, okA := a.(float64)
	valueB, okB := b.(float64)

	switch {
	case okA && okB:
		return int(valueA - valueB)
	case okA:
		return compare([]interface{}{a}, b)
	case okB:
		return compare(a, []interface{}{b})
	}

	// Both are lists
	x, y := a.([]interface{}), b.([]interface{})
	for i, n, m := 0, len(x), len(y); i < n && i < m; i++ {
		if res := compare(x[i], y[i]); res != 0 {
			return res
		}
	}

	// The shortest list wins
	return len(x) - len(y)
}

func Part1(input string) (indices int) {
	nums := parseInput(input)
	for i, n := 0, len(nums); i < n; i += 2 {
		if compare(nums[i], nums[i+1]) < 0 {
			indices += i/2 + 1
		}
	}
	return
}

func Part2(input string) (key int) {
	nums := parseInput(input + " [[2]] [[6]]")
	sort.Slice(nums, func(i, j int) bool {
		return compare(nums[i], nums[j]) < 0
	})

	key = 1
	for i, n := range nums {
		if v := fmt.Sprint(n); v == "[[2]]" || v == "[[6]]" {
			key *= i + 1
		}
	}

	return
}

func parseInput(input string) (nums []interface{}) {
	for _, n := range strings.Fields(input) {
		var num interface{}
		json.Unmarshal([]byte(n), &num)
		nums = append(nums, num)
	}
	return
}
