package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func maximumScore(a int, b int, c int) (ans int) {
	sort3 := func(a ...int) (x, y, z int) { sort.Ints(a); return a[0], a[1], a[2] }
	for {
		a, b, c = sort3(a, b, c)
		if b == 0 {
			break
		}
		b--
		c--
		ans++
	}
	return
}
