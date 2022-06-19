package main

import "math/bits"

// https://space.bilibili.com/206214/dynamic
func greatestLetter(s string) string {
	masks := [2]uint{}
	for _, c := range s {
		masks[c>>5&1] |= 1 << (c & 31)
	}
	mask := masks[0] & masks[1]
	if mask == 0 {
		return ""
	}
	return string(byte('A' + bits.Len(mask) - 2))
}

func greatestLetter2(s string) string {
	has := map[rune]bool{}
	for _, v := range s {
		has[v] = true
	}
	for i := 'Z'; i >= 'A'; i-- {
		if has[i] && has[i-'A'+'a'] {
			return string(i)
		}
	}
	return ""
}
