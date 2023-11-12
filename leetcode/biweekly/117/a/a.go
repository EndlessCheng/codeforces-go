package main

// https://space.bilibili.com/206214
func c2(n int) int {
	if n < 2 {
		return 0
	}
	return n * (n - 1) / 2
}

func distributeCandies(n int, limit int) int {
	return c2(n+2) - 3*c2(n-limit+1) + 3*c2(n-2*limit) - c2(n-3*limit-1)
}
