package main

import "strconv"

func subtractProductAndSum(n int) int {
	mul, sum := 1, 0
	for _, c := range strconv.Itoa(n) {
		c -= '0'
		mul *= int(c)
		sum += int(c)
	}
	return mul - sum
}
