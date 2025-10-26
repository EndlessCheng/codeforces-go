package main

// https://space.bilibili.com/206214
func removeZeros(n int64) (ans int64) {
	pow10 := int64(1)
	for ; n > 0; n /= 10 {
		d := n % 10
		if d > 0 {
			ans += d * pow10
			pow10 *= 10
		}
	}
	return
}
