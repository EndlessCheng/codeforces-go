package main

// github.com/EndlessCheng/codeforces-go
const mx = 250

var c2 [mx*mx + 1]bool

func init() {
	for i := 1; i <= mx; i++ {
		c2[i*i] = true
	}
}

func countTriples(n int) (ans int) {
	for i := 1; i <= n; i++ {
		for j := 1; j < i; j++ {
			if s := i*i + j*j; s <= n*n && c2[s] {
				ans += 2
			}
		}
	}
	return
}
