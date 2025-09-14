package main

// https://space.bilibili.com/206214
func countStableSubsequences(nums []int) int {
	const mod = 1_000_000_007
	f := [2][2]int{}
	for _, x := range nums {
		x %= 2
		f[x][1] = (f[x][1] + f[x][0]) % mod
		f[x][0] = (f[x][0] + f[x^1][0] + f[x^1][1] + 1) % mod
	}
	return (f[0][0] + f[0][1] + f[1][0] + f[1][1]) % mod
}
