package main

// github.com/EndlessCheng/codeforces-go
func checkPartitioning(s string) (ans bool) {
	var maxLen []int
	manacher := func(origin []byte) {
		n := len(origin)
		m := 2*n + 2
		s := make([]byte, m+1)
		s[0] = '^'
		for i, c := range origin {
			s[2*i+1] = '#'
			s[2*i+2] = c
		}
		s[2*n+1] = '#'
		s[2*n+2] = '$'
		maxLen = make([]int, m+1)
		mid, r := 0, 0
		for i := 2; i < m; i++ {
			mx := 1
			if i < r {
				mx = min(maxLen[2*mid-i], r-i)
			}
			for ; s[i-mx] == s[i+mx]; mx++ {
			}
			if i+mx > r {
				mid, r = i, i+mx
			}
			maxLen[i] = mx
		}
	}
	// [l,r]  0<=l<=r<n
	isP := func(l, r int) bool { return maxLen[l+r+2]-1 >= r-l+1 }

	manacher([]byte(s))
	n := len(s)
	for i := 1; i < n-1; i++ {
		if !isP(0, i-1) {
			continue
		}
		for j := i; j < n-1; j++ {
			if isP(i, j) && isP(j+1, n-1) {
				return true
			}
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
