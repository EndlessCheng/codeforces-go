package main

import "unicode"

// https://space.bilibili.com/206214
func clearDigits(s string) string {
	st := []rune{}
	for _, c := range s {
		if unicode.IsDigit(c) {
			st = st[:len(st)-1]
		} else {
			st = append(st, c)
		}
	}
	return string(st)
}
