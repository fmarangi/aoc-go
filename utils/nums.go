package utils

import "strconv"

func ToInt(num string) (value int) {
	value, _ = strconv.Atoi(num)
	return
}
