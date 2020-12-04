package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func string2(K int, s string) (ans int) {
	d := make([]int, len(s))
	for t := 'a'; t <= 'z'; t++ {
		for i, b := range s {
			d[i] = int(abs(t - b))
		}
		sort.Ints(d)
		c, k := 0, K
		for _, v := range d {
			if v <= k {
				k -= v
				c++
			} else {
				break
			}
		}
		if c > ans {
			ans = c
		}
	}
	return
}

func abs(x rune) rune {
	if x < 0 {
		return -x
	}
	return x
}
