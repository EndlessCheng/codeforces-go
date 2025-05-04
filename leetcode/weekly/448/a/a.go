package main

// https://space.bilibili.com/206214
func maxProduct(n int) int {
	mx, mx2 := 0, 0
	for ; n > 0; n /= 10 {
		d := n % 10
		if d > mx {
			mx2 = mx
			mx = d
		} else if d > mx2 {
			mx2 = d
		}
	}
	return mx * mx2
}
