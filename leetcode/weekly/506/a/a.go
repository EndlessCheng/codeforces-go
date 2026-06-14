package main

// https://space.bilibili.com/206214
func checkGoodInteger(n int) bool {
	diff := 0
	for ; n > 0; n /= 10 {
		d := n % 10
		diff += d*d - d
	}
	return diff >= 50
}
