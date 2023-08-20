package main

// https://space.bilibili.com/206214
func minimumSum(n, k int) int {
	m := min(k/2, n)
	return (m*(m+1) + (k*2+n-m-1)*(n-m)) / 2
}

func min(a, b int) int { if b < a { return b }; return a }
