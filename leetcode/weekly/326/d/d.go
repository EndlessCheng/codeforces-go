package main

import "sort"

// https://space.bilibili.com/206214
const mx = 1e6 + 1
var primes = make([]int, 0, 78500)

func init() {
	np := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
	primes = append(primes, mx, mx) // 保证下面下标不会越界
}

func closestPrimes(left, right int) []int {
	p, q := -1, -1
	for i := sort.SearchInts(primes, left); primes[i+1] <= right; i++ {
		if p < 0 || primes[i+1]-primes[i] < q-p {
			p, q = primes[i], primes[i+1]
		}
	}
	return []int{p, q}
}
