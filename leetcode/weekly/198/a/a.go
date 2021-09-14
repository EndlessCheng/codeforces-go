package main

func numWaterBottles(n, e int) (ans int) {
	for ; n >= e; n -= e - 1 {
		ans += e
	}
	ans += n
	return
}
