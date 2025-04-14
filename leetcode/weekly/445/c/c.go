package main

import (
	"bytes"
	"slices"
)

// https://space.bilibili.com/206214
func smallestPalindrome(s string, k int) string {
	n := len(s)
	m := n / 2

	total := [26]int{}
	for _, b := range s[:m] {
		total[b-'a']++
	}

	cnt := make([]int, 26)
	perm := 1
	i, j := m-1, 25
	// 倒着计算排列数
	for ; i >= 0 && perm < k; i-- {
		for cnt[j] == total[j] {
			j--
		}
		cnt[j]++
		perm = perm * (m - i) / cnt[j]
	}

	if perm < k {
		return ""
	}

	ans := make([]byte, 0, n)
	// 已经有足够的排列数了，<= i 的位置直接填字典序最小的排列
	for ch, c := range cnt[:j+1] {
		ans = append(ans, bytes.Repeat([]byte{'a' + byte(ch)}, total[ch]-c)...)
	}

	// 试填法
	j0 := j
	for i++; i < m; i++ {
		for j := j0; j < 26; j++ {
			if cnt[j] == 0 {
				continue
			}
			// 假设填字母 j，根据 perm = p * (m - i) / cnt[j] 倒推 p
			p := perm * cnt[j] / (m - i)
			if p >= k {
				ans = append(ans, 'a'+byte(j))
				cnt[j]--
				perm = p
				break
			}
			k -= p
		}
	}

	rev := slices.Clone(ans)
	if n%2 > 0 {
		ans = append(ans, s[n/2])
	}
	slices.Reverse(rev)
	ans = append(ans, rev...)
	return string(ans)
}

func smallestPalindrome1(s string, k int) string {
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
	ans := make([]byte, m, n)
	for i := range m {
		for j := range cnt {
			if cnt[j] == 0 {
				continue
			}
			cnt[j]--             // 假设填字母 j，看是否有足够的排列
			p := perm(m - i - 1) // 剩余位置的排列个数
			if p >= k {          // 有足够的排列
				ans[i] = 'a' + byte(j)
				break
			}
			k -= p // k 太大，要填更大的字母（类似搜索树剪掉了一个大小为 p 的子树）
			cnt[j]++
		}
	}

	rev := slices.Clone(ans)
	if n%2 > 0 {
		ans = append(ans, s[n/2])
	}
	slices.Reverse(rev)
	ans = append(ans, rev...)
	return string(ans)
}
