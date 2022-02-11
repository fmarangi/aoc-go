package utils

import "strconv"

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
