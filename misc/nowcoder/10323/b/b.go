package main

// github.com/EndlessCheng/codeforces-go
func wwork(n int) (c int64) {
	// or 母函数 1/(1-x^3) => C(n+2,2)
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			v := n - i - j
			if v < 0 {
				continue
			}
			x := (v + 4) * (v + 4)
			if x%20 > 9 {
				x = x/20 + 1
			} else {
				x /= 20
			}
			c += int64(x)
		}
	}
	return
}
