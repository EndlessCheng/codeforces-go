package main

// https://space.bilibili.com/206214
func minimumPossibleSum(n, k int) int64 {
	m := min(k/2, n)
	return int64((m*(m+1) + (k*2+n-m-1)*(n-m)) / 2)
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
