package main

import "slices"

// https://space.bilibili.com/206214
func decimalRepresentation(n int) (ans []int) {
	pow10 := 1
	for ; n > 0; n /= 10 {
		d := n % 10
		if d > 0 {
			ans = append(ans, d*pow10)
		}
		pow10 *= 10
	}
	slices.Reverse(ans)
	return
}
