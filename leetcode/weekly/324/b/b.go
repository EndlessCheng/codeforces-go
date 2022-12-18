package main

// https://space.bilibili.com/206214
func smallestValue(n int) int {
	for {
		x, s := n, 0
		for i := 2; i*i <= x; i++ {
			for ; x%i == 0; x /= i {
				s += i
			}
		}
		if x > 1 {
			s += x
		}
		if s == n {
			return n
		}
		n = s
	}
}
