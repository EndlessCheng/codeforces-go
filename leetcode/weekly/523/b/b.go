package main

import "sort"

// https://space.bilibili.com/206214
const mx = 1_000_001

var np = [mx]bool{}
var primes []int
var sum []int

func init() {
	s := 0
	for i := 2; i < mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			s += i
			sum = append(sum, s)
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
}

func maxPrimes(n int, s int) []int {
	i := sort.SearchInts(primes, n+1)
	j := sort.SearchInts(sum[:i], s+1)
	return primes[:j]
}
