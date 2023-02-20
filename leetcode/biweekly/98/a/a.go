package main

import (
	"strconv"
	"strings"
)

// https://space.bilibili.com/206214
func minMaxDifference(num int) int {
	mx := num
	s := strconv.Itoa(num)
	for _, c := range s {
		if c != '9' {
			mx, _ = strconv.Atoi(strings.ReplaceAll(s, string(c), "9"))
			break
		}
	}
	mn, _ := strconv.Atoi(strings.ReplaceAll(s, s[:1], "0"))
	return mx - mn
}
