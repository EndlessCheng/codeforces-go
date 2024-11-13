package main

import "math"

// https://space.bilibili.com/206214
func numberOfWays(n, x int) int {
	f := make([]int, n+1)
	f[0] = 1
	for i := 1; pow(i, x) <= n; i++ {
		v := pow(i, x)
		for s := n; s >= v; s-- {
			f[s] += f[s-v]
		}
	}
	return f[n] % 1_000_000_007
}

// 本题数据范围小，计算结果一定准确
func pow(i, x int) int {
	return int(math.Pow(float64(i), float64(x)))
}
