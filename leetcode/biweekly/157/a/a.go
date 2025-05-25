package main

import "slices"

// https://space.bilibili.com/206214
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2
}

func sumOfLargestPrimes(s string) (ans int64) {
	primes := []int{}
	n := len(s)
	for i := range n {
		x := 0
		for j := i; j < n; j++ {
			x = x*10 + int(s[j]-'0')
			if isPrime(x) {
				primes = append(primes, x)
			}
		}
	}

	slices.Sort(primes)
	primes = slices.Compact(primes) // 去重

	for _, p := range primes[max(len(primes)-3, 0):] {
		ans += int64(p)
	}
	return
}
