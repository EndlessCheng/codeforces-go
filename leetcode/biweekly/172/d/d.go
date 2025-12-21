package main

// https://space.bilibili.com/206214
func lastInteger(n int64) int64 {
	start, d := int64(1), int64(1) // 等差数列首项，公差
	for ; n > 1; n = (n + 1) / 2 {
		start += (n - 2 + n%2) * d
		d *= -2
	}
	return start
}
