package main

// https://space.bilibili.com/206214
func validDigit(n, x int) bool {
	hasX := false
	for ; n >= 10; n /= 10 {
		if n%10 == x {
			hasX = true
		}
	}
	return hasX && n != x
}
