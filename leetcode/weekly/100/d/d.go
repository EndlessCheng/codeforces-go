package main

import "sort"

// github.com/EndlessCheng/codeforces-go
func orderlyQueue(s string, k int) (ans string) {
	smallestRepresentation := func(s []byte) []byte {
		n := len(s)
		s = append(s, s...)
		i := 0
		for j := 1; j < n; {
			k := 0
			for ; k < n && s[i+k] == s[j+k]; k++ {
			}
			if k >= n {
				break
			}
			if s[i+k] < s[j+k] {
				j += k + 1
			} else {
				i, j = j, max(j, i+k)+1
			}
		}
		return s[i : i+n]
	}
	t := []byte(s)
	if k == 1 {
		return string(smallestRepresentation(t))
	}
	sort.Slice(t, func(i, j int) bool { return t[i] < t[j] })
	return string(t)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
