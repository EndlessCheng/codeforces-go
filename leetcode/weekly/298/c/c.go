package main

import (
	"math/bits"
	"strconv"
	"strings"
)

// https://space.bilibili.com/206214/dynamic
func longestSubsequence(s string, k int) int {
	n, m := len(s), bits.Len(uint(k))
	if n < m {
		return n // 全选
	}
	ans := m
	sufVal, _ := strconv.ParseInt(s[n-m:], 2, 0) // 找后缀
	if int(sufVal) > k {
		ans--
	}
	return ans + strings.Count(s[:n-m], "0") // 添加前导零
}
