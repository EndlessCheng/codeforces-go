package main

// github.com/EndlessCheng/codeforces-go
func mirrorReflection(p int, q int) (ans int) {
	if q == 0 {
		return
	}
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	g := gcd(p, q)
	p /= g
	q /= g
	if p&1 == 0 {
		return 2
	}
	return q & 1
}
