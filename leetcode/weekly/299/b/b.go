package main

// https://space.bilibili.com/206214/dynamic
const mod = 1_000_000_007

var f = [1e4 + 1]int{1, 2}

func init() {
	for i := 2; i <= 10_000; i++ {
		f[i] = (f[i-1] + f[i-2]) % mod
	}
}

func countHousePlacements(n int) int {
	return f[n] * f[n] % mod
}
