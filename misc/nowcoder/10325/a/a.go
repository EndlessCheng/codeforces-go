package main

// github.com/EndlessCheng/codeforces-go
const mod = 1e9 + 7

func countTriplets(a []int, A, B int) (s int) {
	for i := 1; i < len(a)-1; i++ {
		v, c1, c2 := a[i], 0, 0
		for _, w := range a[:i] {
			if abs(v-w) <= A {
				c1++
			}
		}
		for _, w := range a[i+1:] {
			if abs(v-w) <= B {
				c2++
			}
		}
		s += c1 * c2 % mod
	}
	return s % mod
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
