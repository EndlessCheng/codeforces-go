package main

// https://space.bilibili.com/206214
func numberOfCuts(n int) int {
	if n == 1 || n%2 == 0 {
		return n / 2
	}
	return n
}
