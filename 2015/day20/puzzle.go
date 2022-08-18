package day20

import "github.com/fmarangi/aoc-go/utils"

func Part1(input string) (house int) {
	target := utils.ToInt(input) / 10
	for house = 665000; divisorSum(house) < target; house++ {
	}
	return
}

func Part2(input string) (house int) {
	target := utils.ToInt(input) / 11
	for house = 700000; divisorSum50(house) < target; house++ {
	}
	return
}

func divisorSum(num int) (sum int) {
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			sum += i
			if div := num / i; div != i {
				sum += div
			}
		}
	}
	return
}

func divisorSum50(num int) (sum int) {
	for i := 1; i*i <= num; i++ {
		if num%i == 0 {
			if num/i <= 50 {
				sum += i
			}

			if div := num / i; div != i && num/div <= 50 {
				sum += div
			}
		}
	}
	return
}
