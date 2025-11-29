package main

import "math"

// github.com/EndlessCheng/codeforces-go
func countTriples1(n int) (ans int) {
	for a := 1; a < n; a++ {
		for b := 1; b < a && a*a+b*b <= n*n; b++ {
			c2 := a*a + b*b
			rt := int(math.Sqrt(float64(c2)))
			if rt*rt == c2 {
				ans++
			}
		}
	}
	return ans * 2 // (a,b,c) 和 (b,a,c) 各算一次
}

func countTriples(n int) (ans int) {
	for u := 3; u*u < n*2; u += 2 {
		for v := 1; v < u && (u*u+v*v)/2 <= n; v += 2 {
			if gcd(u, v) == 1 {
				c0 := (u*u + v*v) / 2
				ans += n / c0
			}
		}
	}
	return ans * 2
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
