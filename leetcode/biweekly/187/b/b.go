package main

// https://space.bilibili.com/206214
func maximumValue(n, s, m int) int64 {
	if n == 1 {
		return int64(s)
	}
	return int64(s + m + (m-1)*(n/2-1))
}
