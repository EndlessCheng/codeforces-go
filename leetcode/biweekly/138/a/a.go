package main

// https://space.bilibili.com/206214
func generateKey(x, y, z int) (ans int) {
	for pow10 := 1; x > 0 && y > 0 && z > 0; pow10 *= 10 {
		ans += min(x%10, y%10, z%10) * pow10
		x /= 10
		y /= 10
		z /= 10
	}
	return
}
