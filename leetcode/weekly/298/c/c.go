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
		return n
	}
	ans := m
	v, _ := strconv.ParseInt(s[n-m:], 2, 64)
	if int(v) > k {
		ans--
	}
	return ans + strings.Count(s[:n-m], "0")
}
