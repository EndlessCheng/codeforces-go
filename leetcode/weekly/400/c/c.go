package main

import "math/bits"

// https://space.bilibili.com/206214
func clearStars(S string) string {
	s := []byte(S)
	st := make([][]int, 26)
	mask := 0
	for i, c := range s {
		if c != '*' {
			c -= 'a'
			st[c] = append(st[c], i)
			mask |= 1 << c
		} else {
			k := bits.TrailingZeros(uint(mask))
			p := st[k]
			s[p[len(p)-1]] = '*'
			st[k] = p[:len(p)-1]
			if len(st[k]) == 0 {
				mask ^= 1 << k
			}
		}
	}

	t := s[:0]
	for _, c := range s {
		if c != '*' {
			t = append(t, c)
		}
	}
	return string(t)
}
