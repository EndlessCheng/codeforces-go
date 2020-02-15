package main

import "sort"

func nthUglyNumber(n int, a int, b int, c int) int {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	lcm := func(a, b int) int { return a / gcd(a, b) * b }
	return sort.Search(2e9+1, func(x int) bool { return x/a+x/b+x/c-x/lcm(a, b)-x/lcm(b, c)-x/lcm(a, c)+x/lcm(lcm(a, b), c) >= n })
}
