package main

import "bytes"

// https://space.bilibili.com/206214
func removeDuplicates(s string, k int) string {
	type pair struct {
		b byte
		c int
	}
	st := []pair{{}}
	for i := range s {
		b := s[i]
		if st[len(st)-1].b == b && st[len(st)-1].c == k-1 {
			st = st[:len(st)-1]
		} else if st[len(st)-1].b == b {
			st[len(st)-1].c++
		} else {
			st = append(st, pair{b, 1})
		}
	}
	ans := []byte{}
	for _, p := range st {
		ans = append(ans, bytes.Repeat([]byte{p.b}, p.c)...)
	}
	return string(ans)
}
