package utils

import "strconv"

func ToInt(num string) (value int) {
	value, _ = strconv.Atoi(num)
	return
}

func Abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
