package main

import "slices"

// https://space.bilibili.com/206214
const mx = 100_000

var np = [mx + 1]bool{true, true}
var primeNumbers []int

func init() {
	for i := 2; i <= mx; i++ {
		if !np[i] {
			primeNumbers = append(primeNumbers, i)
			for j := i * i; j <= mx; j += i {
				np[j] = true
			}
		}
	}
}

func isPrime(n int) bool {
	if n <= mx {
		return !np[n]
	}
	for _, p := range primeNumbers {
		if p*p > n {
			break
		}
		if n%p == 0 {
			return false
		}
	}
	return true
}

func sumOfLargestPrimes(s string) (ans int64) {
	primes := []int{}
	n := len(s)
	for i := range n {
		x := 0
		for _, b := range s[i:] {
			x = x*10 + int(b-'0')
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
