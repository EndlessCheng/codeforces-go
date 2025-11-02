package main

import (
	"slices"
	"strings"
)

// https://space.bilibili.com/206214
func lexPalindromicPermutation(s, target string) string {
	left := make([]int, 26)
	for _, b := range s {
		left[b-'a']++
	}

	midCh := byte(0)
	for i, c := range left {
		if c%2 == 0 {
			continue
		}
		// s 不能有超过一个字母出现奇数次
		if midCh > 0 {
			return ""
		}
		// 记录填在正中间的字母
		midCh = 'a' + byte(i)
		left[i]--
	}

	n := len(s)
	ans := []byte(target)
	// 先假设答案左半与 t 的左半（不含正中间）相同
	for i, b := range target[:n/2] {
		left[b-'a'] -= 2
		ans[n-1-i] = byte(b) // 把 target 左半翻转到右半
	}
	// 正中间只能填那个出现奇数次的字母
	if midCh > 0 {
		ans[n/2] = midCh
	}

	neg, leftMax := 0, byte(0)
	for i, cnt := range left {
		if cnt < 0 {
			neg++ // 统计 left 中的负数个数
		} else if cnt > 0 {
			leftMax = max(leftMax, byte(i)) // 剩余可用字母的最大值
		}
	}

	if neg == 0 {
		// 把 target 左半翻转到右半，能否比 target 大？
		t := string(ans)
		if t > target {
			return t
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		b := target[i] - 'a'
		left[b] += 2 // 撤销消耗

		if left[b] == 0 {
			neg--
		} else if left[b] == 2 {
			leftMax = max(leftMax, b)
		}

		// left 有负数 or 没有大于 target[i] 的字母
		if neg > 0 || leftMax <= b {
			continue
		}

		// 找到答案（下面的循环在整个算法中只会跑一次）
		j := b + 1
		for left[j] == 0 {
			j++
		}

		// 把 target[i] 和 target[n-1-i] 增大到 j
		left[j] -= 2
		ans[i] = 'a' + j
		ans[n-1-i] = ans[i]

		// 中间的空位可以随便填
		t := make([]byte, 0, n-(i+1)*2)
		for k, c := range left {
			ch := string('a' + byte(k))
			t = append(t, strings.Repeat(ch, c/2)...)
		}

		// 把 t、midCh、Reverse(t) 依次填在 ans[i] 的右边
		a := append(ans[:i+1], t...)
		if midCh > 0 {
			a = append(a, midCh)
		}
		slices.Reverse(t)
		a = append(a, t...)

		return string(ans)
	}
	return ""
}
