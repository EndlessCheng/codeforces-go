package main

// https://space.bilibili.com/206214
func countDistinct(n int64) (ans int64) {
	pow9 := int64(1)
	for ; n > 0; n /= 10 {
		d := n % 10
		if d == 0 {
			ans = 0
		} else {
			if pow9 > 1 {
				d--
			}
			ans += d * pow9
		}
		pow9 *= 9
	}
	return ans + (pow9-9)/8
}
