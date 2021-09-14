package main

import "bytes"

// github.com/EndlessCheng/codeforces-go
func getSmallestString(n int, k int) (ans string) {
	s := bytes.Repeat([]byte{'a'}, n)
	k -= n
	for i := n - 1; i >= 0; i-- {
		if k < 26 {
			s[i] += byte(k)
			break
		}
		s[i] = 'z'
		k -= 25
	}
	return string(s)
}
