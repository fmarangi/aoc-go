package utils

import (
	"math"
	"strconv"
)

func ToInt(num string) (value int) {
	value, _ = strconv.Atoi(num)
	return
}

func Ints(nums []string) (ints []int) {
	for _, n := range nums {
		ints = append(ints, ToInt(n))
	}
	return
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

func Sum(nums []int) (sum int) {
	for _, n := range nums {
		sum += n
	}
	return
}

func Max(nums ...int) int {
	max := math.MinInt64
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func Min(nums ...int) int {
	min := math.MaxInt32
	for _, n := range nums {
		if n < min {
			min = n
		}
	}
	return min
}
