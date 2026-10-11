package main

// https://space.bilibili.com/206214
var fibSet = map[int]bool{}

func init() {
	a, b := 0, 1
	for b < 1e9 {
		a, b = b, a+b
		fibSet[b] = true
	}
}

func threeFibonacciSum(n int) bool {
	return n%2 == 0 && fibSet[n/2]
}
