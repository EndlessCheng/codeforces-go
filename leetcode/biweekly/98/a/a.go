package main

import (
	"strconv"
	"strings"
)

// https://space.bilibili.com/206214
func minMaxDifference(num int) int {
	s := strconv.Itoa(num)
	mx := num
	for _, c := range s {
		if c != '9' { // 第一个不等于 9 的字符，替换成 9
			mx, _ = strconv.Atoi(strings.ReplaceAll(s, string(c), "9"))
			break
		}
	}
	// 第一个不等于 0 的字符，替换成 0
	mn, _ := strconv.Atoi(strings.ReplaceAll(s, s[:1], "0"))
	return mx - mn
}
