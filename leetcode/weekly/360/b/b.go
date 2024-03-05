package main

// https://space.bilibili.com/206214
func minimumPossibleSum(n, k int) int {
	m := min(k/2, n)
	return (m*(m+1) + (k*2+n-m-1)*(n-m)) / 2 % 1_000_000_007
}
