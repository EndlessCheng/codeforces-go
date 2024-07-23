package main

import (
	"strings"
)

// https://space.bilibili.com/206214
func minimumOperations(num string) int {
	n := len(num)
	var found0, found5 bool
	for i := n - 1; i >= 0; i-- {
		c := num[i]
		if found0 && (c == '0' || c == '5') || 
		   found5 && (c == '2' || c == '7') {
			return n - i - 2
		}
		if c == '0' {
			found0 = true
		} else if c == '5' {
			found5 = true
		}
	}
	if found0 {
		return n - 1
	}
	return n
}

func minimumOperations2(num string) int {
	ans := len(num)
	if strings.Contains(num, "0") { // 删除 len(num)-1 次得到 "0"
		ans--
	}
	f := func(tail string) {
		i := strings.LastIndexByte(num, tail[1])
		if i < 0 {
			return
		}
		i = strings.LastIndexByte(num[:i], tail[0])
		if i < 0 {
			return
		}
		ans = min(ans, len(num)-i-2)
	}
	f("00")
	f("25")
	f("50")
	f("75")
	return ans
}
