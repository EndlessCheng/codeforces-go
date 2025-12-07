package main

import "strconv"

// https://space.bilibili.com/206214
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}

func completePrime(num int) bool {
	s := strconv.Itoa(num)
	for i := range len(s) {
		// 前缀
		x, _ := strconv.Atoi(s[:i+1])
		if !isPrime(x) {
			return false
		}

		// 后缀
		x, _ = strconv.Atoi(s[i:])
		if !isPrime(x) {
			return false
		}
	}
	return true
}
