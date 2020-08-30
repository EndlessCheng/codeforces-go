package main

// github.com/EndlessCheng/codeforces-go
func getMaxLen(a []int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	nz, neg, p := 0, 0, -1
	for i, v := range a {
		if v == 0 {
			nz, neg, p = 0, 0, -1
			continue
		}
		nz++
		if v < 0 {
			neg++
		}
		if neg&1 == 0 {
			ans = max(ans, nz)
		} else if p == -1 {
			p = i
		} else {
			ans = max(ans, i-p)
		}
	}
	return
}
