package main

import "sort"

// https://space.bilibili.com/206214
func makeSubKSumEqual(arr []int, k int) (ans int64) {
	k = gcd(k, len(arr))
	g := make([][]int, k)
	for i, x := range arr {
		g[i%k] = append(g[i%k], x)
	}
	for _, b := range g {
		sort.Ints(b)
		for _, x := range b {
			ans += int64(abs(x - b[len(b)/2]))
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
