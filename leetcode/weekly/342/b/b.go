package main

// https://space.bilibili.com/206214
func s(n, m int) int {
	return n / m * (n/m + 1) / 2 * m
}

func sumOfMultiples(n int) int {
	return s(n, 3) + s(n, 5) + s(n, 7) - s(n, 15) - s(n, 21) - s(n, 35) + s(n, 105)
}
