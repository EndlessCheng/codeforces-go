package main

import "strconv"

func findNumbers(nums []int) (ans int) {
	for _, x := range nums {
		for x >= 100 {
			x /= 100
		}
		if x >= 10 {
			ans++
		}
	}
	return
}

func findNumbers1(nums []int) int {
	ans := len(nums)
	for _, x := range nums {
		ans -= len(strconv.Itoa(x)) % 2
	}
	return ans
}
