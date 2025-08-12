package main

import (
	"strconv"
	"strings"
)

func maximum69Number1(num int) int {
	s := strconv.Itoa(num)
	s = strings.Replace(s, "6", "9", 1) // 替换第一个 6
	ans, _ := strconv.Atoi(s)
	return ans
}

func maximum69Number(num int) int {
	maxBase := 0
	base := 1
	for x := num; x > 0; x /= 10 {
		if x%10 == 6 {
			maxBase = base
		}
		base *= 10
	}
	return num + maxBase*3
}
