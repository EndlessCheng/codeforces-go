package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func nthMagicalNumber(n, a, b int) int {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := a / gcd(a, b) * b
	return sort.Search(4e13, func(x int) bool { return x/a+x/b-x/lcm >= n }) % (1e9 + 7)
}
