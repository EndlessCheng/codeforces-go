package main

import "math"

// https://space.bilibili.com/206214
func maxBottlesDrunk1(numBottles, numExchange int) (ans int) {
	for numBottles >= numExchange {
		ans += numExchange
		numBottles -= numExchange - 1
		numExchange++
	}
	return ans + numBottles
}

func maxBottlesDrunk(n, e int) int {
	b := e*2 - 1
	k := (int(math.Sqrt(float64(b*b+(n-e)*8))) - b + 2) / 2
	return n + k
}
