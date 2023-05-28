package main

import "math"

// https://space.bilibili.com/206214
func buyChoco(prices []int, money int) int {
	mn1, mn2 := math.MaxInt, math.MinInt
	for _, p := range prices {
		if p < mn1 {
			mn2 = mn1
			mn1 = p
		} else if p < mn2 {
			mn2 = p
		}
	}
	if mn1+mn2 <= money {
		return money - mn1 - mn2
	}
	return money
}
