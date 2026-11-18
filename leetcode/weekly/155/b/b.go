package main

import "sort"

func nthUglyNumber(n, a, b, c int) int {
	lcmAB := lcm(a, b)
	lcmAC := lcm(a, c)
	lcmBC := lcm(b, c)
	lcmABC := lcm(lcmAB, c)
	right := min(min(a, b, c)*n, 2e9)
	return sort.Search(right, func(x int) bool {
		return x/a+x/b+x/c-x/lcmAB-x/lcmAC-x/lcmBC+x/lcmABC >= n
	})
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}
