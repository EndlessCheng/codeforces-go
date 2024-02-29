package main

import "slices"

// https://space.bilibili.com/206214
func finalString(s string) string {
	qs := [2][]rune{}
	dir := 1
	for _, c := range s {
		if c == 'i' {
			dir ^= 1
		} else {
			qs[dir] = append(qs[dir], c)
		}
	}
	slices.Reverse(qs[dir^1])
	return string(append(qs[dir^1], qs[dir]...))
}
