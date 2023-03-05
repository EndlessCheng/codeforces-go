package main

// https://space.bilibili.com/206214
func passThePillow(n, time int) int {
	t := time % (n - 1)
	if time/(n-1)%2 > 0 {
		return n - t
	}
	return 1 + t
}
