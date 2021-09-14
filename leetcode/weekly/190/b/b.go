package main

import "strings"

// 我的憨憨写法
func maxVowels(s string, k int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(s)

	c := ['z' + 1]int{}
	for _, b := range s[:k] {
		c[b]++
	}
	for _, b := range "aeiou" {
		ans += c[b]
	}
	i, j := 0, k
	for ; j < n; j++ {
		c[s[i]]--
		c[s[j]]++
		s := 0
		for _, b := range "aeiou" {
			s += c[b]
		}
		ans = max(ans, s)
		i++
	}
	return
}

// 优雅的写法：前缀和
func maxVowels2(s string, k int) (ans int) {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	n := len(s)
	sum := make([]int, n+1)
	for i := range s {
		sum[i+1] = sum[i]
		if strings.IndexByte("aeiou", s[i]) >= 0 {
			sum[i+1]++
		}
	}
	for i := k; i <= n; i++ {
		ans = max(ans, sum[i]-sum[i-k])
	}
	return
}
