package main

import "math"

// https://space.bilibili.com/206214
const keys = "elohtcd"
const full = 2012 // 0b11111011100

// pos：字母在二进制上的起始位置
// limit：这个字母能选择的上限
// mask：位掩码
var rules = ['z' + 1]struct{ pos, limit, mask int }{
	'e': {0, 4, 7},
	'l': {3, 3, 3},
	'o': {5, 2, 3},
	'h': {7, 1, 1},
	't': {8, 1, 1},
	'c': {9, 1, 1},
	'd': {10, 1, 1},
}

func merge(cur, add int) int {
	for _, c := range keys {
		r := rules[c]
		c1 := cur >> r.pos & r.mask
		c2 := add >> r.pos & r.mask
		if c1+c2 > r.limit {
			return -1
		}
		cur += c2 << r.pos
	}
	return cur
}

func Leetcode(words []string) int {
	const inf = math.MaxInt32 / 2
	n := len(words)
	costs := make([][1 << 11]int, n)
	for i, word := range words {
		for j := range costs[i] {
			costs[i][j] = inf
		}
		var f func(string, int, int)
		f = func(s string, mask, tot int) {
			costs[i][mask] = min(costs[i][mask], tot)
			for j, c := range s {
				r := rules[c]
				if mask>>r.pos&r.mask < r.limit {
					f(s[:j]+s[j+1:], mask+1<<r.pos, tot+j*(len(s)-1-j))
				}
			}
		}
		f(word, 0, 0)
	}

	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, 1<<11)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	var f func(int, int) int
	f = func(i, mask int) int {
		if i == n {
			if mask == full {
				return 0
			}
			return inf
		}
		ptr := &dp[i][mask]
		if *ptr != -1 {
			return *ptr
		}
		res := inf
		for add, tot := range costs[i] {
			if tot >= res { // 剪枝
				continue
			}
			m2 := merge(mask, add)
			if m2 >= 0 {
				res = min(res, f(i+1, m2)+tot)
			}
		}
		*ptr = res
		return res
	}
	ans := f(0, 0)
	if ans == inf {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
