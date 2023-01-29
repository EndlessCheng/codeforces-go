package main

// https://space.bilibili.com/206214
const mod int = 1e9 + 7

func monkeyMove(n int) int {
	return (pow(2, n) - 2 + mod) % mod
}

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
