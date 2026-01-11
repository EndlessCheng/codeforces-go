package main

import "math/bits"

// https://space.bilibili.com/206214
func residuePrefixes1(s string) (ans int) {
	set := map[rune]struct{}{}
	for i, ch := range s {
		set[ch] = struct{}{}
		if len(set) == (i+1)%3 {
			ans++
		}
	}
	return
}

func residuePrefixes(s string) (ans int) {
	set := 0
	for i, ch := range s {
		set |= 1 << (ch - 'a')
		if bits.OnesCount(uint(set)) == (i+1)%3 {
			ans++
		}
	}
	return
}
