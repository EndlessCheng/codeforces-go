package main

import (
	"slices"
	"strings"
)

// https://space.bilibili.com/206214
func lexPalindromicPermutation1(s, target string) string {
	left := make([]int, 26)
	for _, b := range s {
		left[b-'a']++
	}
	valid := func() bool {
		for _, c := range left {
			if c < 0 {
				return false
			}
		}
		return true
	}

	midCh := ""
	for i, c := range left {
		if c%2 == 0 {
			continue
		}
		// s 不能有超过一个字母出现奇数次
		if midCh != "" {
			return ""
		}
		// 记录填在正中间的字母
		midCh = string('a' + byte(i))
		left[i]--
	}

	n := len(s)
	// 先假设答案左半与 t 的左半（不含正中间）相同
	for _, b := range target[:n/2] {
		left[b-'a'] -= 2
	}

	if valid() {
		// 特殊情况：把 target 左半翻转到右半，能否比 target 大？
		leftS := target[:n/2]
		tmp := []byte(leftS)
		slices.Reverse(tmp)
		rightS := midCh + string(tmp)
		if rightS > target[n/2:] { // 由于左半是一样的，所以只需比右半
			return leftS + rightS
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		b := target[i] - 'a'
		left[b] += 2  // 撤销消耗
		if !valid() { // [0,i-1] 无法做到全部一样
			continue
		}

		// 把 target[i] 增大到 j
		for j := b + 1; j < 26; j++ {
			if left[j] == 0 {
				continue
			}

			// 找到答案（下面的循环在整个算法中只会跑一次）
			left[j] -= 2
			ans := []byte(target[:i+1])
			ans[i] = 'a' + j

			// 中间可以随便填
			for k, c := range left {
				ch := string('a' + byte(k))
				ans = append(ans, strings.Repeat(ch, c/2)...)
			}

			// 镜像翻转
			rightS := slices.Clone(ans)
			slices.Reverse(rightS)
			ans = append(ans, midCh...)
			ans = append(ans, rightS...)

			return string(ans)
		}
		// 增大失败，继续枚举
	}
	return ""
}

func lexPalindromicPermutation(s, target string) string {
	left := make([]int, 26)
	for _, b := range s {
		left[b-'a']++
	}

	midCh := ""
	for i, c := range left {
		if c%2 == 0 {
			continue
		}
		// s 不能有超过一个字母出现奇数次
		if midCh != "" {
			return ""
		}
		// 记录填在正中间的字母
		midCh = string('a' + byte(i))
		left[i]--
	}

	n := len(s)
	// 先假设答案左半与 t 的左半（不含正中间）相同
	for _, b := range target[:n/2] {
		left[b-'a'] -= 2
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
		// 特殊情况：把 target 左半翻转到右半，能否比 target 大？
		leftS := target[:n/2]
		tmp := []byte(leftS)
		slices.Reverse(tmp)
		rightS := midCh + string(tmp)
		if rightS > target[n/2:] { // 由于左半是一样的，所以只需比右半
			return leftS + rightS
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
		ans := []byte(target[:i+1])
		ans[i] = 'a' + j

		// 中间可以随便填
		for k, c := range left {
			ch := string('a' + byte(k))
			ans = append(ans, strings.Repeat(ch, c/2)...)
		}

		// 镜像翻转
		rightS := slices.Clone(ans)
		slices.Reverse(rightS)
		ans = append(ans, midCh...)
		ans = append(ans, rightS...)

		return string(ans)
	}
	return ""
}
