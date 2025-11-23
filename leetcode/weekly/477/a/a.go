package main

// https://space.bilibili.com/206214
func sumAndMultiply(n int) int64 {
	x, sum, pow10 := 0, 0, 1
	for ; n > 0; n /= 10 {
		d := n % 10
		if d > 0 {
			x += d * pow10
			sum += d
			pow10 *= 10
		}
	}
	return int64(x) * int64(sum)
}
