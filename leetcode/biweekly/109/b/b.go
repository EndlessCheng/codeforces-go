package main

import (
	"slices"
)

// https://space.bilibili.com/206214
func sortVowels(s string) string {
	a := []byte{}
	for _, c := range s {
		if 2130466>>(c&31)&1 > 0 {
			a = append(a, byte(c))
		}
	}
	slices.Sort(a)

	t, j := []byte(s), 0
	for i, c := range t {
		if 2130466>>(c&31)&1 > 0 {
			t[i] = a[j]
			j++
		}
	}
	return string(t)
}
