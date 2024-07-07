package main

import "fmt"

// https://space.bilibili.com/206214
func validStrings(n int) (ans []string) {
	for i := 0; i < 1<<n; i++ {
		x := 1<<n - 1 ^ i
		if x>>1&x == 0 {
			ans = append(ans, fmt.Sprintf("%0*b", n, i))
		}
	}
	return
}
