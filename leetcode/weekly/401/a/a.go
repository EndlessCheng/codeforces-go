package main

// https://space.bilibili.com/206214
func numberOfChild(n, k int) int {
	t := k % (n - 1)
	if k/(n-1)%2 > 0 {
		return n - t - 1
	}
	return t
}
