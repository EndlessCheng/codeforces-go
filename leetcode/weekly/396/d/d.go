package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func minCostToEqualizeArray(nums []int, c1 int, c2 int) int {
	const mod = 1_000_000_007
	n := len(nums)
	m := slices.Min(nums)
	M := slices.Max(nums)
	base := n * M
	for _, x := range nums {
		base -= x
	}
	if n <= 2 || c1*2 <= c2 {
		return base * c1 % mod
	}

	f := func(x int) int {
		s := base + (x-M)*n
		d := x - m
		if d*2 <= s {
			return s/2*c2 + s%2*c1
		}
		return (s-d)*c2 + (d*2-s)*c1
	}

	i := (n*M - m*2 - base + n - 3) / (n - 2)
	if i <= M {
		return min(f(M), f(M+1)) % mod
	}
	return min(f(M), f(i-1), f(i), f(i+1)) % mod
}
