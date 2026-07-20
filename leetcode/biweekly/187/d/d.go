package main

import (
	"math"
	"strings"
)

// https://space.bilibili.com/206214
func minCost1(source, target string, rules [][]string, costs []int) int {
	// 把成本加上 '*' 的个数，这样在 DP 中无需反复统计 '*' 的个数
	for i, rule := range rules {
		costs[i] += strings.Count(rule[0], "*")
	}

	isMatch := func(pattern, s string) bool {
		for j, ch := range pattern {
			if ch != '*' && byte(ch) != s[j] {
				return false
			}
		}
		return true
	}

	n := len(source)
	memo := make([]int, n)
	for i := range memo {
		memo[i] = -1 // -1 表示该状态没有计算过
	}

	var dfs func(int) int
	dfs = func(i int) int {
		if i < 0 {
			return 0
		}
		p := &memo[i]
		if *p != -1 { // 之前算过了
			return *p
		}

		res := math.MaxInt / 2 // 防止加法溢出

		// 不替换下标 i
		if source[i] == target[i] {
			res = dfs(i - 1)
		}

		// 替换子串 [i-m+1, i]
		for k, rule := range rules {
			replacement := rule[1]
			left := i - len(replacement) + 1 // 子串左端点
			if left >= 0 && replacement == target[left:i+1] && isMatch(rule[0], source[left:]) {
				res = min(res, dfs(left-1)+costs[k])
			}
		}

		*p = res // 记忆化
		return res
	}

	ans := dfs(n - 1)
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

func minCost2(source, target string, rules [][]string, costs []int) int {
	// 把成本加上 '*' 的个数，这样在 DP 中无需反复统计 '*' 的个数
	for i, rule := range rules {
		costs[i] += strings.Count(rule[0], "*")
	}

	isMatch := func(pattern, s string) bool {
		for j, ch := range pattern {
			if ch != '*' && byte(ch) != s[j] {
				return false
			}
		}
		return true
	}

	n := len(source)
	f := make([]int, n+1)
	for i := range n {
		res := math.MaxInt / 2 // 防止加法溢出

		// 不替换下标 i
		if source[i] == target[i] {
			res = f[i]
		}

		// 替换子串 [i-m+1, i]
		for k, rule := range rules {
			replacement := rule[1]
			left := i - len(replacement) + 1 // 子串左端点
			if left >= 0 && f[left]+costs[k] < res && replacement == target[left:i+1] && isMatch(rule[0], source[left:]) {
				res = f[left] + costs[k]
			}
		}

		f[i+1] = res
	}

	ans := f[n]
	if ans < math.MaxInt/2 {
		return ans
	}
	return -1
}

func minCost(source, target string, rules [][]string, costs []int) int {
	// 把成本加上 '*' 的个数，这样在 DP 中无需反复统计 '*' 的个数
	for i, rule := range rules {
		costs[i] += strings.Count(rule[0], "*")
	}

	isMatch := func(pattern, s string) bool {
		for j, ch := range pattern {
			if ch != '*' && byte(ch) != s[j] {
				return false
			}
		}
		return true
	}

	n := len(source)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		f[i] = math.MaxInt
	}

	for i := range n {
		if f[i] == math.MaxInt {
			continue
		}

		// 不替换下标 i
		if source[i] == target[i] {
			f[i+1] = min(f[i+1], f[i])
		}

		// 替换子串 [i, i+len(replacement)-1]
		for k, rule := range rules {
			replacement := rule[1]
			r := i + len(replacement)
			if r <= n && f[i]+costs[k] < f[r] && replacement == target[i:r] && isMatch(rule[0], source[i:]) {
				f[r] = f[i] + costs[k]
			}
		}
	}

	if f[n] < math.MaxInt {
		return f[n]
	}
	return -1
}
