package main

// https://space.bilibili.com/206214
func checkDivisibility(n int) bool {
	s, m := 0, 1
	for x := n; x > 0; x /= 10 {
		d := x % 10
		s += d
		m *= d
	}
	return n%(s+m) == 0
}
