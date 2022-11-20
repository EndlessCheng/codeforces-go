package main

import "strings"

// https://space.bilibili.com/206214
func beautifulPartitions(s string, k, l int) (ans int) {
	const mod int = 1e9 + 7
	isPrime := func(c byte) bool { return strings.IndexByte("2357", c) >= 0 }
	n := len(s)
	if l == 1 {
		l = 2
	}
	if k*l > n || !isPrime(s[0]) || isPrime(s[n-1]) { // 剪枝
		return
	}
	// 判断是否可以在 j-1 和 j 之间分割（开头和末尾也算）
	canPartition := func(j int) bool { return j == 0 || j == n || !isPrime(s[j-1]) && isPrime(s[j]) }
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, n+1)
	}
	f[0][0] = 1 // 假设开头有一条分割线
	for i := 1; i <= k; i++ {
		sum := 0
		for j := i * l; j+(k-i)*l <= n; j++ {
			if canPartition(j - l) { // 长度至少是 l
				sum = (sum + f[i-1][j-l]) % mod
			}
			if canPartition(j) {
				f[i][j] = sum
			}
		}
	}
	return f[k][n]
}
