package main

// https://space.bilibili.com/206214
func lastInteger1(n int64) int64 {
	start, d := int64(1), int64(1) // 等差数列首项，公差
	for ; n > 1; n = (n + 1) / 2 {
		start += (n - 2 + n%2) * d
		d *= -2
	}
	return start
}

// https://oeis.org/A090569
func lastInteger(n int64) int64 {
	const mask = 0xAAAAAAAAAAAAAAA // ...1010
	return (n-1)&mask + 1 // 取出 n-1 的从低到高第 2,4,6,... 位，最后再加一（从 1 开始）
}
