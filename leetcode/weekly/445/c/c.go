package main

import "slices"

// https://space.bilibili.com/206214
func smallestPalindrome(s string, k int) string {
	n := len(s)
	m := n / 2

	cnt := make([]int, 26)
	for _, b := range s[:m] {
		cnt[b-'a']++
	}

	// 为什么这样做是对的？见 62. 不同路径 我的题解
	comb := func(n, m int) int {
		m = min(m, n-m)
		res := 1
		for i := 1; i <= m; i++ {
			res = res * (n + 1 - i) / i
			if res >= k { // 太大了
				return k
			}
		}
		return res
	}

	// 计算长为 sz 的字符串的排列个数
	perm := func(sz int) int {
		res := 1
		for _, c := range cnt {
			if c == 0 {
				continue
			}
			// 先从 sz 个里面选 c 个位置填当前字母
			res *= comb(sz, c)
			if res >= k { // 太大了
				return k
			}
			// 从剩余位置中选位置填下一个字母
			sz -= c
		}
		return res
	}

	// k 太大
	if perm(m) < k {
		return ""
	}

	// 构造回文串的左半部分
	leftS := make([]byte, m)
	for i := range leftS {
		for j := range cnt {
			if cnt[j] == 0 {
				continue
			}
			cnt[j]-- // 假设填字母 j，看是否有足够的排列
			p := perm(m - i - 1) // 剩余位置的排列个数
			if p >= k { // 有足够的排列
				leftS[i] = 'a' + byte(j)
				break
			}
			k -= p // k 太大，要填更大的字母
			cnt[j]++
		}
	}

	ans := string(leftS)
	if n%2 > 0 {
		ans += string(s[n/2])
	}
	slices.Reverse(leftS)
	return ans + string(leftS)
}
