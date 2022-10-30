package main

// https://space.bilibili.com/206214
func makeIntegerBeautiful(n int64, target int) int64 {
	for tail := int64(1); ; tail *= 10 {
		m := n + (tail-n%tail)%tail // 进位后的数字
		sum := 0
		for x := m; x > 0; x /= 10 {
			sum += int(x % 10)
		}
		if sum <= target {
			return m - n
		}
	}
}
